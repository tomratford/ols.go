package matrix

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	type TestCase struct {
		desc string
		input, want matrix
	}

	test_cases := []TestCase{
		TestCase{
			desc: "transpose of the identity is the identity",
			input: matrix{
				{1,0,0},
				{0,1,0},
				{0,0,1},
			},
			want: matrix{
				{1,0,0},
				{0,1,0},
				{0,0,1},
			},
		},
		TestCase{
			desc: "transpose of a simple 2x2 matrix",
			input: matrix{
				{2,3},
				{8,1},
			},
			want: matrix{
				{2,8},
				{3,1},
			},
		},
		TestCase{
			desc: "transpose of a tall 4x1 matrix",
			input: matrix{
				{1.4},
				{3.2},
				{2.9},
				{0.3},
			},
			want: matrix{
				{1.4,3.2,2.9,0.3},
			},
		},
		TestCase{
			desc: "transpose of a wide two by four matrix",
			input: matrix{
				{1.4,3.2,2.9,0.3},
				{4.4,2.0,9.3,3.8},
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
	t.Run("fail on invalid multiplication (pxn * nxq)", func(t *testing.T) {
		m1 := matrix{
			{2,3},
			{2,3},
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
}
