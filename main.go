package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"gomoku/api/suggest_move"
	"gomoku/board"
	"io/ioutil"
	"log"
	"net/http"
	"gomoku/minimax/heuristic"
)

const PORT = ":5555"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/board", func(res http.ResponseWriter, req *http.Request) {
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
			//threat := []heuristic.Threat{}
			fmt.Println("START")
			fmt.Println(heuristic.IsTerminate(brd, 0,0))
			//fmt.Println(heuristic.IsTerminate(brd, 0, 0))
			//c := board.Coords{6, 7}
			//fmt.Println(heuristic.IsCorrectMove(brd, 1, c))
		}
	})
	mux.HandleFunc("/suggest-move", suggest_move.Handler)
	println("Starting server at " + PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(PORT), cors.Default().Handler(mux)))
}
