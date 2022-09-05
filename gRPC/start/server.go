package main

import (
	"log"
	"net"

	"example/chat"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("9080 portu dinlenirken hata oluştu: %v \n", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpcServer oluşturulurken hata : %v", err)
	}
}
