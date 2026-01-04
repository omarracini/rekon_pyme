package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/omarracini/rekon_pyme/src/banking/domain"
)

type PostgresBankRepository struct {
	db *sql.DB
}

// Constructor para PostgresBankRepository
func NewPostgresBankRepository(db *sql.DB) *PostgresBankRepository {
	return &PostgresBankRepository{db: db}
}

// Guardar un movimiento bancario
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

// Obtener todos los movimientos de una cuenta
func (r *PostgresBankRepository) FindAllMovements(accountID string) ([]domain.BankMovement, error) {
	return []domain.BankMovement{}, nil
}

// Guardar una factura
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

// Obtener todas las facturas
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

// Conciliar un movimiento con una factura
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

	// Insertar en la tabla de auditoría
	_, err = tx.Exec("INSERT INTO conciliations (movement_id, invoice_id) VALUES ($1, $2)", movementID, invoiceID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Si ambos salieron bien, confirmamos los cambios en la DB
	return tx.Commit()
}

// Movimientos pendientes de conciliación
func (r *PostgresBankRepository) FindPendingMovements() ([]domain.BankMovement, error) {
	query := `SELECT id, account_id, date, concept, amount, currency, type
          FROM movements
          WHERE is_conciliated = false
          ORDER BY date DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []domain.BankMovement
	for rows.Next() {
		var m domain.BankMovement
		err := rows.Scan(&m.ID, &m.AccountID, &m.Date, &m.Concept, &m.Amount.Amount, &m.Amount.Currency, &m.Type, &m.IsConciliated)
		if err != nil {
			return nil, err
		}
		movements = append(movements, m)
	}
	return movements, nil
}

// Facturas pendientes de conciliación
func (r *PostgresBankRepository) FindPendingInvoices() ([]domain.Invoice, error) {
	query := `SELECT id, number, provider, date, due_date, amount, currency, status 
			  FROM invoices WHERE status = $1`

	rows, err := r.db.Query(query, domain.InvoicePending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []domain.Invoice
	for rows.Next() {
		var i domain.Invoice
		err := rows.Scan(&i.ID, &i.Number, &i.Provider, &i.Date, &i.DueDate, &i.Amount.Amount, &i.Amount.Currency, &i.Status)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, i)
	}
	return invoices, nil
}

func (r *PostgresBankRepository) FindMovementByID(id string) (*domain.BankMovement, error) {
	var m domain.BankMovement
	cleanID := strings.TrimSpace(id)
	err := r.db.QueryRow("SELECT id, amount, currency FROM movements WHERE id = $1", cleanID).
		Scan(&m.ID, &m.Amount.Amount, &m.Amount.Currency)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("DB DEBUG: No se encontró el ID [%s]\n", cleanID)
			return nil, errors.New("movimiento no encontrado")
		}
		return nil, err
	}
	return &m, nil
}

// Buscar factura por ID
func (r *PostgresBankRepository) FindInvoiceByID(id string) (*domain.Invoice, error) {
	var i domain.Invoice
	err := r.db.QueryRow("SELECT id, amount, currency, status FROM invoices WHERE id = $1", id).Scan(&i.ID, &i.Amount.Amount, &i.Amount.Currency, &i.Status)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (r *PostgresBankRepository) GetUnconciliatedMovements() ([]*domain.BankMovement, error) {
	query := `SELECT id, account_id, date, concept, amount, currency, type
          FROM movements
          WHERE is_conciliated = false
          ORDER BY date DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []*domain.BankMovement
	for rows.Next() {
		m := &domain.BankMovement{}
		err := rows.Scan(
			&m.ID,
			&m.AccountID,
			&m.Date,
			&m.Concept,
			&m.Amount.Amount,
			&m.Amount.Currency,
			&m.Type,
		)
		if err != nil {
			return nil, err
		}
		movements = append(movements, m)
	}

	return movements, nil
}

func (r *PostgresBankRepository) GetDashboardSummary() (*domain.DashboardSummary, error) {
    summary := &domain.DashboardSummary{Currency: "USD"} // Asumimos USD por ahora

    // Consulta para sumar totales por estado
    query := `
        SELECT 
            COALESCE(SUM(CASE WHEN is_conciliated = true THEN amount ELSE 0 END), 0) as reconciled,
            COALESCE(SUM(CASE WHEN is_conciliated = false THEN amount ELSE 0 END), 0) as pending_mov
        FROM movements`
    
    err := r.db.QueryRow(query).Scan(&summary.TotalReconciled, &summary.PendingMovements)
    if err != nil {
        return nil, err
    }

    // Consulta para sumar facturas pendientes
    queryInv := `SELECT COALESCE(SUM(amount), 0) FROM invoices WHERE status = 'PENDING'`
    err = r.db.QueryRow(queryInv).Scan(&summary.PendingInvoices)
    if err != nil {
        return nil, err
    }

    return summary, nil
}