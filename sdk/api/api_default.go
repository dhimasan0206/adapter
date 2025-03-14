package api

import (
	"net/http"

	server "github.com/dhimasan0206/adapter/sdk/server/go"
	"github.com/gin-gonic/gin"
)

type defaultAPI struct{}

func NewDefaultAPI() server.DefaultAPI {
	return &defaultAPI{}
}

func (d *defaultAPI) Accounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "not implemented"})
}

func (d *defaultAPI) DeliveryOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "not implemented"})
}

func (d *defaultAPI) Products(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "not implemented"})
}

func (d *defaultAPI) SalesOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "not implemented"})
}
