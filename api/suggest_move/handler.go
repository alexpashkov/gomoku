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

	minimax.Minimax(state, 1)

	//moves := make([]board.Coords, 3)
	//for coords, i := reqBody.Board.RandomCoords(), 0; i < len(moves); {
	//	if !reqBody.Board.CellIsOccupied(coords) {
	//		moves[i] = coords
	//		i++
	//	}
	//	coords = reqBody.Board.RandomCoords()
	//}
	//coordsJSON, err := json.Marshal(moves)
	//if err != nil {
	//	res.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//res.Header().Set("Content-Type", "application/json")
	//_, err = res.Write(coordsJSON)
}
