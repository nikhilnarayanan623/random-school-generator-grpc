//go:build wireinject
// +build wireinject

package di

import (
	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/api"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/api/handler"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/client"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/config"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	wire.Build(
		client.NewSchoolClient,
		handler.NewSchoolHandler,
		api.NewServerHTTP,
	)

	return &api.Server{}, nil
}
