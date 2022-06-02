package domain

import "testing"

func TestOperationType_Validate(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		o       OperationType
		args    args
		wantErr bool
	}{
		{
			name:    "valid PAYMENT Operation",
			o:       PAYMENT,
			args:    args{amount: 100.00},
			wantErr: false,
		},
		{
			name:    "valid WITHDRAWAL Operation",
			o:       WITHDRAWAL,
			args:    args{amount: -1.00},
			wantErr: false,
		},
		{
			name:    "valid INSTALLMENT_PURCHASE Operation",
			o:       INSTALLMENT_PURCHASE,
			args:    args{amount: -100.00},
			wantErr: false,
		},
		{
			name:    "valid CASH_PURCHASE Operation",
			o:       CASH_PURCHASE,
			args:    args{amount: -100.00},
			wantErr: false,
		},
		{
			name:    "invalid PAYMENT Operation",
			o:       PAYMENT,
			args:    args{amount: -100.00},
			wantErr: true,
		},
		{
			name:    "invalid PAYMENT Operation",
			o:       PAYMENT,
			args:    args{amount: 0.00},
			wantErr: true,
		},
		{
			name:    "invalid WITHDRAWAL Operation",
			o:       WITHDRAWAL,
			args:    args{amount: 0.00},
			wantErr: true,
		},
		{
			name:    "invalid INSTALLMENT_PURCHASE Operation",
			o:       INSTALLMENT_PURCHASE,
			args:    args{amount: 1.00},
			wantErr: true,
		},
		{
			name:    "invalid CASH_PURCHASE Operation",
			o:       CASH_PURCHASE,
			args:    args{amount: 100.00},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.Validate(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("OperationType.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
