package minimax

import (
	"gomoku/board"
	"gomoku/game"
	"log"
	"math"
	"sort"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Minimax(state game.State, maxWidth, depth int,
	heuristic func(b board.Board, blackScore, whiteScore int8) int64) Moves {
	if depth == 0 {
		return Moves{
			{
				State:      state,
				Evaluation: heuristic(state.Board, state.BlackScore, state.WhiteScore),
			},
		}
	}
	// get all cells adjacent to occupied cells
	cellsAdjacentToOccupied := CellsAdjacentToOccupied(state.Board, 1)
	moves := make(Moves, 0, len(cellsAdjacentToOccupied))

	for _, coords := range cellsAdjacentToOccupied {
		state, err := state.Move(coords)
		if err != nil {
			log.Printf("%v", err)
			continue
		}
		move := &Move{
			Coords: coords,
			State:  state,
		}
		if state.Player == state.Winner {
			if state.Player == board.BLACK_PLAYER {
				move.Evaluation = math.MinInt64
			}
			if state.Player == board.WHITE_PLAYER {
				move.Evaluation = math.MaxInt64
			}
			return []*Move{move}
		}
		move.Evaluation = heuristic(state.Board, state.BlackScore, state.WhiteScore)
		moves = append(moves, move)
	}
	sort.Sort(moves)
	moves = moves[:Min(maxWidth, moves.Len())]
	// take best moves and proceed with them, discard the rest
	for _, move := range moves {
		opponentMoves := Minimax(move.State, maxWidth, depth-1, heuristic)
		if len(opponentMoves) != 0 {
			move.Evaluation = opponentMoves[0].Evaluation
		}
	}
	sort.Sort(moves)
	return moves
}
