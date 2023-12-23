package matrix

import (
	"fmt"
	"maps"
	"math"
	"reflect"
)

/*
 A sparse matrix contains a map of values, indexed by a 2-dim array representing
 the coordinates of the matrix. Where `Values[{0,0}]` is the top leftmost
 element and `Values[{N,M}]` is the bottome rightmost element.
*/
type matrix struct {
	Values map[[2]int]float64
	// Number of Rows and Columns
	N, M int
}

// Create a `n` x `m` matrix of zeros
func Zero(n, m int) matrix {
	values := make(map[[2]int]float64, n*m) // Allocate the maximum possible memory

	return matrix{
		Values: values,
		N: n,
		M: m,
	}
}

// Helper for consistent error messaging
func (x *matrix) isSquare() (bool, error) {
	if x.N == x.M {
		return true, nil
	} else {
		return false, fmt.Errorf("Expected a square matrix, got a %d x %d", x.N, x.M)
	} 
}

// Get the value in a matrix at a point
func (x *matrix) Get(i,j int) float64 {
	return x.Values[[2]int{i,j}]
}

// Set a matrix value at a point
func (x *matrix) Set(i,j int, v float64) error {
	if i > x.N || j > x.M || i < 0 || j < 0 {
		return fmt.Errorf("Invalid address")
	}
	if v == 0 {
		delete(x.Values,[2]int{i,j})
	}
	x.Values[[2]int{i,j}] = v
	return nil
}

// Add to a value at a point
func (x *matrix) Update(i,j int, v float64) error {
	if i > x.N || j > x.M || i < 0 || j < 0 {
		return fmt.Errorf("Invalid address")
	}
	if v == 0 {
		return nil // Don't want this to make new empty elements
	}
	x.Values[[2]int{i,j}] += v
	return nil
}

// Copy a matrix in it's entirety
func (x *matrix) Copy() matrix {
	return matrix{
		Values: maps.Clone(x.Values), // OK to use as we are a shallow map
		N: x.N,
		M: x.M,
	}
}

// Check is two matrices are equal
func Equal(x, y matrix) bool {
	if !(x.N == y.N && x.M == y.M) {
		return false
	}

	// number of non-zero elements
	if len(x.Values) != len(y.Values) {
		return false
	}

	for xk, xv := range x.Values {
		yv, ok := y.Values[xk]
		if !ok { // Does y have the same elements as x?
			return false
		}

		if diff := math.Abs(xv - yv); diff > Fuzz {
			fmt.Println(diff)
			return false
		}
	}
	return true
}

// This value is used for comparisons to check for fuzz
var Fuzz = 1.0e-14

// WARN: modifies matrix `x` in place
func fuzzCheck(x matrix) {
	for i := 0; i < x.N; i++ {
		for j := 0; j < x.M; j++ {
			if math.Abs(x.Get(i,j)) < Fuzz {
				x.Set(i,j,0)
			}
		}
	}
}

// Returns the transpose of matrix `x`
func Transpose(x matrix) matrix {
	// Empty matrix
	if reflect.DeepEqual(x, matrix{}) {
		return x
	}

	z := Zero(x.M, x.N)

	for k, v := range x.Values {
		z.Set(k[1],k[0],v)
	}

	return z
}

// Performs matrix multiplication between matrices `x` and `y`
func Multiply(x matrix, y matrix) (matrix, error) {
	if x.M != y.N {
		return matrix{}, fmt.Errorf("Expected the number of `x` columns to be the same as the number of `y` rows")
	}

	z := Zero(x.N, y.M)
	for i := 0; i < x.N; i++ {
		for j := 0; j < y.M; j++ {
			for k := 0; k < x.M; k++ {
				z.Update(i,j, x.Get(i,k) * y.Get(k,j))
			}
		}
	}

	fuzzCheck(z)

	return z, nil

}

// Returns matrices `x` and `y` added together
func Add(x, y matrix) (matrix, error) {
	if x.N != y.N && x.M != y.M {
		return matrix{}, fmt.Errorf("`x` and `y` must be of the same dimension")
	}

	z := Zero(x.N, x.M)

	for i := 0; i < x.N; i++ {
		for j := 0; j < x.M; j++ {
			z.Set(i,j, x.Get(i,j) + y.Get(i,j))
		}
	}

	return z, nil
}

// Returns matrix `x` scaled by factor `a`
func Scale(x matrix, a float64) matrix {
	z := Zero(x.N, x.M)

	for i := 0; i < x.N; i++ {
		for j := 0; j < x.M; j++ {
			z.Set(i,j, x.Get(i,j) * a)
		}
	}

	return z
}

// returns the determinant of matrix `x`
func Det(x matrix) (float64, error) {
	if b, err := x.isSquare(); !b {
		return 0.0, err
	}

	// some simple cases we can account for easily
	switch x.N {
	case 2:
		return x.Get(0,0)*x.Get(1,1) - x.Get(1,0)*x.Get(1,1), nil
	case 3:
		aei := x.Get(0,0) * x.Get(1,1) * x.Get(2,2)
		bfg := x.Get(0,1) * x.Get(1,2) * x.Get(2,0)
		cdh := x.Get(0,2) * x.Get(1,0) * x.Get(2,1)

		ceg := x.Get(0,2) * x.Get(1,1) * x.Get(2,0)
		bdi := x.Get(0,1) * x.Get(1,0) * x.Get(2,2)
		afh := x.Get(0,0) * x.Get(1,2) * x.Get(2,1)

		return aei + bfg + cdh - ceg - bdi - afh, nil
	}

	// We are going to try to implement this in the following way:
	// LU decomposition, this is probably worthy of a new file at this point

	return 1, nil
}

// Returns the inverse of matrix `x`
func Inverse(x matrix) (matrix, error) {
	if b, err := x.isSquare(); !b {
		return matrix{}, err
	}	

	det, err := Det(x)
	if err != nil {
		return matrix{}, err
	}
	if det == 0 {
		return matrix{}, fmt.Errorf("No inverse (determinant = 0)")
	}

	// some simple cases we can account
	switch x.M {
	case 2:
		z := Zero(x.N, x.M)
		z.Set(0,0,x.Get(1,1))
		z.Set(0,1,-x.Get(0,1))
		z.Set(1,0,-x.Get(1,0))
		z.Set(1,1,x.Get(0,0))
		return Scale(z, 1/math.Abs(det)), nil
	}

	return x, nil
}
