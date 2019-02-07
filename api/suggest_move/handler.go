package suggest_move

import (
	"encoding/json"
	"gomoku/game"
	"gomoku/minimax"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	if json.NewDecoder(req.Body).Decode(&state) != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	moves := minimax.Minimax(state, 500, 2)
	resBody, err := json.Marshal(moves)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(resBody)
}
