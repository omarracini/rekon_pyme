package application

import "github.com/omarracini/rekon_pyme/src/banking/domain"

type SuggestCategoryUseCase struct {
	aiService domain.AIService
}

func NewSuggestCategoryUseCase(ai domain.AIService) *SuggestCategoryUseCase {
	return &SuggestCategoryUseCase{aiService: ai}
}

func (uc *SuggestCategoryUseCase) Execute(concept string) (*domain.AICategorySuggestion, error) {
	return uc.aiService.CategorizeMovement(concept)
}
