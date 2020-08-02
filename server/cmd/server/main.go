package main

import (
	"log"
	"net/http"

	"github.com/BorsaTeam/jams-manager/server/database"
	"github.com/BorsaTeam/jams-manager/server/database/repository"
	"github.com/BorsaTeam/jams-manager/server/http/category"
	"github.com/BorsaTeam/jams-manager/server/http/riders"
	"github.com/BorsaTeam/jams-manager/server/http/score"
)

func main() {
	dbConnection := database.NewPgManager()
	dbConnection.TestConnection()

	riderRepository := repository.NewRiderRepository(dbConnection)
	scoreRepository := repository.NewScoreRepository(dbConnection)
	categoryRepo := repository.NewCategory(dbConnection)

	categoryManager := category.NewHandler(categoryRepo)
	riderHandler := rider.NewHandler(riderRepository)
	scoreHandler := score.NewHandler(scoreRepository)

	http.Handle("/categories", categoryManager.Handle())
	http.Handle("/categories/", categoryManager.Handle())

	http.Handle("/riders", riderHandler.Handle())
	http.Handle("/riders/", riderHandler.Handle())

	http.Handle("/scores", scoreHandler.Handle())

	log.Println("Running jams-manager at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
