package api

import (
	"fmt"
	"log"
	"net"
	"school-service/pkg/config"
	"school-service/pkg/pb"

	"google.golang.org/grpc"
)

type Server struct {
	lis  net.Listener
	gsr  *grpc.Server
	port string
}

func NewServerGRPC(cfg config.Config, srv pb.SchoolServiceServer) (*Server, error) {

	addr := fmt.Sprintf("%s:%s", cfg.SchoolServiceHost, cfg.SchoolServicePort)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	gsr := grpc.NewServer()

	pb.RegisterSchoolServiceServer(gsr, srv)

	return &Server{
		lis:  lis,
		gsr:  gsr,
		port: cfg.SchoolServicePort,
	}, err
}

func (c *Server) Start() error {
	log.Println("School service listening on port: ", c.port)
	return c.gsr.Serve(c.lis)
}
