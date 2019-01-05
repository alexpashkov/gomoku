package main

import (
	"fmt"
	"gomoku/board"
	"gomoku/minimax"
)

func main() {
	board := board.Board{}
	fmt.Printf("%v", minimax.PossibleMoves(board, 0))
}
