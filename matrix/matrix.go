package matrix

import (
	"fmt"
	"reflect"
	"math"
)

type matrix = [][]float64

// This value is used for comparisons to check for fuzz
var Fuzz = 1.0e-15

// WARN: modifies matrix `x` in place
func fuzzCheck(x matrix) {
	p := len(x)
	q := len(x[0])

	for i:=0;i<p;i++ {
		for j:=0;j<q;j++ {
			if math.Abs(x[i][j]) < Fuzz {
				x[i][j] = 0.0
			}
		}
	}
}

// Create a `p` x `q` matrix of zeros
func Zero(p, q int) matrix {
	z := make(matrix, p, p)
	for i:=0;i<p;i++ {
		temp := make([]float64, q, q)
		z[i] = temp
	}
	return z
}

// Returns the transpose of matrix `x`
func Transpose(x matrix) matrix {
	// Empty matrix
	if reflect.DeepEqual(x, matrix{}) {
		return x
	}

	nrow := len(x)
	ncol := len(x[0])

	m2 := make(matrix, ncol, max(ncol, nrow))

	for _, v1 := range x {
		for j, v2 := range v1 {
			m2[j] = append(m2[j], v2)
		}
	}
	return m2
}

// Performs matrix multiplication between matrices `x` and `y`
func Multiply(x matrix, y matrix) (matrix, error) {
	// Number of rows in x
	m := len(x)
	// Numbers of cols in y
	p := len(y[0])

	// Number of cols in x/rows in y
	n := len(x[0])
	n_check := len(y)

	if n != n_check {
		return matrix{}, fmt.Errorf("Expected the number of `x` columns to be the same as the number of `y` rows")
	}

	z := Zero(m, p)
	for i:=0;i<m;i++ {
		for j:=0;j<p;j++ {
			for k:=0;k<n;k++ {
				z[i][j] += x[i][k] * y[k][j]
			}
		}
	}

	fuzzCheck(z)

	return z, nil
}

// Returns the inverse of matrix `x`
func Inverse(x matrix) (matrix, error) {
	p := len(x)
	q := len(x[0])

	if p != q {
		return matrix{}, fmt.Errorf("Input must be a square matrix")
	}
	
	return x, nil
}
