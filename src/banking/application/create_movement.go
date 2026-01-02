package application

import (
	"time"
	"github.com/google/uuid"
	"github.com/omarracini/rekon_pyme/src/banking/domain"
	sharedDomain "github.com/omarracini/rekon_pyme/src/shared/domain"
)

type CreateMovementRequest struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Concept   string `json:"concept"`
	Type      string `json:"type"`
}

type CreateMovementUseCase struct {
	repo domain.BankRepository
}

func NewCreateMovementUseCase(repo domain.BankRepository) *CreateMovementUseCase {
	return &CreateMovementUseCase{repo: repo}
}

func (uc *CreateMovementUseCase) Execute(req CreateMovementRequest) error {
	movement := &domain.BankMovement{
		ID:            uuid.New().String(),
		AccountID:     req.AccountID,
		Date:          time.Now(),
		Concept:       req.Concept,
		Amount:        sharedDomain.Money{Amount: req.Amount, Currency: sharedDomain.Currency(req.Currency)},
		Type:          domain.MovementType(req.Type),
		IsConciliated: false,
	}
	return uc.repo.Save(movement)
}