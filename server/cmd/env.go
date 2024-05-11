package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	pgConnString string
	password     string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		pgConnStringEnv = "PG_CONN"
		password     = "PASSWORD"
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

	return cfg, nil
}