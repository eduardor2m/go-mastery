package main

import (
	"sync"
)

func CalcGoRoutine(qtdTimes, value int) (result int) {
	var wg sync.WaitGroup

	for i := 0; i < qtdTimes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result += Fibonnaci(value)
		}()
	}

	return
}

func Calc(qtdTimes, value int) (result int) {
	for i := 0; i < qtdTimes; i++ {
		result += Fibonnaci(value)
	}

	return
}

func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonnaci(n-1) + Fibonnaci(n-2)
}
