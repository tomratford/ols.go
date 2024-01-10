package matrix

import (
	"testing"
)

func TestIsDiagonal(t *testing.T) {
	t.Run("fail on non square matrices", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{1},
			{2},
			{3},
		})
		_, err := IsDiagonal(m)
		if err == nil {
			t.Errorf("expected IsDiagonal to fail")
		}
	})

	t.Run("Passed on diagonal matrices", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{3, 0, 0},
			{0, 89, 0},
			{0, 0, 7},
		})
		got, err := IsDiagonal(m)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if got != true {
			t.Errorf("Wanted true, got false")
		}
	})

	t.Run("Fails on non-diagonal matrices", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{3, 4, 8},
			{0, 8, 3},
			{1, 3, 7},
		})
		got, err := IsDiagonal(m)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if got != false {
			t.Errorf("Wanted false, got true")
		}
	})
}

func TestIsUpperTriangular(t *testing.T) {
	t.Run("fail on non square matrices", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{1},
			{2},
			{3},
		})
		_, err := IsUpperTriangular(m)
		if err == nil {
			t.Errorf("expected IsDiagonal to fail")
		}
	})

	t.Run("Passed on upper triangular matrices", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{3, 2, 0},
			{0, 89, 8},
			{0, 0, 7},
		})
		got, err := IsUpperTriangular(m)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if got != true {
			t.Errorf("Wanted true, got false")
		}
	})

	t.Run("Fails on non upper triangular matrices", func(t *testing.T) {
		m := fromSliceOfSlices([][]float64{
			{3, 4, 8},
			{0, 8, 3},
			{1, 3, 7},
		})
		got, err := IsUpperTriangular(m)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
		if got != false {
			t.Errorf("Wanted false, got true")
		}
	})
}
