package main

import (
	"log"
	"net/http"

	"github.com/BorsaTeam/jams-manager/server/http/category"
	"github.com/BorsaTeam/jams-manager/server/http/riders"
)

func main() {
	categoryManager := category.NewCategoryHandler()
	riderHandler := rider.NewHandler()

	http.Handle("/categories", categoryManager.Handle())
	http.Handle("/categories/", categoryManager.Handle())

	http.Handle("/riders", riderHandler.Handle())
	http.Handle("/riders/", riderHandler.Handle())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
