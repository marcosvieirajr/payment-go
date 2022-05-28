package operations

import "testing"

func TestWithdraw_Validate(t *testing.T) {
	type fields struct {
		Amount float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid Amount",
			fields: fields{
				Amount: -1,
			},
			wantErr: false,
		},
		{
			name: "Valid Amount 2",
			fields: fields{
				Amount: -99.99,
			},
			wantErr: false,
		},
		{
			name: "Invalid Amount",
			fields: fields{
				Amount: 0.0,
			},
			wantErr: true,
		},
		{
			name: "Valid Amount 2",
			fields: fields{
				Amount: 1,
			},
			wantErr: true,
		},
		{
			name: "Invalid Amount 3",
			fields: fields{
				Amount: 99.99,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Withdraw{
				Amount: tt.fields.Amount,
			}
			if err := o.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Withdraw.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
