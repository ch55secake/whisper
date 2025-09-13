package client

import "github.com/charmbracelet/lipgloss"

// HeaderStyle styles the header which currently just contains the title
var HeaderStyle = lipgloss.NewStyle().
	Padding(0, 0).
	MarginTop(0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("white")).
	Bold(true)

// MessageBoxStyle style for the box which contains the messages
var MessageBoxStyle = lipgloss.NewStyle().
	Padding(0, 0).
	MarginTop(0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("white"))

// InputBoxStyle style for the box which handles input
var InputBoxStyle = lipgloss.NewStyle().
	Padding(0, 2).
	MarginTop(0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("white"))

var loginBoxStyle = lipgloss.NewStyle().
	Padding(0, 2).
	MarginTop(0).
	Bold(true)

var PinkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4"))
var CyanStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
