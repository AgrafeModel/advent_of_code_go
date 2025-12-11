package utils

import (
	"strings"
)

type DSU struct {
	parent []int
	Size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, Size: size}
}

// Find with path compression
// Returns the root of the set in which element a is located
func (d *DSU) Find(a int) int {
	root := a
	for root != d.parent[root] {
		root = d.parent[root]
	}

	for d.parent[a] != a {
		next := d.parent[a]
		d.parent[a] = root
		a = next
	}
	return root
}

// Union by size
// Merges the sets that contain elements a and b
func (d *DSU) Union(a, b int) {
	rootA := d.Find(a)
	rootB := d.Find(b)

	if rootA != rootB {
		if d.Size[rootA] < d.Size[rootB] {
			rootA, rootB = rootB, rootA
		}
		d.parent[rootB] = rootA
		d.Size[rootA] += d.Size[rootB]
	}
}

///////////////////////////////

type GraphString struct {
	Nodes map[string]*GraphStringNode
}

type GraphStringNode struct {
	Value string
	Edges []*GraphStringNode
}

func NewGraph() *GraphString {
	return &GraphString{Nodes: make(map[string]*GraphStringNode)}
}

func (gs *GraphString) AddEdge(from, to string) {
	from = strings.TrimSpace(from)
	to = strings.TrimSpace(to)
	if from == "" || to == "" {
		return
	}
	fromNode, ok := gs.Nodes[from]
	if !ok {
		fromNode = &GraphStringNode{Value: from}
		gs.Nodes[from] = fromNode
	}
	toNode, ok := gs.Nodes[to]
	if !ok {
		toNode = &GraphStringNode{Value: to}
		gs.Nodes[to] = toNode
	}
	// avoid duplicate edges
	for _, e := range fromNode.Edges {
		if e.Value == toNode.Value {
			return
		}
	}
	fromNode.Edges = append(fromNode.Edges, toNode)
}

// Get all the paths from 'from' to 'to'
// Using DFS
func (gs *GraphString) PathFromTo(from, to string) [][]string {
	var paths [][]string
	var dfs func(current *GraphStringNode, target string, visited map[string]bool, path []string)
	dfs = func(current *GraphStringNode, target string, visited map[string]bool, path []string) {
		visited[current.Value] = true
		path = append(path, current.Value)
		if current.Value == target {
			// Found a path
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			paths = append(paths, pathCopy)
		} else {
			for _, neighbor := range current.Edges {
				if !visited[neighbor.Value] {
					dfs(neighbor, target, visited, path)
				}
			}
		}
		visited[current.Value] = false
		path = path[:len(path)-1]
	}
	startNode, ok := gs.Nodes[from]
	if !ok {
		return paths
	}
	visited := make(map[string]bool)
	dfs(startNode, to, visited, []string{})
	return paths
}
