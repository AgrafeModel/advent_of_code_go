package utils

import (
	"math"
)

type Position2D struct {
	X, Y int
}

var Directions = []Position2D{
	UP, DOWN, LEFT, RIGHT, UP_LEFT, DOWN_LEFT, UP_RIGHT, DOWN_RIGHT,
}

var LinesDirections = []Position2D{
	UP, DOWN, LEFT, RIGHT,
}

var DiagonalsDirections = []Position2D{
	UP_LEFT, DOWN_LEFT, UP_RIGHT, DOWN_RIGHT,
}

var (
	UP         = Position2D{X: 0, Y: 1}
	DOWN       = Position2D{X: 0, Y: -1}
	LEFT       = Position2D{X: -1, Y: 0}
	RIGHT      = Position2D{X: 1, Y: 0}
	UP_LEFT    = Position2D{X: -1, Y: 1}
	UP_RIGHT   = Position2D{X: 1, Y: 1}
	DOWN_LEFT  = Position2D{X: -1, Y: -1}
	DOWN_RIGHT = Position2D{X: 1, Y: -1}
)

func (p Position2D) ReverseDirection() Position2D {
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

func (p Position2D) Negate() Position2D {
	return Position2D{
		X: p.X * -1,
		Y: p.Y * -1,
	}
}

func (p Position2D) RotateRight90() Position2D {
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

func (p Position2D) RotateLeft90() Position2D {
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

func (p Position2D) Mul(value int) Position2D {
	return Position2D{
		X: p.X * value,
		Y: p.Y * value,
	}
}

func (p Position2D) Add(value int) Position2D {
	return Position2D{
		X: p.X + value,
		Y: p.Y + value,
	}
}

func (p Position2D) Sub(value int) Position2D {
	return Position2D{
		X: p.X - value,
		Y: p.Y - value,
	}
}

func (p Position2D) SubPos(pos Position2D) Position2D {
	return Position2D{
		X: p.X - pos.X,
		Y: p.Y - pos.Y,
	}
}

func (p Position2D) Abs() Position2D {
	return Position2D{
		X: p.X * -1,
		Y: p.Y * -1,
	}
}

func (p Position2D) AddPos(pos Position2D) Position2D {
	return Position2D{
		X: p.X + pos.X,
		Y: p.Y + pos.Y,
	}
}

func (p Position2D) RemovePos(pos Position2D) Position2D {
	return Position2D{
		X: p.X - pos.X,
		Y: p.Y - pos.Y,
	}
}

func (p Position2D) GetArea(other Position2D) int {
	return Abs(p.X-other.X) * Abs(p.Y-other.Y)
}

func (p Position2D) Dist(p2 Position2D) Position2D {
	return p.SubPos(p2).Abs()
}

func Opposite(dir Position2D) Position2D {
	return dir.Negate()
}

type Position3D struct {
	X, Y, Z int
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Distance3D(p1, p2 Position3D) float64 {
	return math.Sqrt(float64(Abs(p1.X-p2.X)*Abs(p1.X-p2.X) +
		Abs(p1.Y-p2.Y)*Abs(p1.Y-p2.Y) +
		Abs(p1.Z-p2.Z)*Abs(p1.Z-p2.Z)))
}

func LinesIntersect(p1, p2, p3, p4 *Position2D) bool {
	ccw := func(a, b, c *Position2D) int {
		val := (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
		if val > 0 {
			return 1
		}
		if val < 0 {
			return -1
		}
		return 0
	}

	d1 := ccw(p3, p4, p1)
	d2 := ccw(p3, p4, p2)
	d3 := ccw(p1, p2, p3)
	d4 := ccw(p1, p2, p4)

	return ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0))
}

type Polygon2DEdge struct {
	Vertices []Position2D
	edges    []Edge
}

func (p *Polygon2DEdge) IsOnEdge(point Position2D) bool {
	for _, edge := range p.edges {
		if edge.start == point || edge.end == point {
			return true
		}
		if edge.horizontal {
			if point.Y == edge.start.Y && point.X >= min(edge.start.X, edge.end.X) && point.X <= max(edge.start.X, edge.end.X) {
				return true
			}
		} else {
			if point.X == edge.start.X && point.Y >= min(edge.start.Y, edge.end.Y) && point.Y <= max(edge.start.Y, edge.end.Y) {
				return true
			}
		}
	}
	return false
}

func (p *Polygon2DEdge) buildEdges() {
	n := len(p.Vertices)
	for i, v := range p.Vertices {
		next := p.Vertices[(i+1)%n]
		edge := NewEdge(v, next)

		p.edges = append(p.edges, edge)
	}
}

func NewPolygon2DEdge(vertices []Position2D) *Polygon2DEdge {
	p := &Polygon2DEdge{
		Vertices: vertices,
		edges:    make([]Edge, 0),
	}
	p.buildEdges()
	return p
}

func (p *Polygon2DEdge) Intersects(other *Polygon2DEdge) bool {
	for _, e1 := range p.edges {
		for _, e2 := range other.edges {
			if e1.Intersect(&e2) {
				return true
			}
		}
	}
	return false
}

type Edge struct {
	start, end Position2D
	horizontal bool
}

func NewEdge(start, end Position2D) Edge {
	horizontal := false
	if start.Y == end.Y {
		horizontal = true
	}

	if horizontal && start.X > end.X {
		start, end = end, start
	}
	if !horizontal && start.Y > end.Y {
		start, end = end, start
	}

	return Edge{
		start:      start,
		end:        end,
		horizontal: horizontal,
	}
}

func (e *Edge) Intersect(other *Edge) bool {
	if e.horizontal == other.horizontal {
		// parallel lines
		return false
	}

	var hEdge, vEdge *Edge
	if e.horizontal {
		hEdge = e
		vEdge = other
	} else {
		hEdge = other
		vEdge = e
	}

	if vEdge.start.X >= hEdge.start.X && vEdge.start.X <= hEdge.end.X &&
		hEdge.start.Y >= vEdge.start.Y && hEdge.start.Y <= vEdge.end.Y {
		return true
	}
	return false
}

func (p *Polygon2DEdge) Contains(point Position2D) bool {
	n := len(p.Vertices)
	count := 0

	for i := range n {
		v1 := p.Vertices[i]
		v2 := p.Vertices[(i+1)%n]

		if (v1.Y > point.Y) != (v2.Y > point.Y) &&
			(point.X < (v2.X-v1.X)*(point.Y-v1.Y)/(v2.Y-v1.Y)+v1.X) {
			count++
		}

	}
	return count%2 == 1
}
