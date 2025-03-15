package api

import (
	"context"

	"github.com/dhimasan0206/tracing"
)

//go:generate mockgen -source ./store.go -destination ./test/mocks/mock_store.go -package=mocks
type StoreRepository interface {
	GetByID(ctx context.Context, id string) (*Store, error)
	GetByCode(ctx context.Context, code string) (*Store, error)
}

type StoreService interface {
	GetByID(ctx context.Context, id string) (*Store, error)
	GetByCode(ctx context.Context, code string) (*Store, error)
}

type storeService struct {
	tracing.Tracing

	StoreRepository StoreRepository
}

func NewStoreService(StoreRepository StoreRepository) StoreService {
	tracer, meter, logger := tracing.InitTracing("StoreService")
	return &storeService{
		Tracing: tracing.Tracing{
			Tracer: tracer,
			Meter:  meter,
			Logger: logger,
		},
		StoreRepository: StoreRepository,
	}
}

func (s *storeService) GetByID(ctx context.Context, id string) (*Store, error) {
	s.Logger.Info("StoreService: get store by id: start", "request", id)
	ctx, span := s.Tracer.Start(ctx, "StoreService: get store by id")
	defer span.End()
	resp, err := s.StoreRepository.GetByID(ctx, id)
	if err != nil {
		s.Logger.Error("StoreService: get store by id: failed", "error", err)
		tracing.RecordError(ctx, err)
		return nil, err
	}
	s.Logger.Info("StoreService: get store by id: success", "response", resp)
	return resp, nil
}

func (s *storeService) GetByCode(ctx context.Context, code string) (*Store, error) {
	s.Logger.Info("StoreService: get store by code: start", "request", code)
	ctx, span := s.Tracer.Start(ctx, "StoreService: get store by code")
	defer span.End()
	resp, err := s.StoreRepository.GetByCode(ctx, code)
	if err != nil {
		s.Logger.Error("StoreService: get store by code: failed", "error", err)
		tracing.RecordError(ctx, err)
		return nil, err
	}
	s.Logger.Info("StoreService: get store by code: success", "response", resp)
	return resp, nil
}
