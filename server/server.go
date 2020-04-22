package server

import "net/http"

type (
	Category struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	Categories []Category
)

type CategoryHandler interface {
	Handle() http.HandlerFunc
}
