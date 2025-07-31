/*
Amazon Pharmacy APIs for ingesting Prescriptions, Patients, Orders and Payment Instruments

Amazon Pharmacy APIs provide the following functionalities for external partners:  - Patient creation, update, and view - Prescription creation, update, and view - Order creation, update, and view - Payment instrument creation and update - Prescription transfer request creation  ## Changelog for the Amazon Pharmacy APIs (from version 1.15.3):  ### 2025-07-14 v1.20.1 #### Changed: - Updated the regex pattern for `patientDetails.name`, `prescriber.prescriberName`, `pharmacist.pharmacistName` field to accept special characters like `-`, `'`, `.` for `firstName`, `middleName`, `lastName` subfields.  ### 2025-07-07 v1.20.0 #### Added: - `INVALID_DELIVERY_ADDRESS` to `OrderStatusReasonCode`.  #### Changed: - Updated `MedicationPrescribed` structure to make the `daysSupply` field optional in `PutPrescription API`, allowing prescriber partners to omit this field when calling the PutPrescription API  ### 2025-07-01 v1.19.0 #### Added: - Asynchronous notifications for orders - Added documentation in `PutOrder` callback with `EventNotification` schema which signifies notification structure  #### Notes: - Order notifications are available for below order statuses for orders created through `PutOrder` API   - ORDER_CREATED   - ORDER_PROCESSING   - ORDER_ON_HOLD   - ORDER_SHIPPED   - ORDER_CANCELLED   - ORDER_FAILED  ### 2025-06-19 v1.18.2 #### Changed: - Updated `PatientDetails` structure to make the `sexAssignedAtBirth` field optional in `PutPatient API`, allowing partners to omit this field when calling the PutPatient API; however, this field will continue to be mandatory for hub partners who manage orders on behalf of customers  ### 2025-06-10 v1.18.1 #### Changed: - Removed mandatory requirement for the deprecated `unitOfMeasureCode` field in `Quantity`. - Updated documentation for quantity and refill fields. - Added comprehensive examples covering various prescription transfer scenarios. - Updated documentation to reflect optional status of `refillsTransferred` field.  ### 2025-06-03 v1.18.0 #### Added: - Added optional field `Observation` in PutPrescriptionRequest to support vital signs data. - Added new structures:   - `Observation` with measurement list and notes   - `Measurement` with vital sign, value, unit, and date fields   - `VitalSign` enum supporting HEIGHT and WEIGHT measurements - Updated `PutPrescription` API examples to demonstrate usage of Observation structure with height and weight measurements.   #### Changed:  - Updated `PatientDetails` structure to make the `smsConsent` field optional, allowing partners to omit this field when calling the PutPatient API. When not provided, `smsConsent` defaults to `false`.  ### 2025-03-24 v1.17.0 #### Added: - Asynchronous notifications for prescriptions. - Added documentation in `PutPrescription` callback with `EventNotification` schema which signifies notification structure  #### Changed: - `Insurance` now supports `cardholderId`. - `Insurance` and `OrderInsurance` field `encryptedCardholderId` is now deprecated. Please use `cardholderId` instead for `PutPatient` and `PutOrder` calls.  ### 2025-03-26 v1.16.4 #### Added: - Added prescriber-direct prescription example to the `PutPrescription` API documentation.  #### Changed: - Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.  ### 2025-03-18 v1.16.3 #### Changed: - Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema. - For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation. - For prescriber-originated prescriptions: This structure is not required and should be omitted. - Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.   ### 2025-01-06 v1.16.2 #### Added: - The changes below are applied to `PutPrescription` API. - Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release. - Added a new required field `ndc` in `PreviousDispensedMedication`. - Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release. - The changes below are applied to `PutOrder` API. - Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.  #### Changed: - The changes below are applied to `PutPrescription` API. - Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation. - Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`. - Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`  ### 2025-01-06 v1.16.1 #### Added: - The changes below are applied to `PutPrescription` API. - Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`. - Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release. - Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`. - Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`. - Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`. - Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`. #### Changed: - The changes below are applied to `PutPrescription` API. - Deprecated `rxcui`. use `DrugDbCode` structure instead. - Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead. - Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead. - Deprecated `diagnosis`. Use `medicalDiagnosis` instead. - Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`. - Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.  ### 2024-12-13 v1.16.0 #### Changed: - Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances. - Updated `InsuranceList` max size to 27 from 50.  ### 2024-12-09 v1.15.10 #### Changed: - Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.  ### 2024-11-04 v1.15.9 #### Added: - New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request. #### Changed: - Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.  ### 2024-10-15 v1.15.8 #### Added: - Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model. - Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.  ### 2024-09-23 v1.15.7 #### Added: - Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues. - Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues. - Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.  ### 2024-09-03 v1.15.6 #### Added: - Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.   If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor. - The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field. - The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure. - A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException. - `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API. - Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.  ### 2024-08-06 v1.15.5 #### Added: - Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities. - Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions. - Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead. - Added documentation for enum fields for all APIs. - Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.  ### 2024-07-31 v1.15.4 #### Improved: - Updated the `directions` field to allow all strings for greater flexibility. - Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.  ### 2024-07-14 v1.15.3 #### Added: - All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs.

API version: 1.20.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amazonpharmacy

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiCancelOrderRequest struct {
	ctx                       context.Context
	ApiService                *DefaultAPIService
	partnerOrderId            string
	clientIntegrationId       *string
	clientRequestId           *string
	xAmzSecurityToken         *string
	cancelOrderRequestContent *CancelOrderRequestContent
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiCancelOrderRequest) ClientIntegrationId(clientIntegrationId string) ApiCancelOrderRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiCancelOrderRequest) ClientRequestId(clientRequestId string) ApiCancelOrderRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiCancelOrderRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiCancelOrderRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiCancelOrderRequest) CancelOrderRequestContent(cancelOrderRequestContent CancelOrderRequestContent) ApiCancelOrderRequest {
	r.cancelOrderRequestContent = &cancelOrderRequestContent
	return r
}

func (r ApiCancelOrderRequest) Execute() (*CancelOrderResponseContent, *http.Response, error) {
	return r.ApiService.CancelOrderExecute(r)
}

/*
CancelOrder Method for CancelOrder

This API is responsible to cancel a patient's order.
Invoking this API will make an attempt to stop the dispense of medication to a patient.
Depending on whether or not the order has shipped, cancellation is not always guaranteed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param partnerOrderId A unique Order identifier provided by the partner.
	@return ApiCancelOrderRequest
*/
func (a *DefaultAPIService) CancelOrder(ctx context.Context, partnerOrderId string) ApiCancelOrderRequest {
	return ApiCancelOrderRequest{
		ApiService:     a,
		ctx:            ctx,
		partnerOrderId: partnerOrderId,
	}
}

