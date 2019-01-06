package board

type Row [19]int8
type Board [19]Row

// IsEmpty if the board is completely empty
func (b Board) IsEmpty(c Coords) bool {
	return b == Board{}
}

// GetCell returns a cell under coordinates
func (b Board) GetCell(c Coords) int8 {
	return b[c.Y][c.X]
}

// CellIsEmpty tells if a particular cell is not occupied
func (b Board) CellIsEmpty(c Coords) bool {
	return b.GetCell(c) == 0
}

// ForEach calls a provided function on every board cell with its values and coordinates
func (b Board) ForEach(f func(int8, Coords)) {
	for y, row := range b {
		for x, cell := range row {
			f(cell, Coords{
				X: x,
				Y: y,
			})
		}
	}
}
