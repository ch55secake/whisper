package server

import (
	"context"
	messenger "github.com/ch55secake/whisper/pkg/server/generated"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	messenger.UnimplementedMessengerServer
}

func (s *server) SendMessage(ctx context.Context, msg *messenger.ChatMessage) (*messenger.Ack, error) {
	log.Printf("Received message from %s: %s", msg.GetSender().GetUsername(), msg.GetContent())
	return &messenger.Ack{
		MessageId: msg.GetMessageId(),
		Receiver:  msg.GetReceiver(),
		Timestamp: msg.GetTimestamp(),
	}, nil
}

// StartServer spin up the grpc server and listen for messages on 41002
func StartServer() {
	listener, err := net.Listen("tcp", ":41002")
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	messenger.RegisterMessengerServer(grpcServer, &server{})

	log.Printf("Started gRPC server on port 41002")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
