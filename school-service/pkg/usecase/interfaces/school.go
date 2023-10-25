package interfaces

import "github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/domain"

type SchoolUseCase interface {
	Create(name string) domain.School
}
