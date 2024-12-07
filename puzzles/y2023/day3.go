package y2023

import (
	"github.com/AgrafeModel/advent_of_code/utils"
)

type Grid struct {
	Data []string
}

func isSymbol(c byte) bool {
	return !utils.IsInt(c) && c != '.'
}

func getFullNumber(line string, x int) int {
	start := x
	end := x
	for start-1 >= 0 && utils.IsInt(line[start-1]) {
		start--
	}
	for end+1 < len(line) && utils.IsInt(line[end+1]) {
		end++
	}
	if start == end {
		return utils.ParseInt(line[start])
	}
	return utils.ParseInt(line[start : end+1])
}

func Day3Part1() int {
	res := 0
	grid := Grid{}
	utils.ReadFilePerLines(utils.GetInputPath(2023, 3), func(line string) {
		grid.Data = append(grid.Data, line)
	})

	for y := 0; y < len(grid.Data); y++ {
		for x := 0; x < len(grid.Data[y]); x++ {
			if isSymbol(grid.Data[y][x]) {
				// Get all arround it
				var num []int
				for col := -1; col < 2; col++ {
					for row := -1; row < 2; row++ {
						if col == 0 && row == 0 {
							continue
						}
						if utils.IsInt(grid.Data[y+col][x+row]) {
							v := getFullNumber(grid.Data[y+col], x+row)
							if !utils.Contains(num, v) { // Add only if not already present (temporary)
								num = append(num, v)
								res += v

							}
						}
					}
				}

			}
		}
	}

	return res
}

func Day3Part2() int {
	res := 0
	grid := Grid{}
	utils.ReadFilePerLines(utils.GetInputPath(2023, 3), func(line string) {
		grid.Data = append(grid.Data, line)
	})

	for y := 0; y < len(grid.Data); y++ {
		for x := 0; x < len(grid.Data[y]); x++ {
			if grid.Data[y][x] == '*' {
				// Get all arround it
				var num []int
				for col := -1; col < 2; col++ {
					for row := -1; row < 2; row++ {
						if col == 0 && row == 0 {
							continue
						}
						if utils.IsInt(grid.Data[y+col][x+row]) {
							v := getFullNumber(grid.Data[y+col], x+row)
							if !utils.Contains(num, v) { // Add only if not already present (temporary)
								num = append(num, v)
							}
						}
					}
				}

				if len(num) > 1 {
					res += utils.SliceMultiplyTogether(num)
				}

			}
		}
	}

	return res
}
