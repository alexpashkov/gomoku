package board

import "fmt"

func SetChip(maps [19][19]int, coordinates []int, status []int) [19][19]int {
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

func IsTerminate(board [19][19]int) int {
	y := 0
	x := 0
	i := 0
	for y < 19 {
		x = 0
		for x < 19 {
			i = 1
			amountX := 0
			amountY := 0
			amountRightZ := 0
			amountLeftZ := 0
			if board[y][x] != 0 {
				fmt.Println("y = ", y, " x = ", x)
				for i <= 4 {
					if x+i < 19 && board[y][x+i] == board[y][x] {
						amountX++
					}
					if y+i < 19 && board[y+i][x] == board[y][x] {
						amountY++
					}
					if x+i < 19 && y+i < 19 && board[y+i][x+i] == board[y][x] {
						amountRightZ++
					}
					if x-i > 1 && y+i < 19 && board[y+i][x-i] == board[y][x] {
						amountLeftZ++
					}
					i++
				}
				//fmt.Println(amountLeftZ, " y = ", y, " x = ", x)
				if amountX == 4 || amountY == 4 || amountRightZ == 4 || amountLeftZ == 4 {
					return board[y][x]
				}
			}
			x++
		}
		y++
	}
	return 0
}

//func main() {
//	board := [19][19]int{}
//	//coordinates := []int{1,0,2,0,3,0,4,0,5,0} // y in started
//	//coordinates := []int{1,18,2,18,3,18,4,18,5,18} // y in end
//	//coordinates := []int{0,1,0,2,0,3,0,4,0,5} // x in started
//	//coordinates := []int{5,14,5,15,5,16,5,17,5,18} // x in started
//	//coordinates := []int{0,1,1,2,2,3,3,4,4,5} // z right in started
//	//coordinates := []int{5,7, 6,6, 7,5, 8,4, 9,3} // z left in started
//	status := []int {2,2,2,2,2}
//	board = setChip(board, coordinates, status)
//
//	i := 0
//	for (i < 19) {
//		fmt.Println(board[i])
//		i++
//	}
//	fmt.Println(isTerminate(board))
//	//fmt.Println(board)
//}

// func main() {
// 	board := [19][19]int{}
// 	coordinates := []int{0,0,0,1,0,2,0,3}
// 	status := []int {3,3,3,3}
// 	board = setChip(board, coordinates, status)
// 	fmt.Println(board)
// }
