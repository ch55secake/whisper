package client

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

type GRPCMessage struct {
	Envelope *messenger.Envelope
	Err      error
}

// connect create a connection with the running grpc server
func connect(serverAddr string) (*grpc.ClientConn, error) {
	connection, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	return connection, nil
}

// SendMessage send an envelope into the stream present on the model, the stream is created at init
func (m *model) SendMessage(msg Message) error {
	if m.stream == nil {
		return fmt.Errorf("Couldn't find stream available to send message")
	}

	env := &messenger.Envelope{
		Payload: &messenger.Envelope_ChatMessage{
			ChatMessage: &messenger.ChatMessage{
				MessageId: uuid.NewString(),
				Sender: &messenger.Peer{
					Id:       uuid.NewString(),
					Username: m.username,
				},
				Receiver: &messenger.Peer{
					Id:       uuid.NewString(),
					Username: msg.from,
				},
				Content:   msg.content,
				Timestamp: time.Now().Unix(),
			},
		},
	}
	return m.stream.Send(env)
}
