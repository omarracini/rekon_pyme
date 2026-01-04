package application

import "github.com/omarracini/rekon_pyme/src/banking/domain"

type GetDashboardUseCase struct {
	repo domain.BankRepository
}

func NewGetDashboardUseCase(repo domain.BankRepository) *GetDashboardUseCase {
	return &GetDashboardUseCase{repo: repo}
}

func (uc *GetDashboardUseCase) Execute() ([]domain.DashboardSummary, error) {
	return uc.repo.GetDashboardSummary()
}
