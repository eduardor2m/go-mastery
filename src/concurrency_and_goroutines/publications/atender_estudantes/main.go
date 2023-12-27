package main

import (
	"fmt"
	"sync"
	"time"
)

// Simula o atendimento a um estudante na fila
func atenderEstudante(nome string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s está comprando um ingresso para o evento...\n", nome)
	// Simula o tempo de atendimento
	time.Sleep(2 * time.Second)
	fmt.Printf("%s comprou o ingresso!\n", nome)
}

func main() {
	// Lista de estudantes na fila
	estudantes := []string{"Eduardo", "Bob", "Charlie", "David", "Eve"}

	// WaitGroup para esperar que todas as goroutines terminem
	var wg sync.WaitGroup

	// Inicia o atendimento para cada estudante na fila
	for _, estudante := range estudantes {
		wg.Add(1)
		go atenderEstudante(estudante, &wg)
	}

	// Aguarda todas as goroutines terminarem
	wg.Wait()

	fmt.Println("Todos os ingressos foram vendidos. O evento está pronto para começar!")
}
