/*
Amazon Pharmacy APIs for ingesting Prescriptions, Patients, Orders and Payment Instruments

Amazon Pharmacy APIs provide the following functionalities for external partners:  - Patient creation, update, and view - Prescription creation, update, and view - Order creation, update, and view - Payment instrument creation and update - Prescription transfer request creation  ## Changelog for the Amazon Pharmacy APIs (from version 1.15.3): ### 2025-03-26 v1.16.4 #### Added: - Added prescriber-direct prescription example to the `PutPrescription` API documentation.  #### Changed: - Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.  ### 2025-03-18 v1.16.3 #### Changed: - Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema. - For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation. - For prescriber-originated prescriptions: This structure is not required and should be omitted. - Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.  ### 2025-01-06 v1.16.2 #### Added: - The changes below are applied to `PutPrescription` API. - Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release. - Added a new required field `ndc` in `PreviousDispensedMedication`. - Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release. - The changes below are applied to `PutOrder` API. - Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.  #### Changed: - The changes below are applied to `PutPrescription` API. - Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation. - Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`. - Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`  ### 2025-01-06 v1.16.1 #### Added: - The changes below are applied to `PutPrescription` API. - Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`. - Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release. - Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`. - Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`. - Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`. - Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`. #### Changed: - The changes below are applied to `PutPrescription` API. - Deprecated `rxcui`. use `DrugDbCode` structure instead. - Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead. - Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead. - Deprecated `diagnosis`. Use `medicalDiagnosis` instead. - Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`. - Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.  ### 2024-12-13 v1.16.0 #### Changed: - Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances. - Updated `InsuranceList` max size to 27 from 50.  ### 2024-12-09 v1.15.10 #### Changed: - Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.  ### 2024-11-04 v1.15.9 #### Added: - New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request. #### Changed: - Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.  ### 2024-10-15 v1.15.8 #### Added: - Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model. - Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.  ### 2024-09-23 v1.15.7 #### Added: - Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues. - Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues. - Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.  ### 2024-09-03 v1.15.6 #### Added: - Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.   If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor. - The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field. - The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure. - A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException. - `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API. - Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.  ### 2024-08-06 v1.15.5 #### Added: - Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities. - Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions. - Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead. - Added documentation for enum fields for all APIs. - Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.  ### 2024-07-31 v1.15.4 #### Improved: - Updated the `directions` field to allow all strings for greater flexibility. - Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.  ### 2024-07-14 v1.15.3 #### Added: - All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs. 

API version: 1.16.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amazonpharmacy

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RejectedInsuranceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectedInsuranceDetails{}

// RejectedInsuranceDetails Structure representing Insurance details used for a rejected claim
type RejectedInsuranceDetails struct {
	// Processor Control Number assigned by the processor. PBMs use it to distinguish between plans. It's a secondary identifier used in addition to the binNumber.
	Pcn *string `json:"pcn,omitempty" validate:"regexp=^[a-zA-Z0-9 ]+$"`
	// The insurance group number found on the insurance card that signifies which group the cardholder belongs to.
	GroupNumber *string `json:"groupNumber,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
	// A 6-digit number that identifies the Pharmacy Benefit Manager (PBM) that this insurance uses for pharmacy claims.
	BinNumber string `json:"binNumber" validate:"regexp=^\\\\d{6}$"`
	PayerType *PayerType `json:"payerType,omitempty"`
	// The name of the payer associated with the insurance plan.
	PayerName *string `json:"payerName,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
	// Optional state code related to the insurance plan, used for state-specific insurance information.
	StateCode *string `json:"stateCode,omitempty"`
}

type _RejectedInsuranceDetails RejectedInsuranceDetails

// NewRejectedInsuranceDetails instantiates a new RejectedInsuranceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectedInsuranceDetails(binNumber string) *RejectedInsuranceDetails {
	this := RejectedInsuranceDetails{}
	this.BinNumber = binNumber
	return &this
}

// NewRejectedInsuranceDetailsWithDefaults instantiates a new RejectedInsuranceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectedInsuranceDetailsWithDefaults() *RejectedInsuranceDetails {
	this := RejectedInsuranceDetails{}
	return &this
}

// GetPcn returns the Pcn field value if set, zero value otherwise.
func (o *RejectedInsuranceDetails) GetPcn() string {
	if o == nil || IsNil(o.Pcn) {
		var ret string
		return ret
	}
	return *o.Pcn
}

// GetPcnOk returns a tuple with the Pcn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectedInsuranceDetails) GetPcnOk() (*string, bool) {
	if o == nil || IsNil(o.Pcn) {
		return nil, false
	}
	return o.Pcn, true
}

// HasPcn returns a boolean if a field has been set.
func (o *RejectedInsuranceDetails) HasPcn() bool {
	if o != nil && !IsNil(o.Pcn) {
		return true
	}

	return false
}

// SetPcn gets a reference to the given string and assigns it to the Pcn field.
func (o *RejectedInsuranceDetails) SetPcn(v string) {
	o.Pcn = &v
}

// GetGroupNumber returns the GroupNumber field value if set, zero value otherwise.
func (o *RejectedInsuranceDetails) GetGroupNumber() string {
	if o == nil || IsNil(o.GroupNumber) {
		var ret string
		return ret
	}
	return *o.GroupNumber
}

// GetGroupNumberOk returns a tuple with the GroupNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectedInsuranceDetails) GetGroupNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GroupNumber) {
		return nil, false
	}
	return o.GroupNumber, true
}

// HasGroupNumber returns a boolean if a field has been set.
func (o *RejectedInsuranceDetails) HasGroupNumber() bool {
	if o != nil && !IsNil(o.GroupNumber) {
		return true
	}

	return false
}

// SetGroupNumber gets a reference to the given string and assigns it to the GroupNumber field.
func (o *RejectedInsuranceDetails) SetGroupNumber(v string) {
	o.GroupNumber = &v
}

// GetBinNumber returns the BinNumber field value
func (o *RejectedInsuranceDetails) GetBinNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinNumber
}

// GetBinNumberOk returns a tuple with the BinNumber field value
// and a boolean to check if the value has been set.
func (o *RejectedInsuranceDetails) GetBinNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinNumber, true
}

// SetBinNumber sets field value
func (o *RejectedInsuranceDetails) SetBinNumber(v string) {
	o.BinNumber = v
}

// GetPayerType returns the PayerType field value if set, zero value otherwise.
func (o *RejectedInsuranceDetails) GetPayerType() PayerType {
	if o == nil || IsNil(o.PayerType) {
		var ret PayerType
		return ret
	}
	return *o.PayerType
}

// GetPayerTypeOk returns a tuple with the PayerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectedInsuranceDetails) GetPayerTypeOk() (*PayerType, bool) {
	if o == nil || IsNil(o.PayerType) {
		return nil, false
	}
	return o.PayerType, true
}

// HasPayerType returns a boolean if a field has been set.
func (o *RejectedInsuranceDetails) HasPayerType() bool {
	if o != nil && !IsNil(o.PayerType) {
		return true
	}

	return false
}

// SetPayerType gets a reference to the given PayerType and assigns it to the PayerType field.
func (o *RejectedInsuranceDetails) SetPayerType(v PayerType) {
	o.PayerType = &v
}

// GetPayerName returns the PayerName field value if set, zero value otherwise.
func (o *RejectedInsuranceDetails) GetPayerName() string {
	if o == nil || IsNil(o.PayerName) {
		var ret string
		return ret
	}
	return *o.PayerName
}

// GetPayerNameOk returns a tuple with the PayerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectedInsuranceDetails) GetPayerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayerName) {
		return nil, false
	}
	return o.PayerName, true
}

// HasPayerName returns a boolean if a field has been set.
func (o *RejectedInsuranceDetails) HasPayerName() bool {
	if o != nil && !IsNil(o.PayerName) {
		return true
	}

	return false
}

// SetPayerName gets a reference to the given string and assigns it to the PayerName field.
func (o *RejectedInsuranceDetails) SetPayerName(v string) {
	o.PayerName = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *RejectedInsuranceDetails) GetStateCode() string {
	if o == nil || IsNil(o.StateCode) {
		var ret string
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectedInsuranceDetails) GetStateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StateCode) {
		return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *RejectedInsuranceDetails) HasStateCode() bool {
	if o != nil && !IsNil(o.StateCode) {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given string and assigns it to the StateCode field.
func (o *RejectedInsuranceDetails) SetStateCode(v string) {
	o.StateCode = &v
}

func (o RejectedInsuranceDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RejectedInsuranceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pcn) {
		toSerialize["pcn"] = o.Pcn
	}
	if !IsNil(o.GroupNumber) {
		toSerialize["groupNumber"] = o.GroupNumber
	}
	toSerialize["binNumber"] = o.BinNumber
	if !IsNil(o.PayerType) {
		toSerialize["payerType"] = o.PayerType
	}
	if !IsNil(o.PayerName) {
		toSerialize["payerName"] = o.PayerName
	}
	if !IsNil(o.StateCode) {
		toSerialize["stateCode"] = o.StateCode
	}
	return toSerialize, nil
}

func (o *RejectedInsuranceDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"binNumber",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRejectedInsuranceDetails := _RejectedInsuranceDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRejectedInsuranceDetails)

	if err != nil {
		return err
	}

	*o = RejectedInsuranceDetails(varRejectedInsuranceDetails)

	return err
}

type NullableRejectedInsuranceDetails struct {
	value *RejectedInsuranceDetails
	isSet bool
}

func (v NullableRejectedInsuranceDetails) Get() *RejectedInsuranceDetails {
	return v.value
}

func (v *NullableRejectedInsuranceDetails) Set(val *RejectedInsuranceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectedInsuranceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectedInsuranceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectedInsuranceDetails(val *RejectedInsuranceDetails) *NullableRejectedInsuranceDetails {
	return &NullableRejectedInsuranceDetails{value: val, isSet: true}
}

func (v NullableRejectedInsuranceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectedInsuranceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


