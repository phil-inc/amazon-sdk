# Go API client for openapi

Amazon Pharmacy APIs provide the following functionalities for external partners:

- Patient creation, update, and view
- Prescription creation, update, and view
- Order creation, update, and view
- Payment instrument creation and update
- Prescription transfer request creation

## Changelog for the Amazon Pharmacy APIs (from version 1.15.3):
### 2025-03-26 v1.16.4
#### Added:
- Added prescriber-direct prescription example to the `PutPrescription` API documentation.

#### Changed:
- Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.

### 2025-03-18 v1.16.3
#### Changed:
- Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema.
- For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation.
- For prescriber-originated prescriptions: This structure is not required and should be omitted.
- Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.

### 2025-01-06 v1.16.2
#### Added:
- The changes below are applied to `PutPrescription` API.
- Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release.
- Added a new required field `ndc` in `PreviousDispensedMedication`.
- Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release.
- The changes below are applied to `PutOrder` API.
- Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.

#### Changed:
- The changes below are applied to `PutPrescription` API.
- Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation.
- Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`.
- Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`

### 2025-01-06 v1.16.1
#### Added:
- The changes below are applied to `PutPrescription` API.
- Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`.
- Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release.
- Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`.
- Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`.
- Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`.
- Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`.
#### Changed:
- The changes below are applied to `PutPrescription` API.
- Deprecated `rxcui`. use `DrugDbCode` structure instead.
- Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead.
- Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead.
- Deprecated `diagnosis`. Use `medicalDiagnosis` instead.
- Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`.
- Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.

### 2024-12-13 v1.16.0
#### Changed:
- Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances.
- Updated `InsuranceList` max size to 27 from 50.

### 2024-12-09 v1.15.10
#### Changed:
- Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.

### 2024-11-04 v1.15.9
#### Added:
- New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request.
#### Changed:
- Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.

### 2024-10-15 v1.15.8
#### Added:
- Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model.
- Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.

### 2024-09-23 v1.15.7
#### Added:
- Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues.
- Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues.
- Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.

### 2024-09-03 v1.15.6
#### Added:
- Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.
  If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor.
