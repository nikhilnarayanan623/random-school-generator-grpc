package usecase

import "school-service/pkg/usecase/interfaces"

type schoolUseCase struct {
}

func NewSchoolUseCase() interfaces.SchoolUseCase {
	return &schoolUseCase{}
}
