package taskexchange

import (
	"time"
)

type Event struct {
	ID        int        `json:"id" db:"id"`
	UserId    int        `json:"user_id" db:"user_id"`
	Message   string     `json:"message" db:"message"`
	Link      string     `json:"link" db:"link"`
	ViewedAt  *time.Time `json:"viewed_at" db:"viewed_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
