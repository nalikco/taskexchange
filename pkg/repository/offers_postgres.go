package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"taskexchange"
)

type OffersPostgres struct {
	db *sqlx.DB
}

func NewOffersPostgres(db *sqlx.DB) *OffersPostgres {
	return &OffersPostgres{
		db: db,
	}
}

func (r *OffersPostgres) Create(performerId, taskId int) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (performer_id, task_id) VALUES ($1, $2) RETURNING id", offersTable)
	row := r.db.QueryRow(query, performerId, taskId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OffersPostgres) Update(offerId int, input taskexchange.UpdateOfferInput) error {
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE id=$2", offersTable)
	_, err := r.db.Exec(query, input.Status, offerId)

	return err
}

func (r *OffersPostgres) FindByTaskAndStatus(taskId, status int) ([]taskexchange.Offer, error) {
	var offers []taskexchange.Offer

	query := fmt.Sprintf("SELECT * FROM %s WHERE task_id=$1 AND status=$2 AND deleted_at IS NULL", offersTable)
	err := r.db.Select(&offers, query, taskId, status)

	return offers, err
}

func (r *OffersPostgres) FindAllByTask(taskId int) ([]taskexchange.Offer, error) {
	var offers []taskexchange.Offer
	fields := []string{
		"u.id as \"performer.id\"",
		"u.username as \"performer.username\"",
		"u.type as \"performer.type\"",
		"u.last_online as \"performer.last_online\"",
		"u.created_at as \"performer.created_at\"",

		"o.id as id",
		"o.task_id as task_id",
		"o.status as status",
		"o.created_at as created_at",
		"o.deleted_at as deleted_at",
	}
	fieldsQuery := strings.Join(fields, ", ")

	query := fmt.Sprintf("SELECT %s FROM %s AS o JOIN %s AS u ON o.performer_id = u.id WHERE o.deleted_at is null AND o.task_id = $1 ORDER BY o.created_at DESC", fieldsQuery, offersTable, usersTable)
	err := r.db.Select(&offers, query, taskId)

	return offers, err
}

func (r *OffersPostgres) FindPerformerActiveOffers(performerId int) ([]taskexchange.Offer, error) {
	var offers []taskexchange.Offer

	query := fmt.Sprintf("SELECT * FROM %s WHERE performer_id=$1 AND status=0 AND deleted_at IS NULL ORDER BY created_at DESC", offersTable)
	err := r.db.Select(&offers, query, performerId)

	return offers, err
}

func (r *OffersPostgres) FindOneById(offerId int) (taskexchange.Offer, error) {
	var offer taskexchange.Offer

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND deleted_at IS NULL", offersTable)
	err := r.db.Get(&offer, query, offerId)

	return offer, err
}

func (r *OffersPostgres) FindOneByPerformerIdAndTaskIdAndStatus(performerId, taskId, status int) (taskexchange.Offer, error) {
	var offer taskexchange.Offer

	query := fmt.Sprintf("SELECT * FROM %s WHERE performer_id=$1 AND task_id=$2 AND status=$3 AND deleted_at IS NULL", offersTable)
	err := r.db.Get(&offer, query, performerId, taskId, status)

	return offer, err
}
