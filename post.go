package taskexchange

import (
	"time"
)

type Post struct {
	ID         int            `json:"id"`
	Categories []PostCategory `json:"categories"`
	Author     User           `json:"author"`
	MainImage  string         `json:"main_image"`
	Status     int            `json:"status"`
	Title      string         `json:"title"`
	Short      string         `json:"short"`
	Text       string         `json:"text"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  *time.Time     `json:"deleted_at"`
}

type UpdatePostInput struct {
	Categories *[]int `json:"categories"`
	MainImage  *string
	Status     *int    `json:"status"`
	Title      *string `json:"title"`
	Short      *string `json:"short"`
	Text       *string `json:"text"`
}

type PostCategory struct {
	ID        int        `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type UpdatePostCategoryInput struct {
	Title *string `json:"title"`
}
