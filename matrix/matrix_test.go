package matrix

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	t.Run("transpose of the identity is the identity", func(t *testing.T) {
		m := matrix{
			{1,0,0},
				{0,1,0},
				{0,0,1},
			}

		got := Transpose(m)
		want := m

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected transpose of the identity to be the same, got %v", got)
		}
	})

	t.Run("transpose of a simple two by two", func(t *testing.T) {
		m := matrix{
			{2,3},
			{8,1},
		}
		got := Transpose(m)
		want := matrix{
			{2,8},
			{3,1},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected to be the same, got %v, want %v", got, want)
		}
	})

	t.Run("transpose of a tall one by four", func(t *testing.T) {
		m := matrix{
			{1.4},
			{3.2},
			{2.9},
			{0.3},
		}
		got := Transpose(m)
		want := matrix{
			{1.4,3.2,2.9,0.3},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected to be the same, got %v, want %v", got, want)
		}
	})

	t.Run("transpose of a wide four by two", func(t *testing.T) {
		m := matrix{
			{1.4,3.2,2.9,0.3},
			{4.4,2.0,9.3,3.8},
		}
		got := Transpose(m)
		want := matrix{
			{1.4, 4.4},
			{3.2, 2.0},
			{2.9, 9.3},
			{0.3, 3.8},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected to be the same, got %v, want %v", got, want)
		}
	})
}
