package client

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"log"

	tea "github.com/charmbracelet/bubbletea"
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
			if input == "" {
				break
			}

			switch m.phase {
			case login:
				m.username = input
				m.phase = chat
				m.input.SetValue("")
			case chat:
				if input != "" {
					message := Message{
						from:    m.username,
						at:      time.Now().Format("15:04"),
						content: input,
						mine:    true,
					}

					m.SendMessage(message)

					m.messages = append(m.messages, message)
					m.input.SetValue("")
					if input == "disconnect" {
						return m, tea.Quit
					}
				}
			}
		}

	case tea.WindowSizeMsg:
		docStyle := lipgloss.NewStyle().Margin(1, 2)
		h, v := docStyle.GetFrameSize()
		m.height = msg.Height - h
		m.width = msg.Width - v

	case GRPCMessage:
		if msg.Err != nil {
			log.Printf("grpc error: %v", msg.Err)
			return m, nil
		}

		cm := msg.Envelope.GetChatMessage()
		if cm != nil {
			timestamp := cm.GetTimestamp()
			t := time.Unix(timestamp, 0)
			timeFormatted := t.Format("15:04")
			message := Message{
				from:    cm.Sender.Username,
				content: cm.Content,
				at:      timeFormatted,
			}
			if message.from != m.username {
				m.messages = append(m.messages, message)
			}
		}

		return m, startChatListener(m.stream)
	}

	m.viewport, _ = m.viewport.Update(msg)

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
