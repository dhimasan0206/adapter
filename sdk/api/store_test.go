package api_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dhimasan0206/adapter/sdk/api"
	"github.com/dhimasan0206/adapter/sdk/api/test/mocks"
	"github.com/dhimasan0206/tracing"
	"github.com/golang/mock/gomock"
)

func Test_storeService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockStoreRepository(ctrl)
	type fields struct {
		Tracing tracing.Tracing
		repo    api.StoreRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.Store
		wantErr bool
		expect  func()
	}{
		// TODO: Add test cases.
		// OK
		{
			name: "OK",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  "store-id",
			},
			want:    &api.Store{},
			wantErr: false,
			expect: func() {
				mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&api.Store{}, nil)
			},
		},
		// Error
		{
			name: "Error",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  "store-id",
			},
			want:    nil,
			wantErr: true,
			expect: func() {
				mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expect()
			s := api.NewStoreService(tt.fields.repo)
			got, err := s.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("storeService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storeService.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeService_GetByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockStoreRepository(ctrl)
	type fields struct {
		Tracing tracing.Tracing
		repo    api.StoreRepository
	}
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.Store
		wantErr bool
		expect  func()
	}{
		// TODO: Add test cases.
		// OK
		{
			name: "OK",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				code: "store-code",
			},
			want:    &api.Store{},
			wantErr: false,
			expect: func() {
				mockRepo.EXPECT().GetByCode(gomock.Any(), gomock.Any()).Return(&api.Store{}, nil)
			},
		},
		// Error
		{
			name: "Error",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				code: "store-code",
			},
			want:    nil,
			wantErr: true,
			expect: func() {
				mockRepo.EXPECT().GetByCode(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expect()
			s := api.NewStoreService(tt.fields.repo)
			got, err := s.GetByCode(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("storeService.GetByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storeService.GetByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
