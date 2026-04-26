package client

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ch55secake/whisper/pkg/config"
	messenger "github.com/ch55secake/whisper/pkg/server/generated"
)

// buildMenuList constructs the bubbles list used on the main menu screen.
func buildMenuList() list.Model {
	items := []list.Item{
		menuItem{title: "Connect", desc: "Connect to the whisper server and start chatting"},
		menuItem{title: "Quit", desc: "Exit whisper"},
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = MenuSelectedStyle
	delegate.Styles.SelectedDesc = MenuNormalStyle.Faint(true)
	delegate.Styles.NormalTitle = MenuNormalStyle
	delegate.Styles.NormalDesc = MenuNormalStyle.Faint(true)

	l := list.New(items, delegate, 40, 10)
	l.Title = ""
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)

	return l
}

// StartClient starts the client and initializes the model, will most likely have to move this
// TODO: Move this code at some point
func StartClient() {
	input := textinput.New()
	input.Prompt = "> "
	input.Placeholder = "Type a message…"
	input.CharLimit = 256

	serverAddr := config.ServerAddress()
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}

	defer conn.Close()

	client := messenger.NewMessengerClient(conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Failed to open chat stream: %v", err)
	}

	m := Model{
		input:      input,
		menuList:   buildMenuList(),
		messages:   []Message{},
		client:     client,
		stream:     stream,
		phase:      menu,
		connected:  true,
		serverAddr: serverAddr,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
