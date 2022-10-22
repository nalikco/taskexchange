package taskexchange

import (
	"errors"
	"time"
)

type Option struct {
	Id        int        `json:"id" gorm:"id"`
	ParentID  *int       `json:"parent_id"`
	Parent    *Option    `json:"parent"`
	Options   []Option   `json:"options" gorm:"foreignkey:parent_id"`
	Short     *string    `json:"short" gorm:"short"`
	Title     string     `json:"title" gorm:"title"`
	Price     float64    `json:"price" gorm:"price"`
	CreatedAt time.Time  `json:"created_at" gorm:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
}

type UpdateOptionInput struct {
	ParentId *int     `json:"parent_id"`
	Short    *string  `json:"short"`
	Title    *string  `json:"title"`
	Price    *float64 `json:"price"`
}

func (i UpdateOptionInput) Validate() error {
	if i.ParentId == nil && i.Title == nil && i.Price == nil && i.Short == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type SortOptions struct {
	Deleted bool
}
