package domain

import (
	"time"
	sharedDomain "github.com/omarracini/rekon_pyme/src/shared/domain"
)

type InvoiceStatus string

const (
	InvoicePending   InvoiceStatus = "PENDIENTE"
	InvoicePaid      InvoiceStatus = "PAGADA"
	InvoiceCancelled InvoiceStatus = "CANCELADA"
)

type Invoice struct {
	ID        string        // Usaremos UUID como string
	Number    string        // Ej: "FAC-2024-001"
	Provider  string        
	Date      time.Time
	DueDate   time.Time     
	Amount    sharedDomain.Money         // Reutilizamos el objeto Money de shared
	Status    InvoiceStatus
}