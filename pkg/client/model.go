// Package client is the logical responsible for creating the TUI
package client

import (
	"fmt"
	"io"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

// model is the current model of the ui, all it contains is the input and the list of messages, alongside the base
// height and width
type model struct {
	height      int
	width       int
	currentTime string
	input       textinput.Model
	messages    list.Model
	username    string
	client      messenger.MessengerClient
	stream      messenger.Messenger_ChatClient
}

// Message is a struct which represents, who sent the message, if it has been seen and what it contains along with the actual content of the message
type Message struct {
	from string
	at   string
	// seen    bool
	content string
}

// Title represents the person that sent the message/is the title of the list item
func (item Message) Title() string {
	return item.from
}

// Description represents the actual message content, or is a description of the list item
func (item Message) Description() string {
	return item.content
}

// FilterValue if the user wants to search for specific content of a message
func (item Message) FilterValue() string {
	return item.content
}

type messageItemDelegate struct{}

func (d messageItemDelegate) Height() int                             { return 1 }
func (d messageItemDelegate) Spacing() int                            { return 0 }
func (d messageItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d messageItemDelegate) Render(w io.Writer, m list.Model, _ int, listItem list.Item) {
	i, ok := listItem.(Message)
	if !ok {
		return
	}

	username := lipgloss.NewStyle().Bold(true).Render(i.from)

	timeSent := lipgloss.NewStyle().Bold(true).Render(i.at)

	var str string
	for j := range m.VisibleItems() {
		item := m.VisibleItems()[j].(Message)
		str = fmt.Sprintf("%s %s\n%s", username, timeSent, item.content)
		_, err := fmt.Fprint(w, SelectedItemStyle.Render(str)+"\n")
		if err != nil {
			return
		}
	}
}

type GRPCMessage struct {
	Envelope *messenger.Envelope
	Err      error
}

// TODO Need to make use of this in the update.go and then configure the model to have the stream on it, should
// also have the sendMessage method on the model
func startChatListener(stream messenger.Messenger_ChatClient) tea.Cmd {
	return func() tea.Msg {
		// fmt.Print("waiting to receive.........")
		env, err := stream.Recv()
		if err != nil {
			return GRPCMessage{Err: err}
		}
		return GRPCMessage{Envelope: env}
	}
}

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

// Init create the model and return the relevant tea cmd, also sets the window title and ticks for the time
func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("whisper"),
		startChatListener(m.stream),
	)
}
