package client

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
}

// Message is a struct which represents, who sent the message, if it has been seen and what it contains along with the actual content of the message
type Message struct {
	from    string
	at      string
	seen    bool
	content string
}

func (item Message) Title() string {
	return item.from
}

func (item Message) Description() string {
	return item.content
}

func (item Message) FilterValue() string {
	return item.content
}

type messageItemDelegate struct{}

func (d messageItemDelegate) Height() int                             { return 1 }
func (d messageItemDelegate) Spacing() int                            { return 0 }
func (d messageItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d messageItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Message)
	if !ok {
		return
	}

	head := lipgloss.NewStyle().Bold(true).Render(i.from)

	var str string

	str = fmt.Sprintf("%s: %s", head, i.content)

	fn := ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return SelectedItemStyle.Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

// Init create the model and return the relevant tea cmd, also sets the window title and ticks for the time
func (m model) Init() tea.Cmd {
	m.currentTime = time.Now().Format("15:04:05")
	tea.SetWindowTitle("whisper")
	return tea.Batch(textinput.Blink, tickCmd())
}
