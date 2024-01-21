package virtualwallet

import "github.com/google/uuid"

type Balance struct {
	ID     uuid.UUID `json:"id"`
	Amount int       `json:"amount"`
	UserID uuid.UUID `json:"user_id"`
}
