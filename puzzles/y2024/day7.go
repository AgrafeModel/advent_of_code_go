package y2024

import (
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

func getPossiblesResults(values []int) []int {
	possibles := []int{values[0]}
	for i := 1; i < len(values); i++ {
		res := []int{}
		for _, p := range possibles {
			res = append(res, p*values[i])
			res = append(res, p+values[i])
		}
		possibles = res
	}
	return possibles
}

func getPossiblesResultsWithConcat(values []int) []int {
	possibles := []int{values[0]}
	for i := 1; i < len(values); i++ {
		res := []int{}
		for _, p := range possibles {
			res = append(res, p*values[i])
			res = append(res, p+values[i])
			res = append(res, utils.ConcatInt(p, values[i]))
		}

		possibles = res
	}

	return possibles
}

func Day7Part1() int {

	res := 0

	utils.ReadFilePerLines(utils.GetInputPath(2024, (7)), func(line string) {
		splt := strings.Split(line, ": ")
		result := utils.ParseInt(splt[0])
		values := utils.ParseSliceInt(strings.Split(splt[1], " "))
		possibles := getPossiblesResults(values)
		for _, p := range possibles {
			if p == result {
				res += p
				return
			}
		}
	})
	return res
}

func Day7Part2() int {

	res := 0

	utils.ReadFilePerLines(utils.GetInputPath(2024, (7)), func(line string) {
		splt := strings.Split(line, ": ")
		result := utils.ParseInt(splt[0])
		values := utils.ParseSliceInt(strings.Split(splt[1], " "))
		possibles := getPossiblesResultsWithConcat(values)
		for _, p := range possibles {
			if p == result {
				res += p
				return
			}
		}
	})
	return res
}
