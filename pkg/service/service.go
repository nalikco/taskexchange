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

type Options interface {
	Create(parentId int, option taskexchange.Option) (int, error)
	GetAll() ([]taskexchange.Option, error)
	GetById(id int) (taskexchange.Option, error)
	Update(id int, input taskexchange.UpdateOptionInput) error
	Delete(id int) error
}

type Events interface {
	Create(userId int, message, link string) (int, error)
	Polling(userId, id int) ([]taskexchange.Event, error)
	GetNew(userId int) ([]taskexchange.Event, error)
	GetAll(userId int, pagination taskexchange.Pagination) ([]taskexchange.Event, taskexchange.Pagination, error)
	GetLastId(userId int) (int, error)
	ViewAll(userId int) error
	View(userId, id int) error
	Delete(userId, id int) error
}

type Tasks interface {
	Create(task taskexchange.Task) (int, error)
	CreateFromExcelFile(userId int, filename string) error
	Update(id int, input taskexchange.UpdateTaskInput) error
	GetById(id int) (taskexchange.Task, error)
	GetAll(userId int, pagination taskexchange.Pagination) ([]taskexchange.Task, taskexchange.Pagination, error)
	CountActiveByUser(userId int) (int, error)
	Delete(id int, task taskexchange.Task, customerId int) error
}

type Offers interface {
	GetPerformerActive(performerId int) ([]taskexchange.Offer, error)
	Make(performerId, taskId int) (int, error)
	ChangeStatus(offerId, customerId, status int) error
}

type Orders interface {
	FindAllByPerformerId(performerId int) ([]taskexchange.Order, error)
	FindActiveByPerformerId(performerId int) ([]taskexchange.Order, error)
	FindAllByCustomerId(customerId int) ([]taskexchange.Order, error)
	Update(orderId int, userId int, input taskexchange.UpdateOrderInput) error
}

type Service struct {
	Authorization
	Users
	Events
	Options
	Tasks
	Offers
	Orders
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Users),
		Users:         NewUsersService(repos.Users),
		Options:       NewOptionService(repos.Options),
		Events:        NewEventsService(repos.Events),
		Tasks:         NewTasksService(repos.Tasks, repos.TaskOptions, repos.Users, repos.Options, repos.Offers),
		Offers:        NewOffersService(repos.Offers, repos.Tasks, repos.Users, repos.Events, repos.Orders),
		Orders:        NewOrdersService(repos.Orders, repos.Users, repos.Options, repos.Tasks, repos.TaskOptions, repos.Events),
	}
}
