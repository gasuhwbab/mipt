package second

import (
	"fmt"
	"testing"
)

func TestTasks(t *testing.T) {
	t.Run("TaskOne", func(t *testing.T) {
		a := []int{1, 17, 2, 6, 3}
		want := 4
		got := TaskOne(a)
		fmt.Println(a)
		assertElements(t, got, want)
	})
}

//	func assertSlices[T comparable](t *testing.T, got, want []T) {
//		for i := range got {
//			if got[i] != want[i] {
//				t.Errorf("got: %v, want: %v", got, want)
//				return
//			}
//		}
//	}
func assertElements[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
