package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	pgConnString string
	password     string
	httpAddr     string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		pgConnStringEnv = "PG_CONN"
		password        = "PASSWORD"
		httpAddr        = "HTTP_ADDR"
	)

	var ok bool

	cfg := &envConfig{}

	cfg.pgConnString, ok = os.LookupEnv(pgConnStringEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, pgConnStringEnv)
	}

	cfg.password, ok = os.LookupEnv(password)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, password)
	}

	cfg.httpAddr, ok = os.LookupEnv(httpAddr)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, httpAddr)
	}

	return cfg, nil
}
