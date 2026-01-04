package infrastructure

import (
	"strings"

	"github.com/omarracini/rekon_pyme/src/banking/domain"
)

type AIMockService struct{}

func NewAIClient() *AIMockService {
	return &AIMockService{}
}

func (s *AIMockService) CategorizeMovement(concept string) (*domain.AICategorySuggestion, error) {
	conceptLower := strings.ToLower(concept)

	// Implementación de la interfaz, Simulación de lógica de IA (Pattern Matching avanzado)
	if strings.Contains(conceptLower, "starbucks") || strings.Contains(conceptLower, "restaurante") {
		return &domain.AICategorySuggestion{Category: "Gastos de Representación", Confidence: 0.98, Reason: "Detectado establecimiento de alimentos"}, nil
	}
	if strings.Contains(conceptLower, "aws") || strings.Contains(conceptLower, "cloud") || strings.Contains(conceptLower, "server") {
		return &domain.AICategorySuggestion{Category: "Infraestructura TI", Confidence: 0.95, Reason: "Identificado proveedor de servicios en la nube"}, nil
	}

	// Categoría por defecto si la "IA" no está segura
	return &domain.AICategorySuggestion{Category: "Otros Operativos", Confidence: 0.45, Reason: "Concepto genérico detectado"}, nil
}
