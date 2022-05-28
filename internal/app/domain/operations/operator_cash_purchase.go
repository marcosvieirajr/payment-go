package operations

import "github.com/marcosvieirajr/payment/internal/app"

type CashPurchase struct {
	Amount float64
}

func (o *CashPurchase) Validate() error {
	if o.Amount >= 0 {
		return app.ErrInvalidAmount
	}
	return nil
}

func (o *CashPurchase) GetAmount() float64 {
	return o.Amount
}

func (CashPurchase) GetOperationType() OperationType {
	return CASH_PURCHASE
}
