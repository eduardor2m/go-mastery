package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func (u User) show() {
	fmt.Printf("Meu nome é %s e tenho %d anos", u.name, u.age)
}

type Student struct {
	User
	classroom string
	semester  int
}

func (s Student) show() {
	fmt.Printf("Meu nome é %s e tenho %d anos, estudo na escola %s e estou no %d semestre\n", s.name, s.age, s.classroom, s.semester)
}

func main() {
	u := User{"dudu", 20}
	std := Student{User{"dudu", 20}, "teste", 5}
	std.show()
	u.show()
}
