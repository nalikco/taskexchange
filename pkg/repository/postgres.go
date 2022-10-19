package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable                       = "users"
	eventsTable                      = "events"
	optionsTable                     = "options"
	tasksTable                       = "tasks"
	taskOptionsTable                 = "task_options"
	offersTable                      = "offers"
	ordersTable                      = "orders"
	messagesConversationsTable       = "messages_conversations"
	messagesConversationMembersTable = "messages_conversation_members"
	messagesTable                    = "messages"
	paymentsTable                    = "payments"
	postsTable                       = "posts"
	postCategoriesTable              = "post_categories"
	postCategoryTable                = "post_category"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
