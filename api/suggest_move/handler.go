package suggest_move

import (
	"encoding/json"
	"net/http"
	"gomoku/game"
	"gomoku/minimax"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	if json.NewDecoder(req.Body).Decode(&state) != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	println("=========================================")
	moves := minimax.Minimax(state, 3,2)
	println("=========================================")
	resBody, err := json.Marshal(moves)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(resBody)
}
