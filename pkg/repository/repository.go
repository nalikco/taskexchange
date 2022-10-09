package repository

import (
	"github.com/jmoiron/sqlx"
	"taskexchange"
)

type Users interface {
	Create(user taskexchange.User) (int, error)
	GetAll() ([]taskexchange.User, error)
	GetById(id int, full bool) (taskexchange.User, error)
	GetByEmail(email string) (taskexchange.User, error)
	GetByEmailAndPassword(email, password string) (taskexchange.User, error)
	Update(id int, input taskexchange.UpdateUserInput) error
	Delete(id int) error
	UpdateOnline(id int) error
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

type Repository struct {
	Users
	Option
	Task
	TaskOption
	Offer
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users:  NewUsersPostgres(db),
		Option: NewOptionPostgres(db),
	}
}
