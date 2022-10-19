package taskexchange

import "time"

type Payment struct {
	ID        int        `json:"id"`
	User      User       `json:"user"`
	Type      int        `json:"type"`
	Comment   string     `json:"comment"`
	Sum       float64    `json:"sum"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
