# \DefaultAPI

All URIs are relative to *https://partner.pharmacy.api.health.amazon/ingestion*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](DefaultAPI.md#CancelOrder) | **Put** /order/{partnerOrderId}/cancel | 
[**GetOrder**](DefaultAPI.md#GetOrder) | **Get** /order/{partnerOrderId} | 
[**GetPatient**](DefaultAPI.md#GetPatient) | **Get** /patient/{partnerPatientId} | 
[**GetPing**](DefaultAPI.md#GetPing) | **Get** /ping | 
[**GetPrescription**](DefaultAPI.md#GetPrescription) | **Get** /prescription/{partnerPrescriptionId} | 
[**PostPing**](DefaultAPI.md#PostPing) | **Post** /ping | 
[**PostPrescriptionTransfer**](DefaultAPI.md#PostPrescriptionTransfer) | **Post** /prescription-transfer | 
[**PutOrder**](DefaultAPI.md#PutOrder) | **Put** /order | 
[**PutPatient**](DefaultAPI.md#PutPatient) | **Put** /patient | 
[**PutPaymentInstrument**](DefaultAPI.md#PutPaymentInstrument) | **Post** /payment-instrument | 
[**PutPing**](DefaultAPI.md#PutPing) | **Put** /ping | 
[**PutPrescription**](DefaultAPI.md#PutPrescription) | **Put** /prescription | 



## CancelOrder

> CancelOrderResponseContent CancelOrder(ctx, partnerOrderId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).CancelOrderRequestContent(cancelOrderRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partnerOrderId := "test-order-1" // string | A unique Order identifier provided by the partner.
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
	cancelOrderRequestContent := *openapiclient.NewCancelOrderRequestContent("CancellationReason_example") // CancelOrderRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CancelOrder(context.Background(), partnerOrderId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).CancelOrderRequestContent(cancelOrderRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: CancelOrderResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerOrderId** | **string** | A unique Order identifier provided by the partner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 
 **cancelOrderRequestContent** | [**CancelOrderRequestContent**](CancelOrderRequestContent.md) |  | 

### Return type

[**CancelOrderResponseContent**](CancelOrderResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> GetOrderResponseContent GetOrder(ctx, partnerOrderId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partnerOrderId := "test-order-1" // string | A unique Order identifier provided by the partner.
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOrder(context.Background(), partnerOrderId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: GetOrderResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerOrderId** | **string** | A unique Order identifier provided by the partner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 

### Return type

[**GetOrderResponseContent**](GetOrderResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatient

> GetPatientResponseContent GetPatient(ctx, partnerPatientId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partnerPatientId := "2b4f8e1a-9c3d-6f7b-8a2e-d5c1e0a3b2f2" // string | A unique Patient identifier provided by the partner.
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPatient(context.Background(), partnerPatientId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPatient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatient`: GetPatientResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPatient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerPatientId** | **string** | A unique Patient identifier provided by the partner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 

### Return type

[**GetPatientResponseContent**](GetPatientResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPing

> GetPing(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetPing(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 

### Return type

 (empty response body)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrescription

> GetPrescriptionResponseContent GetPrescription(ctx, partnerPrescriptionId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	partnerPrescriptionId := "test-rx-1" // string | A unique Prescription identifier provided by the partner to create the prescription.
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrescription(context.Background(), partnerPrescriptionId).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrescription`: GetPrescriptionResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerPrescriptionId** | **string** | A unique Prescription identifier provided by the partner to create the prescription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 

### Return type

[**GetPrescriptionResponseContent**](GetPrescriptionResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPing

> PostPing(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostPing(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 

### Return type

 (empty response body)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPrescriptionTransfer

> PostPrescriptionTransferResponseContent PostPrescriptionTransfer(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PostPrescriptionTransferRequestContent(postPrescriptionTransferRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
	postPrescriptionTransferRequestContent := *openapiclient.NewPostPrescriptionTransferRequestContent("PartnerPrescriptionId_example", *openapiclient.NewTransferToPharmacy("NcpdpId_example", "Npi_example", "PharmacyName_example", "DeaNumber_example", *openapiclient.NewName("FirstName_example", "LastName_example"), *openapiclient.NewAddress("AddressLine1_example", "City_example", "State_example", "ZipCode_example", "CountryCode_example"), "PrimaryTelephone_example", "Fax_example")) // PostPrescriptionTransferRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostPrescriptionTransfer(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PostPrescriptionTransferRequestContent(postPrescriptionTransferRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostPrescriptionTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPrescriptionTransfer`: PostPrescriptionTransferResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostPrescriptionTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPrescriptionTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 
 **postPrescriptionTransferRequestContent** | [**PostPrescriptionTransferRequestContent**](PostPrescriptionTransferRequestContent.md) |  | 

### Return type

[**PostPrescriptionTransferResponseContent**](PostPrescriptionTransferResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrder

> PutOrderResponseContent PutOrder(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutOrderRequestContent(putOrderRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
	putOrderRequestContent := *openapiclient.NewPutOrderRequestContent("PartnerOrderId_example", "PartnerPrescriptionId_example", *openapiclient.NewAddress("AddressLine1_example", "City_example", "State_example", "ZipCode_example", "CountryCode_example")) // PutOrderRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutOrder(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutOrderRequestContent(putOrderRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrder`: PutOrderResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 
 **putOrderRequestContent** | [**PutOrderRequestContent**](PutOrderRequestContent.md) |  | 

### Return type

[**PutOrderResponseContent**](PutOrderResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPatient

> PutPatientResponseContent PutPatient(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutPatientRequestContent(putPatientRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
	putPatientRequestContent := *openapiclient.NewPutPatientRequestContent("PartnerPatientId_example", *openapiclient.NewPatientDetails(*openapiclient.NewName("FirstName_example", "LastName_example"), "DateOfBirth_example", openapiclient.SexAssignedAtBirth("Female"), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.ContactSystemType("Mobile"), float32(123), "PhoneNumber_example")}, false, *openapiclient.NewAllergyDetails(), *openapiclient.NewExistingMedicalConditionsDetails(), openapiclient.PregnancyStatus("PREGNANT")), *openapiclient.NewAddress("AddressLine1_example", "City_example", "State_example", "ZipCode_example", "CountryCode_example")) // PutPatientRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutPatient(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutPatientRequestContent(putPatientRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutPatient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPatient`: PutPatientResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutPatient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPatientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 
 **putPatientRequestContent** | [**PutPatientRequestContent**](PutPatientRequestContent.md) |  | 

### Return type

[**PutPatientResponseContent**](PutPatientResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPaymentInstrument

> PutPaymentInstrumentResponseContent PutPaymentInstrument(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutPaymentInstrumentRequestContent(putPaymentInstrumentRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
	putPaymentInstrumentRequestContent := *openapiclient.NewPutPaymentInstrumentRequestContent("PartnerPatientId_example", "PartnerPaymentInstrumentId_example", openapiclient.PaymentMethod("CARD"), *openapiclient.NewCardDetails("EncryptedCardNumber_example", "CardHolderName_example", float32(123), float32(123)), *openapiclient.NewAddress("AddressLine1_example", "City_example", "State_example", "ZipCode_example", "CountryCode_example")) // PutPaymentInstrumentRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutPaymentInstrument(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutPaymentInstrumentRequestContent(putPaymentInstrumentRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutPaymentInstrument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPaymentInstrument`: PutPaymentInstrumentResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutPaymentInstrument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPaymentInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 
 **putPaymentInstrumentRequestContent** | [**PutPaymentInstrumentRequestContent**](PutPaymentInstrumentRequestContent.md) |  | 

### Return type

[**PutPaymentInstrumentResponseContent**](PutPaymentInstrumentResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPing

> PutPing(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PutPing(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 

### Return type

 (empty response body)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPrescription

> PutPrescriptionResponseContent PutPrescription(ctx).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutPrescriptionRequestContent(putPrescriptionRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientIntegrationId := "clientIntegrationId_example" // string | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
	clientRequestId := "clientRequestId_example" // string | This can be provided by clients and is used to identify client's request. Clients can assign UUID and log it along with their response to track a particular API request.
	xAmzSecurityToken := "xAmzSecurityToken_example" // string | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
	putPrescriptionRequestContent := *openapiclient.NewPutPrescriptionRequestContent("PartnerPrescriptionId_example", "PartnerPatientId_example", *openapiclient.NewPrescriber("Npi_example", *openapiclient.NewName("FirstName_example", "LastName_example"), *openapiclient.NewAddress("AddressLine1_example", "City_example", "State_example", "ZipCode_example", "CountryCode_example"), "PrimaryTelephone_example"), *openapiclient.NewMedicationPrescribed("DrugDescription_example", *openapiclient.NewQuantity(float64(123), "UnitOfMeasureCode_example"), float64(123), "WrittenDate_example", float64(123), "Directions_example", "Ndc_example", "DispenseAsWritten_example"), *openapiclient.NewPharmacy("NcpdpId_example", "Npi_example", "PharmacyName_example", *openapiclient.NewName("FirstName_example", "LastName_example"), *openapiclient.NewAddress("AddressLine1_example", "City_example", "State_example", "ZipCode_example", "CountryCode_example"), "PrimaryTelephone_example")) // PutPrescriptionRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutPrescription(context.Background()).ClientIntegrationId(clientIntegrationId).ClientRequestId(clientRequestId).XAmzSecurityToken(xAmzSecurityToken).PutPrescriptionRequestContent(putPrescriptionRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutPrescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPrescription`: PutPrescriptionResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutPrescription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPrescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIntegrationId** | **string** | This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration. | 
 **clientRequestId** | **string** | This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request. | 
 **xAmzSecurityToken** | **string** | AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService | 
 **putPrescriptionRequestContent** | [**PutPrescriptionRequestContent**](PutPrescriptionRequestContent.md) |  | 

### Return type

[**PutPrescriptionResponseContent**](PutPrescriptionResponseContent.md)

### Authorization

[awsSigv4](../README.md#awsSigv4), [mutualTLS](../README.md#mutualTLS)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

