package minimax

import "gomoku/board"

// isWithinNFrom checks whether any of the coords withing maximum n steps from coords c
// satisfy predicate pred
func isWithinNFrom(c board.Coords, n int, pred func(coords board.Coords) bool) bool {
	for y := -n; y <= n; y++ {
		for x := -n; x <= n; x++ {
			neighboringCoords := board.Coords{
				X: c.X + x,
				Y: c.Y + y,
			}
			if neighboringCoords.AreWithin() && pred(neighboringCoords) {
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
			if b.CellIsEmpty(coords) && isWithinNFrom(coords, 2, b.CellIsOccupied) {
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
