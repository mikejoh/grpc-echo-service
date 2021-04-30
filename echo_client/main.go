package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/mikejoh/grpc-echo-service/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "0.0.0.0:8443"
	defaultName = "echoes"
)

func main() {
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		log.Fatal(fmt.Errorf("failed to add server CA certificate"))
	}

	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoesClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{Message: "ping"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echoes service replied: %s", r.GetMessage())
}
