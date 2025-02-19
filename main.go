package main

import (
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type player struct {
	x           int
	y           int
	height      int
	width       int
	displayChar rune
}

type screen struct {
	height int
	width  int
}

type model struct {
	screen screen
	player player
}

func initialModel() model {
	screenWidth, screenHeight, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Failed to get terminal size: %s", err)
	}
	return model{
		screen: screen{
			height: screenHeight,
			width:  screenWidth,
		},
		player: player{
			15,
			0,
			5,
			4,
			'*',
		},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		case "d":
			m.player.moveRight(m.screen.width)
		case "q":
			m.player.moveLeft(m.screen.width)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := m.player.View()
	return s
}

func (p player) sprite() string {
	row := strings.Repeat(" ", p.x) + strings.Repeat(string(p.displayChar), p.width)
	return strings.Repeat(row+"\n", p.height)
}

func (p player) View() string {
	y_offset := strings.Repeat("\n", p.y)
	return y_offset + p.sprite()
}

func (p *player) moveLeft(screenWidth int) {
	p.x = ((p.x-1)%screenWidth + screenWidth) % screenWidth
}

func (p *player) moveRight(screenWidth int) {
	p.x = (p.x + 1) % screenWidth
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
