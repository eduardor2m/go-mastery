package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// Método associado ao tipo Pessoa
func (p Pessoa) Apresentar() {
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos.\n", p.Nome, p.Idade)
}

type Estudante struct {
	Pessoa
	Curso string
}

// Método associado ao tipo Estudante

func (e Estudante) Apresentar() {
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos. Estou no curso de %s.\n", e.Nome, e.Idade, e.Curso)
}

func main() {
	pessoa := Pessoa{"Alice", 30}
	pessoa.Apresentar()
	// Olá, meu nome é Alice e eu tenho 30 anos.

	estudante := Estudante{Pessoa{"Eduardo", 20}, "Sistemas de Informação"}
	estudante.Apresentar()
	// Olá, meu nome é Eduardo e eu tenho 20 anos. Estou no curso de Sistemas de Informação.
}
