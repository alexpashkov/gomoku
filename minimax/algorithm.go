package minimax

import (
	"gomoku/board"
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

func Minimax(state game.State, width, depth int) (int, *board.Coords) {
	if depth == 0 {
		return heuristic.Evaluation(state.Board, state.BlackScore, state.WhiteScore), nil
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
	// take n best moves and proceed with them, discard the rest
	for i := 0; i < bestMovesLen; i++ {
		move := heap.Pop(&pq).(Move)
		evaluation, _ := Minimax(move.State, width, depth-1)
		move.Evaluation = evaluation
		bestMoves[i] = move
	}
	pq.Moves = bestMoves
	heap.Init(&pq)
	move := heap.Pop(&pq).(Move)
	return 0, &move.Coords
}
