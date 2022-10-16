package taskexchange

import "time"

type User struct {
	Id         int        `json:"id" db:"id"`
	Email      string     `json:"email" db:"email"`
	Password   string     `json:"password"`
	Username   string     `json:"username" db:"username"`
	Type       int        `json:"type" db:"type"`
	Balance    float64    `json:"balance" db:"balance"`
	Points     int        `json:"points" db:"points"`
	LastOnline time.Time  `json:"last_online" db:"last_online"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}

type UpdateUserInput struct {
	Email    *string  `json:"email" db:"email"`
	Password *string  `json:"password"`
	Username *string  `json:"username" db:"username"`
	Type     *int     `json:"type" db:"type"`
	Balance  *float64 `json:"balance" db:"balance"`
	Points   *int     `json:"points" db:"points"`
}

type SortUsersCount struct {
	Performers bool
	Customers  bool
}
