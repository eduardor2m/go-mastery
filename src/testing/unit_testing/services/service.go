package services

import "github.com/eduardor2m/go-mastery/src/testing/unit_testing/interfaces"

type Service struct {
	repository interfaces.Repository
}

func (s *Service) Create(str string) error {
	return s.repository.Create(str)
}

func (s *Service) GetAll() []string {
	return s.repository.GetAll()
}

func NewService(repository interfaces.Repository) *Service {
	return &Service{repository}
}
