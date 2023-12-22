package matrix

import (
	"fmt"
	"math"
	"reflect"
)

/*
 A sparse matrix contains a map of values, indexed by a 2-dim array representing
 the coordinates of the matrix. Where `Values[{0,0}]` is the top leftmost
 element and `Values[{N,M}]` is the bottome rightmost element.
*/
type matrix struct {
	Values map[[2]uint]float64
	N, M int
}

// This value is used for comparisons to check for fuzz
var Fuzz = 1.0e-15

// WARN: modifies matrix `x` in place
func fuzzCheck(x matrix) {
	p := len(x)
	q := len(x[0])

	for i := 0; i < p; i++ {
		for j := 0; j < q; j++ {
			if math.Abs(x[i][j]) < Fuzz {
				x[i][j] = 0.0
			}
		}
	}
}

// Create a `p` x `q` matrix of zeros
func Zero(p, q int) matrix {
	z := make(matrix, p, p)
	for i := 0; i < p; i++ {
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
	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			for k := 0; k < n; k++ {
				z[i][j] += x[i][k] * y[k][j]
			}
		}
	}

	fuzzCheck(z)

	return z, nil

}

// Returns matrices `x` and `y` added together
func Add(x, y matrix) (matrix, error) {
	// Number of rows/cols in x
	m := len(x)
	n := len(x[0])
	// Numbers of rows/cols in y
	p := len(y)
	q := len(y[0])

	if m != p && n != q {
		return matrix{}, fmt.Errorf("`x` and `y` must be of the same dimension")
	}

	z := Zero(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			z[i][j] = x[i][j] + y[i][j]
		}
	}

	return z, nil
}

// Returns matrix `x` scaled by factor `a`
func Scale(x matrix, a float64) matrix {
	// Number of rows/cols in x
	m := len(x)
	n := len(x[0])

	z := Zero(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			z[i][j] = x[i][j] * a
		}
	}

	return z
}

// returns the determinant of matrix `x`
func Det(x matrix) (float64, error) {
	p := len(x)
	q := len(x[0])

	if p != q {
		return 0.0, fmt.Errorf("Input must be a square matrix")
	}

	// some simple cases we can account for easily
	switch p {
	case 2:
		return x[0][0]*x[1][1] - x[1][0]*x[1][1], nil
	case 3:
		aei := x[0][0] * x[1][1] * x[2][2]
		bfg := x[0][1] * x[1][2] * x[2][0]
		cdh := x[0][2] * x[1][0] * x[2][1]

		ceg := x[0][2] * x[1][1] * x[2][0]
		bdi := x[0][1] * x[1][0] * x[2][2]
		afh := x[0][0] * x[1][2] * x[2][1]

		return aei + bfg + cdh - ceg - bdi - afh, nil
	}

	// We are going to try to implement this in the following way:
	// LU decomposition, this is probably worthy of a new file at this point

	return 1, nil
}

// Returns the inverse of matrix `x`
func Inverse(x matrix) (matrix, error) {
	p := len(x)
	q := len(x[0])

	if p != q {
		return matrix{}, fmt.Errorf("Input must be a square matrix")
	}

	det, err := Det(x)
	if err != nil {
		return matrix{}, err
	}
	if det == 0 {
		return matrix{}, fmt.Errorf("No inverse (determinant = 0)")
	}

	// some simple cases we can account
	switch p {
	case 2:
		return Scale(matrix{{x[1][1], -x[0][1]}, {-x[1][0], x[0][0]}}, 1/math.Abs(det)), nil
	}

	return x, nil
}
