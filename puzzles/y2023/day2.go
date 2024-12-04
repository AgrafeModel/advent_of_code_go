package y2023

import (
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

type game struct {
	ID               int
	Red, Green, Blue int
}

type Color string

const RED Color = "red"
const GREEN Color = "green"
const BLUE Color = "blue"

func (g *game) SetMaxColor(c Color, v int) {
	switch c {
	case RED:
		if g.Red < v {
			g.Red = v
		}
	case BLUE:
		if g.Blue < v {
			g.Blue = v
		}
	case GREEN:
		if g.Green < v {
			g.Green = v
		}
	}
}

func processGame(line string) game {
	l := strings.Split(line, ":")
	id := utils.ParseInt(strings.Split(l[0], " ")[1])
	sets := strings.Split(l[1], ";")

	game := game{
		ID:    id,
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, set := range sets {

		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cubesplit := strings.Split(cube, " ")
			cubesplit = utils.RemoveFirstSlice(cubesplit)

			nbr := utils.ParseInt(cubesplit[0])
			color := cubesplit[1]

			game.SetMaxColor(Color(color), nbr)
		}
	}
	return game
}

func Day2Part1() int {
	res := 0

	utils.ReadFilePerLines(utils.GetInputPath(2023, 2), func(line string) {
		game := processGame(line)
		if game.Red <= 12 && game.Green <= 13 && game.Blue <= 14 {
			res += game.ID
		}

	})
	return res
}

func Day2Part2() int {
	res := 0

	utils.ReadFilePerLines(utils.GetInputPath(2023, 2), func(line string) {
		game := processGame(line)
		res += game.Red * game.Green * game.Blue

	})
	return res
}
