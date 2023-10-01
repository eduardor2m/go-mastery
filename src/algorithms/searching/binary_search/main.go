package main

func binarySearch(list []int, target int) int {
	left, right := 0, len(list)-1

	for left <= right {
		mid := (left + right) / 2

		if list[mid] == target {
			return mid
		} else if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

// A complexidade do algoritmo de busca binária é O(log n), onde "n" é o número
// de elementos na lista. Isso significa que o tempo de execução do algoritmo
// cresce logaritmicamente com o tamanho da lista.

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	target := 7

	index := binarySearch(list, target)

	if index == -1 {
		println("Element not found")
	} else {
		println("Element found at index", index)
	}
}
