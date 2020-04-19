package main

import (
	"github.com/BorsaTeam/jams-manager/server/http/category"
	"github.com/BorsaTeam/jams-manager/server/http/pilot"
	"log"
	"net/http"
)

func main() {
	pilotHandler := pilot.NewHandler()

	http.Handle("/categories", category.Handle())
	http.Handle("/categories/", category.Handle())

	http.Handle("/pilots", pilotHandler.Handle())
	http.Handle("/pilots/", pilotHandler.Handle())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
