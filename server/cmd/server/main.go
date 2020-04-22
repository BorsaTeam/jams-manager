package main

import (
	"log"
	"net/http"

	"github.com/BorsaTeam/jams-manager/server/http/category"
)

func main() {
	categoryManager := category.NewCategoryHandler()

	http.Handle("/categories", categoryManager.Handle())
	http.Handle("/categories/", categoryManager.Handle())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
