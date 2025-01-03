package client

import "github.com/charmbracelet/lipgloss"

// HeaderStyle styles the header which currently just contains the title
var HeaderStyle = lipgloss.NewStyle().
	Padding(0, 0).
	MarginTop(0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("141")).
	Bold(true)

// MessageBoxStyle style for the box which contains the messages
var MessageBoxStyle = lipgloss.NewStyle().
	Padding(0, 0).
	MarginTop(0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

// InputBoxStyle style for the box which handles input
var InputBoxStyle = lipgloss.NewStyle().
	Padding(0, 2).
	MarginTop(0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("141"))

// ItemStyle style for each message sent
var ItemStyle = lipgloss.NewStyle().
	MarginTop(1).
	PaddingLeft(1).
	Foreground(lipgloss.Color("#689d6a"))

// SelectedItemStyle for each selected message
var SelectedItemStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	Foreground(lipgloss.Color("#fabd2f"))
