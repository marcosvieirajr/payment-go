package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marcosvieirajr/payment/internal/app/usecases"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

type TransactionRepository interface {
	usecases.CreateTransactionGateway
}

type transaction struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transaction {
	return &transaction{db}
}

func (rep *transaction) CreateTransaction(ctx context.Context, t dto.Transaction) (*int64, error) {
	var id int64
	sql := `insert into transactions (account_id, operation_type_id, amount, created_from) values ($1, $2, $3, $4) returning id;`
	err := rep.db.QueryRow(sql, t.Account.ID, t.OperationType, t.Amount, t.CreatedFrom).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot insert transaction: %v", err)
	}
	return &id, nil
}
