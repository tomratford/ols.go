package matrix

import "fmt"

func IsDiagonal(x matrix) (bool, error) {
	if x.N != x.M {
		return false, fmt.Errorf("Expected a square matrix, got a %d x %d", x.N, x.M)
	}

	for i:=0;i<x.N;i++ {
		for j:=0;j<x.M;j++ {
			if i == j { // Skip on diagonal entries
				continue
			}
			if x.Get(i,j) != 0 {
				return false, nil
			}
		}
	}

	return true, nil
}

func IsUpperTriangular(x matrix) (bool, error) {
	if x.N != x.M {
		return false, fmt.Errorf("Expected a square matrix, got a %d x %d", x.N, x.M)
	}

	for i:=0;i<x.N;i++ {
		for j:=0;j<x.M;j++ {
			if i > j {
				if x.Get(i,j) != 0 {
					return false, nil
				}
			}
		}
	}

	return true, nil
}

func IsLowerTriangular(x matrix) (bool, error) {
	return IsUpperTriangular(Transpose(x))
}
