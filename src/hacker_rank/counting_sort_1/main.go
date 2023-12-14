package main

import "fmt"

func countingSort(arr []int32) []int32 {
	frequencyArray := make([]int32, 100)

	for _, num := range arr {
		frequencyArray[num]++
	}

	return frequencyArray
}

func main() {
	arr := []int32{10, 10, 20, 30, 40}

	result := countingSort(arr)

	for index, value := range result {
		fmt.Println("Index: ", index, " Valor: ", value)
	}

}
