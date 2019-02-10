package game

import (
	"gomoku/board"
)

const (
	BLACK_PLAYER int8 = 1
	WHITE_PLAYER int8 = 2
)

type State struct {
	Player     int8        `json:"player"`
	BlackScore int8        `json:"blackScore"`
	WhiteScore int8        `json:"whiteScore"`
	Board      board.Board `json:"board"`
}

// Move returns new state after applying
func (s State) Move(coords board.Coords) (State, error) {
	s.Board.SetCell(coords, s.Player)
	if caputedCells := GetCaptures(s.Board, coords); len(caputedCells) != 0 {
		for _, capturedCell := range caputedCells {
			s.Board.SetCell(capturedCell, 0)
		}
		if s.Player == BLACK_PLAYER {
			s.BlackScore += int8(len(caputedCells))
		} else {
			s.WhiteScore += int8(len(caputedCells))
		}
	}
	s.switchPlayer()
	return s, nil
}

func (s *State) switchPlayer() {
	if s.Player == BLACK_PLAYER {
		s.Player = WHITE_PLAYER
	} else {
		s.Player = BLACK_PLAYER
	}
}
