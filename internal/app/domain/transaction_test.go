package domain

import (
	"testing"
)

func TestTransaction_Validate(t *testing.T) {
	type fields struct {
		OperationType int
		Amount        float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "valid transaction",
			fields:  fields{OperationType: 4, Amount: 100.00},
			wantErr: false,
		},
		{
			name:    "invalid transaction",
			fields:  fields{OperationType: 4, Amount: -100.00},
			wantErr: true,
		},
		{
			name:    "invalid OperationType",
			fields:  fields{OperationType: 0, Amount: 100.00},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				OperationType: tt.fields.OperationType,
				Amount:        tt.fields.Amount,
			}
			if err := tr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
