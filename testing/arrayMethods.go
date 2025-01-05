package dbcontext_test

func indexOf[T any](arr []T, compare func(value T) bool) int {
	for i, value := range arr {
		if compare(value) {
			return i
		}
	}

	return -1
}

func some[T any](arr []T, compare func(value T) bool) bool {
	return indexOf(arr, compare) != -1
}

func mapArr[T any, R any](arr []T, mapFunc func(value T) R) []R {
	response := make([]R, len(arr))
	for _, value := range arr {
		response = append(response, mapFunc(value))
	}

	return response
}

func equals[T any](firstArr, secondArr []T, compare func(first, second T) bool) bool {
	if len(firstArr) != len(secondArr) {
		return false
	}

	for _, value := range firstArr {
		if !some(secondArr, func(secondValue T) bool { return compare(value, secondValue) }) {
			return false
		}
	}

	for _, value := range secondArr {
		if !some(firstArr, func(firstValue T) bool { return compare(value, firstValue) }) {
			return false
		}
	}

	return true
}
