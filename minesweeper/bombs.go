package minesweeper

func (f *Game) parseNeighbourBombs() {
	for i := range f.Positions {
		for j := range f.Positions[0] {
			f.Positions[i][j].nearbyBombs = f.getNeighbourBombCount(i, j)
		}
	}
}

func (f Game) getNeighbourBombCount(row, col int) int {
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
