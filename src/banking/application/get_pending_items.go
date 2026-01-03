// Caso de uso para obtener movimientos yfacturas pendientes por conciliar
package application

import "github.com/omarracini/rekon_pyme/src/banking/domain"

type PendingItemsResponse struct {
	Movements []domain.BankMovement `json:"pending_movements"`
	Invoices  []domain.Invoice      `json:"pending_invoices"`
}

type GetPendingItemsUseCase struct {
	repo domain.BankRepository
}

func NewGetPendingItemsUseCase(repo domain.BankRepository) *GetPendingItemsUseCase {
	return &GetPendingItemsUseCase{repo: repo}
}

func (uc *GetPendingItemsUseCase) Execute() (PendingItemsResponse, error) {
	movements, err := uc.repo.FindPendingMovements()
	if err != nil {
		return PendingItemsResponse{}, err
	}

	invoices, err := uc.repo.FindPendingInvoices()
	if err != nil {
		return PendingItemsResponse{}, err
	}

	return PendingItemsResponse{
		Movements: movements,
		Invoices:  invoices,
	}, nil
}
