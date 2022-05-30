package usecases

import (
	"context"

	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/domain"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

type CreateAccountUseCase interface {
	Execute(ctx context.Context, acc dto.Account) (*int64, error)
}

type createAccountService struct {
	gCountAcc  CountAccountsByDocumentGateway
	gCreateAcc CreateAccountGateway
}

func NewCreateAccountUseCase(
	gtwCountAcc CountAccountsByDocumentGateway,
	gtwCreateAcc CreateAccountGateway,
) CreateAccountUseCase {
	return &createAccountService{gCountAcc: gtwCountAcc, gCreateAcc: gtwCreateAcc}
}

func (s *createAccountService) Execute(ctx context.Context, acc dto.Account) (*int64, error) {
	account := domain.Account{DocumentNumber: acc.DocumentNumber}
	if err := account.Validate(); err != nil {
		return nil, err
	}

	count, err := s.gCountAcc.CountAccountsByDocument(ctx, acc.DocumentNumber)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, app.ErrAccountAlreadyExists
	}

	return s.gCreateAcc.CreateAccount(ctx, acc)
}
