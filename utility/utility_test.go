package utility

import (
	"testing"
)

func TestSliceIndexOfHappy(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	want := 3

	got := SliceIndexOf(slice, func(item int) bool {
		return item == 4
	})

	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}

func TestSliceIndexOfNoSuchEntry(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	want := -1

	got := SliceIndexOf(slice, func(item int) bool {
		return item == 12
	})

	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}
