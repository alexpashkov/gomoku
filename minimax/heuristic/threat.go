package minimax

import "gomoku/board"

type Threat struct {
	positions	[]board.Coords
	owner		int8
	size 		int8
	status		int8
}
