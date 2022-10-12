package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"taskexchange"
)

type TaskOptionsPostgres struct {
	db *sqlx.DB
}

func NewTaskOptionsPostgres(db *sqlx.DB) *TaskOptionsPostgres {
	return &TaskOptionsPostgres{db: db}
}

func (r *TaskOptionsPostgres) Create(taskId, optionId int) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (task_id, option_id) VALUES ($1, $2) RETURNING id", taskOptionsTable)
	row := r.db.QueryRow(query, taskId, optionId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskOptionsPostgres) GetById(taskOptionId int) (taskexchange.TaskOption, error) {
	var taskOption taskexchange.TaskOption

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", taskOptionsTable)

	err := r.db.Get(&taskOption, query, taskOptionId)

	return taskOption, err
}

func (r *TaskOptionsPostgres) GetByTaskId(taskId int) ([]taskexchange.TaskOption, error) {
	var options []taskexchange.TaskOption
	query := fmt.Sprintf("SELECT * FROM %s WHERE task_id=$1 ORDER BY id DESC", taskOptionsTable)
	err := r.db.Select(&options, query, taskId)

	return options, err
}

func (r *TaskOptionsPostgres) Delete(taskOptionId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", taskOptionsTable)
	_, err := r.db.Exec(query, taskOptionId)

	return err
}
