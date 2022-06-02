package dto

type Transaction struct {
	ID            int64
	Account       Account
	OperationType int
	Amount        float64
	CreatedFrom   string
}
