package main

type List struct {
	Head *Node
	Tail *Node
}

type Person struct {
	Name string
	Age  int
}

func (l *List) Append(p Person) {
	node := &Node{Value: p}

	if l.Head == nil {
		l.Head = node
	}

	if l.Tail != nil {
		l.Tail.Next = node
	}

	l.Tail = node
}

func (l *List) Display() {
	current := l.Head

	for current != nil {
		println(current.Value.Name)
		current = current.Next
	}
}

type Node struct {
	Value Person
	Next  *Node
}

func main() {
	list := &List{}

	list.Append(Person{Name: "John", Age: 30})
	list.Append(Person{Name: "Jane", Age: 25})
	list.Append(Person{Name: "Jack", Age: 40})
	list.Append(Person{Name: "Jill", Age: 35})

	list.Display()
}
