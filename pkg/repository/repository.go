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

type Options interface {
	Create(parentId int, option taskexchange.Option) (int, error)
	GetAll() ([]taskexchange.Option, error)
	GetById(id int) (taskexchange.Option, error)
	Update(id int, input taskexchange.UpdateOptionInput) error
	Delete(id int) error
}

type Events interface {
	Create(event taskexchange.Event) (int, error)
	FindPolling(userId, id int) ([]taskexchange.Event, error)
	FindNew(userId int) ([]taskexchange.Event, error)
	FindLastUser(userId int) (taskexchange.Event, error)
	View(userId, id int) error
	Delete(userId, id int) error
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
	Events
	Options
	Task
	TaskOption
	Offer
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users:   NewUsersPostgres(db),
		Options: NewOptionsPostgres(db),
		Events:  NewEventsPostgres(db),
	}
}
