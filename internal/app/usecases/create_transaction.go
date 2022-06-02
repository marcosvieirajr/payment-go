package usecases

import (
	"context"

	"github.com/marcosvieirajr/payment/internal/app/domain"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

type CreateTransactionUseCase interface {
	Execute(ctx context.Context, transaction dto.Transaction) (*int64, error)
}

type createTransactionService struct {
	gtwGetAcc    GetAccountGateway
	gtwCreateAcc CreateTransactionGateway
}

func NewCreateTransactionUseCase(
	gGetAcc GetAccountGateway,
	gCreateAcc CreateTransactionGateway,
) CreateTransactionUseCase {
	return &createTransactionService{gtwGetAcc: gGetAcc, gtwCreateAcc: gCreateAcc}
}

func (uc *createTransactionService) Execute(ctx context.Context, dto dto.Transaction) (*int64, error) {
	_, err := uc.gtwGetAcc.GetAccount(ctx, dto.Account.ID)
	if err != nil {
		return nil, err
	}

	tr := domain.Transaction{
		OperationType: dto.OperationType,
		Amount:        dto.Amount,
	}

	err = tr.Validate()
	if err != nil {
		return nil, err
	}

	return uc.gtwCreateAcc.CreateTransaction(ctx, dto)
}
