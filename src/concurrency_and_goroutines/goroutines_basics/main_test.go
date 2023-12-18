package main

import "testing"

func BenchmarkCalc(b *testing.B) {
	Calc(b.N, 30)
}

func BenchmarkCalcGoRoutine(b *testing.B) {
	CalcGoRoutine(b.N, 30)
}
