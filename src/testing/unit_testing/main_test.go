package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(2, 3)
	if resultado != 5 {
		t.Errorf("Soma(2, 3) = %d; esperado 5", resultado)
	}
}
