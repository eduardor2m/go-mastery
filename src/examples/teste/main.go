package main

import (
	"log"
)

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
	return nil
}

func NewDatabase() (DatabaseSQL, error) {
	return Postgres{}, nil
}

func main() {
	db, err := NewDatabase()
	if err != nil {
		log.Panic(err)
	}

	_, _ = db.getConnection()
}
