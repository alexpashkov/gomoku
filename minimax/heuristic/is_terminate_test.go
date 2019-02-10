package heuristic

import (
	"testing"
	"gomoku/board"
)

var IsTerminateValidTestCases = map[string]struct {
	board.Board
	blackScore, whiteScore, expected int8
}{
	"score black": {
		Board:      board.Board{},
		blackScore: 12,
		whiteScore: 0,
		expected:   board.BLACK_PLAYER,
	},
	"score white": {
		Board:      board.Board{},
		blackScore: 0,
		whiteScore: 12,
		expected:   board.WHITE_PLAYER,
	},
}

func TestIsTerminate(t *testing.T) {
	for descr, testCase := range IsTerminateValidTestCases {
		received := IsTerminate(testCase.Board, testCase.blackScore, testCase.whiteScore)
		if received != testCase.expected {
			t.Errorf("%s, got: %d, expected: %d", descr, received, testCase.expected)
		}
	}
}
