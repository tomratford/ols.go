package matrix

import "fmt"

func IsDiagonal(x matrix) (bool, error) {
	p := len(x)
	q := len(x[0])

	if p != q {
		return false, fmt.Errorf("Expected a square matrix, got a %d x %d", p, q)
	}

	for i:=0;i<p;i++ {
		for j:=0;j<q;j++ {
			if i == j { // Skip on diagonal entries
				continue
			}
			if x[i][j] != 0 {
				return false, nil
			}
		}
	}

	return true, nil
}

func IsUpperTriangular(x matrix) (bool, error) {
	p := len(x)
	q := len(x[0])

	if p != q {
		return false, fmt.Errorf("Expected a square matrix, got a %d x %d", p, q)
	}

	for i:=0;i<p;i++ {
		for j:=0;j<q;j++ {
			if i > j {
				if x[i][j] != 0 {
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
