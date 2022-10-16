package taskexchange

import "time"

type Order struct {
	Id               int        `json:"id"`
	OfferId          int        `json:"offer_id" db:"offer_id"`
	Offer            Offer      `json:"offer" db:"offer"`
	TaskId           int        `json:"task_id" db:"task_id"`
	Task             Task       `json:"task" db:"task"`
	Status           int        `json:"status" db:"status"`
	CanceledUserId   *int       `json:"canceled_user_id" db:"canceled_user_id"`
	ReturnComment    *string    `json:"return_comment" db:"return_comment"`
	SurrenderComment *string    `json:"surrender_comment" db:"surrender_comment"`
	CancelComment    *string    `json:"cancel_comment" db:"cancel_comment"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	DeletedAt        *time.Time `json:"deleted_at" db:"deleted_at"`
}

type UpdateOrderInput struct {
	TaskId           *int    `json:"task_id"`
	Status           *int    `json:"status"`
	CanceledUserId   *int    `json:"canceled_user_id"`
	ReturnComment    *string `json:"return_comment"`
	SurrenderComment *string `json:"surrender_comment"`
	CancelComment    *string `json:"cancel_comment"`
}
