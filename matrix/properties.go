package matrix

func IsDiagonal(x matrix) (bool, error) {
	if b, err := x.isSquare(); !b {
		return false, err
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
	if b, err := x.isSquare(); !b {
		return false, err
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
