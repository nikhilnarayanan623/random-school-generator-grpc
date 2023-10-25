//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/api"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/api/service"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/config"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/repository"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	wire.Build(
		repository.NewSchoolUseCase,
		usecase.NewSchoolUseCase,
		service.NewSchoolService,
		api.NewServerGRPC,
	)

	return &api.Server{}, nil
}
