package suggest_move

import (
	"gomoku/board"
)

type SuggestMoveBody struct {
	Board  *board.Board `json:"board"`
	Player *int8        `json:"player"`
	Scores [3]int8      `json:"scores"`
}
