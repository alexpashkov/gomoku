package minimax

import (
	"gomoku/board"
	"gomoku/heuristic"
	"container/heap"
	"gomoku/game"
)

func Minimax(state game.State, depth uint) (int, *board.Coords) {
	if depth == 0 {
		return heuristic.Evaluation(state.Board, state.BlackScore, state.WhiteScore), nil
	}
	// get all possible moves (cells adjacent to occupied cells
	possibleMoves := PossibleMoves(state.Board, 1)
	pq := NewPriorityQueue(state.Player)
	// traverse all moves and order them by priority
	for i, possibleMove := range possibleMoves {
		stateAfterMove := state.MakeMoveImmut(possibleMove)
		heap.Push(&pq, Move{
			Coords: possibleMove,
			State:  stateAfterMove,
			index:  i,
			Evaluation: heuristic.Evaluation(stateAfterMove.Board,
				stateAfterMove.BlackScore, stateAfterMove.WhiteScore),
		})
	}
	// take n best moves and proceed with them, discard the rest

	return 0, nil
}