// Execute executes the request
//
//	@return CancelOrderResponseContent
func (a *DefaultAPIService) CancelOrderExecute(r ApiCancelOrderRequest) (*CancelOrderResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CancelOrderResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.CancelOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/{partnerOrderId}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerOrderId"+"}", url.PathEscape(parameterValueToString(r.partnerOrderId, "partnerOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}
	if r.cancelOrderRequestContent == nil {
		return localVarReturnValue, nil, reportError("cancelOrderRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	// body params
	localVarPostBody = r.cancelOrderRequestContent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidInputExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrderRequest struct {
	ctx                 context.Context
	ApiService          *DefaultAPIService
	partnerOrderId      string
	clientIntegrationId *string
	clientRequestId     *string
	xAmzSecurityToken   *string
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiGetOrderRequest) ClientIntegrationId(clientIntegrationId string) ApiGetOrderRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiGetOrderRequest) ClientRequestId(clientRequestId string) ApiGetOrderRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiGetOrderRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiGetOrderRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiGetOrderRequest) Execute() (*GetOrderResponseContent, *http.Response, error) {
	return r.ApiService.GetOrderExecute(r)
}

/*
GetOrder Method for GetOrder

This API is responsible to get status information regarding the order.
The API will return status information about the latest version of the submitted order in case it is submitted multiple times.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param partnerOrderId A unique Order identifier provided by the partner.
	@return ApiGetOrderRequest
*/
func (a *DefaultAPIService) GetOrder(ctx context.Context, partnerOrderId string) ApiGetOrderRequest {
	return ApiGetOrderRequest{
		ApiService:     a,
		ctx:            ctx,
		partnerOrderId: partnerOrderId,
	}
}

