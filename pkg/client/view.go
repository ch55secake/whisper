package client

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// renderMessages builds the styled string that is loaded into the viewport.
// It is called from Update so the real model viewport always has fresh content.
func renderMessages(messages []Message) string {
	var rows []string
	for _, msg := range messages {
		senderStyle := MsgSenderOtherStyle
		contentStyle := MsgContentOtherStyle
		if msg.mine {
			senderStyle = MsgSenderMineStyle
			contentStyle = MsgContentMineStyle
		}
		meta := lipgloss.JoinHorizontal(
			lipgloss.Bottom,
			senderStyle.Render(msg.from),
			MsgTimestampStyle.Render("  "+msg.at),
		)
		content := contentStyle.Render(msg.content)
		rows = append(rows, lipgloss.JoinVertical(lipgloss.Left, meta, content, ""))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

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

	dot := ConnectedDotStyle.Render("●")
	if !m.connected {
		dot = DisconnectedDotStyle.Render("●")
	}
	left := lipgloss.JoinHorizontal(lipgloss.Center, dot, " ", HeaderTitleStyle.Render("whisper"))
	centre := HeaderContextStyle.Render(fmt.Sprintf("%s @ %s", m.username, m.serverAddr))
	clock := HeaderClockStyle.Render(time.Now().Format("15:04:05"))

	leftW := lipgloss.Width(left)
	rightW := lipgloss.Width(clock)
	centreW := m.width - leftW - rightW - 4
	if centreW < 0 {
		centreW = 0
	}

	headerContent := lipgloss.JoinHorizontal(
		lipgloss.Center,
		left,
		lipgloss.PlaceHorizontal(centreW, lipgloss.Center, centre),
		clock,
	)
	header := HeaderStyle.Width(m.width).Render(headerContent)

	var badgeLabel string
	if m.unread > 0 {
		badgeLabel = fmt.Sprintf("↓  %d new message", m.unread)
		if m.unread != 1 {
			badgeLabel += "s"
		}
	}
	unreadBadge := UnreadBadgeStyle.Width(m.width).Render(badgeLabel)

	charCount := fmt.Sprintf("%d / %d", len(m.input.Value()), m.input.CharLimit)
	charCountStyle := CharCountNormalStyle
	if m.input.CharLimit-len(m.input.Value()) <= 20 {
		charCountStyle = CharCountWarningStyle
	}
	counter := charCountStyle.Render(charCount)

	inputRow := lipgloss.JoinHorizontal(
		lipgloss.Center,
		lipgloss.NewStyle().Width(m.width-lipgloss.Width(counter)-1).Render(m.input.View()),
		counter,
	)
	inputView := InputBoxStyle.Width(m.width).Render(inputRow)

	viewPortView := MessageBoxStyle.Width(m.width).Render(m.viewport.View())
	
	return lipgloss.JoinVertical(lipgloss.Top, header, viewPortView, unreadBadge, inputView)
}
