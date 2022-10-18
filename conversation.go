package taskexchange

import "time"

type Conversation struct {
	ID        int        `json:"id" db:"id"`
	Members   []User     `json:"members"`
	Messages  []Message  `json:"messages"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