// Execute executes the request
//
//	@return GetOrderResponseContent
func (a *DefaultAPIService) GetOrderExecute(r ApiGetOrderRequest) (*GetOrderResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.GetOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order/{partnerOrderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerOrderId"+"}", url.PathEscape(parameterValueToString(r.partnerOrderId, "partnerOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidInputExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPatientRequest struct {
	ctx                 context.Context
	ApiService          *DefaultAPIService
	partnerPatientId    string
	clientIntegrationId *string
	clientRequestId     *string
	xAmzSecurityToken   *string
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiGetPatientRequest) ClientIntegrationId(clientIntegrationId string) ApiGetPatientRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiGetPatientRequest) ClientRequestId(clientRequestId string) ApiGetPatientRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiGetPatientRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiGetPatientRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiGetPatientRequest) Execute() (*GetPatientResponseContent, *http.Response, error) {
	return r.ApiService.GetPatientExecute(r)
}

/*
GetPatient Method for GetPatient

This API is responsible for getting patient details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param partnerPatientId A unique Patient identifier provided by the partner.
	@return ApiGetPatientRequest
*/
func (a *DefaultAPIService) GetPatient(ctx context.Context, partnerPatientId string) ApiGetPatientRequest {
	return ApiGetPatientRequest{
		ApiService:       a,
		ctx:              ctx,
		partnerPatientId: partnerPatientId,
	}
}

// Execute executes the request
//
//	@return GetPatientResponseContent
func (a *DefaultAPIService) GetPatientExecute(r ApiGetPatientRequest) (*GetPatientResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPatientResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.GetPatient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/patient/{partnerPatientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerPatientId"+"}", url.PathEscape(parameterValueToString(r.partnerPatientId, "partnerPatientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidInputExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPingRequest struct {
	ctx                 context.Context
	ApiService          *DefaultAPIService
	clientIntegrationId *string
	clientRequestId     *string
	xAmzSecurityToken   *string
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiGetPingRequest) ClientIntegrationId(clientIntegrationId string) ApiGetPingRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiGetPingRequest) ClientRequestId(clientRequestId string) ApiGetPingRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiGetPingRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiGetPingRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiGetPingRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetPingExecute(r)
}

/*
GetPing Method for GetPing

This API provides an endpoint to ping. It does nothing except return 200.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPingRequest
*/
func (a *DefaultAPIService) GetPing(ctx context.Context) ApiGetPingRequest {
	return ApiGetPingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) GetPingExecute(r ApiGetPingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.GetPing")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ping"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return nil, reportError("xAmzSecurityToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPrescriptionRequest struct {
	ctx                   context.Context
	ApiService            *DefaultAPIService
	partnerPrescriptionId string
	clientIntegrationId   *string
	clientRequestId       *string
	xAmzSecurityToken     *string
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiGetPrescriptionRequest) ClientIntegrationId(clientIntegrationId string) ApiGetPrescriptionRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiGetPrescriptionRequest) ClientRequestId(clientRequestId string) ApiGetPrescriptionRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiGetPrescriptionRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiGetPrescriptionRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiGetPrescriptionRequest) Execute() (*GetPrescriptionResponseContent, *http.Response, error) {
	return r.ApiService.GetPrescriptionExecute(r)
}

/*
GetPrescription Method for GetPrescription

This API is responsible for getting status information regarding the creation of the prescription.
The API will return status information about the latest version of the submitted prescription
in case it is submitted multiple times.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param partnerPrescriptionId A unique Prescription identifier provided by the partner to create the prescription.
	@return ApiGetPrescriptionRequest
*/
func (a *DefaultAPIService) GetPrescription(ctx context.Context, partnerPrescriptionId string) ApiGetPrescriptionRequest {
	return ApiGetPrescriptionRequest{
		ApiService:            a,
		ctx:                   ctx,
		partnerPrescriptionId: partnerPrescriptionId,
	}
}

// Execute executes the request
//
//	@return GetPrescriptionResponseContent
func (a *DefaultAPIService) GetPrescriptionExecute(r ApiGetPrescriptionRequest) (*GetPrescriptionResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPrescriptionResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.GetPrescription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prescription/{partnerPrescriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerPrescriptionId"+"}", url.PathEscape(parameterValueToString(r.partnerPrescriptionId, "partnerPrescriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidInputExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostPingRequest struct {
	ctx                 context.Context
	ApiService          *DefaultAPIService
	clientIntegrationId *string
	clientRequestId     *string
	xAmzSecurityToken   *string
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiPostPingRequest) ClientIntegrationId(clientIntegrationId string) ApiPostPingRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiPostPingRequest) ClientRequestId(clientRequestId string) ApiPostPingRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiPostPingRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPostPingRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiPostPingRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostPingExecute(r)
}

