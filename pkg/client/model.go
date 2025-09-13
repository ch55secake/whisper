// Package client is the logical responsible for creating the TUI
package client

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

type phase int

const (
	login phase = iota
	chat
)

// TODO add an error log that can be updated the same way as the messages are and also a status log for when that is added
// model is the current model of the ui, all it contains is the input and the list of messages, alongside the base
// height and width
type model struct {
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

// Message is a struct which represents, who sent the message, if it has been seen and what it contains along with the actual content of the message
type Message struct {
	from    string
	at      string
	seen    bool
	content string
	mine    bool
}

// Init create the model and return the relevant tea cmd, also sets the window title and ticks for the time
func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("whisper"),
		StartChatListener(m.stream),
	)
}
