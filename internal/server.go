package internal

import (
	"fmt"
	"log"
	"net"

	pb "github.com/mahauni/serialreader-server/proto"
	"google.golang.org/grpc"
)

type SerialReaderServer struct {
	port              int
	arduinoDevicePath string
	arduinoReader     *ArduinoReader
	grpcServer        *grpc.Server
}

func New(arduinoDevicePath string, port int) *SerialReaderServer {
	return &SerialReaderServer{
		port:              port,
		arduinoDevicePath: arduinoDevicePath,
		arduinoReader:     nil,
		grpcServer:        nil,
	}
}

func (s *SerialReaderServer) RunMainRuntimeLoop() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	arduinoReader := NewArduinoReader(s.arduinoDevicePath)

	grpcServer := grpc.NewServer()

	s.grpcServer = grpcServer
	s.arduinoReader = arduinoReader

	log.Printf("gRPC server is running.")

	pb.RegisterSerialReaderServiceServer(grpcServer, &SerialReaderServerImpl{
		arduinoReader: arduinoReader,
	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *SerialReaderServer) StopMainRuntimeLoop() {
	log.Printf("Starting graceful shutdown now...")

	s.arduinoReader = nil

	s.grpcServer.GracefulStop()
}
