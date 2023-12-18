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
}
