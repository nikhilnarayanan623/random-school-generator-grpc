package interfaces

import "github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/domain"

type SchoolRepo interface {
	GetStudent() domain.Student
}
