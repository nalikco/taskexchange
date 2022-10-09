package taskexchange

import "time"

type Offer struct {
	Id          int       `json:"id"`
	PerformerId int       `json:"performer_id"`
	TaskId      int       `json:"task_id"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
