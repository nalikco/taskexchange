package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"taskexchange"
)

type OptionsPostgres struct {
	db *sqlx.DB
}

func NewOptionsPostgres(db *sqlx.DB) *OptionsPostgres {
	return &OptionsPostgres{db: db}
}

func (r *OptionsPostgres) Create(parentId int, option taskexchange.Option) (int, error) {
	var id int

	if parentId == 0 {
		query := fmt.Sprintf("INSERT INTO %s (title, price, short) VALUES ($1, $2, $3) RETURNING id", optionsTable)
		row := r.db.QueryRow(query, option.Title, option.Price, option.Short)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (parent_id, title, price, short) VALUES ($1, $2, $3, $4) RETURNING id", optionsTable)
		row := r.db.QueryRow(query, parentId, option.Title, option.Price, option.Short)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *OptionsPostgres) GetAll(full bool) ([]taskexchange.Option, error) {
	var options []taskexchange.Option
	var query string

	if full {
		query = fmt.Sprintf("SELECT * FROM %s ORDER BY created_at DESC", optionsTable)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s WHERE deleted_at IS NULL ORDER BY created_at DESC", optionsTable)
	}
	err := r.db.Select(&options, query)

	return options, err
}

func (r *OptionsPostgres) GetCategories() ([]taskexchange.Option, error) {
	var options []taskexchange.Option

	query := fmt.Sprintf("SELECT * FROM %s WHERE parent_id is null ORDER BY created_at DESC", optionsTable)
	err := r.db.Select(&options, query)

	return options, err
}

func (r *OptionsPostgres) GetById(id int, full bool) (taskexchange.Option, error) {
	var option taskexchange.Option
	var query string

	if full {
		query = fmt.Sprintf("SELECT * FROM %s WHERE id=$1", optionsTable)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND deleted_at IS NULL", optionsTable)
	}
	err := r.db.Get(&option, query, id)

	return option, err
}

func (r *OptionsPostgres) GetByTitle(title string, parentId int) (taskexchange.Option, error) {
	var option taskexchange.Option
	var err error

	if parentId == 0 {
		query := fmt.Sprintf("SELECT * FROM %s WHERE LOWER(title)=LOWER($1) AND deleted_at IS NULL", optionsTable)
		err = r.db.Get(&option, query, title)

	} else {
		query := fmt.Sprintf("SELECT * FROM %s WHERE LOWER(title)=LOWER($1) AND deleted_at IS NULL AND parent_id=$2", optionsTable)
		err = r.db.Get(&option, query, title, parentId)
	}

	return option, err
}

func (r *OptionsPostgres) GetByIds(ids []int) ([]taskexchange.Option, error) {
	var options []taskexchange.Option
	var idsQuery []string
	var args []interface{}

	for i, id := range ids {
		idsQuery = append(idsQuery, fmt.Sprintf("$%d", i+1))
		args = append(args, id)
	}

	setQuery := strings.Join(idsQuery, ",")

	query := fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s) ORDER BY created_at DESC", optionsTable, setQuery)

	err := r.db.Select(&options, query, args...)

	return options, err
}

func (r *OptionsPostgres) Update(id int, input taskexchange.UpdateOptionInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

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

	if input.Short != nil {
		setValues = append(setValues, fmt.Sprintf("short=$%d", argId))
		args = append(args, *input.Short)
		argId++
	}

	if input.ParentId != nil {
		if *input.ParentId == 0 {
			setValues = append(setValues, "parent_id=null")
		} else {
			if *input.ParentId != id {
				setValues = append(setValues, fmt.Sprintf("parent_id=$%d", argId))
				args = append(args, *input.ParentId)
				argId++
			}
		}
	}

	args = append(args, id)

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", optionsTable, setQuery, argId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *OptionsPostgres) Delete(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=now() WHERE id=$1", optionsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *OptionsPostgres) Restore(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=null WHERE id=$1", optionsTable)
	_, err := r.db.Exec(query, id)

	return err
}
