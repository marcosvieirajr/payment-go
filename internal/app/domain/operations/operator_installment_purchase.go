package operations

import "github.com/marcosvieirajr/payment/internal/app"

type InstallmentPurchase struct {
	Amount float64
}

func (o *InstallmentPurchase) Validate() error {
	if o.Amount >= 0 {
		return app.ErrInvalidAmount
	}
	return nil
}

func (o *InstallmentPurchase) GetAmount() float64 {
	return o.Amount
}

func (InstallmentPurchase) GetOperationType() OperationType {
	return INSTALLMENT_PURCHASE
}
