package handlers

import (
	"encoding/json"
	"gomoku/game"
	"gomoku/minimax"
	"gomoku/minimax/heuristic"
	"net/http"
)

func SuggestMove(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	if json.NewDecoder(req.Body).Decode(&state) != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	moves := minimax.Minimax(state, 3, 4, heuristic.Evaluation)
	resBody, err := json.Marshal(moves)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(resBody)
}
