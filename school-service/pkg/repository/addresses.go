package repository

import (
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/domain"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/utils"
)

const (
	minHouseNum = 1
	maxHouseNum = 666
)

func (s *schoolRepo) getAddress() domain.Address {

	s.mu.RLock()

	randDisIdx := utils.GetIntBetween(0, len(s.states)-1)
	randCityIdx := utils.GetIntBetween(0, len(s.states[randDisIdx].Cities)-1)

	city := s.states[randDisIdx].Cities[randCityIdx]
	state := s.states[randDisIdx].Name

	s.mu.RUnlock()

	randHouseNo := utils.GetIntBetween(minHouseNum, maxHouseNum)

	return domain.Address{
		HouseNumber: randHouseNo,
		State:       state,
		City:        city,
	}

}

func (s *schoolRepo) getName() string {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.names[utils.GetIntBetween(0, len(s.names)-1)]
}
