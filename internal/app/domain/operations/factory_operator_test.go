package operations

import (
	"reflect"
	"testing"
)

func TestMakeOperator(t *testing.T) {
	type args struct {
		o OperationType
		a float64
	}
	tests := []struct {
		name string
		args args
		want Operator
	}{
		{
			name: "Invalid OperationType",
			args: args{
				o: OperationType(5),
				a: 99.00,
			},
			want: nil,
		},
		{
			name: "Withdraw",
			args: args{
				o: WITHDRAWAL,
				a: -99.00,
			},
			want: &Withdraw{Amount: -99.00},
		},
		{
			name: "Payment",
			args: args{
				o: PAYMENT,
				a: 99.00,
			},
			want: &Payment{Amount: 99.00},
		},
		{
			name: "Cash Purchase",
			args: args{
				o: CASH_PURCHASE,
				a: -99.00,
			},
			want: &CashPurchase{Amount: -99.00},
		},
		{
			name: "Installment Purchase",
			args: args{
				o: INSTALLMENT_PURCHASE,
				a: -99.00,
			},
			want: &InstallmentPurchase{Amount: -99.00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeOperator(tt.args.o, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
