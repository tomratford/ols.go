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

func Multiply(m1 matrix, m2 matrix) (matrix, error) {
	m1_ncol := len(m1[0])
	m2_nrow := len(m2)
	if m1_ncol != m2_nrow {
		return matrix{}, fmt.Errorf("Expected the number of `m1` columns to be the same as the number of `m2` rows")
	}
	return m1, nil
}
