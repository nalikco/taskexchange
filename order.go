package taskexchange

import "time"

type Order struct {
	Id               int       `json:"id"`
	OfferId          int       `json:"offer_id"`
	Status           int       `json:"status"`
	CanceledUserId   int       `json:"canceled_user_id"`
	ReturnComment    string    `json:"return_comment"`
	SurrenderComment string    `json:"surrender_comment"`
	CancelComment    string    `json:"cancel_comment"`
	CreatedAt        time.Time `json:"created_at"`
	DeletedAt        time.Time `json:"deleted_at"`
}
