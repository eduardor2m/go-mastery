package interfaces

type Repository interface {
	Create(string) error
	GetAll() []string
}
