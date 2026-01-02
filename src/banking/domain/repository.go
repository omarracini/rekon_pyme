package domain

type BankRepository interface {
    Save(movement *BankMovement) error
    FindAllMovements(accountID string) ([]BankMovement, error)
    //FindByID(id string) (*BankMovement, error)
}	