package y2025

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

// Intuition on part 1 : Brutforce all combinaisons
// Intuition on part 2 : Solve an equation

type joltageMachine struct {
	targetLights []int
	buttons      [][]int
	counter      []int
}

type joltageMachineBinary struct {
	targetLights int
	buttons      []int
	counter      []int
}

func (j *joltageMachine) toBinary() joltageMachineBinary {
	var targetLightsBinary int
	for i, light := range j.targetLights {
		if light == 1 {
			targetLightsBinary |= (1 << i)
		}
	}
	var buttonsBinary []int
	for _, button := range j.buttons {
		var buttonBinary int
		for _, val := range button {
			buttonBinary |= (1 << val)
		}
		buttonsBinary = append(buttonsBinary, buttonBinary)
	}

	return joltageMachineBinary{
		targetLights: targetLightsBinary,
		buttons:      buttonsBinary,
		counter:      j.counter,
	}
}

func (jn *joltageMachineBinary) bruteForceSolve() []int {
	N := len(jn.buttons)
	totalCombinations := 1 << N
	for np := 1; np < totalCombinations; np++ {
		// np = number of pressed buttons
		found := false
		var fcombination []int
		jn.sequence(np, func(combination []int) bool {
			if jn.tryCombination(combination) {
				found = true
				fcombination = combination
				return true
			}
			return false
		})
		if found {
			return fcombination
		}
	}
	return nil
}

func (jn *joltageMachine) bruteForceSolveCounter() []int {
	N := len(jn.buttons)
	totalCombinations := 1 << N
	for np := 1; np < totalCombinations; np++ {
		// np = number of pressed buttons
		found := false
		var fcombination []int
		jn.sequence(np, func(combination []int) bool {
			if jn.tryCombinationCounter(combination) {
				found = true
				fcombination = combination
				return true
			}
			return false
		})
		if found {
			return fcombination
		}
	}
	return nil
}

func (jn *joltageMachine) sequence(n int, fn func([]int) bool) {
	combination := make([]int, n)
	var backtrack func(pos int, start int) bool
	backtrack = func(pos int, start int) bool {
		if pos == n {
			// copy before d'appeler fn
			cpy := append([]int(nil), combination...)
			return fn(cpy)
		}
		for i := start; i < len(jn.buttons); i++ {
			combination[pos] = i
			// passer i (et non i+1) pour autoriser la répétition
			if backtrack(pos+1, i) {
				return true
			}
		}
		return false
	}
	backtrack(0, 0)
}

func (jn *joltageMachineBinary) sequence(n int, fn func([]int) bool) {
	combination := make([]int, n)
	var backtrack func(pos int, start int)
	backtrack = func(pos int, start int) {
		if pos == n {
			if fn(combination) {
				return
			}
			return
		}
		for i := start; i < len(jn.buttons); i++ {
			combination[pos] = i
			backtrack(pos+1, i+1)
		}
	}
	backtrack(0, 0)
}

func (jn *joltageMachineBinary) tryCombination(combination []int) bool {
	result := 0
	for _, buttonIndex := range combination {
		result ^= jn.buttons[buttonIndex]
	}
	return result == jn.targetLights
}

// instead of look if the combination matches, count how many time a light is lit
// and compare to jn.counter
func (jn *joltageMachine) tryCombinationCounter(combination []int) bool {
	resultCounter := make([]int, len(jn.counter))
	for _, buttonIndex := range combination {
		lt := jn.buttons[buttonIndex]
		for _, lightIndex := range lt {
			if lightIndex >= 0 && lightIndex < len(resultCounter) {
				resultCounter[lightIndex]++
			}
		}
	}
	for i := 0; i < len(resultCounter); i++ {
		if resultCounter[i] != jn.counter[i] {
			return false
		}
	}
	return true
}

func readDay10Input() []joltageMachine {
	var data []joltageMachine
	utils.ReadFilePerLines(utils.GetInputPath(2025, 10), func(line string) {
		line = strings.TrimSpace(line)
		if line == "" {
			return
		}

		var jm joltageMachine

		// find target between [ and ]
		l := strings.IndexByte(line, '[')
		r := strings.IndexByte(line, ']')
		if l == -1 || r == -1 || r <= l {
			// malformed line, skip
			return
		}
		targetStr := line[l+1 : r]
		for _, ch := range targetStr {
			if ch == '#' {
				jm.targetLights = append(jm.targetLights, 1)
			} else {
				jm.targetLights = append(jm.targetLights, 0)
			}
		}

		rest := strings.TrimSpace(line[r+1:])

		// parse zero or more button groups like "(0,3,4)"
		for len(rest) > 0 && rest[0] == '(' {
			end := strings.IndexByte(rest, ')')
			if end == -1 {
				break
			}
			inner := rest[1:end]
			inner = strings.TrimSpace(inner)
			if inner != "" {
				parts := strings.Split(inner, ",")
				var button []int
				for _, p := range parts {
					p = strings.TrimSpace(p)
					if p == "" {
						continue
					}
					v, err := strconv.Atoi(p)
					if err != nil {
						continue
					}
					button = append(button, v)
				}
				jm.buttons = append(jm.buttons, button)
			} else {
				jm.buttons = append(jm.buttons, []int{})
			}
			rest = strings.TrimSpace(rest[end+1:])
		}

		// optional ext section { ... }
		if len(rest) > 0 && rest[0] == '{' {
			end := strings.IndexByte(rest, '}')
			if end != -1 {
				inner := strings.TrimSpace(rest[1:end])
				if inner != "" {
					parts := strings.Split(inner, ",")
					for _, p := range parts {
						p = strings.TrimSpace(p)
						if p == "" {
							continue
						}
						v, err := strconv.Atoi(p)
						if err != nil {
							continue
						}
						jm.counter = append(jm.counter, v)
					}
				}
			}
		}

		data = append(data, jm)
	})
	return data
}

func Day10Part1() int {
	machines := readDay10Input()
	machinesBin := []joltageMachineBinary{}
	for _, m := range machines {
		machinesBin = append(machinesBin, m.toBinary())
	}

	fmt.Println("Total machines to solve:", len(machinesBin))
	fmt.Println("ex: ", machinesBin[0])

	res := 0

	// The goal is to find the combination of button presses to match the targetLights pattern
	for _, mb := range machinesBin {
		combination := mb.bruteForceSolve()
		if combination != nil {
			fmt.Println("Found combination:", combination, "for target lights:", mb.targetLights, "buttons:", mb.buttons)
			res += len(combination)
		}
	}

	return res
}

func Day10Part2() int {

	res := 0

	fmt.Println("We can solve this one using linear algebra but AAAAAAAAAAAAAAAAAAAAAAAAA f it")
	return res
}