/*
PostPing Method for PostPing

This API provides an endpoint to ping. It does nothing except return 200.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostPingRequest
*/
func (a *DefaultAPIService) PostPing(ctx context.Context) ApiPostPingRequest {
	return ApiPostPingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) PostPingExecute(r ApiPostPingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.PostPing")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ping"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return nil, reportError("xAmzSecurityToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostPrescriptionTransferRequest struct {
	ctx                                    context.Context
	ApiService                             *DefaultAPIService
	clientIntegrationId                    *string
	clientRequestId                        *string
	xAmzSecurityToken                      *string
	postPrescriptionTransferRequestContent *PostPrescriptionTransferRequestContent
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiPostPrescriptionTransferRequest) ClientIntegrationId(clientIntegrationId string) ApiPostPrescriptionTransferRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiPostPrescriptionTransferRequest) ClientRequestId(clientRequestId string) ApiPostPrescriptionTransferRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiPostPrescriptionTransferRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPostPrescriptionTransferRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiPostPrescriptionTransferRequest) PostPrescriptionTransferRequestContent(postPrescriptionTransferRequestContent PostPrescriptionTransferRequestContent) ApiPostPrescriptionTransferRequest {
	r.postPrescriptionTransferRequestContent = &postPrescriptionTransferRequestContent
	return r
}

func (r ApiPostPrescriptionTransferRequest) Execute() (*PostPrescriptionTransferResponseContent, *http.Response, error) {
	return r.ApiService.PostPrescriptionTransferExecute(r)
}

