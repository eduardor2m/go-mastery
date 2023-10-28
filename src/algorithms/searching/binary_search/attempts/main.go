package main

import "fmt"

func binarySearch(arr []int, i int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == i {
			return mid
		} else if arr[mid] < i {
			left = mid + 1
		} else if arr[mid] > i {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{10, 5, 26, 17, 1}

	result := binarySearch(arr, 10)

	fmt.Println(result)
}
