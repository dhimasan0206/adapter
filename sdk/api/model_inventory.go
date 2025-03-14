package api

type UpdateStockLevelRequest struct {
	Store Store                          `json:"store,omitempty"`
	Entry []UpdateStockLevelRequestEntry `json:"entry,omitempty"`
}

type UpdateStockLevelRequestEntry struct {
	ChannelLocationId        string `json:"channel_location_id,omitempty"`
	ChannelLocationCode      string `json:"channel_location_code,omitempty"`
	ChannelProductId         string `json:"channel_product_id,omitempty"`
	ChannelProductVariantId  string `json:"channel_product_variant_id,omitempty"`
	ChannelProductVariantSku string `json:"channel_product_variant_sku,omitempty"`
	DeltaStock               int    `json:"delta_stock,omitempty"`
	ExactStock               int    `json:"exact_stock,omitempty"`
	LocationCode             string `json:"location_code,omitempty"`
	LocationId               string `json:"location_id,omitempty"`
	ProductId                string `json:"product_id,omitempty"`
	ProductVariantId         string `json:"product_variant_id,omitempty"`
	ProductVariantSku        string `json:"product_variant_sku,omitempty"`
}

type UpdateStockLevelResponse struct {
	Store     Store                            `json:"store,omitempty"`
	Error     Error                            `json:"error,omitempty"`
	Results   []UpdateStockLevelResponseResult `json:"results,omitempty"`
	Success   bool                             `json:"success,omitempty"`
	Status    string                           `json:"status,omitempty"`
	Message   string                           `json:"message,omitempty"`
	Timestamp int64                            `json:"timestamp,omitempty"`
}

type UpdateStockLevelResponseResult struct {
	ProductId                string `json:"product_id,omitempty"`
	ProductVariantId         string `json:"product_variant_id,omitempty"`
	ProductVariantSku        string `json:"sku,omitempty"`
	ChannelProductId         string `json:"channel_product_id,omitempty"`
	ChannelProductVariantId  string `json:"channel_product_variant_id,omitempty"`
	ChannelProductVariantSku string `json:"channel_product_variant_sku,omitempty"`
	LocationId               string `json:"location_id,omitempty"`
	LocationCode             string `json:"location_code,omitempty"`
	ChannelLocationId        string `json:"channel_location_id,omitempty"`
	ChannelLocationCode      string `json:"channel_location_code,omitempty"`
	ExactStock               int    `json:"exact_stock,omitempty"`
	DeltaStock               int    `json:"delta_stock,omitempty"`
	Success                  bool   `json:"success,omitempty"`
	Message                  string `json:"message,omitempty"`
	Timestamp                int64  `json:"timestamp,omitempty"`
}
