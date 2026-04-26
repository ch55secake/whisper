package client

import "github.com/charmbracelet/lipgloss"

// HeaderStyle styles the header which currently just contains the title
var HeaderStyle = lipgloss.NewStyle().
	Padding(0, 1).
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

// --- Chat message styles ---

// MsgSenderMineStyle styles the sender name for messages sent by the local user
var MsgSenderMineStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF69B4"))

// MsgSenderOtherStyle styles the sender name for messages from other users
var MsgSenderOtherStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#00FFFF"))

// MsgTimestampStyle styles the timestamp shown next to the sender name
var MsgTimestampStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240"))

// MsgContentMineStyle styles the body of messages sent by the local user
var MsgContentMineStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("255")).
	PaddingLeft(2)

// MsgContentOtherStyle styles the body of messages from other users
var MsgContentOtherStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("252")).
	PaddingLeft(2)

// --- Header styles ---

// HeaderTitleStyle styles the "whisper" title text inside the header
var HeaderTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF69B4"))

// HeaderContextStyle styles the "user @ server" centre section of the header
var HeaderContextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("245"))

// HeaderClockStyle styles the clock on the right of the header
var HeaderClockStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("245"))

// ConnectedDotStyle is the green dot shown when the gRPC stream is healthy
var ConnectedDotStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF7F"))

// DisconnectedDotStyle is the red dot shown when the gRPC stream has errored
var DisconnectedDotStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF4444"))

// --- Input styles ---

// CharCountNormalStyle styles the character counter when well below the limit
var CharCountNormalStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240"))

// CharCountWarningStyle styles the character counter when within 20 chars of the limit
var CharCountWarningStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF69B4")).
	Bold(true)

// --- Unread badge ---

// UnreadBadgeStyle styles the "↓ N new messages" indicator
var UnreadBadgeStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF69B4")).
	Bold(true).
	Padding(0, 1)
