package domain

import (
	"testing"
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
