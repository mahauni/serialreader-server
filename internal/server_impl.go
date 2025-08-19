package internal

import (
	"context"
	"log"

	pb "github.com/mahauni/serialreader-server/proto"
)

type SerialReaderServerImpl struct {
	pb.SerialReaderServiceServer
}

func (s *SerialReaderServerImpl) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("Received: %v\n", in.GetName())
	return &pb.SayHelloResponse{Message: "Hello " + in.GetName()}, nil
}
