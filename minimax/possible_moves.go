package minimax

import "board"

func PossibleMoves(b board.Board, _ int8) []board.Coords {
	moves := make([]board.Coords, 0, len(b))
	b.ForEach(func(cell int8, coords board.Coords) {
		moves = append(moves, coords)
	})
	return moves
}
