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

var PinkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4"))
var CyanStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))

// MenuStyle is the outer container for the main menu screen
var MenuStyle = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#FF69B4"))

// MenuTitleStyle styles the large title shown above the menu list
var MenuTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF69B4")).
	MarginBottom(1)

// MenuSelectedStyle highlights the currently selected menu item
var MenuSelectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FFFF")).
	Bold(true)

// MenuNormalStyle styles unselected menu items
var MenuNormalStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("white"))

// MenuHelpStyle styles the help hint pinned to the bottom of the menu box
var MenuHelpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240")).
	Italic(true)

// LoginBoxStyle is the outer container for the login screen
var LoginBoxStyle = lipgloss.NewStyle().
	Padding(1, 4).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#00FFFF"))

// LoginPromptStyle styles the "Enter your username" label
var LoginPromptStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("white")).
	MarginBottom(1)

// LoginHintStyle styles the small hint line at the bottom of the login box
var LoginHintStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240")).
	Italic(true).
	MarginTop(1)
