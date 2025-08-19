package internal

import (
	"fmt"
	"log"
	"net"

	pb "github.com/mahauni/serialreader-server/proto"
	"google.golang.org/grpc"
)

type SerialReaderServer struct {
	port       int
	grpcServer *grpc.Server
}

func New(port int) *SerialReaderServer {
	return &SerialReaderServer{
		port:       port,
		grpcServer: nil,
	}
}

func (s *SerialReaderServer) RunMainRuntimeLoop() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}

	grpcServer := grpc.NewServer()

	s.grpcServer = grpcServer

	log.Printf("gRPC server is running\n")

	pb.RegisterSerialReaderServiceServer(grpcServer, &SerialReaderServerImpl{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

func (s *SerialReaderServer) StopMainRuntimeLoop() {
	log.Printf("Starting graceful shutdown now...")

	s.grpcServer.GracefulStop()
}
