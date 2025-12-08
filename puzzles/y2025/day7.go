package y2025

import (
	"fmt"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type labRow struct {
	splitters map[int]bool
}

func readDay7Input() (int, []labRow) {
	var res []labRow
	start := -1
	//firt 4 lines are int
	//last line is operators
	utils.ReadFilePerLines(utils.GetInputPath(2025, 7), func(line string) {
		if start == -1 {
			//find the S position
			for i, c := range line {
				if c == 'S' {
					start = i
					break
				}
			}
		} else {
			splitters := make(map[int]bool)
			for i, c := range line {
				if c == '^' {
					splitters[i] = true
				}
			}
			res = append(res, labRow{splitters: splitters})
		}

	})

	return start, res
}

func Day7Part1() int {
	start, lab := readDay7Input()
	beams := make(map[int]bool)
	beams[start] = true
	res := 0
	for _, row := range lab {
		var newBeams = make(map[int]bool)
		fmt.Println("beams:", beams)
		for b, _ := range beams {
			if row.splitters[b] {
				res++
				newBeams[b-1] = true
				newBeams[b+1] = true
			} else {
				newBeams[b] = true
			}

		}
		beams = newBeams
	}

	return res
}

func Day7Part2() int {
	start, lab := readDay7Input()
	beams := make(map[int]int)
	beams[start] = 1
	for _, row := range lab {
		var newBeams = make(map[int]int)
		fmt.Println("beams:", beams)
		for b, v := range beams {
			if row.splitters[b] {
				newBeams[b-1] += v
				newBeams[b+1] += v
			} else {
				newBeams[b] += v
			}

		}
		beams = newBeams
	}
	res := 0
	for _, v := range beams {
		res += v
	}
	return res
}
