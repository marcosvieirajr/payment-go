package dto

import "github.com/marcosvieirajr/payment/internal/app/domain/operations"

type Operation struct {
	ID operations.OperationType
}
