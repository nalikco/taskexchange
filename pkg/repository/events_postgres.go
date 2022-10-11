package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"taskexchange"
)

type EventsPostgres struct {
	db *sqlx.DB
}

func NewEventsPostgres(db *sqlx.DB) *EventsPostgres {
	return &EventsPostgres{db: db}
}

func (r *EventsPostgres) Create(event taskexchange.Event) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, message, link, created_at) VALUES ($1, $2, $3, $4) RETURNING id", eventsTable)
	row := r.db.QueryRow(query, event.UserId, event.Message, event.Link, event.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *EventsPostgres) FindPolling(userId, id int) ([]taskexchange.Event, error) {
	var events []taskexchange.Event
	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null AND viewed_at is null AND id > $1 AND user_id=$2 ORDER BY id DESC", eventsTable)
	err := r.db.Select(&events, query, id, userId)

	return events, err
}

func (r *EventsPostgres) FindNew(userId int) ([]taskexchange.Event, error) {
	var events []taskexchange.Event
	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null AND user_id=$1 AND viewed_at is null ORDER BY id DESC", eventsTable)
	err := r.db.Select(&events, query, userId)

	return events, err
}

func (r *EventsPostgres) FindAll(userId, limit, offset int) ([]taskexchange.Event, error) {
	var events []taskexchange.Event
	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null AND user_id=$1 ORDER BY id DESC LIMIT %d OFFSET %d", eventsTable, limit, offset)
	err := r.db.Select(&events, query, userId)

	return events, err
}

func (r *EventsPostgres) CountAll(userId int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE deleted_at is null AND user_id=$1", eventsTable)
	err := r.db.QueryRow(query, userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EventsPostgres) FindLastUser(userId int) (taskexchange.Event, error) {
	var event taskexchange.Event

	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null AND user_id=$1 ORDER BY id DESC", eventsTable)

	err := r.db.Get(&event, query, userId)

	return event, err
}

func (r *EventsPostgres) ViewAll(userId int) error {
	query := fmt.Sprintf("UPDATE %s SET viewed_at=now() WHERE user_id=$1 AND deleted_at is null AND viewed_at is null", eventsTable)
	_, err := r.db.Exec(query, userId)

	return err
}

func (r *EventsPostgres) View(userId, id int) error {
	query := fmt.Sprintf("UPDATE %s SET viewed_at=now() WHERE id=$1 AND user_id=$2 AND deleted_at is null", eventsTable)
	_, err := r.db.Exec(query, id, userId)

	return err
}

func (r *EventsPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=now() WHERE id=$1 AND user_id=$2 AND deleted_at is null", eventsTable)
	_, err := r.db.Exec(query, id, userId)

	return err
}
