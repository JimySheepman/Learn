package main

import (
	"log"

	"example/chat"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("9080 portu dinlenirken hata oluştu: %v", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Client'ten merhabalar!",
	}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("SayHello fonksiyonu çağırılıken hata oluştu: %v", err)
	}

	log.Printf("Server'dan gelen cevap: %s", response.Body)
}
