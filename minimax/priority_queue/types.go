package priority_queue

import (
	"gomoku/board"
	"gomoku/minimax"
	"fmt"
)

type Move struct {
	board.Coords
	Priority int
	index    int
}

type PQ struct {
	Type  int8
	Moves []Move
}

func (pq PQ) Len() int {
	return len(pq.Moves)
}

func (pq PQ) Less(i, j int) bool {
	if pq.Type == minimax.MIN_PLAYER {
		return pq.Moves[i].Priority < pq.Moves[j].Priority
	}
	if pq.Type == minimax.MAX_PLAYER {
		return pq.Moves[i].Priority > pq.Moves[j].Priority
	}
	panic(fmt.Errorf("pq.Type must be one of %d, %d, got %d", minimax.MIN_PLAYER,
		minimax.MAX_PLAYER, pq.Type))
}

func (pq PQ) Swap(i, j int) {
	pq.Moves[i], pq.Moves[j] = pq.Moves[j], pq.Moves[i]
	pq.Moves[i].index = i
	pq.Moves[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(pq.Moves)
	item := x.(Move)
	item.index = n
	pq.Moves = append(pq.Moves, item)
}

func (pq *PQ) Pop() interface{} {
	oldMoves := pq.Moves
	n := len(oldMoves)
	item := oldMoves[n-1]
	item.index = -1 // for safety
	pq.Moves = oldMoves[0 : n-1]
	return item
}
