package selectionSort

func SelectionSort(array []int) []int {
	var l int = len(array)
	sorted := make([]int, l)

	for i := 0; i < l; i++ {
		smallest := findSmallest(array)
		sorted[i] = array[smallest]
		array = append(array[:smallest], array[smallest+1:]...)
	}
	return sorted
}

func findSmallest(array []int) int {
	var smallest int = array[0]
	var smallestIndex int = 0

	for index, value := range array {
		if value < smallest {
			smallest = value
			smallestIndex = index
		}
	}
	return smallestIndex
}
