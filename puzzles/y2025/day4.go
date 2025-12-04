package y2025

import (
	"github.com/AgrafeModel/advent_of_code/utils"
)

func readDay4Input() [][]rune {
	var ra [][]rune

	utils.ReadFilePerLines(utils.GetInputPath(2025, 4), func(line string) {
		var row []rune
		for _, ch := range line {
			row = append(row, ch)
		}
		//remove the last element (newline)
		row = row[:len(row)-1]
		ra = append(ra, row)
	})
	return ra
}

func Day4Part1() int {
	rows := readDay4Input()
	count := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if canBeMoved(rows, i, j) {
				count++
			}
		}
	}
	return count

}

func Day4Part2() int {

	rows := readDay4Input()
	rowsupdates := make(map[[2]int]interface{})
	count := 0
	countthisround := 0

	for countthisround != 0 || count == 0 {
		countthisround = 0
		for i := 0; i < len(rows); i++ {
			for j := 0; j < len(rows[i]); j++ {
				if canBeMoved(rows, i, j) {
					rowsupdates[[2]int{i, j}] = nil
					count++
					countthisround++
				}
			}
		}

		for k := range rowsupdates {
			rows[k[0]][k[1]] = '.'
		}
		rowsupdates = make(map[[2]int]interface{})
	}
	return count
}

func canBeMoved(rows [][]rune, i, j int) bool {
	if rows[i][j] == '@' {
		rollcount := 0
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				if di == 0 && dj == 0 {
					continue
				}
				ni := i + di
				nj := j + dj
				if ni >= 0 && ni < len(rows) && nj >= 0 && nj < len(rows[i]) {
					if rows[ni][nj] == '@' {
						rollcount++
					}
				}
			}
		}

		if rollcount < 4 {
			return true
		}
	}

	return false
}
