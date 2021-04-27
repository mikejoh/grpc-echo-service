package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/mikejoh/simplest-grpc-example/echo"
	"google.golang.org/grpc"
)

const (
	address = "0.0.0.0:8080"
)

// server is used to implement echo.GreeterServer.
type server struct {
	pb.UnimplementedEchoesServer
}

// Echo implements echo.EchoesServer
func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.EchoReply{Message: "pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Starting the Echoes service (%s)..", address)

	s := grpc.NewServer()

	pb.RegisterEchoesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
