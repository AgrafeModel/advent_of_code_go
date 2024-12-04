package utils

type Position struct {
	X, Y int
}

var Directions = []Position{
	UP, DOWN, LEFT, RIGHT, UP_LEFT, DOWN_LEFT, UP_RIGHT, DOWN_RIGHT,
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

func (p Position) AddPos(pos Position) Position {
	return Position{
		X: p.X + pos.X,
		Y: p.Y + pos.Y,
	}
}