- The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field.
- The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure.
- A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException.
- `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API.
- Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.

### 2024-08-06 v1.15.5
#### Added:
- Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities.
- Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions.
- Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead.
- Added documentation for enum fields for all APIs.
- Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.

### 2024-07-31 v1.15.4
#### Improved:
- Updated the `directions` field to allow all strings for greater flexibility.
- Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.

### 2024-07-14 v1.15.3
#### Added:
- All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.16.4
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://partner.pharmacy.api.health.amazon/ingestion*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**CancelOrder**](docs/DefaultAPI.md#cancelorder) | **Put** /order/{partnerOrderId}/cancel | 
*DefaultAPI* | [**GetOrder**](docs/DefaultAPI.md#getorder) | **Get** /order/{partnerOrderId} | 
*DefaultAPI* | [**GetPatient**](docs/DefaultAPI.md#getpatient) | **Get** /patient/{partnerPatientId} | 
*DefaultAPI* | [**GetPing**](docs/DefaultAPI.md#getping) | **Get** /ping | 
*DefaultAPI* | [**GetPrescription**](docs/DefaultAPI.md#getprescription) | **Get** /prescription/{partnerPrescriptionId} | 
*DefaultAPI* | [**PostPing**](docs/DefaultAPI.md#postping) | **Post** /ping | 
*DefaultAPI* | [**PostPrescriptionTransfer**](docs/DefaultAPI.md#postprescriptiontransfer) | **Post** /prescription-transfer | 
*DefaultAPI* | [**PutOrder**](docs/DefaultAPI.md#putorder) | **Put** /order | 
*DefaultAPI* | [**PutPatient**](docs/DefaultAPI.md#putpatient) | **Put** /patient | 
*DefaultAPI* | [**PutPaymentInstrument**](docs/DefaultAPI.md#putpaymentinstrument) | **Post** /payment-instrument | 
*DefaultAPI* | [**PutPing**](docs/DefaultAPI.md#putping) | **Put** /ping | 
*DefaultAPI* | [**PutPrescription**](docs/DefaultAPI.md#putprescription) | **Put** /prescription | 


## Documentation For Models

 - [AdditionalOrderDetails](docs/AdditionalOrderDetails.md)
 - [Address](docs/Address.md)
 - [AllergiesDescriptor](docs/AllergiesDescriptor.md)
 - [Allergy](docs/Allergy.md)
 - [AllergyCodeType](docs/AllergyCodeType.md)
 - [AllergyDetails](docs/AllergyDetails.md)
 - [CancelOrderRequestContent](docs/CancelOrderRequestContent.md)
 - [CancelOrderResponseContent](docs/CancelOrderResponseContent.md)
 - [CardDetails](docs/CardDetails.md)
 - [ClaimRejectionDetails](docs/ClaimRejectionDetails.md)
 - [ClaimRejectionReason](docs/ClaimRejectionReason.md)
 - [ClaimRejectionType](docs/ClaimRejectionType.md)
 - [ClinicalInformationQualifier](docs/ClinicalInformationQualifier.md)
 - [ConditionCodeType](docs/ConditionCodeType.md)
 - [ContactInfo](docs/ContactInfo.md)
 - [ContactSystemType](docs/ContactSystemType.md)
 - [DaysSupplyDetails](docs/DaysSupplyDetails.md)
 - [DeaSchedule](docs/DeaSchedule.md)
 - [Diagnosis](docs/Diagnosis.md)
 - [DiagnosisCode](docs/DiagnosisCode.md)
 - [DrugDbCode](docs/DrugDbCode.md)
 - [DrugDbCodeQualifier](docs/DrugDbCodeQualifier.md)
 - [ErrorField](docs/ErrorField.md)
 - [ExistingMedicalConditionsDescriptor](docs/ExistingMedicalConditionsDescriptor.md)
 - [ExistingMedicalConditionsDetails](docs/ExistingMedicalConditionsDetails.md)
 - [ExistingMedication](docs/ExistingMedication.md)
 - [ExistingMedicationDetails](docs/ExistingMedicationDetails.md)
 - [ExistingMedicationsDescriptor](docs/ExistingMedicationsDescriptor.md)
 - [FailedRule](docs/FailedRule.md)
 - [FieldErrors](docs/FieldErrors.md)
 - [GetOrderResponseContent](docs/GetOrderResponseContent.md)
 - [GetPatientResponseContent](docs/GetPatientResponseContent.md)
 - [GetPrescriptionResponseContent](docs/GetPrescriptionResponseContent.md)
 - [Insurance](docs/Insurance.md)
 - [InsuranceComplianceExceptionResponseContent](docs/InsuranceComplianceExceptionResponseContent.md)
 - [InsuranceComplianceRuleType](docs/InsuranceComplianceRuleType.md)
 - [InternalServerExceptionResponseContent](docs/InternalServerExceptionResponseContent.md)
 - [InvalidInputErrorCode](docs/InvalidInputErrorCode.md)
 - [InvalidInputExceptionResponseContent](docs/InvalidInputExceptionResponseContent.md)
 - [MedicalCondition](docs/MedicalCondition.md)
 - [MedicalDiagnosis](docs/MedicalDiagnosis.md)
 - [MedicationPrescribed](docs/MedicationPrescribed.md)
 - [Name](docs/Name.md)
 - [Order](docs/Order.md)
 - [OrderInsurance](docs/OrderInsurance.md)
 - [OrderInsuranceDetails](docs/OrderInsuranceDetails.md)
 - [OrderStatus](docs/OrderStatus.md)
 - [OrderStatusReasonCode](docs/OrderStatusReasonCode.md)
 - [Patient](docs/Patient.md)
 - [PatientDetails](docs/PatientDetails.md)
 - [PatientRelationship](docs/PatientRelationship.md)
 - [PatientStatus](docs/PatientStatus.md)
 - [PatientStatusReasonCode](docs/PatientStatusReasonCode.md)
 - [PayerType](docs/PayerType.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [Pharmacy](docs/Pharmacy.md)
 - [PossibleDemographicFieldsInError](docs/PossibleDemographicFieldsInError.md)
 - [PossibleInsuranceFieldsInError](docs/PossibleInsuranceFieldsInError.md)
 - [PostPrescriptionTransferRequestContent](docs/PostPrescriptionTransferRequestContent.md)
 - [PostPrescriptionTransferResponseContent](docs/PostPrescriptionTransferResponseContent.md)
 - [PreconditionErrorCode](docs/PreconditionErrorCode.md)
 - [PreconditionFailedExceptionResponseContent](docs/PreconditionFailedExceptionResponseContent.md)
 - [PregnancyStatus](docs/PregnancyStatus.md)
 - [Prescriber](docs/Prescriber.md)
 - [Prescription](docs/Prescription.md)
 - [PrescriptionFill](docs/PrescriptionFill.md)
 - [PrescriptionFillStatus](docs/PrescriptionFillStatus.md)
 - [PrescriptionFillTransferInType](docs/PrescriptionFillTransferInType.md)
 - [PrescriptionStatus](docs/PrescriptionStatus.md)
 - [PrescriptionTransferInDetails](docs/PrescriptionTransferInDetails.md)
 - [PrescriptionTransferred](docs/PrescriptionTransferred.md)
 - [PreviousDispensedMedication](docs/PreviousDispensedMedication.md)
 - [PutOrder400Response](docs/PutOrder400Response.md)
 - [PutOrder502Response](docs/PutOrder502Response.md)
 - [PutOrderRequestContent](docs/PutOrderRequestContent.md)
 - [PutOrderResponseContent](docs/PutOrderResponseContent.md)
 - [PutPatientRequestContent](docs/PutPatientRequestContent.md)
 - [PutPatientResponseContent](docs/PutPatientResponseContent.md)
 - [PutPaymentInstrumentRequestContent](docs/PutPaymentInstrumentRequestContent.md)
 - [PutPaymentInstrumentResponseContent](docs/PutPaymentInstrumentResponseContent.md)
 - [PutPrescriptionRequestContent](docs/PutPrescriptionRequestContent.md)
 - [PutPrescriptionResponseContent](docs/PutPrescriptionResponseContent.md)
 - [Quantity](docs/Quantity.md)
 - [QuantityCodeListQualifier](docs/QuantityCodeListQualifier.md)
 - [QuantityUnitOfMeasureCode](docs/QuantityUnitOfMeasureCode.md)
 - [RejectedInsuranceDetails](docs/RejectedInsuranceDetails.md)
 - [ResourceNotFoundExceptionResponseContent](docs/ResourceNotFoundExceptionResponseContent.md)
 - [SexAssignedAtBirth](docs/SexAssignedAtBirth.md)
 - [ShippingDetails](docs/ShippingDetails.md)
 - [Strength](docs/Strength.md)
 - [StrengthFormCode](docs/StrengthFormCode.md)
 - [StrengthUnitOfMeasureCode](docs/StrengthUnitOfMeasureCode.md)
 - [ThrottlingExceptionResponseContent](docs/ThrottlingExceptionResponseContent.md)
 - [TransferRequestStatus](docs/TransferRequestStatus.md)
 - [TransferToPharmacy](docs/TransferToPharmacy.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### mutualTLS

### awsSigv4

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



