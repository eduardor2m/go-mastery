package main

import "fmt"

func miniMaxSum(arr []int32) {
	for i := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	var min int64 = int64(arr[0]) + int64(arr[1]) + int64(arr[2]) + int64(arr[3])
	var max int64 = int64(arr[1]) + int64(arr[2]) + int64(arr[3]) + int64(arr[4])

	fmt.Printf("%d %d", min, max)
}

func main() {
	arr := []int32{396285104, 573261094, 759641832, 819230764, 364801279}

	miniMaxSum(arr)
}
