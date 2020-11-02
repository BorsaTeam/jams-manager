package server

import (
	"fmt"
	"net/http"
	"time"
)

type (
	CategoryId string

	Category struct {
		Id        CategoryId `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	}
	Categories []Category

	Score struct {
		Score   float32 `json:"score"`
		Id      string  `json:"id,omitempty"`
		RiderId string  `json:"riderId"`
	}

	Rider struct {
		Id               string   `json:"id"`
		Name             string   `json:"name"`
		Age              int      `json:"age"`
		Gender           string   `json:"gender"`
		City             string   `json:"city"`
		Email            string   `json:"email"`
		PaidSubscription bool     `json:"paidSubscription"`
		Sponsors         []string `json:"sponsors"`
		CategoryId       string   `json:"categoryId"`
	}
	Riders []Rider

	Metadata struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		PageCount  int `json:"page_count"`
		TotalCount int `json:"total_count"`
	}

	Pagination struct {
		Meta    Metadata      `json:"_metadata"`
		Results interface{} `json:"results"`
	}
)

type CategoryHandler interface {
	Handle() http.HandlerFunc
}

type RiderHandler interface {
	Handle() http.HandlerFunc
}

type ScoreHandler interface {
	Handle() http.HandlerFunc
}

type JamsError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (j JamsError) Error() string {
	return fmt.Sprintf("[%s] %s", j.Code, j.Message)
}

type PageRequest struct {
	Page    int
	PerPage int
}
