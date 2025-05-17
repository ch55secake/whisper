package client

import (
	"github.com/charmbracelet/lipgloss"
	"time"
)

func (m model) View() string {
	currentTime := time.Now().Format("15:04:05 PM")
	title := " whisper "
	clock := currentTime

	//available width for the header content
	headerContent := lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		title,
	)[:len(title)] + lipgloss.PlaceHorizontal(
		m.width-len(title),
		lipgloss.Right,
		clock,
	)

	header := HeaderStyle.Width(m.width).Render(headerContent)

	messagesView := MessageBoxStyle.Height(m.height - 6).Width(m.width).Render(m.messages.View())
	inputView := InputBoxStyle.Height(m.height / 6).Width(m.width).Render(m.input.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, messagesView, inputView)
}
