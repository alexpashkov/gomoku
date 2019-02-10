package minimax

import (
	"gomoku/minimax/heuristic"
	"gomoku/game"
	"sort"
	"log"
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
	moves := make(Moves, 0, len(cellsAdjacentToOccupied))

	for _, coords := range cellsAdjacentToOccupied {
		state, err := state.Move(coords)
		if err == nil {
			moves = append(moves, &Move{
				Coords:     coords,
				State:      state,
				Evaluation: heuristic.Evaluation(state.Board, state.BlackScore, state.WhiteScore),
			})
		} else {
			log.Printf("%v", err)
		}
	}
	sort.Sort(moves)
	moves = moves[:Min(maxWidth, moves.Len())]
	// take best moves and proceed with them, discard the rest
	for _, move := range moves {
		move.Evaluation = Minimax(move.State, maxWidth, depth-1)[0].Evaluation
	}
	sort.Sort(moves)
	return moves
}
