package minimax

import (
	"gomoku/minimax/heuristic"
	"gomoku/game"
	"sort"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Minimax(state game.State, maxWidth, depth int) Moves {
	if depth == 0 {
		return []*Move{
			{
				State:      state,
				Evaluation: heuristic.Evaluation(state.Board, state.BlackScore, state.WhiteScore),
			},
		}
	}
	// get all cells adjacent to occupied cells
	cellsAdjacentToOccupied := CellsAdjacentToOccupied(state.Board, 1)
	moves := make(Moves, len(cellsAdjacentToOccupied))

	for i, coords := range cellsAdjacentToOccupied {
		state := state.MakeMoveImmut(coords)
		moves[i] = &Move{
			Coords:     coords,
			State:      state,
			Evaluation: heuristic.Evaluation(state.Board, state.BlackScore, state.WhiteScore),
		}
	}
	sort.Sort(moves)
	// take best moves and proceed with them, discard the rest
	for _, move := range moves {
		move.Evaluation = Minimax(move.State, maxWidth, depth-1)[0].Evaluation
	}
	sort.Sort(moves)
	return moves
}
