//go:build wireinject
// +build wireinject

package di

import (
	"school-service/pkg/api"
	"school-service/pkg/api/service"
	"school-service/pkg/config"
	"school-service/pkg/usecase"
	"school-service/pkg/repository"
	"github.com/google/wire"
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
