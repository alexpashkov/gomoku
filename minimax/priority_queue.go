package minimax

import (
	"gomoku/board"
	"fmt"
	"container/heap"
)

type Move struct {
	board.Coords
	BoardAfterMove board.Board
	Priority       int
	index          int
}

type PriorityQueue struct {
	Type  int8
	Moves []Move
}

func NewPriorityQueue(t int8) (PriorityQueue, error) {
	pq := PriorityQueue{
		Type:  t,
		Moves: make([]Move, 0),
	}
	if t != MIN_PLAYER && t != MAX_PLAYER {
		return pq, fmt.Errorf("pq.Type must be one of %d, %d, got %d", MIN_PLAYER,
			MAX_PLAYER, pq.Type)
	}
	heap.Init(&pq)
	return pq, nil
}

func (pq PriorityQueue) Len() int {
	return len(pq.Moves)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.Type == MIN_PLAYER {
		return pq.Moves[i].Priority < pq.Moves[j].Priority
	}
	if pq.Type == MAX_PLAYER {
		return pq.Moves[i].Priority > pq.Moves[j].Priority
	}
	panic(fmt.Errorf("pq.Type must be one of %d, %d, got %d", MIN_PLAYER,
		MAX_PLAYER, pq.Type))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.Moves[i], pq.Moves[j] = pq.Moves[j], pq.Moves[i]
	pq.Moves[i].index = i
	pq.Moves[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.Moves)
	item := x.(Move)
	item.index = n
	pq.Moves = append(pq.Moves, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	oldMoves := pq.Moves
	n := len(oldMoves)
	item := oldMoves[n-1]
	item.index = -1 // for safety
	pq.Moves = oldMoves[0 : n-1]
	return item
}