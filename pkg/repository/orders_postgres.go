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
