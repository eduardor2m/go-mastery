package main

import "fmt"

// func lonelyInteger(a []int32) int32 {
// 	var unique int32

// 	for _, value := range a {
// 		unique ^= value
// 	}

// 	return unique

// }

func lonelyInteger(a []int32) int32 {
	occurrence := make(map[int32]int32)

	for _, num := range a {
		occurrence[num]++
	}

	for key, value := range occurrence {
		if value == 1 {
			return key
		}
	}

	return -1
}

func main() {
	arr := []int32{10, 20, 10, 30, 20}

	result := lonelyInteger(arr)

	fmt.Println(result)
}
