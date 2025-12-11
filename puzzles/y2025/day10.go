package y2025

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

// Intuition on part 1 : Brutforce all combinaisons
// Intuition on part 2 : Solve an equation A*x = b using Gaussian elimination
// Fuck part 2

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

	// We can solve this using matrix algebra.
	// using this form : A * x = b
	// A = matrix of buttons (rows = lights, columns = buttons) (1 if button lights the light, 0 otherwise)
	// x = vector of button presses (xi = number of times button i is pressed)
	// b = vector of target light counts (bi = number of times light i should be lit)
	// to solve it, we can use Gaussian elimination.
	// ex:
	// A = | 1 0 1 |   b = | 2 |
	//     | 1 1 0 |       | 1 |
	//     | 0 1 1 |       | 1 |
	//
	// after diagonalization :
	// A = | 1 0 0 |   b = | 1 |
	//     | 0 1 0 |       | 0 |
	//     | 0 0 1 |       | 1 |
	//
	// so the solution is : we press button 1 once, button 2 zero times, button 3 once.
	machines := readDay10Input()
	for _, m := range machines {
		A := utils.NewMatrix(len(m.targetLights), len(m.buttons))
		for j, button := range m.buttons {
			for _, lightIndex := range button {
				if lightIndex >= 0 && lightIndex < len(m.targetLights) {
					A.Set(lightIndex, j, 1)
				}
			}
		}

		b := make([]float64, len(m.targetLights))
		for i, count := range m.counter {
			b[i] = float64(count)
		}

		fmt.Println("machine: ", m.buttons, " counter: ", m.counter)
		x := utils.GausianEliminationSolve(A, b)
		fmt.Println("solution x: ", x)

	}
	return res
}
