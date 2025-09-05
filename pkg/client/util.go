package client

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

// StartClient starts the client and initializes the model, will most likely have to move this
// TODO: Move this code at some point
func StartClient() {
	input := textinput.New()
	input.Prompt = "> "
	input.Focus()
	input.CharLimit = 256

	conn, err := grpc.NewClient("localhost:41002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}

	defer conn.Close()

	client := messenger.NewMessengerClient(conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Failed to open chat stream: %v", err)
	}

	m := model{
		input:    input,
		messages: []Message{},
		client:   client,
		stream:   stream,
		phase:    login,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
