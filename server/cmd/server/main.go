package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Prokopevs/GoLaniakea/server/db"
	"github.com/Prokopevs/GoLaniakea/server/internal/service"
	"github.com/Prokopevs/GoLaniakea/server/internal/transport/http/servers/post/handler"
	"go.uber.org/zap"
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

	postSvc := service.NewServiceImpl(dbConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	httpServer := handler.NewHTTP(cfg.httpAddr, cfg.password, sugaredLogger, postSvc)

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context) {
		httpServer.Run(ctx)
		wg.Done()
	}(ctx)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan
	cancel()

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitCodeInitError)
	}
}
