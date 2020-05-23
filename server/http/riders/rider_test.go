package rider

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/database/repository"
)

func TestNewRiderHandler(t *testing.T) {
	type Fields struct {
		method string
		rider  server.Rider
	}

	tests := []struct {
		name   string
		fields Fields
		out    http.HandlerFunc
	}{
		{
			name: "success post",
			fields: Fields{
				method: http.MethodPost,
				rider: server.Rider{
					Name:             "Paibic o cara do whip",
					Age:              27,
					Gender:           "Male",
					City:             "Joinville",
					Cpf:              "0933452212",
					PaidSubscription: false,
					Sponsors:         []string{"EcoFun", "Velho Barreiro", "Honda"},
					CategoryId:       "Pro",
				},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(repositoryMock{
				riderId: "",
				error:   nil,
			})

			body, _ := json.Marshal(tt.fields.rider)

			r, _ := http.NewRequest(tt.fields.method, "/riders", bytes.NewReader(body))

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
	riderId string
	error   error
}

func (r repositoryMock) Save(rider repository.RiderEntity) (string, error) {
	return r.riderId, r.error
}
