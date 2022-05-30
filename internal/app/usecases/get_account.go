package usecases

import (
	"context"

	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

type GetAccountUseCase interface {
	Execute(ctx context.Context, accountID int64) (*dto.Account, error)
}

type getAccountService struct {
	gtw GetAccountGateway
}

func NewGetAccountUseCase(g GetAccountGateway) GetAccountUseCase {
	return &getAccountService{gtw: g}
}

func (s *getAccountService) Execute(ctx context.Context, accountID int64) (*dto.Account, error) {
	account, err := s.gtw.GetAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	dto := dto.Account{
		ID:             account.ID,
		DocumentNumber: account.DocumentNumber}

	return &dto, err
}
