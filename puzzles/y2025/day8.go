package y2025

import (
	"fmt"
	"slices"

	"github.com/AgrafeModel/advent_of_code/utils"
)

/////// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
////// That was so hard tbh

func readDay8Input() []utils.Position3D {
	var start []utils.Position3D
	utils.ReadFilePerLines(utils.GetInputPath(2025, 8), func(line string) {
		var p utils.Position3D
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		start = append(start, p)
	})

	return start
}

func Day8Part1() int {
	positions := readDay8Input()

	// pre process positions to have all the distances between each other ordered
	type edge struct {
		p1   int //index in positions
		p2   int
		dist float64
	}

	var edges []edge

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			dist := utils.Distance3D(positions[i], positions[j])
			edges = append(edges, edge{p1: i, p2: j, dist: dist})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.dist - b.dist)
	})

	dsu := utils.NewDSU(len(positions))
	for i, di := range edges {
		if dsu.Find(di.p1) != dsu.Find(di.p2) {
			dsu.Union(di.p1, di.p2)
		}
		if i == 1000-1 {
			break
		}
	}

	seen := make(map[int]bool)
	size := make([]int, 0, len(positions))
	for i := range len(positions) {
		root := dsu.Find(i)
		if !seen[root] {
			seen[root] = true
			size = append(size, dsu.Size[root])
		}
	}

	slices.SortFunc(size, func(a, b int) int {
		return b - a
	})

	fmt.Println("Sizes:", size)
	return size[0] * size[1] * size[2]

}

func Day8Part2() int {
	positions := readDay8Input()

	// pre process positions to have all the distances between each other ordered
	type edge struct {
		p1   int //index in positions
		p2   int
		dist float64
	}

	var edges []edge

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			dist := utils.Distance3D(positions[i], positions[j])
			edges = append(edges, edge{p1: i, p2: j, dist: dist})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.dist - b.dist)
	})

	dsu := utils.NewDSU(len(positions))
	for _, di := range edges {
		if dsu.Find(di.p1) != dsu.Find(di.p2) {
			dsu.Union(di.p1, di.p2)
		}
		if dsu.Size[dsu.Find(0)] == len(positions) {
			return int(positions[di.p1].X * positions[di.p2].X)
		}
	}

	seen := make(map[int]bool)
	size := make([]int, 0, len(positions))
	for i := range len(positions) {
		root := dsu.Find(i)
		if !seen[root] {
			seen[root] = true
			size = append(size, dsu.Size[root])
		}
	}

	slices.SortFunc(size, func(a, b int) int {
		return b - a
	})

	fmt.Println("Sizes:", size)
	return size[0] * size[1] * size[2]

}
