package heuristic

import (
	"gomoku/board"
)

func IsCenter(cord board.Coords, state int8) bool {
	if state == 7 {
		if (cord.X >= MIN_X_7X7_BOARD && cord.X <= MAX_X_7X7_BOARD) &&
			(cord.Y >= MIN_Y_7X7_BOARD && cord.Y <= MAX_Y_7X7_BOARD) {
			return true
		}
	} else if state == 5 {
		if (cord.X >= MIN_X_5X5_BOARD && cord.X <= MAX_X_5X5_BOARD) &&
			(cord.Y >= MIN_Y_5X5_BOARD && cord.Y <= MAX_Y_5X5_BOARD) {
			return true
		}
	}
	return false
}

func IsCorrectMovePro(cord board.Coords, countMove int, state int8) bool {
	if countMove == 1 {
		return IsCenter(cord, state)
	} else if countMove == 3 {
		if IsCenter(cord, state) == true {
			return false
		} else {
			return true
		}
	}
	return true
}
