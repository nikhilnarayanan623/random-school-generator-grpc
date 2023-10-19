package interfaces

import "school-service/pkg/domain"

type SchoolUseCase interface {
	Create(name string) domain.School
}
