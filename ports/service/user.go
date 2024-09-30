package service

import "github.com/Pauloricardo2019/graphql-teste/internal/domain"

type UserServiceIF interface {
	GetUser(id string) (*domain.User, error)
	CreateUser(id string, name string, age int) (*domain.User, error)
	UpdateUser(id string, name string, age int) (*domain.User, error)
	DeleteUser(id string) error
}
