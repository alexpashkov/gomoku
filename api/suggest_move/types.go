package suggest_move

import (
	"gomoku/board"
)

type ReqBody struct {
	Board  *board.Board `json:"board"`
	Player *int8        `json:"player"`
	Scores [2]int8      `json:"scores"`
}
