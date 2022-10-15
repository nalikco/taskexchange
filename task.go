package taskexchange

import "time"

type Task struct {
	Id           int          `json:"id" db:"id"`
	CustomerId   int          `json:"customer_id" db:"customer_id"`
	Customer     User         `json:"customer" db:"customer"`
	Status       int          `json:"status" db:"status"`
	Amount       int          `json:"amount" db:"amount"`
	DeliveryDate time.Time    `json:"delivery_date" db:"delivery_date"`
	Link         string       `json:"link" db:"link"`
	Description  string       `json:"description" db:"description"`
	Options      []Option     `json:"options"`
	TaskOptions  []TaskOption `json:"task_options" db:"task_options"`
	Offers       []Offer      `json:"offers"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	DeletedAt    *time.Time   `json:"deleted_at" db:"deleted_at"`
}

type TaskOption struct {
	Id       int `json:"id" db:"id"`
	TaskId   int `json:"task_id" db:"task_id"`
	OptionId int `json:"option_id" db:"option_id"`
}

type UpdateTaskInput struct {
	Status             *int    `json:"status"`
	Amount             *int    `json:"amount"`
	DeliveryDateString *string `json:"delivery_date"`
	DeliveryDate       *time.Time
	Link               *string `json:"link"`
	Description        *string `json:"description"`
	Options            *[]int  `json:"options"`
}

func (t *Task) CalculatePrice() float64 {
	var price float64 = 0

	for _, option := range t.Options {
		price += option.Price
	}

	return price * float64(t.Amount)
}
