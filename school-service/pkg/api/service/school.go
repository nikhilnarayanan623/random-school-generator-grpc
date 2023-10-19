package service

import (
	"school-service/pkg/pb"
	"school-service/pkg/usecase/interfaces"
)

type StreamService struct {
	pb.UnimplementedSchoolServiceServer
	usecase interfaces.SchoolUseCase
}

func NewSchoolService(usecase interfaces.SchoolUseCase) pb.SchoolServiceServer {
	return &StreamService{
		usecase: usecase,
	}
}
