package main

import (
	"github.com/BorsaTeam/jams-manager/server/http/category"
	"log"
	"net/http"
)

func main() {
	http.Handle("/categories", category.Handle())
	http.Handle("/categories/", category.Handle())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
