package client

import (
	"time"

	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.phase == login {
		return "Enter your username:\n\n" + m.input.View()
	}

	currentTime := time.Now().Format("15:04:05")
	title := " whisper "
	clock := currentTime

	headerContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.PlaceHorizontal(len(title), lipgloss.Left, title),
		lipgloss.PlaceHorizontal(m.width-len(title), lipgloss.Right, clock),
	)
	header := HeaderStyle.Width(m.width).Render(headerContent)

	var rows []string
	for _, msg := range m.messages {
		msgText := fmt.Sprintf("%s - %s\n%s", msg.from, msg.at, msg.content)

		var row string
		if msg.mine {
			row = lipgloss.Place(
				m.width,
				lipgloss.Height(msgText),
				lipgloss.Right,
				lipgloss.Top,
				PinkStyle.Render(msgText),
			)
		} else {
			row = lipgloss.Place(
				m.width,
				lipgloss.Height(msgText),
				lipgloss.Left,
				lipgloss.Top,
				CyanStyle.Render(msgText),
			)
		}
		rows = append(rows, row)
	}

	messagesContent := lipgloss.JoinVertical(lipgloss.Left, rows...)
	m.viewport.SetContent(messagesContent)

	m.viewport.Width = m.width
	m.viewport.Height = m.height - lipgloss.Height(header) - lipgloss.Height(m.input.View())
	m.viewport.GotoBottom()

	inputView := InputBoxStyle.Width(m.width).Render(m.input.View())

	viewPortView := MessageBoxStyle.Width(m.width).Render(m.viewport.View())

	ui := lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		viewPortView,
		inputView,
	)

	return ui
}
