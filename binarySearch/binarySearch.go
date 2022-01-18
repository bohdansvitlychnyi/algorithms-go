package binarySearch

func BinarySearch(array []int, item int) int {
	var low int = 0
	var high int = len(array) - 1
	var mid int
	var guess int

	for low <= high {
		mid = high - low
		guess = array[mid]
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
