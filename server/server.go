package server

import "net/http"

type (
	Category struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	Categories []Category

	Pilot struct {
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
	Pilots []Pilot
)

type CategoryHandler interface {
	Handle() http.HandlerFunc
}

type PilotHandler interface {
	Handle() http.HandlerFunc
}
