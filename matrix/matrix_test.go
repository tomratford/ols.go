package matrix

import (
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {
	type TestCase struct {
		desc string
		p, q int
		want matrix
	}
	test_cases := []TestCase{
		{
			desc: "2 x 2 zero",
			p:    2,
			q:    2,
			want: matrix{
				{0,0},
				{0,0},
			},
		},
		{
			desc: "1 x 3 zero",
			p:    1,
			q:    3,
			want: matrix{
				{0,0,0},
			},
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got := Zero(test_case.p,test_case.q)
			if !reflect.DeepEqual(got, test_case.want) {
				t.Errorf("expected to be equal, got %v, want %v", got, test_case.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	type TestCase struct {
		desc        string
		input, want matrix
	}

	test_cases := []TestCase{
		{
			desc: "transpose of the identity is the identity",
			input: matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			desc: "transpose of a simple 2x2 matrix",
			input: matrix{
				{2, 3},
				{8, 1},
			},
			want: matrix{
				{2, 8},
				{3, 1},
			},
		},
		{
			desc: "transpose of a tall 4x1 matrix",
			input: matrix{
				{1.4},
				{3.2},
				{2.9},
				{0.3},
			},
			want: matrix{
				{1.4, 3.2, 2.9, 0.3},
			},
		},
		{
			desc: "transpose of a wide two by four matrix",
			input: matrix{
				{1.4, 3.2, 2.9, 0.3},
				{4.4, 2.0, 9.3, 3.8},
			},
			want: matrix{
				{1.4, 4.4},
				{3.2, 2.0},
				{2.9, 9.3},
				{0.3, 3.8},
			},
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
		m1 := matrix{
			{2, 3},
			{2, 3},
		}
		m2 := matrix{
			{1},
			{2},
			{3},
		}
		_, err := Multiply(m1, m2)
		if err == nil {
			t.Errorf("expected multiply to fail")
		}
	})

	test_cases := []TestCase{
		{
			desc: "multiplication with the identity is the same",
			input1: matrix{
				{3.2, 3.0, 2.9},
				{0.3, 1.23, 83.3},
				{58.2, 12.1, 100},
			},
			input2: matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: matrix{
				{3.2, 3.0, 2.9},
				{0.3, 1.23, 83.3},
				{58.2, 12.1, 100},
			},
		},
		{
			desc: "multiplication of two 2x2 matrix",
			input1: matrix{
				{2, 3},
				{8, 1},
			},
			input2: matrix{
				{10, 4},
				{11, 3},
			},
			want: matrix{
				{53, 17},
				{91, 35},
			},
		},
		{
			desc: "multiplication of a wide 1x4 matrix by a tall 4x1 matrix",
			input1: matrix{
				{1.4, 2.4, 3.3, 9.1},
			},
			input2: matrix{
				{1.4},
				{3.2},
				{2.9},
				{0.3},
			},
			want: matrix{
				{21.94},
			},
		},
		{
			desc: "multiplication by it's own inverse returns the identity",
			input1: matrix{
				{3, 4},
				{1, 2},
			},
			input2: matrix{
				{1, -2},
				{-0.5, 1.5},
			},
			want: matrix{
				{1, 0},
				{0, 1},
			},
		},
		{
			desc: "multiplication by it's own inverse returns the identity (3x3)",
			input1: matrix{
				{1, 2, 3},
				{1, 2, 1},
				{1, 1, 4},
			},
			input2: matrix{
				{-3.5, 2.5, 2},
				{1.5, -0.5, -1},
				{0.5, -0.5, 0},
			},
			want: matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			desc: "multiplication by it's own inverse returns the identity (fuzz)",
			input1: matrix{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			},
			input2: matrix{
				{-5.0/12, 0.25, 1.0/3},
				{7.0/12, 0.25, -2.0/3},
				{1.0/12, -0.25, 1.0/3},
			},
			want: matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.desc, func(t *testing.T) {
			got, err := Multiply(test_case.input1, test_case.input2)

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
		m := matrix{
			{1},
			{2},
			{3},
		}
		_, err := Inverse(m)
		if err == nil {
			t.Errorf("expected inverse to fail")
		}
	})

	test_cases := []TestCase{
		{
			desc: "inverse of the identity is the identity",
			input: matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
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
