package client

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// StartClient starts the client and initializes the model, will most likely have to move this
// TODO: Move this code at some point
func StartClient() {
	input := textinput.New()
	input.Prompt = "> "
	input.Focus()
	input.CharLimit = 256

	messagelist := list.New([]list.Item{}, messageItemDelegate{}, 0, 0)
	messagelist.Styles = list.Styles{
		NoItems: lipgloss.NewStyle().PaddingLeft(2).PaddingTop(2).Faint(true),
	}

	messagelist.SetShowHelp(false)
	messagelist.SetShowTitle(false)
	messagelist.SetShowStatusBar(false)
	messagelist.SetShowFilter(false)
	// this is the reason that the count of messages shows up, don't want this enabled
	messagelist.SetShowPagination(false)

	m := model{
		input:    input,
		messages: messagelist,
		username: "user",
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
