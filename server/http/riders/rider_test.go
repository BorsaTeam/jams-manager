package rider

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/database/postgres/repository"
)

func TestNewRiderHandler(t *testing.T) {
	type in struct {
		method   string
		url      string
		rider    server.Rider
		repoMock repository.Rider
	}

	tests := []struct {
		name   string
		fields in
		out    http.HandlerFunc
	}{
		{
			name: "success post",
			fields: in{
				method: http.MethodPost,
				url:    "/riders",
				rider: server.Rider{
					Name:             "Paibic o cara do whip",
					Age:              27,
					Gender:           "Male",
					City:             "Joinville",
					Email:            "dogdacg@gmail.com",
					PaidSubscription: false,
					Sponsors:         []string{"EcoFun", "Velho Barreiro", "Honda"},
					CategoryId:       "Pro",
				},
				repoMock: repositoryMock{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "error to create rider",
			fields: in{
				method: http.MethodPost,
				url:    "/riders",
				rider: server.Rider{
					Name:             "Paibic o cara do whip",
					Age:              27,
					Gender:           "Male",
					City:             "Joinville",
					Email:            "dogdacg@gmail.com",
					PaidSubscription: false,
					Sponsors:         []string{"EcoFun", "Velho Barreiro", "Honda"},
					CategoryId:       "Pro",
				},
				repoMock: repositoryMock{
					saveErr: errors.New("error to create rider"),
				},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "success find all",
			fields: in{
				method:   http.MethodGet,
				url:      "/riders",
				rider:    server.Rider{},
				repoMock: repositoryMock{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "success find one",
			fields: in{
				method: http.MethodGet,
				url:    "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
				rider:  server.Rider{},
				repoMock: repositoryMock{
					rider: repository.RiderEntity{
						Id: "461a04e6-607f-40ac-b5ab-8c3490b09187",
					},
				},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "error 422 find one",
			fields: in{
				method: http.MethodGet,
				url:    "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
				rider:  server.Rider{},
				repoMock: repositoryMock{
					findOneErr: errors.New("error to find rider"),
				},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			}(),
		},
		{
			name: "rider not found",
			fields: in{
				method:   http.MethodGet,
				url:      "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
				rider:    server.Rider{},
				repoMock: repositoryMock{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}
			}(),
		},
		{
			name: "delete rider",
			fields: in{
				method:   http.MethodDelete,
				url:      "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
				repoMock: repositoryMock{rider: repository.RiderEntity{Id: "461a04e6-607f-40ac-b5ab-8c3490b09187"}},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "delete rider error find",
			fields: in{
				method:   http.MethodDelete,
				url:      "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
				repoMock: repositoryMock{findOneErr: errors.New("error to find rider")},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			}(),
		},
		{
			name: "rider not found delete",
			fields: in{
				method:   http.MethodDelete,
				url:      "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
				repoMock: repositoryMock{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}
			}(),
		},
		{
			name: "delete rider error",
			fields: in{
				method: http.MethodDelete,
				url:    "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
				repoMock: repositoryMock{
					rider: repository.RiderEntity{
						Id: "461a04e6-607f-40ac-b5ab-8c3490b09187",
					},
					deleteErr: errors.New("error to delete rider"),
				},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			}(),
		},
		{
			name: "method not allowed",
			fields: in{
				method: http.MethodPatch,
				url:    "/riders/461a04e6-607f-40ac-b5ab-8c3490b09187",
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.fields.repoMock)

			body, _ := json.Marshal(tt.fields.rider)

			r, _ := http.NewRequest(tt.fields.method, tt.fields.url, bytes.NewReader(body))

			w := httptest.NewRecorder()

			tt.out.ServeHTTP(w, r)

			g := httptest.NewRecorder()

			h.Handle().ServeHTTP(g, r)

			if g.Code != w.Code {
				t.Errorf("Handler returned wrong status code: got %v want %v", g.Code, w.Code)
			}
		})
	}
}

type repositoryMock struct {
	riderId    string
	rider      repository.RiderEntity
	findOneErr error
	saveErr    error
	deleteErr  error
}

func (r repositoryMock) FindOne(id string) (repository.RiderEntity, error) {
	return r.rider, r.findOneErr
}

func (r repositoryMock) Save(rider repository.RiderEntity) (string, error) {
	return r.riderId, r.saveErr
}

func (r repositoryMock) Delete(id string) error {
	return r.deleteErr
}
