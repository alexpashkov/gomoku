package board

type Captures struct {
	positions [2]Coords
	enemy     int8
}

func (c Captures) GetEnemy() int8{
	return c.enemy
}

func (c Captures) GetPositions() [2]Coords {
	return c.positions
}
