package y2024

import (
	"regexp"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

func Day3Part1() int {

	res := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, 3), func(line string) {

		regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
		regres := regex.FindAllString(line, -1)
		for _, v := range regres {
			l := strings.Split(v, ",")
			l[0] = l[0][4:]
			l[1] = l[1][:len(l[1])-1]
			rs := utils.ParseInt(l[0]) * utils.ParseInt(l[1])
			res += rs
		}
	})

	return res
}

func Day3Part2() int {

	res := 0
	enabled := true
	utils.ReadFilePerLines(utils.GetInputPath(2024, 3), func(line string) {

		regex := regexp.MustCompile(`(?:mul\(\d+,\d+\)|do\(\)|don\'t\(\))`)
		regres := regex.FindAllString(line, -1)

		for _, v := range regres {
			if v == "don't()" {
				enabled = false
				continue
			}
			if v == "do()" {

				enabled = true
				continue
			}

			if !enabled {
				continue
			}

			l := strings.Split(v, ",")
			l[0] = l[0][4:]
			l[1] = l[1][:len(l[1])-1]
			rs := utils.ParseInt(l[0]) * utils.ParseInt(l[1])
			res += rs
		}
	})

	return res
}
