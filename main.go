package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"gomoku/api/suggest_move"
	"gomoku/board"
	"gomoku/minimax/heuristic"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT = ":4444"

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
			fmt.Println(brd)
			c := board.Coords{5, 5}
			fmt.Println(heuristic.IsCorrectMove(brd, 2, c))
		}
	})
	mux.HandleFunc("/suggest-move", suggest_move.Handler)
	println("Starting server at " + PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(PORT), cors.Default().Handler(mux)))
}
