package client

import (
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// Update function to handle messages and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		case "esc":
			m.phase = menu
			return m, nil
		case "enter":
			input := strings.TrimSpace(m.input.Value())

			switch m.phase {
			case menu:
				selected, ok := m.menuList.SelectedItem().(menuItem)
				if !ok {
					break
				}
				switch selected.title {
				case "Connect":
					m.phase = login
					m.input.Focus()
				case "Quit":
					return m, tea.Quit
				}
				return m, nil
			case login:
				if input == "" {
					break
				}
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
					m.viewport.SetContent(renderMessages(m.messages))
					m.viewport.GotoBottom()
					m.input.SetValue("")
					if input == "disconnect" {
						return m, tea.Quit
					}
				}
			}

		case "down", "pgdown", "end":
			if m.phase == chat && m.viewport.AtBottom() {
				m.unread = 0
			}
		}

	case tea.WindowSizeMsg:
		docStyle := lipgloss.NewStyle().Margin(1, 2)
		h, v := docStyle.GetFrameSize()
		m.height = msg.Height - h
		m.width = msg.Width - v
		m.menuList.SetSize(m.width/2, m.height/2)
		m.viewport.Width = m.width
		m.viewport.Height = m.height - 9

	case GRPCMessage:
		if msg.Err != nil {
			log.Printf("grpc error: %v", msg.Err)
			m.connected = false
			return m, nil
		}

		m.connected = true

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
				atBottom := m.viewport.AtBottom()
				m.messages = append(m.messages, message)
				m.viewport.SetContent(renderMessages(m.messages))
				if atBottom {
					m.viewport.GotoBottom()
				} else {
					m.unread++
				}
			}
		}

		return m, startChatListener(m.stream)
	}

	if m.phase == menu {
		var cmd tea.Cmd
		m.menuList, cmd = m.menuList.Update(msg)
		return m, cmd
	}

	m.viewport, _ = m.viewport.Update(msg)
	if m.viewport.AtBottom() {
		m.unread = 0
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
