package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Unit struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (u *Unit) ConvertCelciusToFahrenheit() (string, error) {
	celcius := strings.TrimSuffix(u.From, "°C")
	celciusToNumber, err := strconv.ParseFloat(celcius, 64)
	if err != nil {
		return "", fmt.Errorf("erro ao converter string para float: %v", err)
	}

	return fmt.Sprintf("%.2f °F", (celciusToNumber*9/5)+32), nil
}

func main() {
	celcius := "30°C"

	unit := &Unit{
		From: celcius,
	}

	expr, err := unit.ConvertCelciusToFahrenheit()
	if err != nil {
		fmt.Println(err)
		return
	}

	println(expr)
}
