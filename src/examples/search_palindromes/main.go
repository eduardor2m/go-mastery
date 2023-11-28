package main

import "fmt"

func isPalindrome(word string) bool {
	chars := []rune(word)
	length := len(chars)

	for i := 0; i < length/2; i++ {
		if chars[i] != chars[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	word := "asa"
	// word := "teste"

	fmt.Println(isPalindrome(word))
}
