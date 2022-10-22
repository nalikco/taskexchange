package taskexchange

import "time"

type Offer struct {
	Id          int        `json:"id" db:"id"`
	Performer   UserHidden `json:"performer" db:"performer"`
	PerformerId int        `json:"performer_id" db:"performer_id"`
	TaskId      int        `json:"task_id" db:"task_id"`
	Status      int        `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type CreateOfferInput struct {
	TaskId int `json:"task_id"`
}

type UpdateOfferInput struct {
	Status int `json:"status"`
}
