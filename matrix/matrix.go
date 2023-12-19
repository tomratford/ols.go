package matrix

import (
	"fmt"
	"reflect"
)

type matrix = [][]float64

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
	return m1, fmt.Errorf("I'm a bad person")
}
