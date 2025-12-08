package y2025

import (
	"fmt"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type Operator struct {
	symbol string
}

func (op Operator) Apply(a, b int) int {
	switch op.symbol {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	}
	return 0
}

func readDay6Input() ([][]int, []Operator) {
	var ra [][]int
	var rb []Operator

	//firt 4 lines are int
	//last line is operators
	utils.ReadFilePerLines(utils.GetInputPath(2025, 6), func(line string) {
		if strings.ContainsAny(line, "+*") {
			//can be multiple space to separate operators
			ss := strings.Fields(line)
			for _, s := range ss {
				rb = append(rb, Operator{symbol: s})
			}
		} else {
			var row []int
			ss := strings.Fields(line)

			for _, s := range ss {
				var v int
				fmt.Sscanf(s, "%d", &v)
				row = append(row, v)
			}
			ra = append(ra, row)
		}
	})

	fmt.Println("Values:", rb)

	return ra, rb

}

func Day6Part1() int {
	values, op := readDay6Input()

	count := 0

	for i := 0; i < len(values[0]); i++ {
		res := values[0][i]
		for j := 1; j < len(values); j++ {
			res = op[i].Apply(res, values[j][i])
		}
		count += res
	}
	return count
}

func readDay6Inputp2() []string {
	var res []string
	//firt 4 lines are int
	//last line is operators
	utils.ReadFilePerLines(utils.GetInputPath(2025, 6), func(line string) {
		res = append(res, line)
	})

	return res

}

func Day6Part2() int {
	lines := readDay6Inputp2()
	count := 0

	chunckstart := 0
	chuncksize := 0
	for {
		//Get the op
		opline := lines[len(lines)-1]
		op := rune(opline[chunckstart])
		//find the next op
		nextopindex := strings.IndexAny(lines[len(lines)-1][chunckstart+1:], "+*")
		//get the chunck size
		if nextopindex == -1 {
			chuncksize = len(lines[0]) - chunckstart
		} else {
			chuncksize = nextopindex + 1
		}

		//process the chunck

		var nb []int
		for i := chunckstart; i < chunckstart+chuncksize; i++ {
			n := 0
			for j := 0; j < len(lines)-1; j++ {
				if lines[j][i] != ' ' {
					n = n*10 + int(lines[j][i]-'0')
				}
			}
			nb = append(nb, n)
		}

		oper := Operator{symbol: string(op)}
		res := nb[0]
		for j := 1; j < len(nb)-1; j++ {

			res = oper.Apply(res, nb[j])
		}
		count += res

		chunckstart += chuncksize
		if chunckstart >= len(lines[0]) {
			break
		}
	}
	return count
}
