package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BorsaTeam/jams-manager/server/database/postgres"
	"github.com/BorsaTeam/jams-manager/server/database/postgres/repository"
	"github.com/BorsaTeam/jams-manager/server/http/category"
	"github.com/BorsaTeam/jams-manager/server/http/riders"
	"github.com/BorsaTeam/jams-manager/server/http/score"
)

func main() {
	migration := postgres.NewMigration()
	migration.Apply()

	dbConnection := postgres.NewPgManager()
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Running jams-manager at port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
