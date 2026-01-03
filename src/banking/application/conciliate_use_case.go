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

	// Buscar el movimiento
	movement, err := uc.repo.FindMovementByID(req.MovementID)
	if err != nil {
		return errors.New("movimiento no encontrado")
	}

	// Buscar la factura
	invoice, err := uc.repo.FindInvoiceByID(req.InvoiceID)
	if err != nil {
		return errors.New("factura no encontrada")
	}

	// Validar montos
	if movement.Amount.Amount != invoice.Amount.Amount {
		return errors.New("los montos del movimiento y la factura deben ser iguales")
	}

	// Validar monedas
	if movement.Amount.Currency != invoice.Amount.Currency {
		return errors.New("error: las monedas no coinciden")
	}

	// Validar estado de la conciliación
	if invoice.Status == domain.InvoicePaid {
		return errors.New("la factura ya está pagada")
	}

	// Paso a la DB
	return uc.repo.Conciliate(req.MovementID, req.InvoiceID)
}
