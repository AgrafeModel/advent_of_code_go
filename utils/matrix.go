package utils

import (
	"fmt"
	"math"
)

type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return &Matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func (m *Matrix) Set(row, col int, value float64) {
	m.Data[row][col] = value
}

func (m *Matrix) Get(row, col int) float64 {
	return m.Data[row][col]
}

func (m *Matrix) MultiplyVector(vec []float64) []float64 {
	if len(vec) != m.Cols {
		panic("vector size does not match matrix columns")
	}
	result := make([]float64, m.Rows)
	for i := 0; i < m.Rows; i++ {
		sum := 0.0
		for j := 0; j < m.Cols; j++ {
			sum += m.Data[i][j] * vec[j]
		}
		result[i] = sum
	}
	return result
}

func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.Rows != other.Rows || m.Cols != other.Cols {
		panic("matrix sizes do not match for addition")
	}

	result := NewMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Data[i][j] = m.Data[i][j] + other.Data[i][j]
		}
	}
	return result
}

// Applies the transpose of the matrix to a vector
func (m *Matrix) Transpose() *Matrix {
	result := NewMatrix(m.Cols, m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Data[j][i] = m.Data[i][j]
		}
	}
	return result
}

func (m *Matrix) Dot(other *Matrix) *Matrix {
	if m.Cols != other.Rows {
		panic("matrix sizes do not match for dot product")
	}

	result := NewMatrix(m.Rows, other.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < other.Cols; j++ {
			sum := 0.0
			for k := 0; k < m.Cols; k++ {
				sum += m.Data[i][k] * other.Data[k][j]
			}
			result.Data[i][j] = sum
		}
	}
	return result
}

func (m *Matrix) Print() {
	for i := 0; i < m.Rows; i++ {
		fmt.Println(m.Data[i])
	}
}

func GausianEliminationSolve(A *Matrix, b []float64) []float64 {
	fmt.Println("A:")

	A.Print()

	fmt.Println("b:", b)

	// Require square system
	if len(b) != A.Rows {
		return nil
	}

	// build augmented matrix
	aug := NewMatrix(A.Rows, A.Cols+1)
	for i := 0; i < A.Rows; i++ {
		for j := 0; j < A.Cols; j++ {
			aug.Data[i][j] = A.Data[i][j]
		}
		aug.Data[i][A.Cols] = b[i]
	}

	aug.Print()

	const eps = 1e-12

	// Forward elimination with partial pivoting
	for i := 0; i < A.Rows; i++ {
		// find pivot row (max abs in column i)
		maxRow := i
		maxVal := math.Abs(aug.Data[i][i])
		for r := i + 1; r < A.Rows; r++ {
			if v := math.Abs(aug.Data[r][i]); v > maxVal {
				maxVal = v
				maxRow = r
			}
		}
		// if pivot is (nearly) zero -> singular matrix -> no unique solution
		if maxVal < eps {
			return nil
		}
		// swap if needed
		if maxRow != i {
			aug.Data[i], aug.Data[maxRow] = aug.Data[maxRow], aug.Data[i]
		}

		// normalize pivot row
		pivot := aug.Data[i][i]
		for j := i; j <= A.Cols; j++ {
			aug.Data[i][j] /= pivot
		}

		// eliminate below
		for r := i + 1; r < A.Rows; r++ {
			factor := aug.Data[r][i]
			if factor == 0 {
				continue
			}
			for c := i; c <= A.Cols; c++ {
				aug.Data[r][c] -= factor * aug.Data[i][c]
			}
		}
	}

	//For debug, print the augmented matrix after elimination
	fmt.Println("Augmented matrix after elimination:")
	for i := 0; i < A.Rows; i++ {
		fmt.Println(aug.Data[i])
	}

	// Back substitution
	x := make([]float64, A.Rows)
	for i := A.Rows - 1; i >= 0; i-- {
		// RHS is at column n
		val := aug.Data[i][A.Cols]
		for j := i + 1; j < A.Cols; j++ {
			val -= aug.Data[i][j] * x[j]
		}
		x[i] = val
	}

	return x
}
