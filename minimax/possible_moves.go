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

func PossibleMoves(b board.Board) map[board.Coords]bool {
	moves := make(map[board.Coords]bool)
	if b == (board.Board{}) {
		// Suggest to place into the center of the board if it's empty
		moves[board.Coords{X: 10, Y: 10}] = true
	} else {
		// Go through all cells and add those that are empty and withing 2 cells
		// from occupied to the set of possible moves
		// In golang, set data structures are implemented as map[key]bool
		b.ForEach(func(cell int8, coords board.Coords) {
			if b.CellIsEmpty(coords) && isWithinTwoFromOccupiedCell(b, coords) {
				moves[coords] = true
			}
		})
	}
	return moves
}
