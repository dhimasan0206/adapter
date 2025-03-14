# \InventoriesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateStockLevel**](InventoriesAPI.md#UpdateStockLevel) | **Post** /inventories/update | Update Stock Level



## UpdateStockLevel

> UpdateStockLevel200Response UpdateStockLevel(ctx).Signature(signature).SignatureTime(signatureTime).UpdateStockLevelRequest(updateStockLevelRequest).Execute()

Update Stock Level



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
	signature := "signature" // string | Requests Signature (optional)
	signatureTime := "signature-time" // string | Requests Signature Time (optional)
	updateStockLevelRequest := *openapiclient.NewUpdateStockLevelRequest() // UpdateStockLevelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.UpdateStockLevel(context.Background()).Signature(signature).SignatureTime(signatureTime).UpdateStockLevelRequest(updateStockLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.UpdateStockLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStockLevel`: UpdateStockLevel200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.UpdateStockLevel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStockLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signature** | **string** | Requests Signature | 
 **signatureTime** | **string** | Requests Signature Time | 
 **updateStockLevelRequest** | [**UpdateStockLevelRequest**](UpdateStockLevelRequest.md) |  | 

### Return type

[**UpdateStockLevel200Response**](UpdateStockLevel200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

