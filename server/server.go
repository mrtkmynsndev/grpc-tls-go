package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/mrtkmynsndev/grpc-tls-go/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type greeterService struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *greeterService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("received name: %v", request.GetName())
	return &pb.HelloReply{Message: "Hello " + request.GetName()}, nil
}

func main() {
	// listen port
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("list port err: %v", err)
	}

	// read ca's cert
	caPem, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		log.Fatal(err)
	}

	// read server cert & key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	tlsCredentials := credentials.NewTLS(conf)

	// create grpc server
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	// register service into grpc server
	pb.RegisterGreeterServiceServer(grpcServer, &greeterService{})

	log.Printf("listening at %v", lis.Addr())

	// listen port
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc serve err: %v", err)
	}
}
