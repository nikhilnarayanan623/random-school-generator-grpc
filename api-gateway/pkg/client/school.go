package client

import (
	"api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/config"
	"fmt"
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
