package minesweeper

import (
	"fmt"
	"math/rand"
)

type tile struct {
	nearbyBombs int
	isBomb      bool
	isRevealed  bool
	isFlagged   bool
}

type Field struct {
	Positions [][]tile
	Status    gameStatus
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

func (g gameStatus) toString() string {
	switch g {
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

func NewField(rows, cols int, dif difficulty) Field {
	f := Field{
		Positions: make([][]tile, rows),
	}
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
	for i := range f.Positions {
		f.Positions[i] = make([]tile, cols)
		for j := range f.Positions[i] {
			f.Positions[i][j] = tile{
				isBomb:      rand.Float32() < bombPercent,
				isRevealed:  false,
				isFlagged:   false,
				nearbyBombs: 0,
			}
		}
	}
	f.parseNeighbourBombs()
	return f
}

func (f *Field) ToggleFlag(row, col int) error {
	if f.Status != PLAYING {
		return fmt.Errorf("game finished")
	}
	f.Positions[row][col].isFlagged = !f.Positions[row][col].isFlagged
	f.checkWin()
	return nil
}

func (f *Field) RevealAt(row, col int) error {
	if f.Status != PLAYING {
		return fmt.Errorf("game finished")
	}
	defer f.checkWin()
	if row < 0 || col < 0 || row >= len(f.Positions) || col >= len(f.Positions[0]) {
		return fmt.Errorf("out of bounds")
	}
	f.Positions[row][col].isRevealed = true
	if f.Positions[row][col].isBomb || f.Positions[row][col].nearbyBombs != 0 {
		return nil
	}
	dxList := []int{-1, 0, 1}
	dyList := []int{-1, 0, 1}
	for _, dx := range dxList {
		i := row + dx
		if i < 0 || i > len(f.Positions)-1 {
			continue
		}
		for _, dy := range dyList {
			j := col + dy
			if j < 0 || j > len(f.Positions[i])-1 || f.Positions[i][j].isRevealed {
				continue
			}
			f.RevealAt(i, j)
		}
	}
	return nil
}

func (f *Field) parseNeighbourBombs() {
	for i := range f.Positions {
		for j := range f.Positions[0] {
			f.Positions[i][j].nearbyBombs = f.getNeighbourBombCount(i, j)
		}
	}
}

func (f Field) getNeighbourBombCount(row, col int) int {
	dxList := []int{-1, 0, 1}
	dyList := []int{-1, 0, 1}
	count := 0

	for _, dx := range dxList {
		i := row + dx
		if i < 0 || i > len(f.Positions)-1 {
			continue
		}
		for _, dy := range dyList {
			j := col + dy
			if j < 0 || j > len(f.Positions[i])-1 {
				continue
			}
			if f.Positions[i][j].isBomb {
				count++
			}
		}
	}
	return count
}

func (f *Field) checkWin() {
	stillPlaying := false

	for i := range f.Positions {
		for j := range f.Positions[i] {
			if f.Positions[i][j].isBomb && f.Positions[i][j].isRevealed {
				f.Status = DEFEAT
				return
			}
			if f.Positions[i][j].isBomb && !f.Positions[i][j].isFlagged ||
				!f.Positions[i][j].isBomb && !f.Positions[i][j].isRevealed {
				stillPlaying = true
			}
		}
	}
	if stillPlaying {
		f.Status = PLAYING
		return
	}
	f.Status = VICTORY
}
