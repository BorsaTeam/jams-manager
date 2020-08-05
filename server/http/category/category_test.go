package category

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/database/repository"
)

func TestNewCategoryHandler(t *testing.T) {
	type in struct {
		method   string
		category server.Category
		repo     repository.Category
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
			in: in{
				method:   http.MethodGet,
				category: server.Category{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "success delete",
			in: in{
				method:   http.MethodDelete,
				category: server.Category{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			}(),
		},
		{
			name: "error status 405",
			in: in{
				method:   http.MethodPut,
				category: server.Category{},
			},
			out: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(405)
				}
			}(),
		},
		{
			name: "run method not allowed",
			in: in{
				method: http.MethodTrace,
				category: server.Category{
					Id:        "",
					Name:      "",
					CreatedAt: time.Time{},
					UpdatedAt: nil,
				},
				repo: categoryRepoMock{},
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
			h := NewHandler(categoryRepoMock{})

			body, _ := json.Marshal(tt.in.category)

			r, _ := http.NewRequest(tt.in.method, "/categories", bytes.NewReader(body))

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

type categoryRepoMock struct {
	err error
	id  string
}

func (c categoryRepoMock) Save(category repository.CategoryEntity) (repository.CategoryId, error) {
	return repository.CategoryId(c.id), nil
}
func (c categoryRepoMock) FindAll() ([]repository.CategoryEntity, error) {
	return nil, nil
}

func (c categoryRepoMock) Delete(id repository.CategoryId) error {
	return nil
}
