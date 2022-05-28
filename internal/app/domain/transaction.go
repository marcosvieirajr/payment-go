package domain

import "github.com/marcosvieirajr/payment/internal/app/domain/operations"

type Transaction struct {
	ID            int64
	AccountID     int64
	OperationType operations.OperationType
	Amount        float64
}
