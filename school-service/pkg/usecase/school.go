package usecase

import (
	"fmt"

	repoInterface "github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/repository/interfaces"

	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/domain"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/utils"
)

const (
	minStudents = 30
	maxStudent  = 40

	minClass = 1
	maxClass = 10
)

type schoolUseCase struct {
	schoolRepo repoInterface.SchoolRepo

	startChan chan struct{}
}

func NewSchoolUseCase(schoolRepo repoInterface.SchoolRepo) interfaces.SchoolUseCase {
	return &schoolUseCase{
		schoolRepo: schoolRepo,
		startChan:  make(chan struct{}, 5),
	}
}

func (s *schoolUseCase) Create(name string) domain.School {

	// select random class count
	randClassCount := utils.GetIntBetween(minClass, maxClass)

	// create a slice of class with the rand size
	classes := make([]domain.Class, randClassCount)
	classChan := make(chan domain.Class, 5)

	// fire create classes in a separate goroutines
	go func() {

		for i := 1; i <= randClassCount; i++ {
			s.startChan <- struct{}{} // check firing a new goroutines is allowed or not
			go s.createClass(fmt.Sprintf("class-%d", i), classChan)
		}
	}()

	// receive classes
	for i := range classes {
		classes[i] = <-classChan
		// release the start channel for new goroutines to run
		<-s.startChan
	}

	// return the school
	return domain.School{
		Name:    name,
		Classes: classes,
	}
}
