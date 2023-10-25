package client

import (
	"api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/config"
	"context"
	"fmt"
	"io"
	"school-service/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type schoolClient struct {
	client pb.SchoolServiceClient
}

func NewSchoolClient(cfg config.Config) (interfaces.SchoolClient, error) {

	addr := fmt.Sprintf("%s:%s", cfg.SchoolServiceHost, cfg.SchoolServicePort)

	cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial to school service: %w", err)
	}

	client := pb.NewSchoolServiceClient(cc)

	return &schoolClient{
		client: client,
	}, nil
}

func (s *schoolClient) GetOneInJSON(name string) ([]byte, error) {

	req := &pb.CreateRequest{
		Name: name,
	}

	stream, err := s.client.Create(context.Background(), req)

	if err != nil {
		return nil, fmt.Errorf("failed to call create school: %w", err)
	}

	var shoolData []byte

	for {
		res, err := stream.Recv()
		if err != nil {

			if err == io.EOF {
				return shoolData, nil
			}

			return nil, fmt.Errorf("failed to receive school on stream: %w", err)
		}
		shoolData = append(shoolData, res.Data...)
	}
}
