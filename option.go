package taskexchange

import (
	"errors"
	"time"
)

type Option struct {
	Id        int        `json:"id" db:"id"`
	ParentId  *int       `json:"parent_id" db:"parent_id"`
	Title     string     `json:"title" db:"title"`
	Price     float64    `json:"price" db:"price"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type UpdateOptionInput struct {
	ParentId *int     `json:"parent_id"`
	Title    *string  `json:"title"`
	Price    *float64 `json:"price"`
}

func (i UpdateOptionInput) Validate() error {
	if i.ParentId == nil && i.Title == nil && i.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
