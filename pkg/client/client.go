package client

import (
	"context"
	"github.com/google/uuid"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

// connect create a connection with the running grpc server
func connect(serverAddr string) (*grpc.ClientConn, error) {
	connection, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	return connection, nil
}

// Dont use this
// SendMessage provides utility to be able to send a message from the chat client to the grpc server
func SendMessage(msg Message, serverAddr string) error {
	connection, err := connect(serverAddr)
	if err != nil {
		return nil
	}
	defer connection.Close()

	client := messenger.NewMessengerClient(connection)

	_, err = client.SendMessage(context.Background(), &messenger.ChatMessage{
		MessageId: uuid.NewString(),
		Sender: &messenger.Peer{
			Id:       uuid.NewString(),
			Username: msg.from,
		},
		Receiver: &messenger.Peer{
			Id:       uuid.NewString(),
			Username: msg.from,
		},
		Content:   msg.content,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		log.Fatalf("failed to send message to server: %v", err)
	}

	return nil
}
