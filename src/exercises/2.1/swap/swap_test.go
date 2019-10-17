package swap

import "testing"

const (
	resultA = 1
	resultB = 2
)

func TestExercise(t *testing.T) {
	a, b := 1, 2

	a, b = swap(a, b)

	if !(a != resultA && b != resultB) {
		t.Error("Should be true")
	}

	swap2(&a, &b)
	if !(a == resultA && b == resultB) {
		t.Error("Should be true")
	}
}
