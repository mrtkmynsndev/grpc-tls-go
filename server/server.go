package main

import (
	"context"
	"log"
	"net"

	pb "github.com/mrtkmynsndev/grpc-tls-go/helloworld"
	"google.golang.org/grpc"
)

type greeterService struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *greeterService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("received name: %v", request.GetName())
	return &pb.HelloReply{Message: "Hello " + request.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("list port err: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})

	pb.RegisterGreeterServiceServer(grpcServer, &greeterService{})

	log.Printf("listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc serve err: %v", err)
	}
}
