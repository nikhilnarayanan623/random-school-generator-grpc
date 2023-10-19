package interfaces

import "school-service/pkg/domain"

type SchoolRepo interface {
	GetStudent() domain.Student
}
