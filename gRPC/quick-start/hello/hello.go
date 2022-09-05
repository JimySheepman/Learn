package hello

import (
	context "context"
	"log"
)

type Server struct{}

func (s *Server) Talkie(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Message from client: %s", in.Content)
	return &Message{
		Content: "Hello I am Server",
	}, nil
}
