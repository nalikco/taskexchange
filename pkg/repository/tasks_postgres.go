package repository

import (
	"fmt"
	"strings"
	"taskexchange"

	"github.com/jmoiron/sqlx"
)

type TasksPostgres struct {
	db *sqlx.DB
}

func NewTasksPostgres(db *sqlx.DB) *TasksPostgres {
	return &TasksPostgres{db: db}
}

func (r *TasksPostgres) Create(task taskexchange.Task) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (customer_id, status, amount, delivery_date, link, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", tasksTable)

	row := r.db.QueryRow(query, task.CustomerId, task.Status, task.Amount, task.DeliveryDate, task.Link, task.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TasksPostgres) Update(id int, input taskexchange.UpdateTaskInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	if input.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount=$%d", argId))
		args = append(args, *input.Amount)
		argId++
	}

	if input.Link != nil {
		setValues = append(setValues, fmt.Sprintf("link=$%d", argId))
		args = append(args, *input.Link)
		argId++
	}

	if input.DeliveryDate != nil {
		setValues = append(setValues, fmt.Sprintf("delivery_date=$%d", argId))
		args = append(args, *input.DeliveryDate)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	args = append(args, id)

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", tasksTable, setQuery, argId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *TasksPostgres) GetById(id int) (taskexchange.Task, error) {
	var task taskexchange.Task

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", tasksTable)
	err := r.db.Get(&task, query, id)

	return task, err
}

func (r *TasksPostgres) FindAll(limit, offset int) ([]taskexchange.Task, error) {
	var tasks []taskexchange.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null AND status=1 AND delivery_date + INTERVAL '1 day' > now() AND amount > 0 ORDER BY id DESC LIMIT %d OFFSET %d", tasksTable, limit, offset)
	err := r.db.Select(&tasks, query)

	return tasks, err
}

func (r *TasksPostgres) FindAllForAdmin(limit, offset int) ([]taskexchange.Task, error) {
	var tasks []taskexchange.Task

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT %d OFFSET %d", tasksTable, limit, offset)
	err := r.db.Select(&tasks, query)

	for i, task := range tasks {
		query := fmt.Sprintf("SELECT %s FROM %s WHERE id=$1", usersAllColumns, usersTable)
		err := r.db.Get(&tasks[i].Customer, query, task.CustomerId)
		if err != nil {
			return []taskexchange.Task{}, err
		}

	}

	return tasks, err
}

func (r *TasksPostgres) CountAll() (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE deleted_at is null AND status=1 AND delivery_date + INTERVAL '1 day' > now() AND amount > 0", tasksTable)
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TasksPostgres) CountAllForAdmin() (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tasksTable)
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TasksPostgres) FindAllByUser(userId, limit, offset int) ([]taskexchange.Task, error) {
	var tasks []taskexchange.Task

	query := fmt.Sprintf("SELECT * FROM %s WHERE customer_id=$1 ORDER BY id DESC LIMIT %d OFFSET %d", tasksTable, limit, offset)
	err := r.db.Select(&tasks, query, userId)

	return tasks, err
}

func (r *TasksPostgres) CountAllByUser(userId int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE customer_id=$1", tasksTable)
	err := r.db.QueryRow(query, userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TasksPostgres) CountActiveByUser(userId int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE customer_id=$1 AND status=1", tasksTable)
	err := r.db.QueryRow(query, userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TasksPostgres) Delete(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=now() WHERE id=$1", tasksTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *TasksPostgres) Restore(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=null WHERE id=$1", tasksTable)
	_, err := r.db.Exec(query, id)

	return err
}
