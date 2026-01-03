package application

import "github.com/omarracini/rekon_pyme/src/banking/domain"

type GetPendingMovementsUseCase struct {
	repo domain.BankRepository
}

func NewGetPendingMovementsUseCase(repo domain.BankRepository) *GetPendingMovementsUseCase {
	return &GetPendingMovementsUseCase{repo: repo}
}

func (uc *GetPendingMovementsUseCase) Execute() ([]*domain.BankMovement, error) {
	return uc.repo.GetUnconciliatedMovements()
}
