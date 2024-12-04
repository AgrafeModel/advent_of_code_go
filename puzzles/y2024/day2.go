package y2024

import (
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

func checkOk(line []int) bool {
	last := line[0]
	dir := utils.DirectionSign(line[0], line[1])
	for i := 1; i < len(line); i++ {
		dist := utils.Dist(last, line[i])
		if utils.Between(dist, 1, 3) == false {
			return false
		}
		if dir != utils.DirectionSign(last, line[i]) {
			return false
		}
		last = line[i]
	}

	return true
}

func d2readInputPerLine(fn func(line []int)) {
	utils.ReadFilePerLines(utils.GetInputPath(2024, 2), func(line string) {
		splt := strings.Split(line, " ")
		fn(utils.StrSliceToIntSlice(splt))
	})
}

func Day2Part1() int {
	nbr := 0
	d2readInputPerLine(func(line []int) {
		if checkOk(line) {
			nbr++
		}

	})
	return nbr
}

func Day2Part2() int {
	nbr := 0
	d2readInputPerLine(func(line []int) {
		if checkOk(line) {
			nbr++
		} else {
			for i := range line {
				templn := utils.RemoveAt(line, i)
				if checkOk(templn) {
					nbr++
					break
				}
			}
		}

	})
	return nbr
}
