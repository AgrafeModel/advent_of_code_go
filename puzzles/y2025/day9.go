package y2025

import (
	"fmt"
	"image/color"
	"math"

	"github.com/AgrafeModel/advent_of_code/utils"
)

/////// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
////// That was so hard tbh

func readDay9Input() ([]utils.Position2D, utils.Position2D) {
	var data []utils.Position2D
	var max utils.Position2D
	utils.ReadFilePerLines(utils.GetInputPath(2025, 9), func(line string) {
		var p utils.Position2D
		fmt.Sscanf(line, "%d,%d", &p.X, &p.Y)
		data = append(data, p)
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	})
	return data, max
}

func Day9Part1() int {
	data, _ := readDay9Input()

	//Get the largest square area
	maxArea := 0

	for _, p := range data {
		for _, q := range data {
			if p == q {
				continue
			}
			area := utils.Abs(p.X-q.X+1) * utils.Abs(p.Y-q.Y+1)
			if area > maxArea {
				maxArea = area
			}

		}
	}

	return maxArea

}

func Day9Part2() int {
	data, _ := readDay9Input()
	polygon := utils.NewPolygon2DEdge(data)

	maxArea := 0
	var maxAreaPolygon utils.Polygon2DEdge

	// iterate over pairs
	for _, p := range data {
		for _, q := range data {
			if p == q {
				continue
			}

			area := utils.Abs(p.X-q.X+1) * utils.Abs(p.Y-q.Y+1)

			x1 := math.Min(float64(p.X), float64(q.X)) + 1
			x2 := math.Max(float64(p.X), float64(q.X)) - 1
			y1 := math.Min(float64(p.Y), float64(q.Y)) + 1
			y2 := math.Max(float64(p.Y), float64(q.Y)) - 1

			rect := utils.NewPolygon2DEdge([]utils.Position2D{
				{X: int(x1), Y: int(y1)},
				{X: int(x2), Y: int(y1)},
				{X: int(x2), Y: int(y2)},
				{X: int(x1), Y: int(y2)},
			})

			if !polygon.Intersects(rect) && area > maxArea {
				maxArea = area
				maxAreaPolygon = *rect
			}

		}
	}

	if utils.IsDebugMode() {

		//debug intesect

		fmt.Println(polygon.Intersects(&maxAreaPolygon))

		dmap := make(map[utils.Position2D]struct{})
		for _, p := range data {
			dmap[p] = struct{}{}
		}

		utils.ShowPolygonEBit([]utils.PolygonVisualizer{
			{Polygon: polygon, Color: color.RGBA{0, 255, 0, 255}},
			{Polygon: &maxAreaPolygon, Color: color.RGBA{255, 0, 0, 255}},
		}, dmap)

	}
	return maxArea
}

///////////
// DEBUG TOOLS

func printDay9Grid(red *map[utils.Position2D]struct{}, max utils.Position2D, maxArea utils.Polygon2DEdge) {
	for y := max.Y; y >= 0; y-- {
		for x := 0; x <= max.X; x++ {
			p := utils.Position2D{X: x, Y: y}
			_, isRed := (*red)[p]
			if isRed {
				fmt.Print("\033[31m#\033[0m")
			} else if maxArea.IsOnEdge(p) {
				fmt.Print("\033[32m*\033[0m")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

///// PISTES

func drawLine(start, end utils.Position2D, pointSet *map[utils.Position2D]struct{}) {
	yy := utils.Min(start.Y, end.Y)
	xx := utils.Min(start.X, end.X)
	xy := utils.Max(start.Y, end.Y)
	yx := utils.Max(start.X, end.X)

	for y := yy; y <= xy; y++ {
		for x := xx; x <= yx; x++ {
			(*pointSet)[utils.Position2D{X: x, Y: y}] = struct{}{}
		}
	}

}

func floodFill(start utils.Position2D, greenSet *map[utils.Position2D]struct{}, max utils.Position2D) *map[utils.Position2D]struct{} {
	directions := [4]utils.Position2D{
		utils.UP,
		utils.DOWN,
		utils.LEFT,
		utils.RIGHT,
	}

	visited := make(map[utils.Position2D]struct{})
	queue := []utils.Position2D{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if _, exists := visited[current]; exists {
			continue
		}

		if _, exists := (*greenSet)[current]; exists {
			continue
		}

		visited[current] = struct{}{}

		for _, dir := range directions {
			neighbor := utils.Position2D{X: current.X + dir.X, Y: current.Y + dir.Y}
			if neighbor.X < 0 || neighbor.Y < 0 || neighbor.X > max.X || neighbor.Y > max.Y {
				continue
			}
			if _, exists := visited[neighbor]; exists {
				continue
			}
			queue = append(queue, neighbor)
		}
	}

	return &visited
}
