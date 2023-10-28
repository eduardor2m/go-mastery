package main

import "fmt"

func linearSearch(arr []int, i int) int {
	for index, item := range arr {
		if item == i {
			return index
		}
	}

	return -1
}

func main() {
	arr := []int{10, 15, 1, 2, 16, 9}

	result := linearSearch(arr, 3)

	if result == -1 {
		fmt.Printf("Número não encontrado, resultado = %d", result)
	} else {
		fmt.Printf("Número encontrado, resultado = %d", result)
	}
}
