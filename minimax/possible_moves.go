package minimax

import "board"

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

func PossibleMoves(b board.Board, player int8) map[board.Coords]bool {
	moves := make(map[board.Coords]bool)
	if b == (board.Board{}) {
		// Suggest to place into the center of the board if it's empty
		moves[board.Coords{X: 10, Y: 10}] = true
	} else {
		b.ForEach(func(cell int8, coords board.Coords) {
			moves[coords] = true
		})
	}
	return moves
}
