package main

import (
	"context"
	"log"

	"example/hello"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %+v", err)
	}

	c := hello.NewHelloServiceClient(conn)
	m, err := c.Talkie(context.Background(), &hello.Message{
		Content: "Hello I am a Client",
	})
	if err != nil {
		log.Fatalf("Failed to dial: %+v", err)
	}

	log.Printf(m.Context)
}
