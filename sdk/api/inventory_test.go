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

func Test_inventoryService_UpdateStockLevel(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockInventoryRepository(ctrl)
	type fields struct {
		Tracing             tracing.Tracing
		inventoryRepository api.InventoryRepository
	}
	type args struct {
		ctx context.Context
		req api.UpdateStockLevelRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.UpdateStockLevelResponse
		wantErr bool
		expect  func()
	}{
		// TODO: Add test cases.
		// OK
		{
			name: "OK",
			fields: fields{
				inventoryRepository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				req: api.UpdateStockLevelRequest{},
			},
			want:    &api.UpdateStockLevelResponse{},
			wantErr: false,
			expect: func() {
				mockRepo.EXPECT().UpdateStockLevel(gomock.Any(), gomock.Any()).Return(&api.UpdateStockLevelResponse{}, nil)
			},
		},
		// Error
		{
			name: "Error",
			fields: fields{
				inventoryRepository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				req: api.UpdateStockLevelRequest{},
			},
			want:    nil,
			wantErr: true,
			expect: func() {
				mockRepo.EXPECT().UpdateStockLevel(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expect()
			s := api.NewInventoryService(tt.fields.inventoryRepository)
			got, err := s.UpdateStockLevel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("inventoryService.UpdateStockLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inventoryService.UpdateStockLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
