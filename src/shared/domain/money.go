package domain

import "errors"

type Currency string

type Money struct {
	Amount   int64 		// Guardamos en centavos (ej: 1000 = 10.00)
	Currency Currency	// Ejemplo: "USD", "MXN", "CLP" 
}

func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, errors.New("las divisas no coinciden")
	}
	return Money{Amount: m.Amount + other.Amount, Currency: m.Currency}, nil
}
