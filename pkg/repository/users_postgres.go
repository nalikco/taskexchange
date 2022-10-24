package repository

import (
	"gorm.io/gorm"
	"taskexchange"
	"time"
)

type UsersPostgres struct {
	db *gorm.DB
}

func NewUsersPostgres(db *gorm.DB) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

func (r *UsersPostgres) Create(user taskexchange.User) (taskexchange.User, error) {
	result := r.db.Table("users").Create(&user)

	return user, result.Error
}

func (r *UsersPostgres) GetAll() ([]taskexchange.User, error) {
	var users []taskexchange.User
	result := r.db.Table("users").Order("created_at DESC").Find(&users)

	return users, result.Error
}

func (r *UsersPostgres) GetAllHidden() ([]taskexchange.UserHidden, error) {
	var users []taskexchange.UserHidden
	result := r.db.Table("users").Where("deleted_at is null").Order("created_at DESC").Find(&users)

	return users, result.Error
}

func (r *UsersPostgres) GetById(id int) (taskexchange.User, error) {
	var user taskexchange.User
	result := r.db.Table("users").First(&user, id)

	return user, result.Error
}

func (r *UsersPostgres) GetByIdHidden(id int) (taskexchange.UserHidden, error) {
	var user taskexchange.UserHidden
	result := r.db.Table("users").Where("deleted_at is null").First(&user, id)

	return user, result.Error
}

func (r *UsersPostgres) GetByEmail(email string) (taskexchange.User, error) {
	var user taskexchange.User
	result := r.db.Table("users").Where("email = ?", email).First(&user)

	return user, result.Error
}

func (r *UsersPostgres) GetByUsername(username string) (taskexchange.User, error) {
	var user taskexchange.User
	result := r.db.Table("users").Where("username = ?", username).First(&user)

	return user, result.Error
}

func (r *UsersPostgres) GetByUsernameHidden(username string) (taskexchange.UserHidden, error) {
	var user taskexchange.UserHidden
	result := r.db.Table("users").Where("username = ?", username).Where("deleted_at is null").First(&user)

	return user, result.Error
}

func (r *UsersPostgres) GetByEmailAndPassword(email, password string) (taskexchange.User, error) {
	var user taskexchange.User
	result := r.db.Table("users").Where(&taskexchange.User{Email: email, Password: password}).First(&user)

	return user, result.Error
}

func (r *UsersPostgres) Update(user taskexchange.User) error {
	result := r.db.Table("users").Save(&user)

	return result.Error
}

func (r *UsersPostgres) Delete(user taskexchange.User) error {
	currentTime := time.Now()
	user.DeletedAt = &currentTime
	result := r.db.Table("users").Save(&user)

	return result.Error
}

func (r *UsersPostgres) Restore(user taskexchange.User) error {
	user.DeletedAt = nil
	result := r.db.Table("users").Save(&user)

	return result.Error
}

func (r *UsersPostgres) CountAll(sort taskexchange.SortUsersCount) (int64, error) {
	var count int64
	whereCon := map[string]interface{}{}

	if sort.Performers == true {
		whereCon["type"] = 1
	}
	if sort.Customers == true {
		whereCon["type"] = 2
	}

	result := r.db.Table("users").Where(whereCon).Count(&count)

	return count, result.Error
}
