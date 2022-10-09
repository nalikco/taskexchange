package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
)

type Authorization interface {
	CreateUser(user taskexchange.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, int, error)
	UpdateOnline(id int) error
}

type Users interface {
	CreateUser(user taskexchange.User) (int, error)
	GetAll() ([]taskexchange.User, error)
	GetById(id int, full bool) (taskexchange.User, error)
	Update(id int, input taskexchange.UpdateUserInput) error
	Delete(id int) error
}

type Option interface {
	Create(parentId int, option taskexchange.Option) (int, error)
	GetAll() ([]taskexchange.Option, error)
	GetById(id int) (taskexchange.Option, error)
	Update(id int, input taskexchange.UpdateOptionInput) error
	Delete(id int) error
}

type Task interface {
}

type TaskOption interface {
}

type Offer interface {
}

type Order interface {
}

type Service struct {
	Authorization
	Users
	Option
	Task
	TaskOption
	Offer
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Users),
		Users:         NewUsersService(repos.Users),
		Option:        NewOptionService(repos.Option),
	}
}
