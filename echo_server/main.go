package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	pb "github.com/mikejoh/grpc-echo-service/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// server is used to implement echo.EchoesServer.
type server struct {
	pb.UnimplementedEchoesServer
}

func main() {
	address := "0.0.0.0:8443"

	fmt.Printf("Running the Echoes service (%s)..\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverCert, err := tls.LoadX509KeyPair("cert/echo_server-cert.pem", "cert/echo_server-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	s := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(config)),
	)

	pb.RegisterEchoesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// Echo implements echo.EchoesServer
func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.EchoReply{Message: "pong"}, nil
}
