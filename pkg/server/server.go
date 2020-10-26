package server

import "github.com/bilalkocoglu/go-crud/pkg/config"

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
	return &Server{cfg: cfg}, nil
}
