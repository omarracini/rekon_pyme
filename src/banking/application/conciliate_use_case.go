package application

import (
	"errors"
	"github.com/omarracini/rekon_pyme/src/banking/domain"
)

type ConciliateRequest struct {
	MovementID string `json:"movement_id"`
	InvoiceID  string `json:"invoice_id"`
}

type ConciliateUseCase struct {
	repo domain.BankRepository
}

func NewConciliateUseCase(repo domain.BankRepository) *ConciliateUseCase {
	return &ConciliateUseCase{repo: repo}
}

func (uc *ConciliateUseCase) Execute(req ConciliateRequest) error {
	// Aquí podrías añadir validaciones extra, como:
	// - ¿El movimiento ya estaba conciliado?
	// - ¿La factura ya estaba pagada?
	// - ¿Los montos coinciden? (Opcional por ahora)
    
	if req.MovementID == "" || req.InvoiceID == "" {
		return errors.New("el ID del movimiento y de la factura son obligatorios")
	}

	return uc.repo.Conciliate(req.MovementID, req.InvoiceID)
}