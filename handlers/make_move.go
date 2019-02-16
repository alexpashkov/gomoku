package handlers

import (
	"net/http"
	"gomoku/game"
	"encoding/json"
	"gomoku/board"
	"errors"
	"fmt"
)

func MakeMove(res http.ResponseWriter, req *http.Request) {
	body := struct {
		State  *game.State   `json:"state"`
		Coords *board.Coords `json:"coords"`
	}{}
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		sendError(res, fmt.Errorf("invalid body: %v", err), http.StatusBadRequest)
	}
	if body.State == nil || body.Coords == nil {
		sendError(res, errors.New("state and coords are required"), http.StatusBadRequest)
		return
	}
	newState, err := body.State.Move(*body.Coords)
	if err != nil {
		sendError(res, fmt.Errorf("cannot make move: %v", err), http.StatusForbidden)
		return
	}
	marshaledState, err := json.Marshal(newState)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Write(marshaledState)
}
