package y2024

import (
	"fmt"

	"github.com/AgrafeModel/advent_of_code/utils"
)

func Day8Part1() int {
	// res := 0

	antennas := make(map[rune][]utils.Position)

	var antinodes []utils.Position

	grid := []string{}

	gridHeight := 0
	gridWidth := 0

	y := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, (8)), func(line string) {
		for x, v := range line {
			if v != '.' {
				if _, ok := antennas[v]; !ok {
					antennas[v] = []utils.Position{}
				}
				antennas[v] = append(antennas[v], utils.Position{X: x, Y: y})
			}
		}
		y++

		grid = append(grid, line)
		gridWidth = len(line)
		gridHeight++
	})

	for _, positions := range antennas {
		for i, p := range positions {
			for j, p2 := range positions {
				if i == j {
					continue
				}

				dist := p.Dist(p2)
				fmt.Println(dist)

				a1 := p.RemovePos(dist)
				a2 := p2.AddPos(dist)

				if (a1.X >= 0 && a1.X < gridWidth &&
					a1.Y >= 0 && a1.Y < gridHeight) && !utils.Contains(antinodes, a1) {
					antinodes = append(antinodes, a1)
				}
				if (a2.X >= 0 && a2.X < gridWidth && a2.Y >= 0 && a2.Y < gridHeight) && !utils.Contains(antinodes, a2) {
					antinodes = append(antinodes, a2)
				}
			}

		}
	}

	return len(antinodes)
}

func Day8Part2() int {

	// res := 0

	antennas := make(map[rune][]utils.Position)
	var antinodes []utils.Position

	grid := []string{}

	gridHeight := 0
	gridWidth := 0

	y := 0
	utils.ReadFilePerLines(utils.GetInputPath(2024, (8)), func(line string) {
		for x, v := range line {
			if v != '.' {
				if _, ok := antennas[v]; !ok {
					antennas[v] = []utils.Position{}
				}
				antennas[v] = append(antennas[v], utils.Position{X: x, Y: y})
			}
		}
		y++

		grid = append(grid, line)
		gridWidth = len(line)
		gridHeight++
	})

	for _, positions := range antennas {
		for i, p := range positions {

			for j, p2 := range positions {
				if i == j {
					continue
				}

				dist := p.Dist(p2)
				fmt.Println(dist)

				a1 := p.RemovePos(dist)

				for {

					if a1.X >= 0 && a1.X < gridWidth &&
						a1.Y >= 0 && a1.Y < gridHeight {
						if !utils.Contains(antinodes, a1) {
							antinodes = append(antinodes, a1)
						}
					} else {
						break
					}
					a1 = a1.RemovePos(dist)
				}

				a1 = p.AddPos(dist)

				for {
					if a1.X >= 0 && a1.X < gridWidth &&
						a1.Y >= 0 && a1.Y < gridHeight {
						if !utils.Contains(antinodes, a1) {
							antinodes = append(antinodes, a1)
						}
					} else {
						break
					}
					a1 = a1.AddPos(dist)
				}
			}

		}
	}

	debugAntennaGrid(grid, antinodes)
	return len(antinodes)
}

func debugAntennaGrid(grid []string, antinodes []utils.Position) {
	for y, line := range grid {
		for x, v := range line {
			if v != '.' {
				fmt.Printf("%c", v)
			} else if utils.Contains(antinodes, utils.Position{X: x, Y: y}) {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", v)
			}
		}

		fmt.Println()
	}
}
