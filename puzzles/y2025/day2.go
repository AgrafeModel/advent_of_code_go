package y2025

import (
	"fmt"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type ranges struct {
	start int
	end   int
}

func readDay2nput() []ranges {
	var ra []ranges
	utils.ReadFilePerLines(utils.GetInputPath(2025, 2), func(line string) {
		ss := strings.Split(line, ",")
		for _, part := range ss {
			var r ranges
			var start, end int
			fmt.Sscanf(part, "%d-%d", &start, &end)
			r.start = start
			r.end = end
			ra = append(ra, r)
		}
	})

	return ra
}

func Day2Part1() int {
	ranges := readDay2nput()
	count := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			s := fmt.Sprintf("%d", i)

			// get left half of the string
			left := s[:len(s)/2]
			right := s[len(s)/2:]
			if left == right {
				count += i
			}

		}
	}

	return count
}

func Day2Part2() int {
	ranges := readDay2nput()
	count := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			s := fmt.Sprintf("%d", i)

			//check if their is reapeating pattern
			for l := 1; l <= len(s)/2; l++ {
				pattern := s[:l]
				matched := true
				for j := l; j < len(s); j += l {
					end := j + l
					if end > len(s) {
						end = len(s)
					}
					if s[j:end] != pattern {
						matched = false
						break
					}
				}
				if matched {
					count += i
					break
				}
			}

		}
	}

	return count
}
