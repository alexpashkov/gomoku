package minimax

import "gomoku/board"

func isWithinTwoFromOccupiedCell(b board.Board, c board.Coords) bool {
	for y := -2; y <= 2; y++ {
		for x := -2; x <= 2; x++ {
			neighboringCoords := board.Coords{
				X: c.X + x,
				Y: c.Y + y,
			}
			if neighboringCoords.AreWithin(b) &&
				!b.CellIsEmpty(neighboringCoords) {
				return true
			}
		}
	}
	return false
}

func PossibleMoves(b board.Board) []board.Coords {
	moves := make([]board.Coords, 0, len(b))
	if b == (board.Board{}) {
		// Suggest to place into the center of the board if it's empty
		moves = append(moves, board.Coords{X: 10, Y: 10})
	} else {
		// Go through all cells and add those that are empty and withing 2 cells
		// from occupied to the set of possible moves
		// In golang, set data structures are implemented as map[key]bool
		b.ForEach(func(cell int8, coords board.Coords) {
			if b.CellIsEmpty(coords) && isWithinTwoFromOccupiedCell(b, coords) {
				moves = append(moves, coords)
			}
		})
	}
	return moves
}

func PossibleMovesEq(a, b []board.Coords) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
