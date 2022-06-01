package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/domain"
	"github.com/marcosvieirajr/payment/internal/app/usecases"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

type AccountRepository interface {
	usecases.CreateAccountGateway
	usecases.GetAccountGateway
	usecases.CountAccountsByDocumentGateway
}

type account struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *account {
	return &account{db}
}

func (rep *account) GetAccount(ctx context.Context, accountID int64) (*domain.Account, error) {
	var acc domain.Account
	sql := `select id, document_number from accounts where id = $1`
	err := rep.db.QueryRow(sql, accountID).
		Scan(&acc.ID, &acc.DocumentNumber)
	if err != nil {
		return nil, app.ErrAccountNotFound //fmt.Errorf("cannot scan account: %v", err)
	}
	return &acc, nil
}

func (rep *account) CountAccountsByDocument(ctx context.Context, doc string) (int, error) {
	var count int
	sql := `select count(1) from accounts where document_number = $1`
	err := rep.db.QueryRow(sql, doc).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("cannot scan count account: %v", err)
	}
	return count, nil
}

func (rep *account) CreateAccount(ctx context.Context, a dto.Account) (*int64, error) {
	var id int64
	sql := `insert into accounts (document_number, created_from) values ($1, $2) returning id;`
	err := rep.db.QueryRow(sql, a.DocumentNumber, a.CreatedFrom).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot insert account: %v", err)
	}
	return &id, nil
}
