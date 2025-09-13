package client

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

// StartClient starts the client and initializes the model, will most likely have to move this
func StartClient() {
	input := textinput.New()
	input.Prompt = ""
	input.Focus()
	input.CharLimit = 256

	conn, err := connect("localhost:41002")

	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	defer conn.Close()

	client := messenger.NewMessengerClient(conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatal("Unable to connect to server, are you sure that its running? Or that the address you provided is correct?")
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

func StartChatListener(stream messenger.Messenger_ChatClient) tea.Cmd {
	return func() tea.Msg {
		env, err := stream.Recv()
		if err != nil {
			return GRPCMessage{Err: err}
		}
		return GRPCMessage{Envelope: env}
	}
}
