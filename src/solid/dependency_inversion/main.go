package main

import "database/sql"

type User struct {
	Name      string
	CPF       string
	Birthdate string
}

type Service interface {
	CreateUser(user User) error
	GetUserByCPF(cpf string) (*User, error)
}

type Repository interface {
	CreateUser(user User) error
	GetUserByCPF(cpf string) (*User, error)
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{repo: repo}
}

func (instance *UserService) CreateUser(user User) error {
	return instance.repo.CreateUser(user)
}

func (instance *UserService) GetUserByCPF(cpf string) (*User, error) {
	return instance.repo.GetUserByCPF(cpf)
}

type Database interface {
	GetConnection() (*sql.Conn, error)
	CloseConnection(conn *sql.Conn) error
}

type Sqlite struct{}

func (instance *Sqlite) GetConnection() (*sql.Conn, error) {
	return &sql.Conn{}, nil
}

func (instance *Sqlite) CloseConnection(conn *sql.Conn) error {
	// conn.Close()
	return nil
}

type UserRepository struct {
	db Database
}

func NewUserRepository(db Database) *UserRepository {
	return &UserRepository{db: db}
}

func (instance *UserRepository) CreateUser(user User) error {
	conn, err := instance.db.GetConnection()
	if err != nil {
		return err
	}

	defer instance.db.CloseConnection(conn)

	return nil
}

func (instance *UserRepository) GetUserByCPF(cpf string) (*User, error) {
	conn, err := instance.db.GetConnection()
	if err != nil {
		return nil, err
	}

	defer instance.db.CloseConnection(conn)

	return &User{"Eduardo", "000.000.000-00", "08/03/2003"}, nil
}

func main() {
	sqlite := &Sqlite{}
	repo := NewUserRepository(sqlite)
	service := NewUserService(repo)
	service.CreateUser(User{})
	user, err := service.GetUserByCPF("12345678910")

	if err != nil {
		panic(err)
	}

	println(user.Name)
}
