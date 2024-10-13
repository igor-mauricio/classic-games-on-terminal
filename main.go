package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/igor-mauricio/classic-games-on-terminal/minesweeper"
)

type model struct {
	choices   []string         // items on the to-do list
	cursorX   int              // which to-do list item our cursor is pointing at
	cursorY   int              // which to-do list item our cursor is pointing at
	selected  map[int]struct{} // which to-do items are selected
	minefield minesweeper.Field
}

func initialModel() model {

	return model{
		choices:   []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected:  make(map[int]struct{}),
		minefield: minesweeper.NewField(5, 5, minesweeper.EASY),
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
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right":
			if m.cursorX < len(m.minefield.Positions[0])-1 {
				m.cursorX++
			}
		case "up":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "down":
			if m.cursorY < len(m.minefield.Positions)-1 {
				m.cursorY++
			}
		case "e", "f":
			m.minefield.ToggleFlag(m.cursorY, m.cursorX)
		case "enter", " ":
			m.minefield.RevealAt(m.cursorY, m.cursorX)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Minesweeper\n\n"
	s += m.minefield.Render(m.cursorX, m.cursorY)
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
