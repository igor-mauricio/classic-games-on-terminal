package minesweeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNeighbourBombs(t *testing.T) {
	f := newEmptyTestField(3, 3)
	addBombsToPositions(&f, [][]int{{0, 0}, {2, 0}, {2, 1}, {2, 2}})
	f.parseNeighbourBombs()
	expectedResult := Field{
		Positions: [][]tile{
			{
				{nearbyBombs: 1, isBomb: true},
				{nearbyBombs: 1, isBomb: false},
				{nearbyBombs: 0, isBomb: false},
			},
			{
				{nearbyBombs: 3, isBomb: false},
				{nearbyBombs: 4, isBomb: false},
				{nearbyBombs: 2, isBomb: false},
			},
			{
				{nearbyBombs: 2, isBomb: true},
				{nearbyBombs: 3, isBomb: true},
				{nearbyBombs: 2, isBomb: true},
			},
		},
	}
	for i := range f.Positions {
		for j := range f.Positions[i] {
			assert.Equal(t, expectedResult.Positions[i][j], f.Positions[i][j], "nearbyBombs mismatch at (%d, %d)", i, j)
		}
	}

}

func TestShouldWin(t *testing.T) {
	f := newEmptyTestField(3, 3)
	/*
		[o,_,_]
		[o,_,_]
		[_,_,_]
	*/

	addBombsToPositions(&f, [][]int{{0, 0}, {1, 0}})
	f.RevealAt(0, 1)
	f.RevealAt(1, 1)
	f.RevealAt(2, 0)
	f.RevealAt(2, 1)
	f.RevealAt(2, 2)
	f.ToggleFlag(0, 0)
	f.ToggleFlag(1, 0)
	assert.Equal(t, VICTORY, f.Status, f.Status.toString())
}

func TestShouldNotWinWhenMissingFlags(t *testing.T) {
	f := newEmptyTestField(3, 3)
	/*
		[o,_,_]
		[o,_,_]
		[_,_,_]
	*/

	addBombsToPositions(&f, [][]int{{0, 0}, {1, 0}})
	f.RevealAt(0, 1)
	f.RevealAt(1, 1)
	f.RevealAt(2, 0)
	f.RevealAt(2, 1)
	f.RevealAt(2, 2)
	f.ToggleFlag(0, 0)
	assert.Equal(t, PLAYING, f.Status, f.Status.toString())
}
func TestShouldLoosingWhenRevealingBomb(t *testing.T) {
	f := newEmptyTestField(3, 3)
	/*
		[o,_,_]
		[o,_,_]
		[_,_,_]
	*/

	addBombsToPositions(&f, [][]int{{0, 0}, {1, 0}})
	f.RevealAt(0, 1)
	f.RevealAt(1, 1)
	f.RevealAt(2, 0)
	f.RevealAt(2, 1)
	f.RevealAt(2, 2)
	f.ToggleFlag(0, 0)
	f.RevealAt(1, 0)
	assert.Equal(t, DEFEAT, f.Status, f.Status.toString())
}

func TestShouldToggleFlag(t *testing.T) {
	f := newEmptyTestField(4, 4)
	f.ToggleFlag(2, 2)
	f.ToggleFlag(1, 1)

	assert.True(t, f.Positions[2][2].isFlagged)
	assert.True(t, f.Positions[1][1].isFlagged)
	f.ToggleFlag(1, 1)
	assert.False(t, f.Positions[1][1].isFlagged)
}

func TestShouldCascadeRevealEmptyTiles(t *testing.T) {
	f := newEmptyTestField(4, 4)
	/*
		[0,0,0,0]
		[x,0,0,0]
		[x,0,0,x]
		[0,0,x,x]
	*/

	addBombsToPositions(&f, [][]int{{1, 0}, {2, 0}, {3, 2}, {3, 3}, {2, 3}})
	f.parseNeighbourBombs()
	row, col := 0, 3
	err := f.RevealAt(row, col)
	assert.Nil(t, err)
	expectedRevealed := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 1}, {1, 2}, {1, 3}}
	expectedUnrevealed := [][]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}, {2, 3}, {3, 0}, {3, 1}, {3, 2}, {3, 3}}

	for _, position := range expectedRevealed {
		assert.True(t, f.Positions[position[0]][position[1]].isRevealed, "position %d %d", position[0], position[1])
	}
	for _, position := range expectedUnrevealed {
		assert.False(t, f.Positions[position[0]][position[1]].isRevealed, "position %d %d", position[0], position[1])
	}

	row, col = 0, 0
	err = f.RevealAt(row, col)
	assert.Nil(t, err)
	expectedRevealed = [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 1}, {1, 2}, {1, 3}, {0, 0}}
	expectedUnrevealed = [][]int{{1, 0}, {2, 0}, {2, 1}, {2, 2}, {2, 3}, {3, 0}, {3, 1}, {3, 2}, {3, 3}}

	for _, position := range expectedRevealed {
		assert.True(t, f.Positions[position[0]][position[1]].isRevealed, "position %d %d", position[0], position[1])
	}
	for _, position := range expectedUnrevealed {
		assert.False(t, f.Positions[position[0]][position[1]].isRevealed, "position %d %d", position[0], position[1])
	}

	row, col = 1, 0
	err = f.RevealAt(row, col)
	assert.Nil(t, err)
	expectedRevealed = [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 1}, {1, 2}, {1, 3}, {0, 0}, {1, 0}}
	expectedUnrevealed = [][]int{{2, 0}, {2, 1}, {2, 2}, {2, 3}, {3, 0}, {3, 1}, {3, 2}, {3, 3}}

	for _, position := range expectedRevealed {
		assert.True(t, f.Positions[position[0]][position[1]].isRevealed, "position %d %d", position[0], position[1])
	}
	for _, position := range expectedUnrevealed {
		assert.False(t, f.Positions[position[0]][position[1]].isRevealed, "position %d %d", position[0], position[1])
	}

}

func newEmptyTestField(rows, cols int) Field {
	f := Field{
		Positions: make([][]tile, rows),
	}
	for i := range f.Positions {
		f.Positions[i] = make([]tile, cols)
		for j := range f.Positions[i] {
			f.Positions[i][j] = tile{
				isBomb:      false,
				isRevealed:  false,
				nearbyBombs: 0,
			}
		}
	}
	return f
}

func addBombsToPositions(f *Field, positions [][]int) {

	for _, position := range positions {
		f.Positions[position[0]][position[1]].isBomb = true
	}
	f.parseNeighbourBombs()
}
