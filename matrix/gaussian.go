package matrix

import "fmt"

// Swap two rows
func SwapRows(x matrix, row1, row2 int) (matrix, error) {
	if row1 > x.N {
		return matrix{}, fmt.Errorf("row1 out of range. %d > %d", row1, x.N)
	}
	if row2 > x.N {
		return matrix{}, fmt.Errorf("row2 out of range. %d > %d", row2, x.N)
	}

	// Implemented via matrix multiplication - less efficient but should be OK
	z := Identity(x.N)

	z.Set(row1,row1,0)
	z.Set(row2,row2,0)

	z.Set(row1,row2,1)
	z.Set(row2,row1,1)

	return Multiply(z,x)
}

// Swap two rows
func ScaleRow(x matrix, row int, scale float64) (matrix, error) {
	if row > x.N {
		return matrix{}, fmt.Errorf("row1 out of range. %d > %d", row, x.N)
	}

	// Implemented via matrix multiplication - less efficient but should be OK
	z := Identity(x.N)

	z.Set(row,row,scale)

	return Multiply(z,x)
}

// Add a multiple of one row to the other
func AddToRow(x matrix, row1, row2 int, scale float64) (matrix, error) {
	if row1 > x.N {
		return matrix{}, fmt.Errorf("row1 out of range. %d > %d", row1, x.N)
	}
	if row2 > x.N {
		return matrix{}, fmt.Errorf("row2 out of range. %d > %d", row2, x.N)
	}

	// Implemented via matrix multiplication - less efficient but should be OK
	z := Identity(x.N)

	z.Set(row1,row2,scale)

	return Multiply(z,x)
}

// Return a matrix in echelon form
func GaussianElimination(x matrix) (matrix, matrix, error) {
	z := x.Copy()
	p := Identity(x.N) // Permutation matrix

	// Current Row
	i := 0
	// Current Column
	j:= 0

	for {
		// Check to see if we actually need to do anything
		if b, _ := IsUpperTriangular(z); b {
			return z,p, nil
		}
		// If we are past x.M/x.N
		if i >= x.N {
			return z, p, nil
		}
		if i > x.M {
			return z, p, nil
		}

		// Get pivot if unable to
		if z.Get(i,j) == 0 {
			for k:=i;k<=x.N;k++ {
				if k == x.N {
					j += 1 // All zero, move to next column
					break
				}
				if z.Get(k,j) != 0 {
					var err error
					z, err = SwapRows(z, i, k)
					if err != nil {
						return matrix{},matrix{}, err
					}
					p, err = SwapRows(p, i, k) // We need to track permutations too
					if err != nil {
						return matrix{},matrix{}, err
					}
					break
				}
			}
			continue // Restart whole loop with new i/j
		}

		// Make all other column entries zero
		for k:=i+1;k<x.N;k++ {
			var err error
			scale := - (z.Get(k,j) / z.Get(i,j))
			z, err = AddToRow(z, k, i, scale) // Add row scaled row i to row k
			if err != nil {
				return matrix{}, matrix{}, err
			}
		}

		i += 1 // Move onto the next row to reduce

		fuzzCheck(z) // Make any 'almost zeros' zero
	}
}
