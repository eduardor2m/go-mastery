package main

import "fmt"

// Componente de Café (Interface)
type Coffee interface {
	GetCost() int
	GetDescription() string
}

// Implementação Básica do Café
type SimpleCoffee struct{}

func (c *SimpleCoffee) GetCost() int {
	return 5
}

func (c *SimpleCoffee) GetDescription() string {
	return "Café simples"
}

// Decorador Base (Abstração)
type CoffeeDecorator struct {
	CoffeeComponent Coffee
}

func (cd *CoffeeDecorator) GetCost() int {
	return cd.CoffeeComponent.GetCost()
}

func (cd *CoffeeDecorator) GetDescription() string {
	return cd.CoffeeComponent.GetDescription()
}

// Decorador de Leite
type MilkDecorator struct {
	CoffeeDecorator
}

func (md *MilkDecorator) GetCost() int {
	return md.CoffeeDecorator.GetCost() + 2
}

func (md *MilkDecorator) GetDescription() string {
	return md.CoffeeDecorator.GetDescription() + ", Leite"
}

// Decorador de Açúcar
type SugarDecorator struct {
	CoffeeDecorator
}

func (sd *SugarDecorator) GetCost() int {
	return sd.CoffeeDecorator.GetCost() + 1
}

func (sd *SugarDecorator) GetDescription() string {
	return sd.CoffeeDecorator.GetDescription() + ", Açúcar"
}

func main() {
	// Criando um café simples
	simpleCoffee := &SimpleCoffee{}
	fmt.Printf("Descrição: %s, Custo: $%d\n", simpleCoffee.GetDescription(), simpleCoffee.GetCost())

	// Adicionando leite ao café
	milkCoffee := &MilkDecorator{CoffeeDecorator: CoffeeDecorator{CoffeeComponent: simpleCoffee}}
	fmt.Printf("Descrição: %s, Custo: $%d\n", milkCoffee.GetDescription(), milkCoffee.GetCost())

	// Adicionando açúcar ao café
	sugarCoffee := &SugarDecorator{CoffeeDecorator: CoffeeDecorator{CoffeeComponent: simpleCoffee}}
	fmt.Printf("Descrição: %s, Custo: $%d\n", sugarCoffee.GetDescription(), sugarCoffee.GetCost())

	// Adicionando leite e açúcar ao café
	milkAndSugarCoffee := &SugarDecorator{CoffeeDecorator: CoffeeDecorator{CoffeeComponent: milkCoffee}}
	fmt.Printf("Descrição: %s, Custo: $%d\n", milkAndSugarCoffee.GetDescription(), milkAndSugarCoffee.GetCost())
}
