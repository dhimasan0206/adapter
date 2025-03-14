package api

import (
	"context"

	"github.com/dhimasan0206/tracing"
)

type InventoryRepository interface{}

type InventoryService interface {
	UpdateStockLevel(ctx context.Context, req UpdateStockLevelRequest) (*UpdateStockLevelResponse, error)
}

type inventoryService struct {
	tracing.Tracing

	inventoryRepository InventoryRepository
}

func NewInventoryService(inventoryRepository InventoryRepository) InventoryService {
	tracer, meter, logger := tracing.InitTracing("InventoryService")
	return &inventoryService{
		Tracing: tracing.Tracing{
			Tracer: tracer,
			Meter:  meter,
			Logger: logger,
		},
		inventoryRepository: inventoryRepository,
	}
}

func (s *inventoryService) UpdateStockLevel(ctx context.Context, req UpdateStockLevelRequest) (*UpdateStockLevelResponse, error) {
	return nil, nil
}
