package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	for index := range arr {
		for i := 0; i < len(arr)-1-index; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return arr
}

func main() {
	arr := []int{10, 5, 20, -10, 3}

	arrF := bubbleSort(arr)

	fmt.Println(arrF)
}
