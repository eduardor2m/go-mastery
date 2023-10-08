package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct{}

func (instance SQLite) Connect() (*sql.Conn, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = conn.ExecContext(context.Background(), "CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, name TEXT, email TEXT, password TEXT, birth_date TEXT)")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (instance SQLite) Disconnect(conn *sql.Conn) error {
	err := conn.Close()
	if err != nil {
		return err
	}

	return nil
}
