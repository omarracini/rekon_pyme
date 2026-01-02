package domain

type BankRepository interface {
    Save(movement *BankMovement) error
    FindAllMovements(accountID string) ([]BankMovement, error)
    
    SaveInvoice(invoice *Invoice) error
	FindAllInvoices() ([]Invoice, error)
}	