package main

import "fmt"

func minMax(arr []int) {
	l := len(arr)

	for index := range arr {
		for j := 0; j < l-index-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr[0], arr[l-1])
}

func main() {
	arr := []int{10, 20, 5, 50, 2}

	minMax(arr)
}
