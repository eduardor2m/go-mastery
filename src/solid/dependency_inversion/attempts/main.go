package main

import (
	"fmt"
)

type Student struct {
	id        string
	name      string
	birthdate string
}

var students []Student

type Database interface {
	getConnection() string
	closeConnection() string
}

type MySQL struct{}

func (m *MySQL) getConnection() string {
	return "MySQL connection"
}

func (m *MySQL) closeConnection() string {
	return "MySQL connection closed"
}

func NewMySQL() *MySQL {
	return &MySQL{}
}

type StudentMySQLRepository interface {
	GetAll() []Student
	GetById(id string) Student
	Create(student Student) Student
}

type StudentMySQLRepositoryImpl struct {
	db Database
}

func NewStudentMySQLRepository(db Database) *StudentMySQLRepositoryImpl {
	return &StudentMySQLRepositoryImpl{db: db}
}

func (s *StudentMySQLRepositoryImpl) GetAll() []Student {
	s.db.getConnection()
	defer s.db.closeConnection()
	return students
}

func (s *StudentMySQLRepositoryImpl) GetById(id string) Student {
	s.db.getConnection()
	var student Student

	for _, s := range students {
		if s.id == id {
			student = s
		}
	}
	defer s.db.closeConnection()
	return student
}

func (s *StudentMySQLRepositoryImpl) Create(student Student) Student {
	s.db.getConnection()
	students = append(students, student)
	defer s.db.closeConnection()
	return student
}

type StudentService interface {
	GetAll() []Student
	GetById(id string) Student
	Create(student Student) Student
}

type StudentServiceImpl struct {
	repository StudentMySQLRepository
}

func (s *StudentServiceImpl) GetAll() []Student {
	return s.repository.GetAll()
}

func (s *StudentServiceImpl) GetById(id string) Student {
	return s.repository.GetById(id)
}

func (s *StudentServiceImpl) Create(student Student) Student {
	return s.repository.Create(student)
}

func NewStudentService(repository StudentMySQLRepository) *StudentServiceImpl {
	return &StudentServiceImpl{repository: repository}
}

func main() {
	db := NewMySQL()
	repository := NewStudentMySQLRepository(db)
	service := NewStudentService(repository)
	s := Student{id: "1", name: "John", birthdate: "01/01/2000"}
	service.Create(s)
	searchedStudent := service.GetById("1")
	fmt.Println("Nome do aluno: ", searchedStudent.name)
	students := service.GetAll()
	for _, student := range students {
		fmt.Println("Nome do aluno: ", student.name)
		fmt.Println("Data de nascimento do aluno: ", student.birthdate)
		fmt.Println("ID do aluno: ", student.id)
	}
}
