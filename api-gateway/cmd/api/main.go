package main

import (
	"log"

	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatalf("failed to initialize api : %v", err)
	}

	server.Start()
}
