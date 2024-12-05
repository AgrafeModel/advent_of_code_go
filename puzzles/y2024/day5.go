package y2024

import (
	"fmt"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type pageRules struct {
	value         int
	before, after []int
}

func newPage(value int) pageRules {
	return pageRules{
		value: value,
	}
}

func (p pageRules) addBefore(value int) pageRules {
	p.before = append(p.before, value)
	return p
}

func (p pageRules) addAfter(value int) pageRules {
	p.after = append(p.after, value)
	return p
}

func orderPageRules(rules map[int]pageRules, page []int) []int {
	var res []int
	for _, element := range page {
		if len(res) == 0 {
			res = append(res, element)
			continue
		}
		done := false
		//we order based on the rules
		for i, v := range res {
			if utils.Contains(rules[v].after, element) {
				res = utils.InsertBefore(res, i, element)
				done = true
				break
			}
		}

		if !done {
			res = append(res, element)
		}

	}
	return res
}

func isValidRuleUpdate(rules map[int]pageRules, splt []int) bool {
	for i, v := range splt {
		//check left rules
		for lefti := i - 1; lefti >= 0; lefti-- {
			if utils.Contains(rules[v].after, (splt[lefti])) {
				return false
			}
			if !utils.Contains(rules[v].before, (splt[lefti])) {
				return false
			}
		}

		//check right rules
		for righti := i + 1; righti < len(splt); righti++ {
			if utils.Contains(rules[v].before, (splt[righti])) {
				return false
			}
			if !utils.Contains(rules[v].after, (splt[righti])) {
				return false
			}
		}
	}
	return true
}

func Day5Part1() int {
	res := 0

	rules := make(map[int]pageRules)

	parsingRules := true

	utils.ReadFilePerLines(utils.GetInputPath(2024, (5)), func(line string) {
		if line == "" {
			parsingRules = false
			return
		}

		if parsingRules {
			// rules
			splt := strings.Split(line, "|")
			left := utils.ParseInt(splt[0])
			right := utils.ParseInt(splt[1])
			if _, exists := rules[left]; !exists {
				rules[left] = newPage(left)
			}
			if _, exists := rules[right]; !exists {
				rules[right] = newPage(right)
			}
			rules[right] = rules[right].addBefore(left)
			rules[left] = rules[left].addAfter(right)
		} else {
			splt := utils.ParseSliceInt(strings.Split(line, ","))
			if isValidRuleUpdate(rules, splt) {
				//get the middle value
				mid := len(splt) / 2

				if mid == 0 {
					mid = 1
				}
				res += (splt[mid])
			}
		}
	})

	return res
}

func Day5Part2() int {
	res := 0

	rules := make(map[int]pageRules)

	parsingRules := true

	n := 0

	utils.ReadFilePerLines(utils.GetInputPath(2024, (5)), func(line string) {
		if line == "" {
			parsingRules = false
			return
		}
		// if n > 1 {
		// 	return
		// }

		if parsingRules {
			// rules
			splt := strings.Split(line, "|")
			left := utils.ParseInt(splt[0])
			right := utils.ParseInt(splt[1])
			if _, exists := rules[left]; !exists {
				rules[left] = newPage(left)
			}
			if _, exists := rules[right]; !exists {
				rules[right] = newPage(right)
			}
			rules[right] = rules[right].addBefore(left)
			rules[left] = rules[left].addAfter(right)
		} else {
			splt := utils.ParseSliceInt(strings.Split(line, ","))
			if !isValidRuleUpdate(rules, splt) { // beurk

				order := orderPageRules(rules, splt)
				fmt.Println(splt, order)
				res += order[len(order)/2]
				n++

			}
		}
	})

	return res
}
