package y2024

import (
	"fmt"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type lavamap struct {
	grid     [][]int
	startpos []utils.Position
}

// return the number of paths possibles to reach 9 with step = 1
func (lv *lavamap) LookForTrail(from utils.Position, elevation int, walked *map[utils.Position]bool) int {
	if lv.grid[from.Y][from.X] != elevation || (walked != nil && (*walked)[from]) {
		return 0
	}
	if lv.grid[from.Y][from.X] == 9 {
		if walked != nil {
			(*walked)[from] = true
		}
		return 1
	}

	res := 0
	for _, dir := range utils.LinesDirections {
		to := from.AddPos(dir)
		//if out of bounds
		if to.X < 0 || to.Y < 0 || to.Y >= len(lv.grid) || to.X >= len(lv.grid[0]) {
			continue
		}
		if lv.grid[to.Y][to.X] == elevation+1 {
			res += lv.LookForTrail(to, elevation+1, walked)
		}
	}

	return res
}

func Day10Part1() int {
	res := 0
	var g lavamap
	y := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, (10)), func(line string) {
		var ln []int
		for xid, x := range line {
			v := utils.ParseInt(x)
			if v == 0 {
				g.startpos = append(g.startpos, utils.Position{
					X: xid,
					Y: y,
				})
			}
			ln = append(ln, v)
		}
		g.grid = append(g.grid, ln)
		y++
	})

	for _, p := range g.startpos {
		fmt.Println(p)
		mp := make(map[utils.Position]bool)
		res += g.LookForTrail(p, 0, &mp)
	}

	return res
}

func Day10Part2() int {
	res := 0
	var g lavamap
	y := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, (10)), func(line string) {
		var ln []int
		for xid, x := range line {
			v := utils.ParseInt(x)
			if v == 0 {
				g.startpos = append(g.startpos, utils.Position{
					X: xid,
					Y: y,
				})
			}
			ln = append(ln, v)
		}
		g.grid = append(g.grid, ln)
		y++
	})

	for _, p := range g.startpos {
		fmt.Println(p)
		res += g.LookForTrail(p, 0, nil)
	}

	return res
}
