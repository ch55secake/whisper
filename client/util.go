package client

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

func StartClient() {
	input := textinput.New()
	input.Prompt = " > "
	input.Placeholder = "Say something..."
	input.Focus()
	input.CharLimit = 256
	input.PlaceholderStyle = lipgloss.NewStyle().Italic(true)

	m := model{
		input: input,
		//messages: []string{},
		username: "You",
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
