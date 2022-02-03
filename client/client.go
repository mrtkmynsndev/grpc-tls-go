package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mrtkmynsndev/grpc-tls-go/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Mert Kimyonsen"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())
}
