package main

import (
	"log"
	"time"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/jacenr/gRPCinDockerDemo/grpcDemo"
)

func main() {
	// connect to server
	addr := "localhost:8001"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err!=nil {
		log.Fatalf("can not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSumClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetSum(ctx, &pb.Req{M:3,N:2})
	if err!=nil {
		log.Fatalf("could not get resp: %v", err)
	}

	log.Printf("the sum: %v", r.S)
}

