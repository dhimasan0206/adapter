# \DefaultAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Accounts**](DefaultAPI.md#Accounts) | **Post** /accounts | Accounts
[**DeliveryOrders**](DefaultAPI.md#DeliveryOrders) | **Post** /delivery-orders | Delivery Orders
[**Products**](DefaultAPI.md#Products) | **Post** /products | Products
[**SalesOrders**](DefaultAPI.md#SalesOrders) | **Post** /sales-orders | Sales Orders



## Accounts

> Accounts(ctx).Signature(signature).SignatureTime(signatureTime).AccountsRequest(accountsRequest).Execute()

Accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhimasan0206/adapter/sdk/client"
)

func main() {
	signature := "signature_example" // string | Request's signature (optional)
	signatureTime := "signatureTime_example" // string | Request's signature time (optional)
	accountsRequest := *openapiclient.NewAccountsRequest() // AccountsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.Accounts(context.Background()).Signature(signature).SignatureTime(signatureTime).AccountsRequest(accountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Accounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signature** | **string** | Request&#39;s signature | 
 **signatureTime** | **string** | Request&#39;s signature time | 
 **accountsRequest** | [**AccountsRequest**](AccountsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryOrders

> DeliveryOrders(ctx).Signature(signature).DeliveryOrdersRequest(deliveryOrdersRequest).Execute()

Delivery Orders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhimasan0206/adapter/sdk/client"
)

func main() {
	signature := "signature_example" // string |  (optional)
	deliveryOrdersRequest := *openapiclient.NewDeliveryOrdersRequest() // DeliveryOrdersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeliveryOrders(context.Background()).Signature(signature).DeliveryOrdersRequest(deliveryOrdersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeliveryOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signature** | **string** |  | 
 **deliveryOrdersRequest** | [**DeliveryOrdersRequest**](DeliveryOrdersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Products

> Products(ctx).ProductsRequest(productsRequest).Execute()

Products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhimasan0206/adapter/sdk/client"
)

func main() {
	productsRequest := *openapiclient.NewProductsRequest() // ProductsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.Products(context.Background()).ProductsRequest(productsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Products``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productsRequest** | [**ProductsRequest**](ProductsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalesOrders

> SalesOrders(ctx).SalesOrdersRequest(salesOrdersRequest).Execute()

Sales Orders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhimasan0206/adapter/sdk/client"
)

func main() {
	salesOrdersRequest := *openapiclient.NewSalesOrdersRequest() // SalesOrdersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SalesOrders(context.Background()).SalesOrdersRequest(salesOrdersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SalesOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSalesOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **salesOrdersRequest** | [**SalesOrdersRequest**](SalesOrdersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

