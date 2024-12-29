package client

import (
	"github.com/charmbracelet/lipgloss"
	"time"
)

var now = time.Now()

// View function to render the TUI
func (m model) View() string {
	headerView := HeaderStyle.Width(m.width).Height(m.height / 20).Render(" whisper ")

	messagesView := MessageBoxStyle.Height(m.height - 6).Width(m.width).Render(m.messages.View())

	inputView := InputBoxStyle.Height(m.height / 16).Width(m.width).Render(m.input.View())

	return lipgloss.JoinVertical(lipgloss.Top, headerView, messagesView, inputView)
}
