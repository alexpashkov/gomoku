package main

import (
	"github.com/rs/cors"
	"gomoku/handlers"
	"log"
	"net/http"
)

const PORT = "0.0.0.0:5555"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/board", handlers.Board)
	mux.HandleFunc("/make-move", handlers.MakeMove)
	mux.HandleFunc("/suggest-move", handlers.SuggestMove)

	println("Starting server at " + PORT)
	log.Fatalln(http.ListenAndServe(PORT, cors.Default().Handler(mux)))
}
