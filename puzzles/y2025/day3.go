package y2025

import (
	"fmt"
	"math"
	"slices"

	"github.com/AgrafeModel/advent_of_code/utils"
)

// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

func readDay3nput() [][]int {
	var ra [][]int
	utils.ReadFilePerLines(utils.GetInputPath(2025, 3), func(line string) {
		var row []int
		for _, ch := range line {
			row = append(row, int(ch-'0'))
		}
		//remove the last element (newline)
		row = row[:len(row)-1]
		ra = append(ra, row)
	})

	return ra
}

func Day3Part1() int {
	rows := readDay3nput()
	count := 0

	for i := 0; i < len(rows); i++ {
		d1 := slices.Max(rows[i][:len(rows[i])-1])
		d1Index := slices.Index(rows[i], d1)
		d2 := slices.Max(rows[i][d1Index+1:])
		count += d1*10 + d2
	}

	return count
}

func Day3Part2() int {
	rows := readDay3nput()
	count := 0
	fmt.Println(rows)
	for i := 0; i < len(rows); i++ {
		last := 0
		fmt.Println("row", rows[i])
		for j := 0; j < 12; j++ {
			d := slices.Max(rows[i][last : len(rows[i])-(12-j-1)])
			//find the first element with this value in the range and get index
			for k := last; k < len(rows[i])-(12-j-1); k++ {
				if rows[i][k] == d {
					last = k + 1
					break
				}
			}

			fmt.Println("  d", d, "last", last)
			count += int(math.Pow(10, float64(11-j)) * float64(d))
		}
	}

	return count
}
