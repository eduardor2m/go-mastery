package main

import "fmt"

// Observer interface que os observadores devem implementar.
type Observer interface {
	Update(message string)
}

// Subject interface que o sujeito (objeto observável) deve implementar.
type Subject interface {
	RegisterObserver(observer Observer)
	NotifyObservers(message string)
}

// ConcreteSubject implementação concreta de Subject.
type ConcreteSubject struct {
	observers []Observer
}

func (cs *ConcreteSubject) RegisterObserver(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) NotifyObservers(message string) {
	for _, observer := range cs.observers {
		observer.Update(message)
	}
}

// ConcreteObserver implementação concreta de Observer.
type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("[%s] Recebeu uma atualização: %s\n", co.name, message)
}

func main() {
	// Criando um sujeito
	subject := &ConcreteSubject{}

	// Criando observadores
	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}

	// Registrando observadores no sujeito
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	// Notificando os observadores sobre uma mudança
	subject.NotifyObservers("Nova atualização no sujeito!")
}
