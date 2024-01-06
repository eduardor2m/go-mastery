package main

import "fmt"

// model

type Todo struct {
	id          string
	title       string
	description string
	completed   bool
}

type Builder struct {
	todo Todo
	err  error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Id(id string) *Builder {
	b.todo.id = id
	return b
}

func (b *Builder) Title(title string) *Builder {
	b.todo.title = title
	return b
}

func (b *Builder) Description(description string) *Builder {
	b.todo.description = description
	b.err = fmt.Errorf("error")
	return b
}

func (b *Builder) Completed(completed bool) *Builder {
	b.todo.completed = completed
	return b
}

func (b *Builder) Build() (Todo, error) {
	return b.todo, b.err
}

// controller

type TodoController interface {
	GetAll() []Todo
	GetById(id string) Todo
	Create(todo Todo) Todo
}

type TodoControllerImpl struct {
	todos []Todo
}

func (t *TodoControllerImpl) GetAll() []Todo {
	return t.todos
}

func (t *TodoControllerImpl) GetById(id string) Todo {
	for _, todo := range t.todos {
		if todo.id == id {
			return todo
		}
	}
	return Todo{}
}

func (t *TodoControllerImpl) Create(todo Todo) Todo {
	t.todos = append(t.todos, todo)
	return todo
}

// view

type TodoView interface {
	Render(todos []Todo)
}

type TodoViewImpl struct{}

func (t *TodoViewImpl) Render(todos []Todo) {
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

// main

func main() {
	controller := TodoControllerImpl{}
	view := TodoViewImpl{}
	t, err := NewBuilder().Id("1").Title("title").Description("description").Completed(false).Build()
	if err != nil {
		panic(err)
	}
	controller.Create(t)
	view.Render(controller.GetAll())
}
