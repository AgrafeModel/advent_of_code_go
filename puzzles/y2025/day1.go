package y2025

import (
	"fmt"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type rotation struct {
	direction int // 1 for right, -1 for left
	steps     int
}

func readDay1Input() []rotation {
	var rotations []rotation

	utils.ReadFilePerLines(utils.GetInputPath(2025, 1), func(line string) {
		var rot rotation
		if line[0] == 'L' {
			rot.direction = -1
		} else {
			rot.direction = 1
		}
		fmt.Sscanf(line[1:], "%d", &rot.steps)
		rotations = append(rotations, rot)
	})

	return rotations
}

func Day1Part1() int {
	rotations := readDay1Input()
	password := 0

	pos := 50
	for _, rot := range rotations {
		pos = (pos + rot.direction*rot.steps + 100) % 100
		if pos == 0 {
			password += 1
		}
	}
	return password
}

// Simple method (not the most efficient)
func Day1Part2() int {
	rotations := readDay1Input()
	password := 0

	pos := 50
	for _, rot := range rotations {

		steps := rot.steps
		for steps > 0 {
			pos = (pos + rot.direction + 100) % 100
			steps--
			if pos == 0 {
				password += 1
			}
		}

	}
	return password
}
