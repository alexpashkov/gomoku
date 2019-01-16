package minimax

import "gomoku/board"

func SearchThreatFourInRow(b board.Board, threat []Threat) []Threat{
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
				for i <= 4 {
					if x + 4 < 19 && x - 1 >= 0 {
						if ((((b.GetCell(board.Coords{x - 1, y}) == 0) && (b.GetCell(board.Coords{x + 4, y}) != 0)) ||
							((b.GetCell(board.Coords{x - 1, y}) != 0) && (b.GetCell(board.Coords{x + 4, y}) == 0))) &&
							(b.GetCell(board.Coords{x + i, y}) == b.GetCell(board.Coords{x, y}))) {
							amountX++
							positionsX = append(positionsX, board.Coords{x + i, y})
						}
						if (
							y - 1 >= 0 && y + 4 < 19 &&
							((b.GetCell(board.Coords{x - 1, y - 1}) == 0) && (b.GetCell(board.Coords{x + 4, y + 4}) != 0) ||
								((b.GetCell(board.Coords{x - 1, y - 1}) != 0) && (b.GetCell(board.Coords{x + 4, y + 4}) == 0))) &&
								(b.GetCell(board.Coords{x + i, y + i}) == b.GetCell(board.Coords{x, y}))) {
							amountRightZ++
							positionsRightZ = append(positionsRightZ, board.Coords{x + i, y + i})
						}
					}
					if y - 1 >= 0 && y + 4 < 19 {
						if (((b.GetCell(board.Coords{x, y - 1}) == 0 && b.GetCell(board.Coords{x, y + 4}) != 0) ||
							(b.GetCell(board.Coords{x, y - 1}) != 0 && b.GetCell(board.Coords{x, y + 4}) == 0)) &&
							(b.GetCell(board.Coords{x, y + i}) == b.GetCell(board.Coords{x, y}))) {
							amountY++
							positionsY = append(positionsY, board.Coords{x, y + i})
						}
						if (x + 1 < 19 && x - 4 >= 0 &&
							((b.GetCell(board.Coords{x + 1, y - 1}) == 0 && b.GetCell(board.Coords{x - 4, y + 4}) != 0) ||
								(b.GetCell(board.Coords{x + 1, y - 1}) != 0 && b.GetCell(board.Coords{x - 4, y + 4}) == 0)) &&
							b.GetCell(board.Coords{x - i, y + i}) == b.GetCell(board.Coords{x, y})) {
							amountLeftZ++
							positionsLeftZ = append(positionsLeftZ, board.Coords{x - i, y + i})
						}
					}
					i++
				}
				if amountX == 4 {
					threat = append(threat, Threat{owner:b.GetCell(board.Coords{x, y}), positions:positionsX, rate:4})
				}
				if amountY == 4 {
					threat = append(threat, Threat{owner:b.GetCell(board.Coords{x, y}), positions:positionsY, rate:5})
				}
				if amountRightZ == 4 {
					threat = append(threat, Threat{owner:b.GetCell(board.Coords{x, y}), positions:positionsRightZ, rate:5})
				}
				if amountLeftZ == 4 {
					threat = append(threat, Threat{owner:b.GetCell(board.Coords{x, y}), positions:positionsLeftZ, rate:5})
				}
			}
		}
	}
	return threat
}