package usecases

import (
	"context"
	"reflect"
	"testing"

	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/domain"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

type mockGetAccountGateway struct {
	expectedData *domain.Account
	expectedErr  error
}

func (fake *mockGetAccountGateway) GetAccount(_ context.Context, _ int64) (*domain.Account, error) {
	return fake.expectedData, fake.expectedErr
}

func Test_getAccountService_Execute(t *testing.T) {
	type fields struct {
		gtw GetAccountGateway
	}
	type args struct {
		ctx       context.Context
		accountID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Account
		wantErr bool
	}{
		{
			name: "get account success",
			fields: fields{
				gtw: &mockGetAccountGateway{
					expectedData: &domain.Account{ID: 1, DocumentNumber: "12345678901"},
					expectedErr:  nil,
				},
			},
			args:    args{ctx: nil, accountID: 1},
			want:    &dto.Account{ID: 1, DocumentNumber: "12345678901"},
			wantErr: false,
		},
		{
			name: "get account error",
			fields: fields{
				gtw: &mockGetAccountGateway{
					expectedData: nil,
					expectedErr:  app.ErrAccountNotFound,
				},
			},
			args:    args{ctx: nil, accountID: 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &getAccountService{
				gtw: tt.fields.gtw,
			}
			got, err := s.Execute(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAccountService.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAccountService.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
