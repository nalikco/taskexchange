package taskexchange

import "time"

type Message struct {
	ID           int          `json:"id" db:"id"`
	Conversation Conversation `json:"conversation"`
	Sender       User         `json:"sender"`
	Text         string       `json:"text" db:"text"`
	ViewedAt     *time.Time   `json:"viewed_at" db:"viewed_at"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	DeletedAt    *time.Time   `json:"deleted_at" db:"deleted_at"`
}

type Conversation struct {
	ID        int        `json:"id" db:"id"`
	Members   []User     `json:"members"`
	Messages  []Message  `json:"messages"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
