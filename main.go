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
			//fmt.Println(brd)
			fmt.Println(heuristic.Evaluation(brd, 6, 8))
			//fmt.Println(heuristic.IsTerminate(brd, 0, 10))
			//fmt.Println(brd)
			//t0 := time.Now()
			//i := 4
			//for i >= 1 {
			//	fmt.Println(heuristic.SearchThreatRowClose(brd, threat, 2))
			//	threat = minimax.SearchThreatRowClose(brd, threat, i)
			//	fmt.Println(threat)
			//	i--
			//}
			//threat = minimax.SearchThreatRowOpen(brd, threat, 4)
			//threat = minimax.SearchThreatRowOpen(brd, threat, 2)
			//t1 := time.Now()
			//fmt.Printf("Elapsed time: %s", t1.Sub(t0).Seconds())
			//fmt.Println(threat)
			//threat = []minimax.Threat{}
		}
	})
	mux.HandleFunc("/suggest-move", suggest_move.Handler)
	mux.Handle("/", http.FileServer(http.Dir("client/build")))
	log.Fatalln(http.ListenAndServe(":4444", cors.Default().Handler(mux)))
}
