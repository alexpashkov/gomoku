package main

import (
	"gomoku/board"
	"gomoku/minimax"
)

func main() {
	x := 0
	for n := 0; n < 19*10; n++ {
		for i := 0; i < 10; i++ {
			b := board.Board{}
			b.FillRandomly(n, 1)
			posMovesLen := len(minimax.PossibleMoves(b))
			println(posMovesLen)
			x += posMovesLen
		}
	}
	println(x / (10 * 10 * 10))
}
