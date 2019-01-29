package main

import (
	"encoding/json"
	"fmt"
	"gomoku/board"
	"io/ioutil"
	"log"
	"net/http"
	"gomoku/minimax/heuristic"
)

func main() {
	//threat := []heuristic.Threat{}

	http.HandleFunc("/board", func(res http.ResponseWriter, req *http.Request) {
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
	http.Handle("/", http.FileServer(http.Dir("client/build")))
	log.Fatalln(http.ListenAndServe(":4444", nil))
}

