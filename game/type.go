package game

import (
	"gomoku/board"
)

const (
	BLACK_PLAYER int8 = 1
	WHITE_PLAYER int8 = 2
)

type State struct {
	Board      board.Board `json:"board"`
	Player     int8        `json:"player"`
	BlackScore int8        `json:"blackScore"`
	WhiteScore int8        `json:"whiteScore"`
}

func (s State) MakeMoveImmut(c board.Coords) State {
	// TODO handle capturing too
	s.Board.SetCell(c, int8(s.Player))
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
