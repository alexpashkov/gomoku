package main

import (
	"fmt"
	"gomoku/board"
)

func main() {
	maps := board.Board{}
	fmt.Println(maps.IsTerminate())
}
