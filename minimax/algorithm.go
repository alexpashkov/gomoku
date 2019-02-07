package minimax

import (
	"gomoku/board"
	"gomoku/heuristic"
	"gomoku/minimax/priority_queue"
)

func Minimax(b board.Board, player int8, minScore, maxScore int8, depth uint) (int, *board.Coords) {
	if depth == 0 {
		return heuristic.Evaluation(b, minScore, maxScore), nil
	}
	possibleMoves := PossibleMoves(b, 1)
	q := priority_queue.New(player)
	for _, possibleMove := range possibleMoves {

	}

	return 0, nil
}
