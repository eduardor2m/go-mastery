package main

import (
	"fmt"
	"log"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

type Book struct {
	Title  string
	Author string
}

var books []Book

func (b Book) Create(wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()

	books = append(books, b)

}

func (b Book) GetQuantity() int {
	return len(books)
}

func main() {
	var book Book
	for i := 0; i < 10; i++ {
		wg.Add(1)
		book = Book{
			Title:  fmt.Sprintf("Clean Architecture %d", i),
			Author: "Robert C. Martin",
		}

		go book.Create(&wg)
	}
	wg.Wait()

	log.Println("Total books:", book.GetQuantity())

	for _, b := range books {
		log.Println(b)
	}

}
