package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mikejoh/grpc-echo-service/echo"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	defaultName = "echoes"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoesClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{Message: "ping"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Echoes service replied: %s", r.GetMessage())
}
