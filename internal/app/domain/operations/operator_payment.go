package operations

import "github.com/marcosvieirajr/payment/internal/app"

type Payment struct {
	Amount float64
}

func (o *Payment) Validate() error {
	if o.Amount <= 0 {
		return app.ErrInvalidAmount
	}
	return nil
}

func (o *Payment) GetAmount() float64 {
	return o.Amount
}

func (Payment) GetOperationType() OperationType {
	return PAYMENT
}
