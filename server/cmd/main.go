package main

import (
	"context"
	"fmt"
	"github/Prokopevs/GoLaniakea/db"
	"github/Prokopevs/GoLaniakea/internal/repository/post"
	"github/Prokopevs/GoLaniakea/internal/services/post"
	"github/Prokopevs/GoLaniakea/internal/transport/http/servers/post/handler"
	"github/Prokopevs/GoLaniakea/internal/transport/http/servers/post/router"
	"os"
)

const (
	exitCodeInitError = 2
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	dbConn, err := db.NewDatabase(context.Background(), cfg.pgConnString)
	if err != nil {
		return err
	}

	postRep := repository.NewPostRepository(dbConn.GetDB())
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitCodeInitError)
	}
	

	
	postSvc := services.NewService(postRep)
	postHandler := handler.NewHandler(postSvc)

	router.InitRouter(postHandler)

	router.Start("0.0.0.0:8080")
}