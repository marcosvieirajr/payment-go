package domain

type Transaction struct {
	ID            int64
	AccountID     int64
	OperationType int
	Amount        float64
}

func (t *Transaction) Validate() error {
	ot := OperationType(t.OperationType)
	err := ot.Validate(t.Amount)
	if err != nil {
		return err
	}

	return nil
}
