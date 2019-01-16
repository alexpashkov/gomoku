package main

import (
	"encoding/json"
	"fmt"
	"gomoku/board"
	"io/ioutil"
	"log"
	"net/http"
	"gomoku/minimax/heuristic"
	"time"
)

func main() {
	threat := []minimax.Threat{}
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
			t0 := time.Now()
			threat = minimax.SearchThreatStraightFour(brd, threat)
			//threat = minimax.SearchThreatFourInRow(brd, threat)
			t1 := time.Now()
			fmt.Printf("Elapsed time: %s", t1.Sub(t0).Seconds())
			fmt.Println(threat)
			threat = []minimax.Threat{}
		}
	})
	http.Handle("/", http.FileServer(http.Dir("client/build")))
	log.Fatalln(http.ListenAndServe(":4444", nil))
}

