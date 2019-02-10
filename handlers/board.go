package handlers

import (
	"net/http"
	"io/ioutil"
	"log"
	"gomoku/board"
	"encoding/json"
	"fmt"
	"gomoku/minimax/heuristic"
)

func Board(_ http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	brd := board.Board{}
	err = json.Unmarshal(body, &brd)
	if err != nil {
		fmt.Println("Invalid board sent")
	} else {
		threat := []heuristic.Threat{}
		fmt.Println("START")
		heuristic.SearchThreatRowClose(brd, threat, 4)
		//fmt.Println(heuristic.IsTerminate(brd, 0, 0))
		//c := board.Coords{6, 7}
		//fmt.Println(heuristic.IsCorrectMove(brd, 1, c))
	}
}
