package main

import (
	"board"
	"fmt"
)

func main() {
	maps := board.Board{}
	fmt.Println(maps.IsTerminate())
}
