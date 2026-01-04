package application

import (
	"github.com/omarracini/rekon_pyme/src/banking/domain"
)

type GetPendingInvoicesUseCase struct {
	repo domain.BankRepository
}

func NewGetPendingInvoicesUseCase(repo domain.BankRepository) *GetPendingInvoicesUseCase {
	return &GetPendingInvoicesUseCase{repo: repo}
}

func (uc *GetPendingInvoicesUseCase) Execute() ([]domain.Invoice, error) {
	return uc.repo.FindPendingInvoices()
}
