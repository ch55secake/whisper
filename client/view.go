package client

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
	"time"
)

var now = time.Now()

// View function to render the TUI
func (m model) View() string {
	headerView := HeaderStyle.Width(m.width).Height(m.height / 20).Render(" whisper v0 - " + m.currentTime)

	messagesView := MessageBoxStyle.Height(m.height - 6).Width(m.width).Render(strings.Join(m.messages, "\n"))

	inputView := InputBoxStyle.Height(m.height / 16).Width(m.width).Render(m.input.View())

	return lipgloss.JoinVertical(lipgloss.Top, headerView, messagesView, inputView)
}
