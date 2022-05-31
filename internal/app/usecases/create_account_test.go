package usecases

import (
	"context"
	"testing"

	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

type mockCountAccountsByDocumentGateway struct {
	expectedData int
	expectedErr  error
}

func (fake *mockCountAccountsByDocumentGateway) CountAccountsByDocument(_ context.Context, _ string) (int, error) {
	return fake.expectedData, fake.expectedErr
}

type mockCreateAccountGateway struct {
	expectedData *int64
	expectedErr  error
}

func (fake *mockCreateAccountGateway) CreateAccount(ctx context.Context, a dto.Account) (*int64, error) {
	return fake.expectedData, fake.expectedErr
}

func Test_createAccountService_Execute(t *testing.T) {
	id := int64(1)

	type fields struct {
		gCountAcc  CountAccountsByDocumentGateway
		gCreateAcc CreateAccountGateway
	}
	type args struct {
		ctx context.Context
		acc dto.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *int64
		wantErr bool
	}{
		{
			name: "create account success",
			fields: fields{
				gCountAcc: &mockCountAccountsByDocumentGateway{
					expectedData: 0,
					expectedErr:  nil,
				},
				gCreateAcc: &mockCreateAccountGateway{
					expectedData: &id,
					expectedErr:  nil,
				},
			},
			args:    args{ctx: nil, acc: dto.Account{DocumentNumber: "12345678901"}},
			want:    &id,
			wantErr: false,
		},
		{
			name: "error document is invalid",
			fields: fields{
				gCountAcc:  nil,
				gCreateAcc: nil,
			},
			args:    args{ctx: nil, acc: dto.Account{DocumentNumber: "abcdefghijkl"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error document already exists",
			fields: fields{
				gCountAcc: &mockCountAccountsByDocumentGateway{
					expectedData: 1,
					expectedErr:  app.ErrAccountAlreadyExists,
				},
				gCreateAcc: nil,
			},
			args:    args{ctx: nil, acc: dto.Account{DocumentNumber: "12345678901"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &createAccountService{
				gCountAcc:  tt.fields.gCountAcc,
				gCreateAcc: tt.fields.gCreateAcc,
			}
			got, err := s.Execute(tt.args.ctx, tt.args.acc)
			if (err != nil) != tt.wantErr {
				t.Errorf("createAccountService.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createAccountService.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
