package usecases

import (
	"context"

	"github.com/marcosvieirajr/payment/internal/app/domain"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

// account gateways
type GetAccountGateway interface {
	GetAccount(ctx context.Context, accountID int64) (*domain.Account, error)
}

type CreateAccountGateway interface {
	CreateAccount(ctx context.Context, a dto.Account) (*int64, error)
}

type CountAccountsByDocumentGateway interface {
	CountAccountsByDocument(ctx context.Context, doc string) (int, error)
}

// transaction gateways
type CreateTransactionGateway interface {
	CreateTransaction(ctx context.Context, t dto.Transaction) (*int64, error)
}
