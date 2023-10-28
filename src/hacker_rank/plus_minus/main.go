package main

import "fmt"

func plusMinus(arr []int32) {
	var negative float32 = 0
	var positive float32 = 0
	var zeros float32 = 0
	l := float32(len(arr))

	for _, v := range arr {
		if v > 0 {
			positive += 1
		} else if v < 0 {
			negative += 1
		} else {
			zeros += 1
		}
	}

	fmt.Printf("%.6f\n%.6f\n%.6f", positive/l, negative/l, zeros/l)
}

func main() {
	arr := []int32{-6, 10, 3, -4, 0, 0}

	plusMinus(arr)
}
