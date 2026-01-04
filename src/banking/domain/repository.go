package domain

type BankRepository interface {
	Save(movement *BankMovement) error
	FindAllMovements(accountID string) ([]BankMovement, error)

	SaveInvoice(invoice *Invoice) error
	FindAllInvoices() ([]Invoice, error)

	Conciliate(movementID string, invoiceID string) error

	//Filtros de busqueda
	FindPendingMovements() ([]BankMovement, error)
	FindPendingInvoices() ([]Invoice, error)
	FindMovementByID(id string) (*BankMovement, error)
	FindInvoiceByID(id string) (*Invoice, error)
	GetUnconciliatedMovements() ([]*BankMovement, error)
	GetDashboardSummary() ([]DashboardSummary, error)
}
