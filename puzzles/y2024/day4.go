package y2024

import (
	"github.com/AgrafeModel/advent_of_code/utils"
)

type xmasGrid struct {
	grid []string
}

const XMAS = "XMAS"

func isXmasInDirection(grid xmasGrid, pos, dir utils.Position) bool {
	ptr := 0
	for ptr < len(XMAS) {
		dr := dir.
			Mul(ptr).
			AddPos(pos)

		if dr.Y < 0 ||
			dr.Y > len(grid.grid)-1 ||
			dr.X > len(grid.grid[dr.Y])-1 ||
			dr.X < 0 {
			return false
		}
		if grid.grid[dr.Y][dr.X] != XMAS[ptr] {
			return false
		}
		ptr++
	}
	return true
}

func isX_mas(grid xmasGrid, pos utils.Position) bool {
	if grid.grid[pos.Y][pos.X] != 'A' {
		return false
	}

	mascount := 0
	for _, dir := range utils.DiagonalsDirections {
		p := pos.AddPos(dir)

		if p.Y < 0 || // check if in bounds
			p.Y > len(grid.grid)-1 ||
			p.X > len(grid.grid[p.Y])-1 ||
			p.X < 0 {
			continue
		}

		if grid.grid[p.Y][p.X] == 'M' {
			b := pos.AddPos(dir.ReverseDirection())

			// Check if in bounds
			if b.Y < 0 ||
				b.Y > len(grid.grid)-1 ||
				b.X > len(grid.grid[b.Y])-1 ||
				b.X < 0 {
				continue
			}

			if grid.grid[b.Y][b.X] == 'S' {
				mascount++
			}
		}
	}

	if mascount != 2 { // we need exactly 2 MAS
		return false
	}

	return true
}

func Day4Part1() int {
	var grid xmasGrid

	var xpos []utils.Position

	res := 0
	y := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, 4), func(line string) {
		for x, v := range line {
			if v == 'X' {
				xpos = append(xpos, utils.Position{X: x, Y: y})
			}
		}
		grid.grid = append(grid.grid, line)
		y++
	})
	for _, pos := range xpos {
		for _, dir := range utils.Directions {
			if isXmasInDirection(grid, pos, dir) {
				res++
			}
		}
	}

	return res
}

func Day4Part2() int {
	var grid xmasGrid

	var apos []utils.Position

	res := 0
	y := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, 4), func(line string) {
		for x, v := range line {
			if v == 'A' {
				apos = append(apos, utils.Position{X: x, Y: y})
			}
		}
		grid.grid = append(grid.grid, line)
		y++
	})
	for _, pos := range apos {
		if isX_mas(grid, pos) {
			res++
		}
	}
	return res
}
