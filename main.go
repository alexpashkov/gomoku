package main

import (
	"encoding/json"
	"fmt"
	"gomoku/board"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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
		}
	})
	http.Handle("/", http.FileServer(http.Dir("client/build")))
	log.Fatalln(http.ListenAndServe(":4444", nil))
}
