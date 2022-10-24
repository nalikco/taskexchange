package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
)

type Authorization interface {
	CreateUser(user taskexchange.User) (taskexchange.User, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	UpdateOnline(user taskexchange.User) error
}

type Users interface {
	CreateUser(user taskexchange.User) (taskexchange.User, error)
	GetAll() ([]taskexchange.User, error)
	GetAllHidden() ([]taskexchange.UserHidden, error)
	GetById(id int) (taskexchange.User, error)
	GetByIdHidden(id int) (taskexchange.UserHidden, error)
	GetByUsername(username string) (taskexchange.User, error)
	GetByUsernameHidden(username string) (taskexchange.UserHidden, error)
	Update(user taskexchange.User, input taskexchange.UpdateUserInput) error
	CountAll(sort taskexchange.SortUsersCount) (int64, error)
	Delete(id int) error
}

type Options interface {
	Create(option taskexchange.Option) (taskexchange.Option, error)
	GetAll(sort taskexchange.SortOptions) ([]taskexchange.Option, error)
	GetCategories() ([]taskexchange.Option, error)
	GetById(id int, deleted bool) (taskexchange.Option, error)
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
	CreateFromExcelFile(userId int, filename string) (float64, error)
	Update(id int, input taskexchange.UpdateTaskInput) (float64, error)
	GetById(id int) (taskexchange.Task, error)
	GetAll(userId int, pagination taskexchange.Pagination) ([]taskexchange.Task, taskexchange.Pagination, error)
	CountActive() (int, error)
	CountActiveByUser(userId int) (int, error)
	Delete(id int, task taskexchange.Task, user taskexchange.User) error
}

type Offers interface {
	GetPerformerActive(performerId int) ([]taskexchange.Offer, error)
	Make(performerId, taskId int) (int, error)
	ChangeStatus(offerId, customerId, status int) error
}

type Orders interface {
	FindAllByPerformerId(performerId int) ([]taskexchange.Order, error)
	FindActiveCountByTaskId(taskId int) (int, error)
	FindActiveByPerformerId(performerId int) ([]taskexchange.Order, error)
	FindActiveByCustomerId(customerId int) ([]taskexchange.Order, error)
	FindAll() ([]taskexchange.Order, error)
	FindAllByCustomerId(customerId int) ([]taskexchange.Order, error)
	Update(orderId int, userId int, input taskexchange.UpdateOrderInput) error
	CountAllActive() (int, error)
	GetAllCompleted() ([]taskexchange.Order, error)
}

type Messages interface {
	GetUserConversations(user taskexchange.User) ([]taskexchange.Conversation, error)
	CreateConversation(members []taskexchange.User) (int, error)
	GetConversationById(id int) (taskexchange.Conversation, error)
	SendMessageToRecipient(sender taskexchange.User, recipient taskexchange.User, text string) (int, error)
	GetMessagesByConversation(conversation taskexchange.Conversation) ([]taskexchange.Message, error)
	CountUserUnViewedMessages(user taskexchange.User) (int, error)
	ViewConversation(conversation taskexchange.Conversation, user taskexchange.User) error
	Polling(user taskexchange.User) (PollingMessage, error)
}

type Payments interface {
	Create(payment taskexchange.Payment) (int, error)
	GetByUser(user taskexchange.User) ([]taskexchange.Payment, error)
}

type Posts interface {
	CreateCategory(category taskexchange.PostCategory) (int, error)
	CreatePost(post taskexchange.Post) (int, error)
	GetById(id int) (taskexchange.Post, error)
	GetCategoryById(id int) (taskexchange.PostCategory, error)
	GetCategoriesById(ids []int) ([]taskexchange.PostCategory, error)
	GetAll(limit, offset int) ([]taskexchange.Post, error)
	GetAllCategories() ([]taskexchange.PostCategory, error)
	UpdatePostImage(id int, filename string) error
	Update(id int, input taskexchange.UpdatePostInput) error
	UpdateCategory(id int, input taskexchange.UpdatePostCategoryInput) error
	Delete(id int) error
	DeleteCategory(id int) error
}

type Queue interface {
	Run() error
}

type Service struct {
	Authorization
	Users
	Events
	Options
	Tasks
	Offers
	Orders
	Messages
	Payments
	Posts
	Queue
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Users),
		Users:         NewUsersService(repos.Users),
		Options:       NewOptionService(repos.Options),
		Events:        NewEventsService(repos.Events),
		Tasks:         NewTasksService(repos.Tasks, repos.TaskOptions, repos.Users, repos.Options, repos.Offers),
		Offers:        NewOffersService(repos.Offers, repos.Tasks, repos.Users, repos.Events, repos.Orders),
		Orders:        NewOrdersService(repos.Orders, repos.Users, repos.Options, repos.Tasks, repos.TaskOptions, repos.Events, repos.Payments),
		Messages:      NewMessagesService(repos.Messages),
		Payments:      NewPaymentsService(repos.Payments),
		Posts:         NewPostsService(repos.Posts),
		Queue:         NewQueueService(repos.Tasks, repos.Orders),
	}
}
