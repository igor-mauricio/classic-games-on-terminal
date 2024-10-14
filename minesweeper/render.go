package minesweeper

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

var (
	unrevealedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#3D3634"))
	revealedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4"))
	flaggedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#FFFF00"))
	bombStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#FF0000"))
)

func (f Game) Render() string {
	switch f.Status {
	case VICTORY:
		return f.victoryScreen()
	case DEFEAT:
		return f.defeatScreen()
	case PLAYING:
		return f.gameScreen()
	}
	return ""
}

func (f Game) victoryScreen() string {
	result := ""
	victoryMessage := lipgloss.NewStyle().
		Width(60).
		Height(20).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Background(lipgloss.Color("#3D3634")).
		Foreground(lipgloss.Color("#00FF00"))
	result += victoryMessage.Render("Victory")
	return result
}

func (f Game) defeatScreen() string {
	result := ""
	victoryMessage := lipgloss.NewStyle().
		Width(60).
		Height(20).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Background(lipgloss.Color("#3D3634")).
		Foreground(lipgloss.Color("#FF0000"))
	result += victoryMessage.Render("Defeat\nPress enter to retry")
	return result
}

func (f Game) gameScreen() string {
	result := ""
	for i := range f.Positions {
		for j := range f.Positions[i] {
			var style lipgloss.Style
			text := ""
			if f.Positions[i][j].isFlagged {
				style = flaggedStyle
				text = "o"
			} else if !f.Positions[i][j].isRevealed {
				style = unrevealedStyle
				text = "?"
			} else if f.Positions[i][j].isBomb {
				style = bombStyle
				text = "o"
			} else {
				style = revealedStyle
				if f.Positions[i][j].nearbyBombs == 0 {
					text = " "
				} else {
					text = strconv.Itoa(f.Positions[i][j].nearbyBombs)
				}
			}
			if f.pos.row == i && f.pos.col == j {
				style = style.Background(lipgloss.Color("#FFFFFF"))
			}
			result += style.Padding(0, 1).Render(text)
		}
		result += "\n"
	}
	return result
}
