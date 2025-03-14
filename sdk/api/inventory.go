package api

import (
	"context"

	"github.com/dhimasan0206/tracing"
)

//go:generate mockgen -source ./inventory.go -destination ./test/mocks/mock_inventory.go -package=mocks
type InventoryRepository interface {
	UpdateStockLevel(ctx context.Context, req UpdateStockLevelRequest) (*UpdateStockLevelResponse, error)
}

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
	s.Logger.Info("InventoryService: update stock level: start", "request", req)
	ctx, span := s.Tracer.Start(ctx, "InventoryService: update stock level")
	defer span.End()
	resp, err := s.inventoryRepository.UpdateStockLevel(ctx, req)
	if err != nil {
		s.Logger.Error("InventoryService: update stock level: failed", "error", err)
		tracing.RecordError(ctx, err)
		return nil, err
	}
	s.Logger.Info("InventoryService: update stock level: success", "response", resp)
	return resp, nil
}
