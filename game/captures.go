package game

import "gomoku/board"

func GetCaptures(b board.Board, coords board.Coords) []board.Coords {
	captures := []board.Coords(nil)
	player := b.GetCell(coords)
	x := coords.X
	y := coords.Y
	if (b.GetCell(board.Coords{x, y}) != 0) {
		if (x+3 < 19 && b.GetCell(board.Coords{x, y}) == b.GetCell(board.Coords{x + 3, y})) {
			if (b.GetCell(board.Coords{x + 1, y}) != 0 &&
				b.GetCell(board.Coords{x + 1, y}) == b.GetCell(board.Coords{x + 2, y})) {
				if b.GetCell(board.Coords{x, y}) == player {
					captures = append(captures, board.Coords{x + 1, y}, board.Coords{x + 2, y})
				}
			}
		}
	}
	if (y+3 < 19 && b.GetCell(board.Coords{x, y}) == b.GetCell(board.Coords{x, y + 3})) {
		if (b.GetCell(board.Coords{x, y + 1}) != 0 &&
			b.GetCell(board.Coords{x, y + 1}) == b.GetCell(board.Coords{x, y + 2})) {
			if b.GetCell(board.Coords{x, y}) == player {
				captures = append(captures, board.Coords{x, y + 1}, board.Coords{x, y + 2})
			}
		}
	}
	if (y+3 < 19 && x+3 < 19 && b.GetCell(board.Coords{x, y}) == b.GetCell(board.Coords{x + 3, y + 3})) {
		if (b.GetCell(board.Coords{x + 1, y + 1}) != 0 &&
			b.GetCell(board.Coords{x + 1, y + 1}) == b.GetCell(board.Coords{x + 2, y + 2})) {
			if b.GetCell(board.Coords{x, y}) == player {
				captures = append(captures, board.Coords{x + 1, y + 1}, board.Coords{x + 2, y + 2})
			}
		}
	}
	if (y+3 < 19 && x-3 > 1 && b.GetCell(board.Coords{x, y}) == b.GetCell(board.Coords{x - 3, y + 3})) {
		if (b.GetCell(board.Coords{x - 1, y + 1}) != 0 &&
			b.GetCell(board.Coords{x - 1, y + 1}) == b.GetCell(board.Coords{x - 2, y + 2})) {
			if b.GetCell(board.Coords{x, y}) == player {
				captures = append(captures, board.Coords{x - 1, y + 1}, board.Coords{x - 2, y + 2})
			}
		}
	}
	return captures
}
