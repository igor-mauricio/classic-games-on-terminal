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

func (f Field) Render(currentX, currentY int) string {
	switch f.Status {
	case VICTORY:
		return f.victoryScreen()
	case DEFEAT:
		return f.defeatScreen()
	case PLAYING:
		return f.gameScreen(currentX, currentY)
	}
	return ""
}

func (f Field) victoryScreen() string {
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

func (f Field) defeatScreen() string {
	result := ""
	victoryMessage := lipgloss.NewStyle().
		Width(60).
		Height(20).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Background(lipgloss.Color("#3D3634")).
		Foreground(lipgloss.Color("#00FF00"))
	result += victoryMessage.Render("Defeat")
	return result
}

func (f Field) gameScreen(currentX, currentY int) string {
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
			if currentY == i && currentX == j {
				style = style.Background(lipgloss.Color("#FFFFFF"))
			}
			result += style.Padding(0, 1).Render(text)
		}
		result += "\n"
	}
	return result
}
