package domain

import (
	"regexp"

	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/domain/operations"
)

type Account struct {
	ID             int64
	DocumentNumber string
}

func (a *Account) Transact(o operations.Operator) (*Transaction, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	transaction := Transaction{
		AccountID:     a.ID,
		OperationType: o.GetOperationType(),
		Amount:        o.GetAmount(),
	}

	return &transaction, nil
}

func (a *Account) Validate() error {
	match, _ := regexp.MatchString("^[0-9]{11}$", a.DocumentNumber)
	if !match {
		return app.ErrDocumentNumberIsInvalid
	}

	return nil
}
