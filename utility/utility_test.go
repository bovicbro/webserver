package utility

import (
	"testing"
)

// TestSliceIndexOfHappy tests SliceIndexOf with a matching element
func TestSliceIndexOfHappy(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4}
	want := 3

	// Act
	got := SliceIndexOf(slice, func(item int) bool {
		return item == 4
	})

	// Assert
	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}

// TestSliceIndexOfNoSuchEntry tests SliceIndexOf when no element matches
func TestSliceIndexOfNoSuchEntry(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4}
	want := -1

	// Act
	got := SliceIndexOf(slice, func(item int) bool {
		return item == 12
	})

	// Assert
	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}

// TestSliceIndexOfEmptySlice tests SliceIndexOf with empty slice
func TestSliceIndexOfEmptySlice(t *testing.T) {
	// Arrange
	slice := []int{}
	want := -1

	// Act
	got := SliceIndexOf(slice, func(item int) bool {
		return item == 1
	})

	// Assert
	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}

// TestSliceIndexOfFirstElement tests SliceIndexOf returns first matching element
func TestSliceIndexOfFirstElement(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 2, 4}
	want := 1

	// Act
	got := SliceIndexOf(slice, func(item int) bool {
		return item == 2
	})

	// Assert
	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}

// TestSliceIndexOfWithStrings tests SliceIndexOf with string slice
func TestSliceIndexOfWithStrings(t *testing.T) {
	// Arrange
	slice := []string{"apple", "banana", "cherry"}
	want := 1

	// Act
	got := SliceIndexOf(slice, func(item string) bool {
		return item == "banana"
	})

	// Assert
	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}

// TestSliceIndexOfWithComplexPredicate tests SliceIndexOf with complex condition
func TestSliceIndexOfWithComplexPredicate(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4, 5}
	want := 2

	// Act
	got := SliceIndexOf(slice, func(item int) bool {
		return item > 2 && item < 4
	})

	// Assert
	if got != want {
		t.Errorf("Got %d, expected %d", got, want)
	}
}
