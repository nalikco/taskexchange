package repository

import (
	"fmt"
	"strings"
	"taskexchange"

	"github.com/jmoiron/sqlx"
)

type OrdersPostgres struct {
	db *sqlx.DB
}

func NewOrdersPostgres(db *sqlx.DB) *OrdersPostgres {
	return &OrdersPostgres{
		db: db,
	}
}

func (r *OrdersPostgres) Create(offerId, taskId int) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (offer_id, task_id) VALUES ($1, $2) RETURNING id", ordersTable)
	row := r.db.QueryRow(query, offerId, taskId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrdersPostgres) FindActiveByTaskId(taskId int) ([]taskexchange.Order, error) {
	var orders []taskexchange.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE task_id=$1 AND status IN (0, 1) AND deleted_at IS NULL", ordersTable)
	err := r.db.Select(&orders, query, taskId)

	return orders, err
}

func (r *OrdersPostgres) FindAllByPerformerId(performerId int) ([]taskexchange.Order, error) {
	var orders []taskexchange.Order
	fields := []string{
		"ofe.id as \"offer.id\"",
		"ofe.performer_id as \"offer.performer_id\"",
		"ofe.status as \"offer.status\"",
		"ofe.created_at as \"offer.created_at\"",

		"tas.id as \"task.id\"",
		"tas.customer_id as \"task.customer_id\"",
		"tas.status as \"task.status\"",
		"tas.amount as \"task.amount\"",
		"tas.delivery_date as \"task.delivery_date\"",
		"tas.link as \"task.link\"",
		"tas.description as \"task.description\"",
		"tas.created_at as \"task.created_at\"",

		"ord.id as id",
		"ord.task_id as task_id",
		"ord.status as status",
		"ord.canceled_user_id as canceled_user_id",
		"ord.return_comment as return_comment",
		"ord.surrender_comment as surrender_comment",
		"ord.cancel_comment as cancel_comment",
		"ord.created_at as created_at",
		"ord.deleted_at as deleted_at",
	}
	fieldsQuery := strings.Join(fields, ", ")

	query := fmt.Sprintf("SELECT %s FROM %s AS ord JOIN %s AS tas ON ord.task_id = tas.id JOIN %s AS ofe ON ord.offer_id = ofe.id WHERE ord.deleted_at is null AND ofe.performer_id=$1 ORDER BY ord.created_at DESC", fieldsQuery, ordersTable, tasksTable, offersTable)
	err := r.db.Select(&orders, query, performerId)

	return orders, err
}

