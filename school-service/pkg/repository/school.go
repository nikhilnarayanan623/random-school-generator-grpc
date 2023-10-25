package repository

import (
	"sync"

	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/domain"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/sheet"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/utils"
)

const (
	MinAge = 12
	MaxAge = 18
)

var (
	namesSheetFileName = "./inputs/names.xls"
	namesSheetName     = "names"

	statesSheetFileName = "./inputs/states.xlsx"
	statesSheetName     = "states"
)

type schoolRepo struct {
	mu     sync.RWMutex
	states []domain.State
	names  []string
}

func NewSchoolUseCase() (interfaces.SchoolRepo, error) {

	names, err := sheet.GetAllNames(namesSheetFileName, namesSheetName)
	if err != nil {
		return nil, err
	}

	states, err := sheet.GetAllStates(statesSheetFileName, statesSheetName)
	if err != nil {
		return nil, err
	}

	return &schoolRepo{
		mu:     sync.RWMutex{},
		states: states,
		names:  names,
	}, nil
}

func (s *schoolRepo) GetStudent() domain.Student {

	// in a class all students age should be almost same(like : 12,13,14)
	randomAge := utils.GetIntBetween(MinAge, MaxAge)

	// new min and max age so in this call all students age will be in this new min and max age range
	newMinAge, newMaxAge := randomAge-1, randomAge+1

	return domain.Student{
		Name:           s.getName(),
		Age:            uint(utils.GetIntBetween(newMinAge, newMaxAge)),
		Gender:         s.getGender(),
		Scores:         s.getAllSubjects(),
		HaveDisability: utils.LowProbability(),
		Address:        s.getAddress(),
	}
}
