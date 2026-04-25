// Package client is the logical responsible for creating the TUI
package client

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

type phase int

const (
	login phase = iota
	chat
)

// Model this Model is the current Model of the ui, all it contains is the input and the list of messages, alongside the base
// height and width
type Model struct {
	height      int
	width       int
	currentTime string
	phase       phase
	input       textinput.Model
	viewport    viewport.Model
	messages    []Message
	username    string
	client      messenger.MessengerClient
	stream      messenger.Messenger_ChatClient
}

// Message is a struct that represents who sent the message if it has been seen, and what it contains along with the actual content of the message
type Message struct {
	from    string
	at      string
	seen    bool
	content string
	mine    bool
}

type GRPCMessage struct {
	Envelope *messenger.Envelope
	Err      error
}

// TODO Need to make use of this in the update.go and then configure the Model to have the stream on it, should
// also have the sendMessage method on the Model
func startChatListener(stream messenger.Messenger_ChatClient) tea.Cmd {
	return func() tea.Msg {
		env, err := stream.Recv()
		if err != nil {
			return GRPCMessage{Err: err}
		}
		return GRPCMessage{Envelope: env}
	}
}

func (m *Model) SendMessage(msg Message) error {
	if m.stream == nil {
		return fmt.Errorf("couldn't find stream available to send message")
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

// Init create the model and return the relevant tea cmd, also sets the window title and ticks for the time
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("whisper"),
		startChatListener(m.stream),
	)
}
