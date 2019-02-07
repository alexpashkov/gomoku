package game

import "gomoku/board"

const (
	BLACK_PLAYER = 1
	WHITE_PLAYER = 2
)

type State struct {
	Board                  board.Board
	Player                 int8
	BlackScore, WhiteScore int8
}

func (s State) MakeMoveImmut(c board.Coords) State {
	// TODO handle capturing too
	s.Board.SetCell(c, s.Player)
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
