package heuristic

import (
	"gomoku/board"
)

func SearchDoubleThreat(b board.Board, threat []Threat, len int) []Threat{
	for y, row := range b {
		for x := range row {
			if (b.GetCell(board.Coords{x, y}) != 0) {
				i := 0
				positionsX := []board.Coords{}
				positionsY := []board.Coords{}
				positionsRightZ := []board.Coords{}
				positionsLeftZ := []board.Coords{}
				amountX := 0
				amountY := 0
				amountRightZ := 0
				amountLeftZ := 0
				current := b.GetCell(board.Coords{x, y})
				for i <= len {
					if x + len < 19 && x - 1 >= 0 {
						if (y - 1 >= 0 && y + len < 19 &&
							(b.GetCell(board.Coords{x - 1, y - 1}) == 0) &&
							(b.GetCell(board.Coords{x + len, y + len}) == 0) &&
							(b.GetCell(board.Coords{x + i, y + i}) == current)) {
							amountRightZ++
							positionsRightZ = append(positionsRightZ, board.Coords{x + i, y + i})
						}
						if ((b.GetCell(board.Coords{x - 1, y}) == 0) &&
							(b.GetCell(board.Coords{x + len, y}) == 0) &&
							(b.GetCell(board.Coords{x + i, y}) == current)) {
							amountX++
							positionsX = append(positionsX, board.Coords{x + i, y})
						}
					}
					if y - 1 >= 0 && y + len < 19 {
						if ((b.GetCell(board.Coords{x, y - 1}) == 0) &&
							(b.GetCell(board.Coords{x, y + len}) == 0) &&
							(b.GetCell(board.Coords{x, y + i}) == current)) {
							amountY++
							positionsY = append(positionsY, board.Coords{x, y + i})
						}
						if (x + 1 < 19 && x - len >= 0 &&
							(b.GetCell(board.Coords{x + 1, y - 1}) == 0) &&
							(b.GetCell(board.Coords{x - len, y + len}) == 0) &&
							(b.GetCell(board.Coords{x - i, y + i}) == current)) {
							amountLeftZ++
							positionsLeftZ = append(positionsLeftZ, board.Coords{x - i, y + i})
						}
					}
					i++
				}
				if amountX == len {
					corner := []board.Coords{{positionsX[0].X - 1, positionsX[0].Y}, {positionsX[2].X + 1, positionsX[2].Y}}
 					threat = append(threat, Threat{owner:current, positions:positionsX, size:int8(len), status:1, corner:corner})
				}
				if amountY == len {
					corner := []board.Coords{{positionsY[0].X, positionsY[0].Y - 1}, {positionsY[2].X, positionsY[2].Y + 1}}
					threat = append(threat, Threat{owner:current, positions:positionsY, size:int8(len), status:1, corner:corner})
				}
				if amountRightZ == len {
					corner := []board.Coords{{positionsRightZ[0].X - 1, positionsRightZ[0].Y - 1}, {positionsRightZ[2].X + 1, positionsRightZ[2].Y + 1}}
					threat = append(threat, Threat{owner:current, positions:positionsRightZ, size:int8(len), status:1, corner:corner})
				}
				if amountLeftZ == len {
					corner := []board.Coords{{positionsLeftZ[0].X + 1, positionsLeftZ[0].Y - 1}, {positionsLeftZ[2].X - 1, positionsLeftZ[2].Y + 1}}
					threat = append(threat, Threat{owner:current, positions:positionsLeftZ, size:int8(len), status:1, corner:corner})
				}
			}
		}
	}
	return threat
}

func FindDoubleThreeThreat(b board.Board, player int8, cord board.Coords) bool {
	b.SetCell(cord, player)
	var threat []Threat
	threat = SearchDoubleThreat(b, threat, 3)
	amountPlayer := 0
	for _, value := range threat {
		if value.owner == player {
			amountPlayer++
		}
	}
	if amountPlayer == 2 {
		i := 0
		for i < 2 {
			j := 0
			for j < 2 {
				if threat[0].corner[j].X == threat[1].corner[i].X && threat[0].corner[j].Y == threat[1].corner[i].Y {
					return true
				}
				j++
			}
			i++
		}
		return false
	}
	return true
}

func IsCorrectMove(b board.Board, player int8, cord board.Coords) bool {
	// base
	current := b.GetCell(cord)
	if current != 0 {
		return false
	}
	// capture
	b.SetCell(cord, player)
	s := b.GetCaptures()
	if s.GetEnemy() != 0 && s.GetEnemy() != player {
		return true
	}
	// double
	if FindDoubleThreeThreat(b, player, cord) == false {
		return false
	}
	return true
}
