package y2024

import (
	"strconv"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

func readDay1Input() ([]int, []int) {
	var left []int
	var right []int
	utils.ReadFilePerLines(utils.GetInputPath(2024, 1), func(line string) {
		ln := strings.Split(line, "   ")
		// ln[1] = ln[1][:len(ln[1])-1]

		n1, err := strconv.ParseInt(ln[0], 10, 64)
		utils.HandleErr(err)

		n2, err := strconv.ParseInt(ln[1], 10, 64)
		utils.HandleErr(err)

		left = append(left, int(n1))
		right = append(right, int(n2))
	})

	return left, right
}

func Day1Part1() int {

	totalDist := 0
	left, right := readDay1Input()

	i := 0
	for len(left) > 0 && len(right) > 0 {
		leftmin := utils.MinIntIndex(left)
		rightmin := utils.MinIntIndex(right)

		dist := utils.Dist(left[leftmin], right[rightmin])

		left = utils.RemoveAt(left, leftmin)
		right = utils.RemoveAt(right, rightmin)
		totalDist += dist
		i++
	}

	return totalDist
}

func Day1Part2() int {

	count := make(map[int]int)
	left, right := readDay1Input()

	for _, v := range right {
		count[v]++
	}

	res := 0
	for _, v := range left {
		res += count[v] * v
	}

	return res
}
