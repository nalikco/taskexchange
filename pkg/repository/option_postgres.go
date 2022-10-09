package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"taskexchange"
)

type OptionPostgres struct {
	db *sqlx.DB
}

func NewOptionPostgres(db *sqlx.DB) *OptionPostgres {
	return &OptionPostgres{db: db}
}

func (r *OptionPostgres) Create(parentId int, option taskexchange.Option) (int, error) {
	var id int

	if parentId == 0 {
		query := fmt.Sprintf("INSERT INTO %s (title, price) VALUES ($1, $2) RETURNING id", optionsTable)
		row := r.db.QueryRow(query, option.Title, option.Price)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (parent_id, title, price) VALUES ($1, $2, $3) RETURNING id", optionsTable)
		row := r.db.QueryRow(query, parentId, option.Title, option.Price)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *OptionPostgres) GetAll() ([]taskexchange.Option, error) {
	var options []taskexchange.Option

	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at IS NULL ORDER BY created_at DESC", optionsTable)
	err := r.db.Select(&options, query)

	return options, err
}

func (r *OptionPostgres) GetById(id int) (taskexchange.Option, error) {
	var option taskexchange.Option

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND deleted_at IS NULL", optionsTable)
	err := r.db.Get(&option, query, id)

	return option, err
}

func (r *OptionPostgres) Update(id int, input taskexchange.UpdateOptionInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.ParentId != nil {
		setValues = append(setValues, fmt.Sprintf("parent_id=$%d", argId))
		args = append(args, *input.ParentId)
		argId++
	}

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	args = append(args, id)

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", optionsTable, setQuery, argId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *OptionPostgres) Delete(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=now() WHERE id=$1", optionsTable)
	_, err := r.db.Exec(query, id)

	return err
}
