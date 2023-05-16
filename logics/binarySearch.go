package logics

func BinarySearch(needle int, arr []int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != needle {
		return false
	}

	return true
}
