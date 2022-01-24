package main

import (
	//bs "algorithms-go/binarySearch"

	ss "algorithms-go/selectionSort"
	"fmt"
)

func main() {
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//fmt.Println(bs.BinarySearch(arr, 16))

	arr := []int{1, 7, 11, 4, 3, 6, 2, 8, 14, 10, 5, 15, 13, 9, 12}
	fmt.Println(ss.SelectionSort(arr))

}
