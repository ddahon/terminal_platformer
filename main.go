package main

// These imports will be used later on the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type player struct {
	x           int
	y           int
	height      int
	width       int
	displayChar rune
}

type model struct {
	player player
}

func initialModel() model {
	return model{
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
			m.player.moveRight()
		case "q":
			m.player.moveLeft()
		}
	}

	return m, nil
}

func (m model) View() string {
	s := m.player.View()
	return s
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (p player) sprite() string {
	row := strings.Repeat(" ", p.x) + strings.Repeat(string(p.displayChar), p.width)
	return strings.Repeat(row+"\n", p.height)
}

func (p player) View() string {
	y_offset := strings.Repeat("\n", p.y)
	return y_offset + p.sprite()
}

func (p *player) moveLeft() {
	p.x-- //todo: find width of terminal and wrap around it
}

func (p *player) moveRight() {
	p.x++
}
