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