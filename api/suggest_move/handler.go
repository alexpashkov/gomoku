package suggest_move

import (
	"encoding/json"
	"gomoku/minimax"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	reqBody := &ReqBody{}
	err := json.NewDecoder(req.Body).Decode(reqBody)
	// check if there is no error decoding the body and do some preliminary
	// validation
	if err != nil ||
		reqBody.Board == nil ||
		reqBody.Player == nil ||
		*reqBody.Player != minimax.MIN_PLAYER && *reqBody.Player != minimax.MAX_PLAYER ||
		reqBody.BlackScore == nil ||
		reqBody.WhiteScore == nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}
