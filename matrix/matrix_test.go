package matrix

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	t.Run("transpose of the identity is the identity", func(t *testing.T) {
		m := [][]int {
			{1,0,0},
				{0,1,},
				{0,0,1},
			}

		got := Transpose(m)
		want := m

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected transpose of the identity to be the same, got %v", got)
		}
	})
}
