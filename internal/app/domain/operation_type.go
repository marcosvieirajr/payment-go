package domain

import "github.com/marcosvieirajr/payment/internal/app"

type OperationType int

const (
	CASH_PURCHASE OperationType = iota + 1
	INSTALLMENT_PURCHASE
	WITHDRAWAL
	PAYMENT
)

var debito = func(amount float64) bool {
	return amount < 0.00
}

var credito = func(amount float64) bool {
	return amount > 0.00
}

func (o OperationType) Validate(amount float64) error {
	var validatorFunc func(amount float64) bool
	switch o {
	case CASH_PURCHASE, INSTALLMENT_PURCHASE, WITHDRAWAL:
		validatorFunc = debito //(amount)
	case PAYMENT:
		validatorFunc = credito //(amount)
	default:
		return app.ErrInvalidOperationType
	}

	if !validatorFunc(amount) {
		return app.ErrInvalidAmount
	}

	return nil
}
