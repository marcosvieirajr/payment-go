package domain

import (
	"reflect"
	"testing"

	"github.com/marcosvieirajr/payment/internal/app/domain/operations"
)

func TestAccount_Validate(t *testing.T) {
	type fields struct {
		DocumentNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "valid account",
			fields:  fields{DocumentNumber: "12345678900"},
			wantErr: false,
		},
		{
			name:    "invalid account",
			fields:  fields{DocumentNumber: "12345"},
			wantErr: true,
		},
		{
			name:    "invalid account 2",
			fields:  fields{DocumentNumber: "123456789001"},
			wantErr: true,
		},
		{
			name:    "invalid account 3",
			fields:  fields{DocumentNumber: "12345 78900"},
			wantErr: true,
		},
		{
			name:    "invalid account 4",
			fields:  fields{DocumentNumber: "a1234578900"},
			wantErr: true,
		},
		{
			name:    "invalid account 5",
			fields:  fields{DocumentNumber: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				DocumentNumber: tt.fields.DocumentNumber,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Account.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_Transact(t *testing.T) {
	type fields struct {
		ID int64
	}
	type args struct {
		o operations.Operator
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
		wantErr bool
	}{
		{
			name:    "valid transaction",
			fields:  fields{ID: 1},
			args:    args{o: operations.MakeOperator(4, 100.00)},
			want:    &Transaction{AccountID: 1, OperationType: operations.PAYMENT, Amount: 100.00},
			wantErr: false,
		},
		{
			name:    "invalid transaction",
			fields:  fields{ID: 1},
			args:    args{o: operations.MakeOperator(4, -100.00)},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				ID: tt.fields.ID,
			}
			got, err := a.Transact(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Account.Transact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
