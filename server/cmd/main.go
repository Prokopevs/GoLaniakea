package main

import (
	"github/Prokopevs/GoLaniakea/db"
	"github/Prokopevs/GoLaniakea/internal/handler/post"
	"github/Prokopevs/GoLaniakea/internal/repository/post"
	"github/Prokopevs/GoLaniakea/internal/router"
	"github/Prokopevs/GoLaniakea/internal/services/post"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	postRep := repository.NewPostRepository(dbConn.GetDB())
	postSvc := services.NewService(postRep)
	postHandler := handler.NewHandler(postSvc)

	router.InitRouter(postHandler)

	router.Start("0.0.0.0:8080")
}