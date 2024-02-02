package main

import (
	"fmt"
)

// 1. Single Responsibility Principle (SRP)
type Logger interface {
	Log(message string)
}

type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	// Implementação para gravar em um arquivo
	fmt.Println("File Log:", message)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	// Implementação para exibir no console
	fmt.Println("Console Log:", message)
}

// 2. Open/Closed Principle (OCP)
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 3. Liskov Substitution Principle (LSP)
func printArea(shape Shape) {
	fmt.Printf("Shape Area: %f\n", shape.Area())
}

// 4. Interface Segregation Principle (ISP)
type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

type Human struct{}

func (h Human) Work() {
	fmt.Println("Working...")
}

func (h Human) Eat() {
	fmt.Println("Eating...")
}

// 5. Dependency Inversion Principle (DIP)
type Notifier interface {
	Notify(message string)
}

type AlertSystem struct {
	logger   Logger
	notifier Notifier
}

func (a *AlertSystem) RaiseAlert(message string) {
	a.logger.Log(message)
	a.notifier.Notify(message)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(message string) {
	fmt.Println("Sending email notification:", message)
}

func main() {
	// SRP
	fileLogger := &FileLogger{}
	fileLogger.Log("File Logger Message")

	consoleLogger := &ConsoleLogger{}
	consoleLogger.Log("Console Logger Message")

	// OCP
	rectangle := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 3}

	printArea(rectangle)
	printArea(circle)

	// LSP
	printArea(rectangle) // LSP is implicitly applied here

	// ISP
	human := Human{}
	human.Work()
	human.Eat()

	// DIP
	emailNotifier := &EmailNotifier{}
	alertSystem := &AlertSystem{
		logger:   consoleLogger,
		notifier: emailNotifier,
	}

	alertSystem.RaiseAlert("Critical Alert!")
}
