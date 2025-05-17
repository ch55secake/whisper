package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
	"time"
)

// Update function to handle messages and commands
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		case "enter":
			input := strings.TrimSpace(m.input.Value())
			if input != "" {
				message := Message{
					from:    m.username,
					at:      time.Now().Format("15:04"),
					content: input,
				}
				m.messages.InsertItem(len(m.messages.Items()), message)
				m.input.SetValue("")
				if input == "disconnect" {
					return m, tea.Quit
				}
			}
		}

	case tea.WindowSizeMsg:
		docStyle := lipgloss.NewStyle().Margin(1, 2)
		h, v := docStyle.GetFrameSize()
		m.height = msg.Height - h
		m.width = msg.Width - v

	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
