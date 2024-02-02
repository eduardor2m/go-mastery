package main

import (
	"fmt"
	"log"
)

type Book struct {
	Title  string
	Author string
}

var books []Book

type DatabaseSQL interface {
	getConnection() (*string, error)
	closeConnection() error
}

type Postgres struct{}

func (p Postgres) getConnection() (*string, error) {
	log.Println("database connected")
	str := new(string)
	return str, nil
}

func (p Postgres) closeConnection() error {
	log.Println("database closed")
	return nil
}

func NewDatabase() (DatabaseSQL, error) {
	return Postgres{}, nil
}

type BookRepositoryInterface interface {
	create(Book) error
	get(string) (Book, error)
	list() ([]Book, error)
	delete(string) error
}

type BookRepository struct {
	db DatabaseSQL
}

func (b BookRepository) create(book Book) error {
	_, err := b.db.getConnection()
	if err != nil {
		return err
	}

	books = append(books, book)

	defer func() {
		err = b.db.closeConnection()
		if err != nil {
			log.Println(err)
		}

	}()

	log.Println("book created")

	return nil
}

func (b BookRepository) get(title string) (Book, error) {
	_, err := b.db.getConnection()
	if err != nil {
		return Book{}, err
	}

	var book Book

	for _, item := range books {
		if book.Title == title {
			book = item
			break
		}
	}

	defer func() {
		err = b.db.closeConnection()
		if err != nil {
			log.Println(err)
		}

	}()

	log.Println("book get")

	return book, nil
}

func (b BookRepository) list() ([]Book, error) {
	_, err := b.db.getConnection()
	if err != nil {
		return []Book{}, err
	}
	defer func() {
		err = b.db.closeConnection()
		if err != nil {
			log.Println(err)
		}

	}()

	log.Println("book list")

	return books, nil
}

func (b BookRepository) delete(title string) error {
	_, err := b.db.getConnection()
	if err != nil {
		return err
	}

	for i, item := range books {
		if item.Title == title {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	defer func() {
		err = b.db.closeConnection()
		if err != nil {
			log.Println(err)
		}

	}()

	log.Println("book deleted")

	return nil
}

func NewBookRepository(db DatabaseSQL) BookRepositoryInterface {
	return BookRepository{db}
}

type BookServiceInterface interface {
	create(Book) error
	get(string) (Book, error)
	list() ([]Book, error)
	delete(string) error
}

type BookService struct {
	repository BookRepositoryInterface
}

func (b BookService) create(book Book) error {
	return b.repository.create(book)
}

func (b BookService) get(title string) (Book, error) {
	return b.repository.get(title)
}

func (b BookService) list() ([]Book, error) {
	return b.repository.list()
}

func (b BookService) delete(title string) error {
	return b.repository.delete(title)
}

func NewBookService(repository BookRepositoryInterface) BookServiceInterface {
	return BookService{repository}
}

func init() {
	log.Printf("init() from %s\n", "src/archive/study/main.go")
}
func main() {
	log.Printf("main() from %s\n", "src/archive/study/main.go")

	db, err := NewDatabase()
	if err != nil {
		log.Panic(err)
	}

	repository := NewBookRepository(db)
	service := NewBookService(repository)

	for {
		fmt.Println(" "+
			"1 - Create book\n",
			"2 - Get book\n",
			"3 - List books\n",
			"4 - Delete book\n",
			"5 - Exit\n",
		)
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Panic(err)
		}
		switch {
		case option == 1:
			var title, author string
			fmt.Println("Title: ")
			_, err := fmt.Scan(&title)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println("Author: ")
			_, err = fmt.Scan(&author)
			if err != nil {
				log.Panic(err)
			}
			err = service.create(Book{title, author})
			if err != nil {
				log.Println(err)
			}
		case option == 2:
			var title string
			fmt.Println("Title: ")
			_, err := fmt.Scan(&title)
			if err != nil {
				log.Panic(err)
			}
			book, err := service.get(title)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(book)

		}
	}
}
