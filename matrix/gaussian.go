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