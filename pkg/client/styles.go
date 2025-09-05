package client

import "github.com/charmbracelet/lipgloss"

// HeaderStyle styles the header which currently just contains the title
var HeaderStyle = lipgloss.NewStyle().
	Padding(0, 0).
	MarginTop(0).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#FF69B4")).
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
	BorderForeground(lipgloss.Color("#00FFFF"))

// ItemStyle style for each message sent
var ItemStyle = lipgloss.NewStyle().
	MarginTop(1).
	PaddingLeft(1).
	Foreground(lipgloss.Color("#800080"))

// SelectedItemStyle for each selected message, this controls the style of the message and user name
// TODO: split the colour of username and messages
var SelectedItemStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	Foreground(lipgloss.Color("#800080"))

var PinkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4"))
var CyanStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
