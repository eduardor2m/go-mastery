package main

import "fmt"

func bubbleSort(list []int) {
	l := len(list)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// A complexidade do algoritmo de ordenação bubble sort é O(n²), onde "n" é o
// número de elementos na lista. Isso significa que o tempo de execução do
// algoritmo cresce quadráticamente com o tamanho da lista.

func main() {
	list := []int{5, 4, 3, 2, 1}

	bubbleSort(list)

	fmt.Println(list)
}
