package taskexchange

import "time"

type Task struct {
	Id           int       `json:"id" db:"id"`
	CustomerId   int       `json:"customer_id" db:"customer_id"`
	Status       int       `json:"status" db:"status"`
	Amount       int       `json:"amount" db:"amount"`
	DeliveryDate time.Time `json:"delivery_date" db:"delivery_date"`
	Link         string    `json:"link" db:"link"`
	Description  string    `json:"description" db:"description"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type TaskOption struct {
	Id       int `db:"id"`
	TaskId   int `db:"task_id"`
	OptionId int `db:"option_id"`
}
