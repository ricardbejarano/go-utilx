package slicex

func Head[T any](slice []T, length int) []T {
	if len(slice) > length {
		return slice[:length]
	}

	return slice
}

func Contains[T comparable](item T, slice []T) bool {
	for i := range slice {
		if slice[i] == item {
			return true
		}
	}
	return false
}
