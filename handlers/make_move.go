package handlers

import (
	"net/http"
	"gomoku/game"
	"encoding/json"
	"gomoku/board"
)

func MakeMove(res http.ResponseWriter, req *http.Request) {
	body := struct {
		State  *game.State   `json:"state"`
		Coords *board.Coords `json:"coords"`
	}{}
	if json.NewDecoder(req.Body).Decode(&body) != nil ||
		body.State == nil || body.Coords == nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	newState, err := body.State.Move(*body.Coords)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	marshaledState, err := json.Marshal(newState)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Write(marshaledState)
}
