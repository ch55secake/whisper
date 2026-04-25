package client

import (
	"time"

	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.phase == menu {
		const margin = 2
		boxWidth := m.width - margin*2
		boxHeight := m.height - margin*2

		const frameV, frameH = 4, 6
		innerWidth := boxWidth - frameH
		innerHeight := boxHeight - frameV

		title := MenuTitleStyle.Render("whisper")
		help := MenuHelpStyle.Render("↑/↓ navigate  •  enter select  •  esc back")

		listHeight := innerHeight - lipgloss.Height(title) - lipgloss.Height(help)
		m.menuList.SetWidth(innerWidth)
		m.menuList.SetHeight(listHeight)
		m.menuList.SetShowHelp(false)

		inner := lipgloss.JoinVertical(lipgloss.Left, title, m.menuList.View(), help)
		box := MenuStyle.Width(boxWidth).Height(boxHeight).Render(inner)
		return lipgloss.NewStyle().Margin(margin).Render(box)
	}

	if m.phase == login {
		title := MenuTitleStyle.Render("whisper")
		prompt := LoginPromptStyle.Render("Enter your username")
		hint := LoginHintStyle.Render("enter to confirm • esc to go back")
		inner := lipgloss.JoinVertical(lipgloss.Left, prompt, m.input.View(), hint)
		box := LoginBoxStyle.Render(inner)
		content := lipgloss.JoinVertical(lipgloss.Center, title, box)
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
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
