package api

import (
	"net/http"

	server "github.com/dhimasan0206/adapter/sdk/server/go"
	"github.com/dhimasan0206/tracing"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type inventoriesAPI struct {
	tracing.Tracing

	inventoryService InventoryService
}

func NewInventoriesAPI(inventoryService InventoryService) server.InventoriesAPI {
	tracer, meter, logger := tracing.InitTracing("InventoriesAPI")
	return &inventoriesAPI{
		Tracing: tracing.Tracing{
			Tracer: tracer,
			Meter:  meter,
			Logger: logger,
		},
		inventoryService: inventoryService,
	}
}

func (i *inventoriesAPI) UpdateStockLevel(c *gin.Context) {
	i.Logger.Info("Inventories API: update stock level: start")
	ctx, span := i.Tracer.Start(c.Request.Context(), "inventoriesAPI.updateStockLevel")
	defer span.End()

	i.Logger.Info("Inventories API: update stock level: bind request")
	var req server.UpdateStockLevelRequest
	err := c.ShouldBind(&req)
	if err != nil {
		i.Logger.Error("Inventories API: update stock level: bind request: failed", "error", err)
		tracing.RecordError(ctx, err)
		var resp server.UpdateStockLevel400Response
		resp.Error = server.UpdateStockLevel400ResponseError{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	entry := make([]UpdateStockLevelRequestEntry, 0)
	for _, r := range req.Requests {
		entry = append(entry, UpdateStockLevelRequestEntry{
			ChannelLocationId:        r.ChannelLocationId,
			ChannelLocationCode:      r.ChannelLocationCode,
			ChannelProductId:         r.ChannelProductId,
			ChannelProductVariantId:  r.ChannelProductVariantId,
			ChannelProductVariantSku: r.ChannelProductVariantSku,
			DeltaStock:               int(r.DeltaStock),
			ExactStock:               int(r.ExactStock),
			LocationCode:             r.LocationCode,
			LocationId:               r.LocationId,
			ProductId:                r.ProductId,
			ProductVariantId:         r.ProductVariantId,
		})
	}
	newReq := UpdateStockLevelRequest{
		Store: Store(req.Store),
		Entry: entry,
	}
	i.Logger.Info("Inventories API: update stock level: execute inventory service update stock level", "request", newReq)
	res, err := i.inventoryService.UpdateStockLevel(ctx, newReq)
	if err != nil {
		i.Logger.Error("Inventories API: update stock level: execute inventory service update stock level: failed", "error", err, "request", newReq)
		tracing.RecordError(ctx, err)
		var resp server.UpdateStockLevel400Response
		resp.Id = req.Id
		resp.Error = server.UpdateStockLevel400ResponseError{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	results := make([]server.UpdateStockLevel200ResponseResultsInner, 0)
	for _, r := range res.Results {
		results = append(results, server.UpdateStockLevel200ResponseResultsInner{
			ChannelLocationCode:      r.ChannelLocationCode,
			ChannelLocationId:        r.ChannelLocationId,
			ChannelProductId:         r.ChannelProductId,
			ChannelProductVariantId:  r.ChannelProductVariantId,
			ChannelProductVariantSku: r.ChannelProductVariantSku,
			LocationCode:             r.LocationCode,
			LocationId:               r.LocationId,
			Message:                  r.Message,
			ProductId:                r.ProductId,
			ProductVariantId:         r.ProductVariantId,
			Success:                  r.Success,
			// ExactStock:               r.ExactStock,
			// DeltaStock:               r.DeltaStock,
			Timestamp: float32(r.Timestamp),
		})
	}
	resp := server.UpdateStockLevel200Response{
		Id:        uuid.NewString(),
		Error:     server.UpdateStockLevel200ResponseError(res.Error),
		Message:   res.Message,
		Results:   results,
		Status:    res.Status,
		Store:     server.UpdateStockLevelRequestStore(res.Store),
		Success:   res.Success,
		Timestamp: float32(res.Timestamp),
	}
	i.Logger.Info("Inventories API: update stock level: done", "result", resp)
	c.JSON(http.StatusOK, resp)
}
