package game

import (
	"gomoku/board"
	"fmt"
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

func (s State) MakeMoveImmut(c board.Coords) State {
	s.Board.SetCell(c, s.Player)
	if caputedCells := GetCaptures(s.Board, c); len(caputedCells) != 0 {
		if len(caputedCells)%2 != 0 {
			panic(fmt.Errorf("odd captures len %d", len(caputedCells)))
		}
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
	return s
}

func (s *State) switchPlayer() {
	if s.Player == BLACK_PLAYER {
		s.Player = WHITE_PLAYER
	} else {
		s.Player = BLACK_PLAYER
	}
}
