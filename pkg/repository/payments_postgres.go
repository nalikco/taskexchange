package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"taskexchange"
	"time"
)

type PaymentsPostgres struct {
	db *sqlx.DB
}

type paymentDb struct {
	ID        int        `db:"id"`
	UserId    int        `db:"user_id"`
	Type      int        `db:"type"`
	Comment   string     `db:"comment"`
	Sum       float64    `db:"sum"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func NewPaymentsPostgres(db *sqlx.DB) *PaymentsPostgres {
	return &PaymentsPostgres{db: db}
}

func (r *PaymentsPostgres) Create(payment taskexchange.Payment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, type, comment, sum) VALUES ($1, $2, $3, $4) RETURNING id", paymentsTable)
	row := r.db.QueryRow(query, payment.User.Id, payment.Type, payment.Comment, payment.Sum)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PaymentsPostgres) GetByUser(user taskexchange.User) ([]taskexchange.Payment, error) {
	var payments []taskexchange.Payment
	var paymentsDb []paymentDb

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 ORDER BY created_at DESC", paymentsTable)
	err := r.db.Select(&paymentsDb, query, user.Id)
	if err != nil {
		return payments, err
	}

	for _, payment := range paymentsDb {
		payments = append(payments, taskexchange.Payment{
			ID:        payment.ID,
			User:      user,
			Type:      payment.Type,
			Comment:   payment.Comment,
			Sum:       payment.Sum,
			CreatedAt: payment.CreatedAt,
			DeletedAt: payment.DeletedAt,
		})
	}

	return payments, nil
}
