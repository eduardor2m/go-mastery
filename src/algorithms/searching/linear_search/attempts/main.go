package main

import "fmt"

func linearSearch(arr []int, item int) int {
	for index, value := range arr {
		if item == value {
			return index
		}
	}

	return -1
}

func main() {
	arr := []int{10, 20, 30}

	position := linearSearch(arr, -10)

	fmt.Println(position)
}
