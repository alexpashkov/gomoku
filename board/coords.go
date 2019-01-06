package board

type Coords struct {
	X, Y int
}

func (c Coords) AreWithin(b Board) bool {
	return c.X >= 0 && c.X < len(b) && c.Y >= 0 && c.Y < len(b)
}
