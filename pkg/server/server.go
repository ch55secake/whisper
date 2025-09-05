package server

import (
	"context"
	"log"
	"net"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
	"google.golang.org/grpc"
	"sync"
)

type server struct {
	messenger.UnimplementedMessengerServer
	mu      sync.Mutex
	clients map[messenger.Messenger_ChatServer]struct{}
}

func (s *server) SendMessage(ctx context.Context, msg *messenger.ChatMessage) (*messenger.Ack, error) {
	log.Printf("Received message from %s: %s", msg.GetSender().GetUsername(), msg.GetContent())
	return &messenger.Ack{
		MessageId: msg.GetMessageId(),
		Receiver:  msg.GetReceiver(),
		Timestamp: msg.GetTimestamp(),
	}, nil
}

// NewServer create a new server object to keep track of connected clients
func NewServer() *server {
	return &server{
		clients: make(map[messenger.Messenger_ChatServer]struct{}),
	}
}

// Chat register multiple chat clients and sync
func (s *server) Chat(stream messenger.Messenger_ChatServer) error {
	s.mu.Lock()
	s.clients[stream] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, stream)
		s.mu.Unlock()
	}()

	for {
		env, err := stream.Recv()
		if err != nil {
			return nil
		}

		s.mu.Lock()
		for c := range s.clients {
			if err := c.Send(env); err != nil {
				log.Printf("failed to send to client: %v", err)
			}
		}
		s.mu.Unlock()
	}
}

// StartServer spin up the grpc server and listen for messages on 41002
func StartServer() {
	listener, err := net.Listen("tcp", ":41002")
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	messenger.RegisterMessengerServer(grpcServer, NewServer())

	log.Printf("Started gRPC server on port 41002")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
