package y2025

import (
	"strings"

	"github.com/AgrafeModel/advent_of_code/utils"
)

func readDay11nput() *utils.GraphString {
	var ra *utils.GraphString = utils.NewGraph()
	utils.ReadFilePerLines(utils.GetInputPath(2025, 11), func(line string) {
		line = strings.TrimSpace(line)
		if line == "" {
			return
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) < 2 {
			return
		}
		from := strings.TrimSpace(parts[0])
		// Fields splits on whitespace and ignores empty tokens
		toParts := strings.Fields(parts[1])
		for _, to := range toParts {
			to = strings.TrimSpace(to)
			if to == "" {
				continue
			}
			ra.AddEdge(from, to)
		}
	})

	return ra
}

func Day11Part1() int {
	graph := readDay11nput()
	count := 0
	paths := graph.PathFromTo("you", "out")
	count = len(paths)

	return count
}

func Day11Part2() int {
	count := 0
	graph := readDay11nput()
	nbrpaths := numberPathFromToPassingBy(graph.Nodes["svr"], "out")

	count = nbrpaths
	return count
}

func numberPathFromToPassingBy(from *utils.GraphStringNode, to string) int {
	type memoKey struct {
		node     string
		fft, dac bool
	}

	memo := make(map[memoKey]int)
	visited := make(map[string]bool)

	var dfs func(n *utils.GraphStringNode, fftPassed, dacPassed bool) int
	dfs = func(n *utils.GraphStringNode, fftPassed, dacPassed bool) int {
		if visited[n.Value] {
			return 0
		}

		if n.Value == "fft" {
			fftPassed = true
		}
		if n.Value == "dac" {
			dacPassed = true
		}

		if n.Value == to {
			if fftPassed && dacPassed {
				return 1
			}
			return 0
		}

		key := memoKey{node: n.Value, fft: fftPassed, dac: dacPassed}
		if v, ok := memo[key]; ok {
			return v
		}

		visited[n.Value] = true
		total := 0
		for _, nb := range n.Edges {
			total += dfs(nb, fftPassed, dacPassed)
		}
		visited[n.Value] = false

		memo[key] = total
		return total
	}

	return dfs(from, false, false)
}
