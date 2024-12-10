package utils

type Position struct {
	X, Y int
}

var Directions = []Position{
	UP, DOWN, LEFT, RIGHT, UP_LEFT, DOWN_LEFT, UP_RIGHT, DOWN_RIGHT,
}

var LinesDirections = []Position{
	UP, DOWN, LEFT, RIGHT,
}

var DiagonalsDirections = []Position{
	UP_LEFT, DOWN_LEFT, UP_RIGHT, DOWN_RIGHT,
}

var (
	UP         = Position{X: 0, Y: 1}
	DOWN       = Position{X: 0, Y: -1}
	LEFT       = Position{X: -1, Y: 0}
	RIGHT      = Position{X: 1, Y: 0}
	UP_LEFT    = Position{X: -1, Y: 1}
	UP_RIGHT   = Position{X: 1, Y: 1}
	DOWN_LEFT  = Position{X: -1, Y: -1}
	DOWN_RIGHT = Position{X: 1, Y: -1}
)

func (p Position) ReverseDirection() Position {
	switch p {
	case UP:
		return DOWN
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	case UP_LEFT:
		return DOWN_RIGHT
	case UP_RIGHT:
		return DOWN_LEFT
	case DOWN_LEFT:
		return UP_RIGHT
	case DOWN_RIGHT:
		return UP_LEFT
	}
	return p
}

func (p Position) Negate() Position {
	return Position{
		X: p.X * -1,
		Y: p.Y * -1,
	}
}

func (p Position) RotateRight90() Position {
	switch p {
	case UP:
		return LEFT
	case DOWN:
		return RIGHT
	case LEFT:
		return DOWN
	case RIGHT:
		return UP
	case UP_LEFT:
		return UP_RIGHT
	case UP_RIGHT:
		return DOWN_LEFT
	case DOWN_LEFT:
		return DOWN_RIGHT
	case DOWN_RIGHT:
		return UP_LEFT
	}
	return p
}

func (p Position) RotateLeft90() Position {
	switch p {
	case UP:
		return RIGHT
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	case UP_LEFT:
		return DOWN_LEFT
	case UP_RIGHT:
		return DOWN_RIGHT
	case DOWN_LEFT:
		return UP_RIGHT
	case DOWN_RIGHT:
		return UP_LEFT
	}
	return p
}

func (p Position) Mul(value int) Position {
	return Position{
		X: p.X * value,
		Y: p.Y * value,
	}
}

func (p Position) Add(value int) Position {
	return Position{
		X: p.X + value,
		Y: p.Y + value,
	}
}

func (p Position) Sub(value int) Position {
	return Position{
		X: p.X - value,
		Y: p.Y - value,
	}
}

func (p Position) SubPos(pos Position) Position {
	return Position{
		X: p.X - pos.X,
		Y: p.Y - pos.Y,
	}
}

func (p Position) Abs() Position {
	return Position{
		X: p.X * -1,
		Y: p.Y * -1,
	}
}

func (p Position) AddPos(pos Position) Position {
	return Position{
		X: p.X + pos.X,
		Y: p.Y + pos.Y,
	}
}

func (p Position) RemovePos(pos Position) Position {
	return Position{
		X: p.X - pos.X,
		Y: p.Y - pos.Y,
	}
}

func (p Position) Dist(p2 Position) Position {
	return p.SubPos(p2).Abs()
}
