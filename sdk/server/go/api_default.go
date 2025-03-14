/*
 * Integration Adapter
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"github.com/gin-gonic/gin"
)

type DefaultAPI interface {


    // Accounts Post /accounts
    // Accounts 
     Accounts(c *gin.Context)

    // DeliveryOrders Post /delivery-orders
    // Delivery Orders 
     DeliveryOrders(c *gin.Context)

    // Products Post /products
    // Products 
     Products(c *gin.Context)

    // SalesOrders Post /sales-orders
    // Sales Orders 
     SalesOrders(c *gin.Context)

}