package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	n := len(arr)
	var primaryDiagonalSum, secondaryDiagonalSum int32

	for i := 0; i < n; i++ {
		primaryDiagonalSum += arr[i][i]
	}

	for i := 0; i < n; i++ {
		secondaryDiagonalSum += arr[i][n-i-1]
	}

	difference := int32(math.Abs(float64(primaryDiagonalSum - secondaryDiagonalSum)))

	return difference

}

func main() {
	matrix := [][]int32{{
		10, 20, 30,
	}, {
		10, 20, 30,
	}, {
		20, 30, 100,
	}}

	result := diagonalDifference(matrix)

	fmt.Println(result)

}