func (r *OrdersPostgres) CountAllByPerformerId(performerId int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s AS ord JOIN %s AS ofe ON ord.offer_id = ofe.id JOIN %s AS tas ON ord.task_id = tas.id WHERE ord.deleted_at is null AND ofe.performer_id=$1", ordersTable, offersTable, tasksTable)
	err := r.db.QueryRow(query, performerId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *OrdersPostgres) FindAllByCustomerId(customerId int) ([]taskexchange.Order, error) {
	var orders []taskexchange.Order
	fields := []string{
		"ofe.id as \"offer.id\"",
		"ofe.performer_id as \"offer.performer_id\"",
		"ofe.status as \"offer.status\"",
		"ofe.created_at as \"offer.created_at\"",

		"tas.id as \"task.id\"",
		"tas.customer_id as \"task.customer_id\"",
		"tas.status as \"task.status\"",
		"tas.amount as \"task.amount\"",
		"tas.delivery_date as \"task.delivery_date\"",
		"tas.link as \"task.link\"",
		"tas.description as \"task.description\"",
		"tas.created_at as \"task.created_at\"",

		"ord.id as id",
		"ord.task_id as task_id",
		"ord.status as status",
		"ord.canceled_user_id as canceled_user_id",
		"ord.return_comment as return_comment",
		"ord.surrender_comment as surrender_comment",
		"ord.cancel_comment as cancel_comment",
		"ord.created_at as created_at",
		"ord.deleted_at as deleted_at",
	}
	fieldsQuery := strings.Join(fields, ", ")

	query := fmt.Sprintf("SELECT %s FROM %s AS ord JOIN %s AS tas ON ord.task_id = tas.id JOIN %s AS ofe ON ord.offer_id = ofe.id WHERE ord.deleted_at is null AND tas.customer_id=$1 ORDER BY ord.created_at DESC", fieldsQuery, ordersTable, tasksTable, offersTable)
	err := r.db.Select(&orders, query, customerId)

	return orders, err
}

func (r *OrdersPostgres) CountAllByCustomerId(customerId int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s AS ord JOIN %s AS tas ON ord.task_id = tas.id WHERE ord.deleted_at is null AND tas.customer_id=$1 ORDER BY ord.created_at DESC", ordersTable, tasksTable)
	err := r.db.QueryRow(query, customerId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *OrdersPostgres) FindActiveByPerformerId(performerId int) ([]taskexchange.Order, error) {
	var orders []taskexchange.Order
	fields := []string{
		"ofe.id as \"offer.id\"",
		"ofe.performer_id as \"offer.performer_id\"",
		"ofe.status as \"offer.status\"",
		"ofe.created_at as \"offer.created_at\"",

		"ord.id as id",
		"ord.task_id as task_id",
		"ord.status as status",
		"ord.canceled_user_id as canceled_user_id",
		"ord.return_comment as return_comment",
		"ord.surrender_comment as surrender_comment",
		"ord.cancel_comment as cancel_comment",
		"ord.created_at as created_at",
		"ord.deleted_at as deleted_at",
	}
	fieldsQuery := strings.Join(fields, ", ")

	query := fmt.Sprintf("SELECT %s FROM %s AS ord JOIN %s AS ofe ON ord.offer_id = ofe.id WHERE ord.deleted_at is null AND ord.status IN (0, 1) AND ofe.performer_id=$1 ORDER BY ord.created_at DESC", fieldsQuery, ordersTable, offersTable)
	err := r.db.Select(&orders, query, performerId)

	return orders, err
}

func (r *OrdersPostgres) Update(id int, input taskexchange.UpdateOrderInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.TaskId != nil {
		setValues = append(setValues, fmt.Sprintf("task_id=$%d", argId))
		args = append(args, *input.TaskId)
		argId++
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	if input.CanceledUserId != nil {
		setValues = append(setValues, fmt.Sprintf("canceled_user_id=$%d", argId))
		args = append(args, *input.CanceledUserId)
		argId++
	}

	if input.ReturnComment != nil {
		setValues = append(setValues, fmt.Sprintf("return_comment=$%d", argId))
		args = append(args, *input.ReturnComment)
		argId++
	}

	if input.SurrenderComment != nil {
		setValues = append(setValues, fmt.Sprintf("surrender_comment=$%d", argId))
		args = append(args, *input.SurrenderComment)
		argId++
	}

	if input.CancelComment != nil {
		setValues = append(setValues, fmt.Sprintf("cancel_comment=$%d", argId))
		args = append(args, *input.CancelComment)
		argId++
	}

	args = append(args, id)

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", ordersTable, setQuery, argId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *OrdersPostgres) FindOneById(orderId int) (taskexchange.Order, error) {
	var order taskexchange.Order
	fields := []string{
		"ofe.id as \"offer.id\"",
		"ofe.performer_id as \"offer.performer_id\"",
		"ofe.status as \"offer.status\"",
		"ofe.created_at as \"offer.created_at\"",

		"tas.id as \"task.id\"",
		"tas.customer_id as \"task.customer_id\"",
		"tas.status as \"task.status\"",
		"tas.amount as \"task.amount\"",
		"tas.delivery_date as \"task.delivery_date\"",
		"tas.link as \"task.link\"",
		"tas.description as \"task.description\"",
		"tas.created_at as \"task.created_at\"",

		"ord.id as id",
		"ord.task_id as task_id",
		"ord.status as status",
		"ord.canceled_user_id as canceled_user_id",
		"ord.return_comment as return_comment",
		"ord.surrender_comment as surrender_comment",
		"ord.cancel_comment as cancel_comment",
		"ord.created_at as created_at",
		"ord.deleted_at as deleted_at",
	}
	fieldsQuery := strings.Join(fields, ", ")

	query := fmt.Sprintf("SELECT %s FROM %s AS ord JOIN %s AS tas ON ord.task_id = tas.id JOIN %s AS ofe ON ord.offer_id = ofe.id WHERE ord.id = $1", fieldsQuery, ordersTable, tasksTable, offersTable)
	err := r.db.Get(&order, query, orderId)

	return order, err
}
