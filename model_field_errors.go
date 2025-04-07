/*
Amazon Pharmacy APIs for ingesting Prescriptions, Patients, Orders and Payment Instruments

Amazon Pharmacy APIs provide the following functionalities for external partners:  - Patient creation, update, and view - Prescription creation, update, and view - Order creation, update, and view - Payment instrument creation and update - Prescription transfer request creation  ## Changelog for the Amazon Pharmacy APIs (from version 1.15.3): ### 2025-03-26 v1.16.4 #### Added: - Added prescriber-direct prescription example to the `PutPrescription` API documentation.  #### Changed: - Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.  ### 2025-03-18 v1.16.3 #### Changed: - Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema. - For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation. - For prescriber-originated prescriptions: This structure is not required and should be omitted. - Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.  ### 2025-01-06 v1.16.2 #### Added: - The changes below are applied to `PutPrescription` API. - Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release. - Added a new required field `ndc` in `PreviousDispensedMedication`. - Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release. - The changes below are applied to `PutOrder` API. - Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.  #### Changed: - The changes below are applied to `PutPrescription` API. - Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation. - Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`. - Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`  ### 2025-01-06 v1.16.1 #### Added: - The changes below are applied to `PutPrescription` API. - Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`. - Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release. - Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`. - Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`. - Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`. - Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`. #### Changed: - The changes below are applied to `PutPrescription` API. - Deprecated `rxcui`. use `DrugDbCode` structure instead. - Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead. - Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead. - Deprecated `diagnosis`. Use `medicalDiagnosis` instead. - Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`. - Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.  ### 2024-12-13 v1.16.0 #### Changed: - Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances. - Updated `InsuranceList` max size to 27 from 50.  ### 2024-12-09 v1.15.10 #### Changed: - Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.  ### 2024-11-04 v1.15.9 #### Added: - New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request. #### Changed: - Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.  ### 2024-10-15 v1.15.8 #### Added: - Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model. - Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.  ### 2024-09-23 v1.15.7 #### Added: - Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues. - Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues. - Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.  ### 2024-09-03 v1.15.6 #### Added: - Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.   If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor. - The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field. - The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure. - A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException. - `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API. - Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.  ### 2024-08-06 v1.15.5 #### Added: - Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities. - Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions. - Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead. - Added documentation for enum fields for all APIs. - Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.  ### 2024-07-31 v1.15.4 #### Improved: - Updated the `directions` field to allow all strings for greater flexibility. - Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.  ### 2024-07-14 v1.15.3 #### Added: - All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs. 

API version: 1.16.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FieldErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldErrors{}

// FieldErrors Structure representing fields in error for a rejected claim  `DemographicErrors` will be present when ClaimRejectionType is INVALID_CUSTOMER_DEMOGRAPHIC_INFORMATION. This is when there is mismatch with customer demographic and Insurance demographic data  `InsuranceErrors` will be present when ClaimRejectionType is INVALID_INSURANCE_INFORMATION
type FieldErrors struct {
	// List of PossibleDemographicFieldsInError enums
	DemographicErrors []PossibleDemographicFieldsInError `json:"demographicErrors,omitempty"`
	// List of PossibleInsuranceFieldsInError enums
	InsuranceErrors []PossibleInsuranceFieldsInError `json:"insuranceErrors,omitempty"`
}

// NewFieldErrors instantiates a new FieldErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldErrors() *FieldErrors {
	this := FieldErrors{}
	return &this
}

// NewFieldErrorsWithDefaults instantiates a new FieldErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldErrorsWithDefaults() *FieldErrors {
	this := FieldErrors{}
	return &this
}

// GetDemographicErrors returns the DemographicErrors field value if set, zero value otherwise.
func (o *FieldErrors) GetDemographicErrors() []PossibleDemographicFieldsInError {
	if o == nil || IsNil(o.DemographicErrors) {
		var ret []PossibleDemographicFieldsInError
		return ret
	}
	return o.DemographicErrors
}

// GetDemographicErrorsOk returns a tuple with the DemographicErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldErrors) GetDemographicErrorsOk() ([]PossibleDemographicFieldsInError, bool) {
	if o == nil || IsNil(o.DemographicErrors) {
		return nil, false
	}
	return o.DemographicErrors, true
}

// HasDemographicErrors returns a boolean if a field has been set.
func (o *FieldErrors) HasDemographicErrors() bool {
	if o != nil && !IsNil(o.DemographicErrors) {
		return true
	}

	return false
}

// SetDemographicErrors gets a reference to the given []PossibleDemographicFieldsInError and assigns it to the DemographicErrors field.
func (o *FieldErrors) SetDemographicErrors(v []PossibleDemographicFieldsInError) {
	o.DemographicErrors = v
}

// GetInsuranceErrors returns the InsuranceErrors field value if set, zero value otherwise.
func (o *FieldErrors) GetInsuranceErrors() []PossibleInsuranceFieldsInError {
	if o == nil || IsNil(o.InsuranceErrors) {
		var ret []PossibleInsuranceFieldsInError
		return ret
	}
	return o.InsuranceErrors
}

// GetInsuranceErrorsOk returns a tuple with the InsuranceErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldErrors) GetInsuranceErrorsOk() ([]PossibleInsuranceFieldsInError, bool) {
	if o == nil || IsNil(o.InsuranceErrors) {
		return nil, false
	}
	return o.InsuranceErrors, true
}

// HasInsuranceErrors returns a boolean if a field has been set.
func (o *FieldErrors) HasInsuranceErrors() bool {
	if o != nil && !IsNil(o.InsuranceErrors) {
		return true
	}

	return false
}

// SetInsuranceErrors gets a reference to the given []PossibleInsuranceFieldsInError and assigns it to the InsuranceErrors field.
func (o *FieldErrors) SetInsuranceErrors(v []PossibleInsuranceFieldsInError) {
	o.InsuranceErrors = v
}

func (o FieldErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DemographicErrors) {
		toSerialize["demographicErrors"] = o.DemographicErrors
	}
	if !IsNil(o.InsuranceErrors) {
		toSerialize["insuranceErrors"] = o.InsuranceErrors
	}
	return toSerialize, nil
}

type NullableFieldErrors struct {
	value *FieldErrors
	isSet bool
}

func (v NullableFieldErrors) Get() *FieldErrors {
	return v.value
}

func (v *NullableFieldErrors) Set(val *FieldErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldErrors(val *FieldErrors) *NullableFieldErrors {
	return &NullableFieldErrors{value: val, isSet: true}
}

func (v NullableFieldErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


