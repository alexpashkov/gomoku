package main

import "fmt"

func setChip(maps [19][19]int, coordinates[]int, status[]int) [19][19] int {
	i := 0
	j := 0
	amountChip := len(status)
	for i < amountChip {
		maps[coordinates[j]][coordinates[j+1]] = status[i]
		i++
		j += 2
	}
	return maps
}

func main() {
	board := [19][19]int{}
	coordinates := []int{0,0,0,1,0,2,0,3}
	status := []int {3,3,3,3}
	board = setChip(board, coordinates, status)
	fmt.Println(board)
}