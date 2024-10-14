package minesweeper

import (
	"fmt"
	"math/rand"
)

type Minesweeper interface {
	NewGame(rows, cols int, dif difficulty) error
	Reveal() (gameStatus, error)
	ToggleFlag() error
	Render() string
	MoveTo(dir direction) error
}

type Game struct {
	Positions [][]cell
	Status    gameStatus
	pos       struct{ row, col int }
	size      struct{ rows, cols int }
}

type cell struct {
	nearbyBombs int
	isBomb      bool
	isRevealed  bool
	isFlagged   bool
}

type difficulty int

const (
	EASY difficulty = iota
	MEDIUM
	HARD
)

type gameStatus int

const (
	PLAYING gameStatus = iota
	VICTORY
	DEFEAT
)

func (status gameStatus) toString() string {
	switch status {
	case PLAYING:
		return "Playing"
	case VICTORY:
		return "Victory"
	case DEFEAT:
		return "Defeat"
	default:
		return ""
	}
}

func Create(rows, cols int, dif difficulty) Minesweeper {
	g := &Game{}
	g.NewGame(rows, cols, dif)
	return g
}

func (g *Game) NewGame(rows, cols int, dif difficulty) error {
	if rows < 0 || cols < 0 {
		return fmt.Errorf("out of bounds")
	}
	g.Positions = make([][]cell, rows)
	g.Status = PLAYING
	g.pos = struct {
		row int
		col int
	}{0, 0}
	g.size = struct {
		rows int
		cols int
	}{rows, cols}

	var bombPercent float32
	switch dif {
	case EASY:
		bombPercent = 0.2
	case MEDIUM:
		bombPercent = 0.4
	case HARD:
		bombPercent = 0.6
	default:
		bombPercent = 0.4
	}
	for i := range g.Positions {
		g.Positions[i] = make([]cell, cols)
		for j := range g.Positions[i] {
			g.Positions[i][j] = cell{
				isBomb:      rand.Float32() < bombPercent,
				isRevealed:  false,
				isFlagged:   false,
				nearbyBombs: 0,
			}
		}
	}
	g.parseNeighbourBombs()
	return nil
}

func (g *Game) ToggleFlag() error {
	return g.toggleFlagAt(g.pos.row, g.pos.col)
}

func (g *Game) toggleFlagAt(row, col int) error {
	if g.Status != PLAYING {
		return fmt.Errorf("game finished")
	}
	g.Positions[row][col].isFlagged = !g.Positions[row][col].isFlagged
	g.checkWin()
	return nil
}

func (g *Game) Reveal() (gameStatus, error) {
	return g.revealAt(g.pos.row, g.pos.col)
}

func (g *Game) revealAt(row, col int) (gameStatus, error) {
	if g.Status != PLAYING {
		return g.Status, fmt.Errorf("game finished")
	}
	defer g.checkWin()
	g.Positions[row][col].isRevealed = true
	if g.Positions[row][col].isBomb || g.Positions[row][col].nearbyBombs != 0 {
		return g.Status, nil
	}
	dxList := []int{-1, 0, 1}
	dyList := []int{-1, 0, 1}
	for _, dx := range dxList {
		i := row + dx
		if i < 0 || i > len(g.Positions)-1 {
			continue
		}
		for _, dy := range dyList {
			j := col + dy
			if j < 0 || j > len(g.Positions[i])-1 || g.Positions[i][j].isRevealed {
				continue
			}
			g.revealAt(i, j)
		}
	}
	return g.Status, nil
}

type direction struct{ dx, dy int }

var (
	UP    = direction{dx: 0, dy: -1}
	DOWN  = direction{dx: 0, dy: 1}
	LEFT  = direction{dx: -1, dy: 0}
	RIGHT = direction{dx: 1, dy: 0}
)

func (g *Game) MoveTo(dir direction) error {
	row := g.pos.row + dir.dy
	col := g.pos.col + dir.dx
	if row < 0 || col < 0 || row >= g.size.rows || col >= g.size.cols {
		return fmt.Errorf("out of bounds")
	}
	g.pos.row = row
	g.pos.col = col
	return nil
}

func (g *Game) checkWin() {
	stillPlaying := false

	for i := range g.Positions {
		for j := range g.Positions[i] {
			if g.Positions[i][j].isBomb && g.Positions[i][j].isRevealed {
				g.Status = DEFEAT
				return
			}
			if g.Positions[i][j].isBomb && !g.Positions[i][j].isFlagged ||
				!g.Positions[i][j].isBomb && !g.Positions[i][j].isRevealed {
				stillPlaying = true
			}
		}
	}
	if stillPlaying {
		g.Status = PLAYING
		return
	}
	g.Status = VICTORY
}
