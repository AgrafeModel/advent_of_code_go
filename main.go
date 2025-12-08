package main

import (
	"flag"
	"fmt"

	"github.com/AgrafeModel/advent_of_code/puzzles/y2023"
	"github.com/AgrafeModel/advent_of_code/puzzles/y2024"
	"github.com/AgrafeModel/advent_of_code/puzzles/y2025"
)

var y2025Funcs = []func() int{
	y2025.Day1Part1,
	y2025.Day1Part2,
	y2025.Day2Part1,
	y2025.Day2Part2,
	y2025.Day3Part1,
	y2025.Day3Part2,
	y2025.Day4Part1,
	y2025.Day4Part2,
	y2025.Day5Part1,
	y2025.Day5Part2,
}

var y2024Funcs = []func() int{
	y2024.Day1Part1,
	y2024.Day1Part2,
	y2024.Day2Part1,
	y2024.Day2Part2,
	y2024.Day3Part1,
	y2024.Day3Part2,
	y2024.Day4Part1,
	y2024.Day4Part2,
	y2024.Day5Part1,
	y2024.Day5Part2,
	y2024.Day6Part1,
	y2024.Day6Part2,
	y2024.Day7Part1,
	y2024.Day7Part2,
	y2024.Day8Part1,
	y2024.Day8Part2,
	y2024.Day9Part1,
	y2024.Day9Part2,
	y2024.Day10Part1,
	y2024.Day10Part2,
}

var y2023Funcs = []func() int{
	y2023.Day1Part1,
	y2023.Day1Part2,
	y2023.Day2Part1,
	y2023.Day2Part2,
	y2023.Day3Part1,
	y2023.Day3Part2,
}

var (
	chosen_year = flag.Int("y", 0, "The year to run")
	chosen_day  = flag.Int("d", 0, "The day to run. You also have to chose a year")
	chosen_part = flag.Int("p", 0, "The part to run. You also have to chose a year and a day")
)

func main() {
	flag.Parse()

	if *chosen_year == 0 {
		allYear2023()
		allYear2024()
		allYear2025()

	} else if *chosen_part != 0 && *chosen_day != 0 && *chosen_year != 0 {
		runDayPart(*chosen_year, *chosen_day, *chosen_part)
	} else if *chosen_day != 0 && *chosen_year != 0 {
		runDay(*chosen_year, *chosen_day)
	} else if *chosen_year != 0 {
		runYear(*chosen_year)
	} else {
		fmt.Errorf("You must choose a year, a day and a part")
	}
}

func formatRunning(fn func() int, year, day, part int) {
	fmt.Println("=-=-=-={ Year ", year, " }--=-=-=")
	fmt.Println("---{ Day ", day, " - part ", part, " }---")
	res := fn()
	fmt.Println("Anwser: ", res)
}

func runYear(year int) {
	switch year {
	case 2023:
		allYear2023()
	case 2024:
		allYear2024()
	case 2025:
		allYear2025()
	default:
		fmt.Errorf("Year %d is not available", year)
	}
}

func runDay(year int, day int) {
	d := (day - 1) * 2
	switch year {
	case 2023:
		formatRunning(y2023Funcs[d], year, day, 1)
		formatRunning(y2023Funcs[d+1], year, day, 2)
	case 2024:
		formatRunning(y2024Funcs[d], year, day, 1)
		formatRunning(y2024Funcs[d+1], year, day, 2)
	case 2025:
		formatRunning(y2025Funcs[d], year, day, 1)
		formatRunning(y2025Funcs[d+1], year, day, 2)
	default:
		fmt.Errorf("Year %d is not available", year)
	}
}

func runDayPart(year int, day int, part int) {
	d := (day-2)*2 + 1
	switch year {
	case 2023:
		formatRunning(y2023Funcs[d+part], year, day, part)
	case 2024:
		formatRunning(y2024Funcs[d+part], year, day, part)
	}
}

func allYear2025() {

	fmt.Println("======[ YEAR 2025 ]======")
	for i, fn := range y2025Funcs {
		fmt.Println("---{ Day ", i/2+1, " - part ", i%2+1, " }---")
		res := fn()
		fmt.Println("Anwser: ", res)
	}

	fmt.Println("\n\n")
}

func allYear2024() {

	fmt.Println("======[ YEAR 2024 ]======")
	for i, fn := range y2024Funcs {
		fmt.Println("---{ Day ", i/2+1, " - part ", i%2+1, " }---")
		res := fn()
		fmt.Println("Anwser: ", res)
	}

	fmt.Println("\n\n")
}

func allYear2023() {

	fmt.Println("======[ YEAR 2023 ]======")
	for i, fn := range y2023Funcs {
		fmt.Println("---{ Day ", i/2+1, " - part ", i%2+1, " }---")
		res := fn()
		fmt.Println("Anwser: ", res)
	}
	fmt.Println("\n\n")
}
