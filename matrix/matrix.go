package matrix

import (
	"fmt"
	"reflect"
)

type matrix = [][]float64

// Create a `p` x `q` matrix of zeros
func Zero(p, q int) matrix {
	m := make(matrix, p, p)
	for i:=0;i<p;i++ {
		temp := make([]float64, q, q)
		m[i] = temp
	}
	return m
}

func Transpose(m matrix) matrix {
	// Empty matrix
	if reflect.DeepEqual(m, matrix{}) {
		return m
	}

	nrow := len(m)
	ncol := len(m[0])

	m2 := make(matrix, ncol, max(ncol, nrow))

	for _, v1 := range m {
		for j, v2 := range v1 {
			m2[j] = append(m2[j], v2)
		}
	}
	return m2
}

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

	return z, nil
}
