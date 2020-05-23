package category

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BorsaTeam/jams-manager/server"
)

func TestNewCategoryHandler(t *testing.T) {
	type Fields struct {
		method   string
		category server.Category
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
				category: server.Category{
					Name: "pro",
				},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "success get",
			fields: Fields{
				method: http.MethodGet,
				category: server.Category{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "success delete",
			fields: Fields{
				method: http.MethodDelete,
				category: server.Category{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "error status 405",
			fields: Fields{
				method: http.MethodPut,
				category: server.Category{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(405)
				}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := NewHandler()

			body, _ := json.Marshal(tt.fields.category)

			r, _ := http.NewRequest(tt.fields.method, "/categories", bytes.NewReader(body))

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