/*
PostPrescriptionTransfer Method for PostPrescriptionTransfer

This API is responsible for requesting prescription transfer from Amazon Pharmacy.
As of now, this request triggers an asynchronous operation.
So, API clients are suggested to invoke GET /prescription API to monitor the status of the transfer request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostPrescriptionTransferRequest
*/
func (a *DefaultAPIService) PostPrescriptionTransfer(ctx context.Context) ApiPostPrescriptionTransferRequest {
	return ApiPostPrescriptionTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PostPrescriptionTransferResponseContent
func (a *DefaultAPIService) PostPrescriptionTransferExecute(r ApiPostPrescriptionTransferRequest) (*PostPrescriptionTransferResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PostPrescriptionTransferResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.PostPrescriptionTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prescription-transfer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}
	if r.postPrescriptionTransferRequestContent == nil {
		return localVarReturnValue, nil, reportError("postPrescriptionTransferRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	// body params
	localVarPostBody = r.postPrescriptionTransferRequestContent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v PreconditionFailedExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutOrderRequest struct {
	ctx                    context.Context
	ApiService             *DefaultAPIService
	clientIntegrationId    *string
	clientRequestId        *string
	xAmzSecurityToken      *string
	putOrderRequestContent *PutOrderRequestContent
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiPutOrderRequest) ClientIntegrationId(clientIntegrationId string) ApiPutOrderRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiPutOrderRequest) ClientRequestId(clientRequestId string) ApiPutOrderRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiPutOrderRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutOrderRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiPutOrderRequest) PutOrderRequestContent(putOrderRequestContent PutOrderRequestContent) ApiPutOrderRequest {
	r.putOrderRequestContent = &putOrderRequestContent
	return r
}

func (r ApiPutOrderRequest) Execute() (*PutOrderResponseContent, *http.Response, error) {
	return r.ApiService.PutOrderExecute(r)
}

/*
PutOrder Method for PutOrder

This API is responsible for placing an order on behalf of the patient (Customer) for the referred prescription.
This will generate a dispense for the medication and a physical shipment to the request contained in the request.
The order can be updated under select status conditions (ORDER_CANCELLED/ORDER_FAILED).

Please look at [notification onboarding document](https://github.com/AmazonHealthServices-External/AmazonPharmacyExternalAPISpecs/tree/main/onboarding/async-notification)
for onboarding to order event notifications.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPutOrderRequest
*/
func (a *DefaultAPIService) PutOrder(ctx context.Context) ApiPutOrderRequest {
	return ApiPutOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PutOrderResponseContent
func (a *DefaultAPIService) PutOrderExecute(r ApiPutOrderRequest) (*PutOrderResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PutOrderResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.PutOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}
	if r.putOrderRequestContent == nil {
		return localVarReturnValue, nil, reportError("putOrderRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	// body params
	localVarPostBody = r.putOrderRequestContent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v PutOrder400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v InsuranceComplianceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutPatientRequest struct {
	ctx                      context.Context
	ApiService               *DefaultAPIService
	clientIntegrationId      *string
	clientRequestId          *string
	xAmzSecurityToken        *string
	putPatientRequestContent *PutPatientRequestContent
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiPutPatientRequest) ClientIntegrationId(clientIntegrationId string) ApiPutPatientRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiPutPatientRequest) ClientRequestId(clientRequestId string) ApiPutPatientRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiPutPatientRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutPatientRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiPutPatientRequest) PutPatientRequestContent(putPatientRequestContent PutPatientRequestContent) ApiPutPatientRequest {
	r.putPatientRequestContent = &putPatientRequestContent
	return r
}

func (r ApiPutPatientRequest) Execute() (*PutPatientResponseContent, *http.Response, error) {
	return r.ApiService.PutPatientExecute(r)
}

/*
PutPatient Method for PutPatient

This API is responsible for setting up a patient profile within the Amazon Pharmacy system.
This profile is unique to each partner, however multiple profiles can match to the same Amazon customer internally.
The API request contains information regarding the patient such as insurance and demographic details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPutPatientRequest
*/
func (a *DefaultAPIService) PutPatient(ctx context.Context) ApiPutPatientRequest {
	return ApiPutPatientRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PutPatientResponseContent
func (a *DefaultAPIService) PutPatientExecute(r ApiPutPatientRequest) (*PutPatientResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PutPatientResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.PutPatient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/patient"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}
	if r.putPatientRequestContent == nil {
		return localVarReturnValue, nil, reportError("putPatientRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	// body params
	localVarPostBody = r.putPatientRequestContent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v PreconditionFailedExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutPaymentInstrumentRequest struct {
	ctx                                context.Context
	ApiService                         *DefaultAPIService
	clientIntegrationId                *string
	clientRequestId                    *string
	xAmzSecurityToken                  *string
	putPaymentInstrumentRequestContent *PutPaymentInstrumentRequestContent
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiPutPaymentInstrumentRequest) ClientIntegrationId(clientIntegrationId string) ApiPutPaymentInstrumentRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiPutPaymentInstrumentRequest) ClientRequestId(clientRequestId string) ApiPutPaymentInstrumentRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiPutPaymentInstrumentRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutPaymentInstrumentRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiPutPaymentInstrumentRequest) PutPaymentInstrumentRequestContent(putPaymentInstrumentRequestContent PutPaymentInstrumentRequestContent) ApiPutPaymentInstrumentRequest {
	r.putPaymentInstrumentRequestContent = &putPaymentInstrumentRequestContent
	return r
}

func (r ApiPutPaymentInstrumentRequest) Execute() (*PutPaymentInstrumentResponseContent, *http.Response, error) {
	return r.ApiService.PutPaymentInstrumentExecute(r)
}

/*
PutPaymentInstrument Method for PutPaymentInstrument

This API is responsible for ingesting a single payment instrument that belongs to a patient.
This payment method will be used for processing orders in cases where a co-pay amount is due from the patient.
The payment-instrument ID must be included as part of the Put Order request when the quoted copay amount is greater than zero.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPutPaymentInstrumentRequest
*/
func (a *DefaultAPIService) PutPaymentInstrument(ctx context.Context) ApiPutPaymentInstrumentRequest {
	return ApiPutPaymentInstrumentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PutPaymentInstrumentResponseContent
func (a *DefaultAPIService) PutPaymentInstrumentExecute(r ApiPutPaymentInstrumentRequest) (*PutPaymentInstrumentResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PutPaymentInstrumentResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.PutPaymentInstrument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-instrument"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}
	if r.putPaymentInstrumentRequestContent == nil {
		return localVarReturnValue, nil, reportError("putPaymentInstrumentRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	// body params
	localVarPostBody = r.putPaymentInstrumentRequestContent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidInputExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutPingRequest struct {
	ctx                 context.Context
	ApiService          *DefaultAPIService
	clientIntegrationId *string
	clientRequestId     *string
	xAmzSecurityToken   *string
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiPutPingRequest) ClientIntegrationId(clientIntegrationId string) ApiPutPingRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiPutPingRequest) ClientRequestId(clientRequestId string) ApiPutPingRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiPutPingRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutPingRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiPutPingRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutPingExecute(r)
}

/*
PutPing Method for PutPing

This API provides an endpoint to ping. It does nothing except return 200.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPutPingRequest
*/
func (a *DefaultAPIService) PutPing(ctx context.Context) ApiPutPingRequest {
	return ApiPutPingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) PutPingExecute(r ApiPutPingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.PutPing")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ping"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return nil, reportError("xAmzSecurityToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPutPrescriptionRequest struct {
	ctx                           context.Context
	ApiService                    *DefaultAPIService
	clientIntegrationId           *string
	clientRequestId               *string
	xAmzSecurityToken             *string
	putPrescriptionRequestContent *PutPrescriptionRequestContent
}

// This will be provided to your Amazon Pharmacy Account Manager. The identifier will be used by Amazon Pharmacy to determine configurations for a specific program or integration.
func (r ApiPutPrescriptionRequest) ClientIntegrationId(clientIntegrationId string) ApiPutPrescriptionRequest {
	r.clientIntegrationId = &clientIntegrationId
	return r
}

// This can be provided by clients and is used to identify client&#39;s request. Clients can assign UUID and log it along with their response to track a particular API request.
func (r ApiPutPrescriptionRequest) ClientRequestId(clientRequestId string) ApiPutPrescriptionRequest {
	r.clientRequestId = &clientRequestId
	return r
}

// AWS temporary session token will be obtained by connecting to AmazonSecurityTokenService
func (r ApiPutPrescriptionRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutPrescriptionRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

func (r ApiPutPrescriptionRequest) PutPrescriptionRequestContent(putPrescriptionRequestContent PutPrescriptionRequestContent) ApiPutPrescriptionRequest {
	r.putPrescriptionRequestContent = &putPrescriptionRequestContent
	return r
}

func (r ApiPutPrescriptionRequest) Execute() (*PutPrescriptionResponseContent, *http.Response, error) {
	return r.ApiService.PutPrescriptionExecute(r)
}

/*
PutPrescription Method for PutPrescription

Prescriptions are required to dispense medications to patients.
This API is responsible for sending the patient's prescription to Amazon Pharmacy.
If the patient does not exist, this request will be rejected.
If the patient exists but in a failed state, the prescription will be ingested but set to failed as well, requiring a new version to be sent.
For a successful request, the prescription will flow through into the patient's (Amazon customer's) dashboard.
The prescription can be updated under select status conditions (PRESCRIPTION_FAILED).

Please look at [notification onboarding document](https://github.com/AmazonHealthServices-External/AmazonPharmacyExternalAPISpecs/tree/main/onboarding/async-notification)
for onboarding to prescription event notifications.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPutPrescriptionRequest
*/
func (a *DefaultAPIService) PutPrescription(ctx context.Context) ApiPutPrescriptionRequest {
	return ApiPutPrescriptionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PutPrescriptionResponseContent
func (a *DefaultAPIService) PutPrescriptionExecute(r ApiPutPrescriptionRequest) (*PutPrescriptionResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PutPrescriptionResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.PutPrescription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prescription"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientIntegrationId == nil {
		return localVarReturnValue, nil, reportError("clientIntegrationId is required and must be specified")
	}
	if r.clientRequestId == nil {
		return localVarReturnValue, nil, reportError("clientRequestId is required and must be specified")
	}
	if r.xAmzSecurityToken == nil {
		return localVarReturnValue, nil, reportError("xAmzSecurityToken is required and must be specified")
	}
	if r.putPrescriptionRequestContent == nil {
		return localVarReturnValue, nil, reportError("putPrescriptionRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-integration-id", r.clientIntegrationId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "client-request-id", r.clientRequestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "simple", "")
	// body params
	localVarPostBody = r.putPrescriptionRequestContent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v PutOrder400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResourceNotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServerExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PutOrder502Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
