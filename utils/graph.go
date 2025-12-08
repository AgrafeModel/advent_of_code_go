package utils

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
