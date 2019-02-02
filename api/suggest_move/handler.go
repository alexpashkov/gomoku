package suggest_move

import (
	"encoding/json"
	"gomoku/board"
	"gomoku/minimax"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	reqBody := SuggestMoveBody{}
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	// check if there is no error decoding the body and do some preliminary
	// validation
	if err != nil ||
		reqBody.Board == nil ||
		reqBody.Player == nil ||
		*reqBody.Player != minimax.MIN_PLAYER && *reqBody.Player != minimax.MAX_PLAYER {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	moves := make([]board.Coords, 3)
	for coords, i := reqBody.Board.RandomCoords(), 0; i < len(moves); {
		if !reqBody.Board.CellIsOccupied(coords) {
			moves[i] = coords
			i++
		}
		coords = reqBody.Board.RandomCoords()
	}
	coordsJSON, err := json.Marshal(moves)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(coordsJSON)
}
