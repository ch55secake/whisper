// Package client is the logical responsible for creating the TUI
package client

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// model is the current model of the ui, all it contains is the input and the list of messages, alongside the base
// height and width
type model struct {
	height int
	width  int
	// currentTime string
	input    textinput.Model
	messages list.Model
	username string
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

// Init create the model and return the relevant tea cmd, also sets the window title and ticks for the time
func (m model) Init() tea.Cmd {
	tea.SetWindowTitle("whisper")
	return nil
}
