package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/mrtkmynsndev/grpc-tls-go/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// read ca's cert
	caCert, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		log.Fatal(caCert)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	// create client connection
	conn, err := grpc.Dial(
		"0.0.0.0:9000",
		grpc.WithTransportCredentials(tlsCredential),
	)
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
