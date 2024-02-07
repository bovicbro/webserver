package utility

// Returns index of the first element of the array for which f returns true.
// If there is no element for which f returns true, -1 will be returned.
func SliceIndexOf[E any](slice []E, f func(E) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
}
