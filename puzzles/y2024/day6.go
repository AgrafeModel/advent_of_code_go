package y2024

import (
	"fmt"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type guardGrid struct {
	grid [][]rune

	guard          utils.Position2D
	guardDirection utils.Position2D

	walkedpos []posDir

	rightLookPos []utils.Position2D
}

type posDir struct {
	pos utils.Position2D
	dir utils.Position2D
}

const (
	void     = '.'
	obstacle = '#'
	guard    = '^'
)

func Day6Part1() int {

	y := 0

	var grid guardGrid

	utils.ReadFilePerLines(utils.GetInputPath(2024, (6)), func(line string) {
		var ln []rune

		for x, c := range line {
			if c == guard {
				grid.guard = utils.Position2D{X: x, Y: y}
				grid.guardDirection = utils.DOWN
			}

			ln = append(ln, c)
		}

		grid.grid = append(grid.grid, ln)

		y++
	})

	for {

		pos := grid.guard
		next := pos.AddPos(grid.guardDirection)

		if next.X < 0 || next.Y < 0 || next.X >= len(grid.grid) || next.Y >= len(grid.grid[0]) {

			grid.walkedpos = append(grid.walkedpos, posDir{pos: pos, dir: grid.guardDirection})

			break
		}

		if grid.grid[next.Y][next.X] == '#' {
			grid.guardDirection = grid.guardDirection.RotateRight90()
			continue
		} else {
			grid.guard = next
			if !utils.Contains(grid.walkedpos, posDir{pos: pos, dir: grid.guardDirection}) {
				grid.walkedpos = append(grid.walkedpos, posDir{pos: pos, dir: grid.guardDirection})

			}
			continue
		}
	}

	return len(grid.walkedpos)
}

func checkGuardLoop(grid guardGrid, pos, before posDir) bool {
	if grid.grid[pos.pos.Y][pos.pos.X] == obstacle {
		return false
	}

	grid.grid[pos.pos.Y][pos.pos.X] = obstacle

	var seen []posDir

	p := before

	for {
		if utils.Contains(seen, p) { // if already seen
			// fmt.Println("already seen m>")
			grid.grid[pos.pos.Y][pos.pos.X] = void
			return true
		}

		seen = append(seen, p)
		// fmt.Println(seen)
		next := posDir{
			dir: p.dir,
			pos: p.pos.AddPos(p.dir),
		}

		//if out of bounds
		if next.pos.X < 0 || next.pos.Y < 0 || next.pos.X >= len(grid.grid) || next.pos.Y >= len(grid.grid[0]) {
			grid.grid[pos.pos.Y][pos.pos.X] = void
			return false
		}

		if grid.grid[next.pos.Y][next.pos.X] == obstacle {
			r := p.dir.RotateRight90()
			p = posDir{
				dir: r,
				pos: p.pos,
			}
		} else {
			p = next
		}
	}

}

func Day6Part2() int {

	y := 0

	var grid guardGrid

	res := 0

	utils.ReadFilePerLines(utils.GetInputPath(2024, (6)), func(line string) {
		var ln []rune

		for x, c := range line {
			if c == guard {
				grid.guard = utils.Position2D{X: x, Y: y}
				grid.guardDirection = utils.DOWN
			}

			ln = append(ln, c)
		}

		grid.grid = append(grid.grid, ln)

		y++
	})

	startPos := grid.guard

	for {

		pos := grid.guard
		next := pos.AddPos(grid.guardDirection)

		if next.X < 0 || next.Y < 0 || next.X >= len(grid.grid) || next.Y >= len(grid.grid[0]) {

			grid.walkedpos = append(grid.walkedpos, posDir{pos: pos, dir: grid.guardDirection})

			break
		}

		if grid.grid[next.Y][next.X] == '#' {
			grid.guardDirection = grid.guardDirection.RotateRight90()
			continue
		} else {
			grid.guard = next
			if !utils.Contains(grid.walkedpos, posDir{pos: pos, dir: grid.guardDirection}) {
				grid.walkedpos = append(grid.walkedpos, posDir{pos: pos, dir: grid.guardDirection})

			}
			continue
		}
	}

	// for each pos that we walked
	for i, p := range grid.walkedpos {
		if p.pos == startPos {
			continue
		}
		if checkGuardLoop(grid, p, grid.walkedpos[i-1]) {
			res++
		}
	}

	// debugGrid(grid)

	return res
}

//// DEBUG ////

func debugGrid(grid guardGrid) {
	//For debug, show the grid with P as the guard path
	for y, ln := range grid.grid {
		for x, c := range ln {
			if c == obstacle {
				//in red
				fmt.Printf("\033[31m %c \033[0m", c)

			} else if ok, dir := ContainsWalkedPos(grid.walkedpos, utils.Position2D{X: x, Y: y}); ok {
				switch dir {
				case utils.DOWN:
					fmt.Printf("%c", '^')
				case utils.UP:
					fmt.Printf("%c", 'v')
				case utils.LEFT:
					fmt.Printf("%c", '>')
				case utils.RIGHT:
					fmt.Printf("%c", '<')
				}
			} else {
				fmt.Printf("%c", c)
			}
		}

		println()
	}
}

func ContainsWalkedPos(walkedPos []posDir, pos utils.Position2D) (bool, utils.Position2D) {
	for _, p := range walkedPos {
		if p.pos == pos {
			return true, p.dir
		}
	}
	return false, utils.Position2D{}
}
