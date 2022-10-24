package repository

import (
	"gorm.io/gorm"
	"taskexchange"

	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(user taskexchange.User) (taskexchange.User, error)
	GetAll() ([]taskexchange.User, error)
	GetAllHidden() ([]taskexchange.UserHidden, error)
	GetById(id int) (taskexchange.User, error)
	GetByIdHidden(id int) (taskexchange.UserHidden, error)
	GetByEmail(email string) (taskexchange.User, error)
	GetByUsername(username string) (taskexchange.User, error)
	GetByUsernameHidden(username string) (taskexchange.UserHidden, error)
	GetByEmailAndPassword(email, password string) (taskexchange.User, error)
	CountAll(sort taskexchange.SortUsersCount) (int64, error)
	Update(user taskexchange.User) error
	Delete(user taskexchange.User) error
	Restore(user taskexchange.User) error
}

type Options interface {
	Create(option taskexchange.Option) (taskexchange.Option, error)
	GetAll(sort taskexchange.SortOptions) ([]taskexchange.Option, error)
	GetCategories() ([]taskexchange.Option, error)
	GetById(id int, deleted bool) (taskexchange.Option, error)
	GetByIds(ids []int) ([]taskexchange.Option, error)
	GetByTitle(title string, parentId int) (taskexchange.Option, error)
	Update(option taskexchange.Option) error
	Delete(option taskexchange.Option) error
	Restore(option taskexchange.Option) error
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
	FindAllForAdmin(limit, offset int) ([]taskexchange.Task, error)
	FindAllByUser(userId, limit, offset int) ([]taskexchange.Task, error)
	CountAll() (int, error)
	CountAllForAdmin() (int, error)
	CountAllByUser(userId int) (int, error)
	CountActiveByUser(userId int) (int, error)
	Delete(id int) error
	Restore(id int) error
}

type TaskOptions interface {
	Create(taskId, optionId int) (int, error)
	GetById(taskOptionId int) (taskexchange.TaskOption, error)
	GetByTaskId(taskId int) ([]taskexchange.TaskOption, error)
	Delete(taskOptionId int) error
}

type Offers interface {
	Create(performerId, taskId int) (int, error)
	Update(offerId int, input taskexchange.UpdateOfferInput) error
	FindPerformerActiveOffers(performerId int) ([]taskexchange.Offer, error)
	FindAllByTask(taskId int) ([]taskexchange.Offer, error)
	FindByTaskAndStatus(taskId, status int) ([]taskexchange.Offer, error)
	FindOneById(offerId int) (taskexchange.Offer, error)
	FindOneByPerformerIdAndTaskIdAndStatus(performerId, taskId, status int) (taskexchange.Offer, error)
}

type Orders interface {
	Create(offerId, taskId int) (int, error)
	FindActiveByTaskId(taskId int) ([]taskexchange.Order, error)
	FindAllByPerformerId(performerId int) ([]taskexchange.Order, error)
	CountAllByPerformerId(performerId int) (int, error)
	FindAll() ([]taskexchange.Order, error)
	FindAllByCustomerId(customerId int) ([]taskexchange.Order, error)
	CountAllByCustomerId(customerId int) (int, error)
	FindActiveByPerformerId(performerId int) ([]taskexchange.Order, error)
	FindActiveByCustomerId(customerId int) ([]taskexchange.Order, error)
	Update(id int, input taskexchange.UpdateOrderInput) error
	FindOneById(orderId int) (taskexchange.Order, error)
	CountAllActive() (int, error)
	GetAllCompleted() ([]taskexchange.Order, error)
	GetActiveCountByTaskId(taskId int) (int, error)
}

type Messages interface {
	GetUserConversations(user taskexchange.User) ([]taskexchange.Conversation, error)
	CreateConversation(members []taskexchange.User) (int, error)
	GetConversationById(id int) (taskexchange.Conversation, error)
	GetMessageById(id int) (taskexchange.Message, error)
	SendMessageToRecipient(sender taskexchange.User, recipient taskexchange.User, text string) (int, error)
	GetMessagesByConversation(conversation taskexchange.Conversation) ([]taskexchange.Message, error)
	CountUserUnViewedMessages(user taskexchange.User) (int, error)
	ViewConversation(conversation taskexchange.Conversation, user taskexchange.User) error
}

type Payments interface {
	Create(payment taskexchange.Payment) (int, error)
	GetByUser(user taskexchange.User) ([]taskexchange.Payment, error)
}

type Posts interface {
	CreateCategory(category taskexchange.PostCategory) (int, error)
	CreatePost(post taskexchange.Post) (int, error)
	GetById(id int, deleted bool) (taskexchange.Post, error)
	GetCategoryById(id int) (taskexchange.PostCategory, error)
	GetCategoriesById(ids []int) ([]taskexchange.PostCategory, error)
	GetAll(limit, offset int) ([]taskexchange.Post, error)
	GetAllCategories() ([]taskexchange.PostCategory, error)
	Update(id int, input taskexchange.UpdatePostInput) error
	UpdateCategory(id int, input taskexchange.UpdatePostCategoryInput) error
	Delete(id int) error
	DeleteCategory(id int) error
	Restore(id int) error
	RestoreCategory(id int) error
}

type Repository struct {
	Users
	Events
	Options
	Tasks
	TaskOptions
	Offers
	Orders
	Messages
	Payments
	Posts
}

func NewRepository(db *sqlx.DB, gorm *gorm.DB) *Repository {
	return &Repository{
		Users:       NewUsersPostgres(gorm),
		Options:     NewOptionsPostgres(gorm),
		Events:      NewEventsPostgres(db),
		Tasks:       NewTasksPostgres(db),
		TaskOptions: NewTaskOptionsPostgres(db),
		Offers:      NewOffersPostgres(db),
		Orders:      NewOrdersPostgres(db, gorm),
		Messages:    NewMessagesPostgres(db, gorm),
		Payments:    NewPaymentsPostgres(db),
		Posts:       NewPostsPostgres(db),
	}
}
