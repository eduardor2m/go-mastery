package main

func linearSearch(list []int, elementToFind int) int {
	for index, item := range list {
		if item == elementToFind {
			return index
		}
	}

	return -1
}

// A complexidade do algoritmo de busca linear é O(n), onde "n" é o número de
// elementos na lista. Isso significa que o tempo de execução do algoritmo
// cresce linearmente com o tamanho da lista.

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	elementToFind := 7

	index := linearSearch(list, elementToFind)

	if index == -1 {
		println("Element not found")
	} else {
		println("Element found at index", index)
	}
}
