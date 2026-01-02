package infrastructure

import (
	"database/sql"

	"github.com/omarracini/rekon_pyme/src/banking/domain"
)

type PostgresBankRepository struct {
	db *sql.DB
}

func NewPostgresBankRepository(db *sql.DB) *PostgresBankRepository {
	return &PostgresBankRepository{db: db}
}

func (r *PostgresBankRepository) Save(m *domain.BankMovement) error {
	query := `INSERT INTO movements (id, account_id, date, concept, amount, currency, type, is_conciliated) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.Exec(query,
		m.ID,
		m.AccountID,
		m.Date,
		m.Concept,
		m.Amount.Amount,
		string(m.Amount.Currency),
		m.Type,
		m.IsConciliated,
	)
	return err
}

func (r *PostgresBankRepository) FindAllMovements(accountID string) ([]domain.BankMovement, error) {
	return []domain.BankMovement{}, nil
}

func (r *PostgresBankRepository) SaveInvoice(i *domain.Invoice) error {
	query := `INSERT INTO invoices (id, number, provider, date, due_date, amount, currency, status) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.Exec(query,
		i.ID,
		i.Number,
		i.Provider,
		i.Date,
		i.DueDate,
		i.Amount.Amount,
		string(i.Amount.Currency),
		i.Status,
	)
	return err
}

func (r *PostgresBankRepository) FindAllInvoices() ([]domain.Invoice, error) {
	rows, err := r.db.Query("SELECT id, number, provider, date, due_date, amount, currency, status FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []domain.Invoice
	for rows.Next() {
		var i domain.Invoice
		// Escaneamos los datos de la fila a la estructura
		err := rows.Scan(&i.ID, &i.Number, &i.Provider, &i.Date, &i.DueDate, &i.Amount.Amount, &i.Amount.Currency, &i.Status)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, i)
	}
	return invoices, nil
}

func (r *PostgresBankRepository) Conciliate(movementID string, invoiceID string) error {
	// Iniciamos la transacción
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Marcar el movimiento como conciliado
	_, err = tx.Exec("UPDATE movements SET is_conciliated = true WHERE id = $1", movementID)
	if err != nil {
		tx.Rollback() // Si falla, deshacemos todo
		return err
	}

	// Marcar la factura como pagada
	_, err = tx.Exec("UPDATE invoices SET status = $1 WHERE id = $2", domain.InvoicePaid, invoiceID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Si ambos salieron bien, confirmamos los cambios en la DB
	return tx.Commit()
}