package handlers

import (
	"encoding/json"
	"gomoku/game"
	"gomoku/minimax"
	"gomoku/minimax/heuristic"
	"net/http"
	"math"
)

func SuggestMove(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	if json.NewDecoder(req.Body).Decode(&state) != nil || state.Winner != 0 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	moves := minimax.Minimax(state, 3, 5, heuristic.Evaluation, math.MinInt64, math.MaxInt64)
	resBody, err := json.Marshal(moves)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(resBody)
}
