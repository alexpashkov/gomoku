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

// SetChips takes an array of coordinates and values ​​and sets them on the board
func (b *Board) SetChips(coordinates []Coords, status []int8) {
	i := 0
	amountChip := len(status)
	for i < amountChip {
		b[coordinates[i].Y][coordinates[i].X] = status[i]
		i++
	}
}

// IsTerminate a function that checks if there is a winner, and returns who won.
func (b Board) IsTerminate() int8 {
	for y, row := range b {
		for x := range row {
			if (b.GetCell(Coords{x, y}) != 0) {
				i := 1
				amountX := 0
				amountY := 0
				amountRightZ := 0
				amountLeftZ := 0
				for i <= 4 {
					if (x+i < 19 && b.GetCell(Coords{x + i, y}) == b.GetCell(Coords{x, y})) {
						amountX++
					}
					if (y+i < 19 && b.GetCell(Coords{x, y + i}) == b.GetCell(Coords{x, y})) {
						amountY++
					}
					if (x+i < 19 && y+i < 19 && b.GetCell(Coords{x + i, y + i}) == b.GetCell(Coords{x, y})) {
						amountRightZ++
					}
					if (x-i > 1 && y+i < 19 && b.GetCell(Coords{x - i, y + i}) == b.GetCell(Coords{x, y})) {
						amountLeftZ++
					}
					i++
				}
				if amountX == 4 || amountY == 4 || amountRightZ == 4 || amountLeftZ == 4 {
					return b.GetCell(Coords{x, y})
				}
			}
		}
	}
	return 0
}
