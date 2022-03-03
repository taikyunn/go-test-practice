package calc

import "testing"

func TestSum_SumMulti(t *testing.T) {

	t.Run("Len=1", func(t *testing.T) {
		t.Log("Len=1")
		if new(Sum).SumMulti(1) != 1 {
			t.Fail()
		}
	})

	t.Run("Len=2", func(t *testing.T) {
		t.Log("Len=2")
		if new(Sum).SumMulti(1, 2) != 3 {
			t.Fail()
		}
	})

	t.Run("Len=3", func(t *testing.T) {
		t.Log("Len=3")
		if new(Sum).SumMulti(1, 2, 3) != 6 {
			t.Fail()
		}
	})
}
