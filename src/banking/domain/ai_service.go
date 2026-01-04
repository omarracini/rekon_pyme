package domain

type AICategorySuggestion struct {
	Category   string  `json:"category"`
	Confidence float64 `json:"confidence"`
	Reason     string  `json:"reason"`
}

type AIService interface {
	CategorizeMovement(concept string) (*AICategorySuggestion, error)
}