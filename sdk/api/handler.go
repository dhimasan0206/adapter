package api

import (
	server "github.com/dhimasan0206/adapter/sdk/server/go"
)

func NewApiHandleFunctions(services Services) server.ApiHandleFunctions {
	return server.ApiHandleFunctions{
		DefaultAPI: NewDefaultAPI(),
		InventoriesAPI: NewInventoriesAPI(
			services.InventoryService,
		),
	}
}
