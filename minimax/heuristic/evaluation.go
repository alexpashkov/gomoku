package heuristic

import (
	"gomoku/board"
	"gomoku/game"
)

const (
	FiveROW      = 10 * 10000000000 + 5 // # # # # #
	TwoRowCloseWIN = FiveROW - 3
	ForRowOpen   = 10 * 100000000 + 4 // # # # #
	ForRowClose  = 10 * 200000 + 4 // * # # # #
	ThreeRowOpen = 10 * 10000 + 3 // # # #
	ThreeRowClose= 10 * 10 + 3 // * # # #
	TwoRowClose  = 10 * 1000 + 2 // * # #
	TwoRowOpen   = 10 * 10 + 2 // # #
	TwoRowCloseSix = 10 * 1000 + 2 // # #
	)

func EvaluationRate(threat []Threat, amountPointMinPlayer int8, amountPointMaxPlayer int8) []Threat{
	for key, value := range threat {
		if value.size == 5 {
			threat[key].rate = FiveROW
		} else if value.size == 3 {
			if value.status == 1 {
				if value.owner != game.BLACK_PLAYER {
					threat[key].rate = ForRowOpen
				} else {
					threat[key].rate = -ForRowOpen
				}
			} else {
				if value.owner != game.BLACK_PLAYER {
					threat[key].rate = ForRowClose
				} else {
					threat[key].rate = -ForRowClose
				}
			}
		} else if value.size == 3 {
			if value.status == 1 {
				if value.owner != game.BLACK_PLAYER {
					threat[key].rate = ThreeRowOpen
				} else {
					threat[key].rate = -ThreeRowOpen
				}
			} else {
				if value.owner != game.BLACK_PLAYER {
					threat[key].rate = ThreeRowClose
				} else {
					threat[key].rate = -ThreeRowClose
				}
			}
		} else if value.size == 2 {
			if value.status == 1 {
				if value.owner != game.BLACK_PLAYER {
					threat[key].rate = TwoRowOpen
				} else {
					threat[key].rate = -TwoRowCloseSix
				}
			} else {
				if value.owner != game.BLACK_PLAYER {
					if amountPointMinPlayer == 8 {
						threat[key].rate = -TwoRowCloseWIN
					} else if amountPointMinPlayer == 6 {
						threat[key].rate = -TwoRowCloseSix
					} else {
						threat[key].rate = -TwoRowClose
					}
				} else if value.owner != game.WHITE_PLAYER {
					if amountPointMaxPlayer == 8 {
						threat[key].rate = TwoRowCloseWIN
					} else if amountPointMaxPlayer == 6 {
						threat[key].rate = TwoRowCloseSix
					} else {
						threat[key].rate = TwoRowClose
					}
				}
			}
		}
	}
	return threat
}

func Evaluation(brd board.Board, amountPointMinPlayer int8, amountPointMaxPlayer int8) int {
	power := 0

	i := 5
	threat := []Threat{}
	for i >= 2 {
		threat = SearchThreatRowClose(brd, threat, i)
		threat = SearchThreatRowOpen(brd, threat, i)
		i--
	}
	threat = EvaluationRate(threat, amountPointMinPlayer, amountPointMaxPlayer)
	for _, value := range threat {
		power = power + value.rate
	}
	return power
}
