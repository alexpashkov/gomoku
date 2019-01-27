package suggest_move

import (
	"gomoku/board"
)

type ReqBody struct {
	Board      *board.Board `json:"board"`
	Player     *int8        `json:"player"`
	BlackScore *int8        `json:"blackScore"`
	WhiteScore *int8        `json:"whiteScore"`
}
