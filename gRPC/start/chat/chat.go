package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Clientten mesaj alındı: %s", message.Body)

	return &Message{Body: "Server'dan merhaba!"}, nil //hata nil olsun
}
