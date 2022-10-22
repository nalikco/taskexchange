package taskexchange

import "time"

type User struct {
	Id         int        `json:"id" gorm:"column:id" db:"id"`
	Email      string     `json:"email" gorm:"column:email" db:"email"`
	Password   string     `json:"password" gorm:"column:password_hash"`
	Username   string     `json:"username" gorm:"column:username" db:"username"`
	FirstName  string     `json:"first_name" gorm:"column:first_name" db:"first_name"`
	LastName   string     `json:"last_name" gorm:"column:last_name" db:"last_name"`
	Type       int        `json:"type" gorm:"column:type" db:"type"`
	Balance    float64    `json:"balance" gorm:"column:balance"`
	Points     int        `json:"points" gorm:"column:points" db:"points"`
	LastOnline time.Time  `json:"last_online" gorm:"column:last_online" db:"last_online"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at" db:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type UserHidden struct {
	Id         int       `json:"id" gorm:"id" db:"id"`
	Email      string    `json:"email" gorm:"email" db:"email"`
	Username   string    `json:"username" gorm:"username" db:"username"`
	FirstName  string    `json:"first_name" gorm:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" gorm:"last_name" db:"last_name"`
	Type       int       `json:"type" gorm:"type" db:"type"`
	Points     int       `json:"points" gorm:"points" db:"points"`
	LastOnline time.Time `json:"last_online" gorm:"last_online" db:"last_online"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at" db:"created_at"`
}

type UpdateUserInput struct {
	Email     *string  `json:"email"`
	Password  *string  `json:"password"`
	Username  *string  `json:"username"`
	FirstName *string  `json:"first_name"`
	LastName  *string  `json:"last_name"`
	Type      *int     `json:"type"`
	Balance   *float64 `json:"balance"`
	Points    *int     `json:"points"`
}

type SortUsersCount struct {
	Performers bool
	Customers  bool
}
