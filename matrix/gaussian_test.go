package matrix

import "testing"

func TestSwapRow(t *testing.T) {
	type TestCase struct {
		desc        string
		input, want matrix
		row1, row2  int
	}

	test_cases := []TestCase{
		{
			desc: "Swap rows in a matrix",
			input: fromSliceOfSlices([][]float64{
				{1.4, 4.4},
				{3.2, 2.0},
				{2.9, 9.3},
				{0.3, 3.8},
			}),
			row1: 0,
			row2: 2,
			want: fromSliceOfSlices([][]float64{
				{2.9, 9.3},
				{3.2, 2.0},
				{1.4, 4.4},
				{0.3, 3.8},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := SwapRows(test_case.input, test_case.row1, test_case.row2)

			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if !Equal(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestScaleRow(t *testing.T) {
	type TestCase struct {
		desc        string
		input, want matrix
		row         int
		scale       float64
	}

	test_cases := []TestCase{
		{
			desc: "Swap rows in a matrix",
			input: fromSliceOfSlices([][]float64{
				{1.4, 4.4},
				{3.2, 2.0},
				{2.9, 9.3},
				{0.3, 3.8},
			}),
			row:   0,
			scale: 2,
			want: fromSliceOfSlices([][]float64{
				{2.8, 8.8},
				{3.2, 2.0},
				{2.9, 9.3},
				{0.3, 3.8},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := ScaleRow(test_case.input, test_case.row, test_case.scale)

			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if !Equal(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestAddToRow(t *testing.T) {
	type TestCase struct {
		desc        string
		input, want matrix
		row1, row2  int
		scale       float64
	}

	test_cases := []TestCase{
		{
			desc: "Swap rows in a matrix",
			input: fromSliceOfSlices([][]float64{
				{1.4, 4.4},
				{3.2, 2.0},
				{2.9, 9.3},
				{0.3, 3.8},
			}),
			row1:  2,
			row2:  0,
			scale: -2,
			want: fromSliceOfSlices([][]float64{
				{1.4, 4.4},
				{3.2, 2.0},
				{0.1, 0.5},
				{0.3, 3.8},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := AddToRow(test_case.input, test_case.row1, test_case.row2, test_case.scale)

			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if !Equal(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestGuassianElimination(t *testing.T) {
	type TestCase struct {
		desc        string
		input, want matrix
	}

	test_cases := []TestCase{
		{
			desc: "Gaussian elimination in a 3x3 matrix",
			input: fromSliceOfSlices([][]float64{
				{2, 1, -1},
				{-3, -1, 2},
				{-2, 1, 2},
			}),
			want: fromSliceOfSlices([][]float64{
				{2, 1, -1},
				{0, 0.5, 0.5},
				{0, 0, -1},
			}),
		},
		{
			desc: "Gaussian elimination in a 3x4 matrix",
			input: fromSliceOfSlices([][]float64{
				{1, 3, 1, 9},
				{1, 1, -1, 1},
				{3, 11, 5, 35},
			}),
			want: fromSliceOfSlices([][]float64{
				{1, 3, 1, 9},
				{0, -2, -2, -8},
				{0, 0, 0, 0},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, _, err := GaussianElimination(test_case.input)

			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if !Equal(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}
