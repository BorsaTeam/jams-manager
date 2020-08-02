package score

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/database/repository"
)

func TestHandle(t *testing.T) {
	type in struct {
		method string
		score  server.Score
		repo   repository.Score
	}

	tests := []struct {
		name string
		in   in
		out  http.HandlerFunc
	}{
		{
			name: "success post",
			in: in{
				method: http.MethodPost,
				score: server.Score{
					Score: 8.5,
				},
				repo: scoreRepoMock{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "run method not allowed",
			in: in{
				method: http.MethodTrace,
				score: server.Score{
					Score: 8.5,
				},
				repo: scoreRepoMock{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			}(),
		},
		{
			name: "error while saving score",
			in: in{
				method: http.MethodPost,
				score: server.Score{
					Score: 8.5,
				},
				repo: scoreRepoMock{err: errors.New("couldn't save data information")},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusUnprocessableEntity)
					_, _ = w.Write([]byte("Error while processing data score"))
				}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.in.repo)

			body, _ := json.Marshal(tt.in.score)

			r, _ := http.NewRequest(tt.in.method, "/scores", bytes.NewReader(body))

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

type scoreRepoMock struct {
	err error
	id  string
}

func (s scoreRepoMock) Save(score repository.ScoreEntity) (string, error) {
	return s.id, s.err
}
