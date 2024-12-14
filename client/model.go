package client

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

// model is the current model of the ui, all it contains is the input and the list of messages, alongside the base
// height and width
type model struct {
	height      int
	width       int
	currentTime string
	input       textinput.Model
	messages    []string
}

// Init create the model and return the relevant tea cmd, also sets the window title and ticks for the time
func (m model) Init() tea.Cmd {
	m.currentTime = time.Now().Format("15:04:05")
	tea.SetWindowTitle("whisper")
	return tea.Batch(textinput.Blink, tickCmd())
}
