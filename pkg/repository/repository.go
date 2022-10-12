package repository

import (
	"taskexchange"

	"github.com/jmoiron/sqlx"
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
	GetByIds(ids []int) ([]taskexchange.Option, error)
	Update(id int, input taskexchange.UpdateOptionInput) error
	Delete(id int) error
}

type Events interface {
	Create(event taskexchange.Event) (int, error)
	FindPolling(userId, id int) ([]taskexchange.Event, error)
	FindNew(userId int) ([]taskexchange.Event, error)
	FindAll(userId, limit, offset int) ([]taskexchange.Event, error)
	CountAll(userId int) (int, error)
	FindLastUser(userId int) (taskexchange.Event, error)
	ViewAll(userId int) error
	View(userId, id int) error
	Delete(userId, id int) error
}

type Tasks interface {
	Create(task taskexchange.Task) (int, error)
	Update(id int, input taskexchange.UpdateTaskInput) error
	GetById(id int) (taskexchange.Task, error)
	FindAll(limit, offset int) ([]taskexchange.Task, error)
	FindAllByUser(userId, limit, offset int) ([]taskexchange.Task, error)
	CountAll() (int, error)
	CountAllByUser(userId int) (int, error)
	Delete(id int) error
}

type TaskOptions interface {
	Create(taskId, optionId int) (int, error)
	GetById(taskOptionId int) (taskexchange.TaskOption, error)
	GetByTaskId(taskId int) ([]taskexchange.TaskOption, error)
	Delete(taskOptionId int) error
}

type Offers interface {
}

type Orders interface {
}

type Repository struct {
	Users
	Events
	Options
	Tasks
	TaskOptions
	Offers
	Orders
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users:       NewUsersPostgres(db),
		Options:     NewOptionsPostgres(db),
		Events:      NewEventsPostgres(db),
		Tasks:       NewTasksPostgres(db),
		TaskOptions: NewTaskOptionsPostgres(db),
	}
}
