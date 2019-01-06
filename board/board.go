package board

import (
	"math/rand"
	"time"
)

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

func (b *Board) SetCell(c Coords, val int8) int8 {
	b[c.Y][c.X] = val
	return val
}

// SetCells takes an array of coordinates and values ​​and sets them on the board
func (b *Board) SetCells(coords []Coords, vals []int8) {
	for i, c := range coords {
		b.SetCell(c, vals[i]) // not very safe, since i may be out of the range
	}
}

// CellIsEmpty tells if a particular cell is not occupied
func (b Board) CellIsEmpty(c Coords) bool {
	return b.GetCell(c) == 0
}

// CellIsEmpty tells if a particular cell is not occupied
func (b Board) CellIsOccupied(c Coords) bool {
	return !b.CellIsEmpty(c)
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

func (b Board) RandomCoords() Coords {
	rand.Seed(time.Now().UnixNano())
	return Coords{
		X: rand.Intn(len(b)),
		Y: rand.Intn(len(b)),
	}
}

// Fill n random cells with val
func (b *Board) FillRandomly(n int, val int8) {
	coords := make(map[Coords]bool)
	for ; n > 0; n-- {
		randomCoords := b.RandomCoords()
		_, ok := coords[randomCoords]
		for ok {
			randomCoords := b.RandomCoords()
			_, ok = coords[randomCoords]
		}
		coords[randomCoords] = true
		b.SetCell(randomCoords, val)
	}
}
