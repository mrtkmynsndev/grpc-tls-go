package main

import (
	"log"
	"net"

	pb "github.com/mrtkmynsndev/grpc-tls-go/helloworld"
	"google.golang.org/grpc"
)

type greeterService struct {
	pb.UnimplementedGreeterServiceServer
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
