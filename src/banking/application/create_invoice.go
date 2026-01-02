package application

import (
	"time"
	"github.com/google/uuid"
	"github.com/omarracini/rekon_pyme/src/banking/domain"
	sharedDomain "github.com/omarracini/rekon_pyme/src/shared/domain"
)

type CreateInvoiceRequest struct {
	Number   string `json:"number"`
	Provider string `json:"provider"`
	DueDate  string `json:"due_date"` // Formato "2006-01-02"
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type CreateInvoiceUseCase struct {
	repo domain.BankRepository
}

func NewCreateInvoiceUseCase(repo domain.BankRepository) *CreateInvoiceUseCase {
	return &CreateInvoiceUseCase{repo: repo}
}

func (uc *CreateInvoiceUseCase) Execute(req CreateInvoiceRequest) error {
	// Parsear la fecha de vencimiento
	dueDate, _ := time.Parse("2006-01-02", req.DueDate)

	invoice := &domain.Invoice{
		ID:       uuid.New().String(),
		Number:   req.Number,
		Provider: req.Provider,
		Date:     time.Now(),
		DueDate:  dueDate,
		Amount:   sharedDomain.Money{Amount: req.Amount, Currency: sharedDomain.Currency(req.Currency)},
		Status:   domain.InvoicePending,
	}

	return uc.repo.SaveInvoice(invoice)
}