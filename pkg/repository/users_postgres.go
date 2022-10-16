package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"taskexchange"
)

const usersVisibleColumns = "id,email,username,type,points,last_online,created_at"
const usersAllColumns = "id,email,username,type,balance,points,last_online,created_at,deleted_at"

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) Create(user taskexchange.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, username, type, balance, points) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Email, user.Password, user.Username, user.Type, user.Balance, user.Points)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UsersPostgres) GetAll(full bool) ([]taskexchange.User, error) {
	var users []taskexchange.User
	var query string

	if full {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at is null ORDER BY id DESC", usersAllColumns, usersTable)
	} else {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at is null ORDER BY id DESC", usersVisibleColumns, usersTable)
	}

	err := r.db.Select(&users, query)

	return users, err
}

func (r *UsersPostgres) GetById(id int, full bool) (taskexchange.User, error) {
	var user taskexchange.User
	var query string

	if full {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE id=$1", usersAllColumns, usersTable)
	} else {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE id=$1", usersVisibleColumns, usersTable)
	}

	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UsersPostgres) GetByEmail(email string) (taskexchange.User, error) {
	var user taskexchange.User
	query := fmt.Sprintf("SELECT %s FROM %s WHERE email=$1 AND deleted_at is null", usersVisibleColumns, usersTable)

	err := r.db.Get(&user, query, email)

	return user, err
}

func (r *UsersPostgres) GetByEmailAndPassword(email, password string) (taskexchange.User, error) {
	var user taskexchange.User
	query := fmt.Sprintf("SELECT id, type FROM %s WHERE email=$1 AND password_hash=$2", usersTable)

	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *UsersPostgres) Update(id int, input taskexchange.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, *input.Password)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.Type != nil {
		setValues = append(setValues, fmt.Sprintf("type=$%d", argId))
		args = append(args, *input.Type)
		argId++
	}

	if input.Balance != nil {
		setValues = append(setValues, fmt.Sprintf("balance=$%d", argId))
		args = append(args, *input.Balance)
		argId++
	}

	if input.Points != nil {
		setValues = append(setValues, fmt.Sprintf("points=$%d", argId))
		args = append(args, *input.Points)
		argId++
	}

	args = append(args, id)

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", usersTable, setQuery, argId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *UsersPostgres) UpdateOnline(id int) error {
	query := fmt.Sprintf("UPDATE %s SET last_online=now() WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *UsersPostgres) Delete(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=now() WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *UsersPostgres) CountAll(sort taskexchange.SortUsersCount) (int, error) {
	var count int

	sortQuery := ""
	if sort.Performers == true {
		sortQuery = "WHERE type=1"
	}
	if sort.Customers == true {
		sortQuery = "WHERE type=2"
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", usersTable, sortQuery)
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
