package minimax

import (
	"gomoku/minimax/heuristic"
	"container/heap"
	"gomoku/game"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Minimax(state game.State, width, depth int) []Move {
	if depth == 0 {
		return []Move{
			{
				State:      state,
				Evaluation: heuristic.Evaluation(state.Board, state.BlackScore, state.WhiteScore),
			},
		}
	}
	// get all possible moves (cells adjacent to occupied cells
	pq := NewPriorityQueue(state.Player)
	moveCoords := PossibleMoves(state.Board, 1)
	pq.Moves = make([]Move, len(moveCoords))
	// traverse all moves and order them by priority
	for i, coords := range moveCoords {
		stateAfterMove := state.MakeMoveImmut(coords)
		pq.Moves[i] = Move{
			Coords: coords,
			State:  stateAfterMove,
			index:  i,
			Evaluation: heuristic.Evaluation(stateAfterMove.Board,
				stateAfterMove.BlackScore, stateAfterMove.WhiteScore),
		}
	}
	heap.Init(&pq)
	bestMovesLen := Min(width, pq.Len())
	bestMoves := make([]Move, bestMovesLen)
	// take best moves and proceed with them, discard the rest
	for i := 0; i < bestMovesLen; i++ {
		move := heap.Pop(&pq).(Move)
		move.Evaluation = Minimax(move.State, width, depth-1)[0].Evaluation
		bestMoves[i] = move
	}
	pq.Moves = bestMoves
	heap.Init(&pq)
	return pq.Slice(pq.Len())
}
