package domain

import (
	"time"

	"github.com/omarracini/rekon_pyme/src/shared/domain"
)

type MovementType string

const (
	Credit MovementType = "ABONO"
	Debit  MovementType = "CARGO"
)

type BankMovement struct {
	ID        string
	AccountID string
	Date      time.Time
	Concept   string
	Amount    domain.Money
	Type      MovementType
	IsConciliated bool
}
