package operations

type OperationType int

const (
	CASH_PURCHASE OperationType = iota + 1
	INSTALLMENT_PURCHASE
	WITHDRAWAL
	PAYMENT
)

type Operator interface {
	Validate() error
	GetAmount() float64
	GetOperationType() OperationType
}

func MakeOperator(ot OperationType, a float64) Operator {
	var o Operator

	switch ot {
	case CASH_PURCHASE:
		o = &CashPurchase{Amount: a}

	case INSTALLMENT_PURCHASE:
		o = &InstallmentPurchase{Amount: a}

	case WITHDRAWAL:
		o = &Withdraw{Amount: a}

	case PAYMENT:
		o = &Payment{Amount: a}

	default:
		return nil
	}

	return o
}
