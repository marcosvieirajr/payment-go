package datastruct

const TransactionTableName = "transactions"

type Transaction struct {
	ID              int64   `db:"id"`
	AccountID       int64   `db:"account_id"`
	OperationTypeID int     `db:"operation_type_id"`
	Amount          float64 `db:"amount"`
	Audit           Audit
}
