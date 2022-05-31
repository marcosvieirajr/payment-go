package usecases

import (
	"context"

	"github.com/marcosvieirajr/payment/internal/app/domain/operations"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
	"github.com/sirupsen/logrus"
)

type CreateTransactionUseCase interface {
	Execute(ctx context.Context, transaction dto.Transaction) (*int64, error)
}

type createTransactionService struct {
	gtwGetAcc    GetAccountGateway
	gtwCreateAcc CreateTransactionGateway
}

func NewCreateTransactionUseCase(l *logrus.Logger,
	gGetAcc GetAccountGateway,
	gCreateAcc CreateTransactionGateway,
) CreateTransactionUseCase {
	return &createTransactionService{gtwGetAcc: gGetAcc, gtwCreateAcc: gCreateAcc}
}

func (uc *createTransactionService) Execute(ctx context.Context, t dto.Transaction) (*int64, error) {
	account, err := uc.gtwGetAcc.GetAccount(ctx, t.Account.ID)
	if err != nil {
		return nil, err
	}

	operation := operations.MakeOperator(t.OperationType, t.Amount)

	transaction, err := account.Transact(operation)
	if err != nil {
		return nil, err
	}

	dto := dto.Transaction{
		Account:       dto.Account{ID: transaction.AccountID},
		OperationType: transaction.OperationType,
		Amount:        transaction.Amount,
		CreatedFrom:   t.CreatedFrom}

	return uc.gtwCreateAcc.CreateTransaction(ctx, dto)
}
