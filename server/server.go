package server

import (
	"net/http"
	"time"
)

type (
	CategoryId string

	Category struct {
		Id        CategoryId `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time  `json:"updatedAt,omitempty"`
	}
	Categories []Category

	Rider struct {
		Id               string   `json:"id"`
		Name             string   `json:"name"`
		Age              int      `json:"age"`
		Gender           string   `json:"gender"`
		City             string   `json:"city"`
		Cpf              string   `json:"cpf"`
		PaidSubscription bool     `json:"paidSubscription"`
		Sponsors         []string `json:"sponsors"`
		CategoryId       string   `json:"categoryId"`
	}
	Riders []Rider
)

type CategoryHandler interface {
	Handle() http.HandlerFunc
}

type RiderHandler interface {
	Handle() http.HandlerFunc
}

type JamsError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
