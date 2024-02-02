package main

import "fmt"

// Interface para definir os algoritmos
type Strategy interface {
	Execute() string
}

// Implementações específicas dos algoritmos
type ConcreteStrategyA struct{}

func (a *ConcreteStrategyA) Execute() string {
	return "Executando estratégia A"
}

type ConcreteStrategyB struct{}

func (b *ConcreteStrategyB) Execute() string {
	return "Executando estratégia B"
}

// Contexto que usa a estratégia
type Context struct {
	strategy Strategy
}

// Método para executar a estratégia atual
func (c *Context) ExecuteStrategy() string {
	return c.strategy.Execute()
}

// Método para definir a estratégia atual
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func main() {
	// Cria instâncias das estratégias
	strategyA := &ConcreteStrategyA{}
	strategyB := &ConcreteStrategyB{}

	// Cria o contexto com a estratégia A
	context := &Context{strategy: strategyA}

	// Executa a estratégia atual
	resultA := context.ExecuteStrategy()
	fmt.Println(resultA) // Saída: Executando estratégia A

	// Muda a estratégia para B
	context.SetStrategy(strategyB)

	// Executa a nova estratégia
	resultB := context.ExecuteStrategy()
	fmt.Println(resultB) // Saída: Executando estratégia B
}
