package priority_queue

import (
	"container/heap"
	"gomoku/minimax"
	"fmt"
)

func New(t int8, moves ...Move) (PQ, error) {
	pq := PQ{
		Type:  t,
		Moves: moves,
	}
	if t != minimax.MIN_PLAYER && t != minimax.MAX_PLAYER {
		return pq, fmt.Errorf("pq.Type must be one of %d, %d, got %d", minimax.MIN_PLAYER,
			minimax.MAX_PLAYER, pq.Type)
	}
	heap.Init(&pq)
	return pq, nil
}
