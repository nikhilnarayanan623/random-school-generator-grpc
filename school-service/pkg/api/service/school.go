package service

import (
	"encoding/json"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/pb"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/usecase/interfaces"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	maxBufferSize = 1024
)

type schoolService struct {
	pb.UnimplementedSchoolServiceServer
	usecase interfaces.SchoolUseCase
}

func NewSchoolService(usecase interfaces.SchoolUseCase) pb.SchoolServiceServer {
	return &schoolService{
		usecase: usecase,
	}
}

func (s *schoolService) Create(req *pb.CreateRequest, stream pb.SchoolService_CreateServer) error {

	school := s.usecase.Create(req.GetName())

	data, err := json.Marshal(school)

	if err != nil {
		return status.Errorf(codes.Internal, "failed to marshal school to json: %v", err)
	}

	var (
		start  = 0
		end    = maxBufferSize
		buffer []byte
	)

	for {

		// check the end is out of bound of data
		if end >= len(data) {
			buffer = data[start:]
			if err := stream.Send(&pb.CreateResponse{Data: buffer}); err != nil {
				return err
			}

			break
		} else {
			// slice out the buffer from data with max buffer size
			buffer = data[start:end]

			// update the start and end
			start += maxBufferSize
			end += maxBufferSize
		}

		// send the buffer as stream
		if err := stream.Send(&pb.CreateResponse{Data: buffer}); err != nil {
			return err
		}
	}
	return nil
}
