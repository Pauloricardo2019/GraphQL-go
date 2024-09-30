package service

import (
	"errors"
	"github.com/Pauloricardo2019/graphql-teste/internal/domain"
)

type Service struct {
	users map[string]*domain.User
}

func NewService() *Service {
	return &Service{
		users: make(map[string]*domain.User),
	}
}

// GetUser busca um usuário pelo ID
func (s *Service) GetUser(id string) (*domain.User, error) {
	if user, ok := s.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

// CreateUser cria um novo usuário
func (s *Service) CreateUser(id string, name string, age int) (*domain.User, error) {
	if _, exists := s.users[id]; exists {
		return nil, errors.New("user already exists")
	}

	user := &domain.User{
		ID:   id,
		Name: name,
		Age:  age,
	}

	s.users[id] = user
	return user, nil
}

// UpdateUser atualiza um usuário existente
func (s *Service) UpdateUser(id string, name string, age int) (*domain.User, error) {
	if user, exists := s.users[id]; exists {
		user.Name = name
		user.Age = age
		return user, nil
	}
	return nil, errors.New("user not found")
}

// DeleteUser deleta um usuário pelo ID
func (s *Service) DeleteUser(id string) error {
	if _, exists := s.users[id]; exists {
		delete(s.users, id)
		return nil
	}
	return errors.New("user not found")
}
