package y2025

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

func readDay5Input() ([]Ranges, []int) {
	var ra []Ranges
	var rb []int

	readingrange := true
	utils.ReadFilePerLines(utils.GetInputPath(2025, 5), func(line string) {
		if readingrange {
			var r Ranges
			//split by -
			ss := strings.Split(line, "-")
			if len(ss) != 2 {
				readingrange = false
				return
			}
			fmt.Sscanf(ss[0], "%d", &r.start)
			fmt.Sscanf(ss[1], "%d", &r.end)
			ra = append(ra, r)
		} else {
			{
				var v int
				fmt.Sscanf(line, "%d", &v)
				rb = append(rb, v)
			}
		}
	})
	return ra, rb

}

func Day5Part1() int {
	ranges, values := readDay5Input()

	count := 0

	for _, v := range values {
		for _, r := range ranges {
			if v >= r.start && v <= r.end {
				count++
				break
			}
		}
	}

	return count
}

func Day5Part2() int {
	ranges, _ := readDay5Input()
	count := 0

	//sort the ranges by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	nrg := []Ranges{}
	curr := ranges[0]
	for i := 0; i < len(ranges); i++ {
		if curr.end >= ranges[i].start-1 {
			if ranges[i].end > curr.end {
				curr.end = ranges[i].end
			}
		} else {
			nrg = append(nrg, curr)
			curr = ranges[i]
		}
	}

	nrg = append(nrg, curr)

	fmt.Println("Merged ranges:", nrg)

	//count the number of fresh vegetables in the ranges
	//reduce the number of range to the minimum number of ranges that cover all the vegetables
	for _, r := range nrg {
		count += r.end - r.start + 1
	}

	return count
}
