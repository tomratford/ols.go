package matrix

type matrix = [][]float64

func Transpose(m matrix) matrix {
	m2 := make(matrix, len(m), len(m))
	for _, v1 := range m {
		for j, v2 := range v1 {
			m2[j] = append(m2[j], v2)
		}
	}
	return m2
}
