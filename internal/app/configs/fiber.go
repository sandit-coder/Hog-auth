package configs

import (
	"errors"
	"os"
	"time"
)

const (
	ReadTimeout  = 15
	WriteTimeout = 15
	IdleTimeout  = 60
)

type Fiber struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewFiber() (*Fiber, error) {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		return nil, errors.New("environment variable HTTP_PORT is not set")
	}

	return &Fiber{
		Port:         ":" + port,
		ReadTimeout:  ReadTimeout * time.Second,
		WriteTimeout: WriteTimeout * time.Second,
		IdleTimeout:  IdleTimeout * time.Second,
	}, nil

}
