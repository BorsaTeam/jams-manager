package main

import (
	"log"
	"net/http"

	"github.com/BorsaTeam/jams-manager/server/http/category"
	"github.com/BorsaTeam/jams-manager/server/http/pilot"
)

func main() {
	categoryManager := category.NewCategoryHandler()
	pilotHandler := pilot.NewHandler()

	http.Handle("/categories", categoryManager.Handle())
	http.Handle("/categories/", categoryManager.Handle())

	http.Handle("/pilots", pilotHandler.Handle())
	http.Handle("/pilots/", pilotHandler.Handle())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
