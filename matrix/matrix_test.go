package matrix

import (
	"reflect"
	"testing"
)

// Helper function to convert slice of slices to a matrix object
func fromSliceOfSlices(x [][]float64) matrix {
	p := len(x)
	q := len(x[0])

	z := Zero(p, q)

	for i := 0; i < p; i++ {
		for j := 0; j < q; j++ {
			z.Set(i,j,x[i][j])
		}
	}
	return z
}

func TestUpdate(t *testing.T) {
	t.Run("what if we update an empty element?", func(t *testing.T) {
		z := Zero(4,4)
		z.Update(1,1,3) // Does this element become 3?
		z.Update(0,0,0) // Does this stay an empty element?

		if z.Get(1,1) != 3 {
			t.Errorf("does not work on empties")
		}
		if _, got := z.Values[[2]int{0,0}]; got {
			t.Errorf("makes an empty value have a value")
		}
	})
}

func TestTranspose(t *testing.T) {
	type TestCase struct {
		desc        string
		input, want matrix
	}

	test_cases := []TestCase{
		{
			desc: "transpose of the identity is the identity",
			input: fromSliceOfSlices([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
			want: fromSliceOfSlices([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
		},
		{
			desc: "transpose of a simple 2x2 matrix",
			input: fromSliceOfSlices([][]float64{
				{2, 3},
				{8, 1},
			}),
			want: fromSliceOfSlices([][]float64{
				{2, 8},
				{3, 1},
			}),
		},
		{
			desc: "transpose of a tall 4x1 matrix",
			input: fromSliceOfSlices([][]float64{
				{1.4},
				{3.2},
				{2.9},
				{0.3},
			}),
			want: fromSliceOfSlices([][]float64{
				{1.4, 3.2, 2.9, 0.3},
			}),
		},
		{
			desc: "transpose of a wide two by four matrix",
			input: fromSliceOfSlices([][]float64{
				{1.4, 3.2, 2.9, 0.3},
				{4.4, 2.0, 9.3, 3.8},
			}),
			want: fromSliceOfSlices([][]float64{
				{1.4, 4.4},
				{3.2, 2.0},
				{2.9, 9.3},
				{0.3, 3.8},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got := Transpose(test_case.input)

			if !reflect.DeepEqual(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type TestCase struct {
		desc                 string
		input1, input2, want matrix
	}

	t.Run("fail on invalid multiplication (pxn * nxq)", func(t *testing.T) {
		m1 := fromSliceOfSlices([][]float64{
			{2, 3},
			{2, 3},
		})
		m2 := fromSliceOfSlices([][]float64{
			{1},
			{2},
			{3},
		})
		_, err := Multiply(m1, m2)
		if err == nil {
			t.Errorf("expected multiply to fail")
		}
	})

	test_cases := []TestCase{
		{
			desc: "multiplication with the identity is the same",
			input1: fromSliceOfSlices([][]float64{
				{3.2, 3.0, 2.9},
				{0.3, 1.23, 83.3},
				{58.2, 12.1, 100},
			}),
			input2: fromSliceOfSlices([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
			want: fromSliceOfSlices([][]float64{
				{3.2, 3.0, 2.9},
				{0.3, 1.23, 83.3},
				{58.2, 12.1, 100},
			}),
		},
		{
			desc: "multiplication of two 2x2 matrix",
			input1: fromSliceOfSlices([][]float64{
				{2, 3},
				{8, 1},
			}),
			input2: fromSliceOfSlices([][]float64{
				{10, 4},
				{11, 3},
			}),
			want: fromSliceOfSlices([][]float64{
				{53, 17},
				{91, 35},
			}),
		},
		{
			desc: "multiplication of a wide 1x4 matrix by a tall 4x1 matrix",
			input1: fromSliceOfSlices([][]float64{
				{1.4, 2.4, 3.3, 9.1},
			}),
			input2: fromSliceOfSlices([][]float64{
				{1.4},
				{3.2},
				{2.9},
				{0.3},
			}),
			want: fromSliceOfSlices([][]float64{
				{21.94},
			}),
		},
		{
			desc: "multiplication by it's own inverse returns the identity",
			input1: fromSliceOfSlices([][]float64{
				{3, 4},
				{1, 2},
			}),
			input2: fromSliceOfSlices([][]float64{
				{1, -2},
				{-0.5, 1.5},
			}),
			want: fromSliceOfSlices([][]float64{
				{1, 0},
				{0, 1},
			}),
		},
		{
			desc: "multiplication by it's own inverse returns the identity (3x3)",
			input1: fromSliceOfSlices([][]float64{
				{1, 2, 3},
				{1, 2, 1},
				{1, 1, 4},
			}),
			input2: fromSliceOfSlices([][]float64{
				{-3.5, 2.5, 2},
				{1.5, -0.5, -1},
				{0.5, -0.5, 0},
			}),
			want: fromSliceOfSlices([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
		},
		{
			desc: "multiplication by it's own inverse returns the identity (fuzz)",
			input1: fromSliceOfSlices([][]float64{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			}),
			input2: fromSliceOfSlices([][]float64{
				{-5.0 / 12, 0.25, 1.0 / 3},
				{7.0 / 12, 0.25, -2.0 / 3},
				{1.0 / 12, -0.25, 1.0 / 3},
			}),
			want: fromSliceOfSlices([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := Multiply(test_case.input1, test_case.input2)

			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			if !Equal(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestDet(t *testing.T) {
	type TestCase struct {
		desc  string
		input matrix
		want  float64
	}

	t.Run("fail on invalid multiplication (non-square matrix)", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{1},
			{2},
			{3},
		})
		_, err := Det(m)
		if err == nil {
			t.Errorf("expected determinant to fail")
		}
	})

	test_cases := []TestCase{
		{
			desc: "determinant of a 2x2 matrix",
			input: fromSliceOfSlices([][]float64{
				{1, 2},
				{3, 2},
			}),
			want: -4,
		},
		{
			desc: "determinant of a 3x3 matrix",
			input: fromSliceOfSlices([][]float64{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			}),
			want: -12,
		},
	}
	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := Det(test_case.input)

			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			if !reflect.DeepEqual(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestInverse(t *testing.T) {
	type TestCase struct {
		desc        string
		input, want matrix
	}

	t.Run("fail on invalid multiplication (non-square matrix)", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{1},
			{2},
			{3},
		})
		_, err := Inverse(m)
		if err == nil {
			t.Errorf("expected inverse to fail")
		}
	})

	test_cases := []TestCase{
		{
			desc: "inverse of the identity is the identity",
			input: fromSliceOfSlices([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
			want: fromSliceOfSlices([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
		},
		{
			desc: "inverse of a 2x2 matrix",
			input: fromSliceOfSlices([][]float64{
				{1, 2},
				{3, 2},
			}),
			want: fromSliceOfSlices([][]float64{
				{0.5, -0.5},
				{-0.75, 0.25},
			}),
		},
		{
			desc: "inverse of a 3x3 matrix",
			input: fromSliceOfSlices([][]float64{
				{1, 2, 3},
				{1, 2, 1},
				{1, 1, 4},
			}),
			want: fromSliceOfSlices([][]float64{
				{-3.5, 2.5, 2},
				{1.5, -0.5, -1},
				{0.5, -0.5, 0},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := Inverse(test_case.input)

			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			if !reflect.DeepEqual(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type TestCase struct {
		desc                 string
		input1, input2, want matrix
	}

	t.Run("fail on invalid addition (pxq + nxm)", func(t *testing.T) {
		m1 := fromSliceOfSlices([][]float64{
			{2, 3},
			{2, 3},
		})
		m2 := fromSliceOfSlices([][]float64{
			{1},
			{2},
			{3},
		})
		_, err := Add(m1, m2)
		if err == nil {
			t.Errorf("expected add to fail")
		}
	})

	test_cases := []TestCase{
		{
			desc: "Add two matrices together",
			input1: fromSliceOfSlices([][]float64{
				{3.2, 3.0, 2.9},
				{0.3, 1.23, 83.3},
				{58.2, 12.1, 100},
			}),
			input2: fromSliceOfSlices([][]float64{
				{1.2, 0.3, 0.1},
				{2.0, 1.0, 7.0},
				{8.3, 48.3, 100},
			}),
			want: fromSliceOfSlices([][]float64{
				{4.4, 3.3, 3.0},
				{2.3, 2.23, 90.3},
				{66.5, 60.4, 200},
			}),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := Add(test_case.input1, test_case.input2)

			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			if !reflect.DeepEqual(got, test_case.want) {
				t.Errorf("expected to be the same, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	t.Run("scale by a factor", func(t *testing.T) {
		x := fromSliceOfSlices([][]float64{
			{1, 2},
			{3, 4},
		})
		a := 0.5
		got := Scale(x, a)
		want := fromSliceOfSlices([][]float64{
			{0.5, 1.0},
			{1.5, 2.0},
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected to be the same, got %v, want %v", got, want)
		}
	})
}
