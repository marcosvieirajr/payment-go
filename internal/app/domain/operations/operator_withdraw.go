package operations

import "github.com/marcosvieirajr/payment/internal/app"

type Withdraw struct {
	Amount float64
}

func (o *Withdraw) Validate() error {
	if o.Amount >= 0 {
		return app.ErrInvalidAmount
	}
	return nil
}

func (o *Withdraw) GetAmount() float64 {
	return o.Amount
}

func (Withdraw) GetOperationType() OperationType {
	return WITHDRAWAL
}
