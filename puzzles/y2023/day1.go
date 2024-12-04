package y2023

import (
	"github.com/AgrafeModel/advent_of_code/utils"
)

func Day1Part1() int {
	res := 0
	utils.ReadFilePerLines(utils.GetInputPath(2023, 1), func(line string) {
		first := utils.NewConst(0)
		last := utils.NewConst(0)

		for i := 0; i < len(line); i++ {
			if !first.IsSet() {
				if utils.IsInt(line[i]) {
					first.Set(utils.ParseInt(line[i]))
				}
			}

			if !last.IsSet() {
				if utils.IsInt(line[len(line)-i-1]) {
					last.Set(utils.ParseInt(line[len(line)-i-1]))
				}
			}

			if first.IsSet() && last.IsSet() {
				break
			}
		}

		cres := utils.ConcatInt(first.Get(), last.Get())
		res += cres
	})

	return res
}

func Day1Part2() int {
	res := 0
	utils.ReadFilePerLines(utils.GetInputPath(2023, 1), func(line string) {
		first := utils.NewConst(0)
		last := utils.NewConst(0)

		for i := 0; i < len(line); i++ {
			if !first.IsSet() {
				if utils.IsInt(line[i]) {
					first.Set(utils.ParseInt(line[i]))
				}

				for _, num := range utils.Numbers {
					if i+len(num.String()) < len(line) && line[i:i+len(num.String())] == num.String() {
						first.Set(num.Number())
					}
				}

			}

			if !last.IsSet() {
				x := len(line) - i - 1

				if utils.IsInt(line[x]) {
					last.Set(utils.ParseInt(line[x]))
				}

				for _, num := range utils.Numbers {
					if x+len(num.String()) < len(line) && line[x:x+len(num.String())] == num.String() {
						last.Set(num.Number())
					}
				}
			}

			if first.IsSet() && last.IsSet() {
				break
			}
		}

		cres := utils.ConcatInt(first.Get(), last.Get())

		res += cres
	})
	return res
}
