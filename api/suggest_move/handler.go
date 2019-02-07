package suggest_move

import (
	"encoding/json"
	"net/http"
	"gomoku/game"
	"gomoku/minimax"
	"gomoku/board"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	if json.NewDecoder(req.Body).Decode(&state) != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	println("=========================================")
	_, bestMove := minimax.Minimax(state, 4,2)
	println("=========================================")
	//println(evaluation)
	resBody, err := json.Marshal([]board.Coords{*bestMove})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(resBody)
}
