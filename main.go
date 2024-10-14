package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/igor-mauricio/classic-games-on-terminal/minesweeper"
)

type model struct {
	choices     []string         // items on the to-do list
	cursorX     int              // which to-do list item our cursor is pointing at
	cursorY     int              // which to-do list item our cursor is pointing at
	selected    map[int]struct{} // which to-do items are selected
	minesweeper minesweeper.Minesweeper
}

func initialModel() model {
	return model{
		choices:     []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected:    make(map[int]struct{}),
		minesweeper: minesweeper.Create(20, 20, minesweeper.EASY),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			m.minesweeper.MoveTo(minesweeper.LEFT)
		case "right":
			m.minesweeper.MoveTo(minesweeper.RIGHT)
		case "up":
			m.minesweeper.MoveTo(minesweeper.UP)

		case "down":
			m.minesweeper.MoveTo(minesweeper.DOWN)
		case "e", "f":
			m.minesweeper.ToggleFlag()
		case "enter", " ":
			status, _ := m.minesweeper.Reveal()
			if status != minesweeper.PLAYING {
				m.minesweeper.NewGame(20, 20, minesweeper.EASY)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Minesweeper\n\n"
	s += m.minesweeper.Render()
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
