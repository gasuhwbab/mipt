package first

import "testing"

func TestTasks(t *testing.T) {
	t.Run("TaskOne", func(t *testing.T) {
		a, b := []int{3, 5, 9, 12}, []int{4, 5, 5, 6}
		want := []int{3, 4, 5, 5, 5, 6, 9, 12}
		got := TaskOne(a, b, len(a), len(b))
		assertSlices(t, got, want)
	})
	t.Run("TaskTwo/TestOne", func(t *testing.T) {
		a, b := []int{3, 4, 5, 10}, []int{1, 2, 3, 4, 5, 16, 22, 10}
		want := true
		got := TaskTwo(a, b, len(a), len(b))
		assertElements(t, got, want)
	})
	t.Run("TaskTwo/TestOne", func(t *testing.T) {
		a, b := []int{3, 4, 5, 10}, []int{1, 2, 3, 4, 10, 5, 16, 22}
		want := false
		got := TaskTwo(a, b, len(a), len(b))
		assertElements(t, got, want)
	})
	t.Run("TaskThree", func(t *testing.T) {
		want := []int{0, 0, 3, 3, 3, 7, 0, 0, 0, 0}
		got := []int{0, 0, 0, 0, 0, 4, 0, 0, 0, 0}
		TaskThree(got, 2, 5, 3)
		assertSlices(t, got, want)
	})
	t.Run("TaskFour", func(t *testing.T) {
		a := []int{1, 0, 1, 2, 3, 1, 1, 1, 1, 0, 1}
		want := 5
		got := TaskFour(a, 4)
		assertElements(t, got, want)
	})

	t.Run("TaskFive", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		m := 7
		want := 3
		got := TaskFive(a, m)
		assertElements(t, got, want)
	})
}

func assertSlices[T comparable](t *testing.T, got, want []T) {
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got: %v, want: %v", got, want)
			return
		}
	}
}
func assertElements[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
