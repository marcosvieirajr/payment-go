package dto

import "github.com/marcosvieirajr/payment/internal/app/domain/operations"

type Transaction struct {
	ID            int64
	Account       Account
	OperationType operations.OperationType
	Amount        float64
	CreatedFrom   string
}
