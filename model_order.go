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

// checks if the Order type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Order{}

// Order Structure representing an order.
type Order struct {
	Status OrderStatus `json:"status"`
	// The partner ID of the prescription this order is associated with.
	PartnerPrescriptionId string `json:"partnerPrescriptionId" validate:"regexp=^[0-9a-zA-Z-]{1,40}$"`
	StatusReasonCode *OrderStatusReasonCode `json:"statusReasonCode,omitempty"`
	// A more detailed explanation of the order's status.
	StatusDescription *string `json:"statusDescription,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
	AdditionalOrderDetails *AdditionalOrderDetails `json:"additionalOrderDetails,omitempty"`
	// The date the fill was created, in the format 'YYYY-MM-DD'.
	FilledDate *string `json:"filledDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// The date the fill was shipped, in the format 'YYYY-MM-DD'.
	ShipDate *string `json:"shipDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// The number of times this prescription has been filled.
	FillNumber *float32 `json:"fillNumber,omitempty"`
	// The number of units in dispense.
	QuantityDispensed *float32 `json:"quantityDispensed,omitempty"`
}

type _Order Order

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(status OrderStatus, partnerPrescriptionId string) *Order {
	this := Order{}
	this.Status = status
	this.PartnerPrescriptionId = partnerPrescriptionId
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetStatus returns the Status field value
func (o *Order) GetStatus() OrderStatus {
	if o == nil {
		var ret OrderStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Order) GetStatusOk() (*OrderStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Order) SetStatus(v OrderStatus) {
	o.Status = v
}

// GetPartnerPrescriptionId returns the PartnerPrescriptionId field value
func (o *Order) GetPartnerPrescriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerPrescriptionId
}

// GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field value
// and a boolean to check if the value has been set.
func (o *Order) GetPartnerPrescriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerPrescriptionId, true
}

// SetPartnerPrescriptionId sets field value
func (o *Order) SetPartnerPrescriptionId(v string) {
	o.PartnerPrescriptionId = v
}

// GetStatusReasonCode returns the StatusReasonCode field value if set, zero value otherwise.
func (o *Order) GetStatusReasonCode() OrderStatusReasonCode {
	if o == nil || IsNil(o.StatusReasonCode) {
		var ret OrderStatusReasonCode
		return ret
	}
	return *o.StatusReasonCode
}

// GetStatusReasonCodeOk returns a tuple with the StatusReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetStatusReasonCodeOk() (*OrderStatusReasonCode, bool) {
	if o == nil || IsNil(o.StatusReasonCode) {
		return nil, false
	}
	return o.StatusReasonCode, true
}

// HasStatusReasonCode returns a boolean if a field has been set.
func (o *Order) HasStatusReasonCode() bool {
	if o != nil && !IsNil(o.StatusReasonCode) {
		return true
	}

	return false
}

// SetStatusReasonCode gets a reference to the given OrderStatusReasonCode and assigns it to the StatusReasonCode field.
func (o *Order) SetStatusReasonCode(v OrderStatusReasonCode) {
	o.StatusReasonCode = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *Order) GetStatusDescription() string {
	if o == nil || IsNil(o.StatusDescription) {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDescription) {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *Order) HasStatusDescription() bool {
	if o != nil && !IsNil(o.StatusDescription) {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *Order) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetAdditionalOrderDetails returns the AdditionalOrderDetails field value if set, zero value otherwise.
func (o *Order) GetAdditionalOrderDetails() AdditionalOrderDetails {
	if o == nil || IsNil(o.AdditionalOrderDetails) {
		var ret AdditionalOrderDetails
		return ret
	}
	return *o.AdditionalOrderDetails
}

// GetAdditionalOrderDetailsOk returns a tuple with the AdditionalOrderDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetAdditionalOrderDetailsOk() (*AdditionalOrderDetails, bool) {
	if o == nil || IsNil(o.AdditionalOrderDetails) {
		return nil, false
	}
	return o.AdditionalOrderDetails, true
}

// HasAdditionalOrderDetails returns a boolean if a field has been set.
func (o *Order) HasAdditionalOrderDetails() bool {
	if o != nil && !IsNil(o.AdditionalOrderDetails) {
		return true
	}

	return false
}

// SetAdditionalOrderDetails gets a reference to the given AdditionalOrderDetails and assigns it to the AdditionalOrderDetails field.
func (o *Order) SetAdditionalOrderDetails(v AdditionalOrderDetails) {
	o.AdditionalOrderDetails = &v
}

// GetFilledDate returns the FilledDate field value if set, zero value otherwise.
func (o *Order) GetFilledDate() string {
	if o == nil || IsNil(o.FilledDate) {
		var ret string
		return ret
	}
	return *o.FilledDate
}

// GetFilledDateOk returns a tuple with the FilledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetFilledDateOk() (*string, bool) {
	if o == nil || IsNil(o.FilledDate) {
		return nil, false
	}
	return o.FilledDate, true
}

// HasFilledDate returns a boolean if a field has been set.
func (o *Order) HasFilledDate() bool {
	if o != nil && !IsNil(o.FilledDate) {
		return true
	}

	return false
}

// SetFilledDate gets a reference to the given string and assigns it to the FilledDate field.
func (o *Order) SetFilledDate(v string) {
	o.FilledDate = &v
}

// GetShipDate returns the ShipDate field value if set, zero value otherwise.
func (o *Order) GetShipDate() string {
	if o == nil || IsNil(o.ShipDate) {
		var ret string
		return ret
	}
	return *o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetShipDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShipDate) {
		return nil, false
	}
	return o.ShipDate, true
}

// HasShipDate returns a boolean if a field has been set.
func (o *Order) HasShipDate() bool {
	if o != nil && !IsNil(o.ShipDate) {
		return true
	}

	return false
}

// SetShipDate gets a reference to the given string and assigns it to the ShipDate field.
func (o *Order) SetShipDate(v string) {
	o.ShipDate = &v
}

// GetFillNumber returns the FillNumber field value if set, zero value otherwise.
func (o *Order) GetFillNumber() float32 {
	if o == nil || IsNil(o.FillNumber) {
		var ret float32
		return ret
	}
	return *o.FillNumber
}

// GetFillNumberOk returns a tuple with the FillNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetFillNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.FillNumber) {
		return nil, false
	}
	return o.FillNumber, true
}

// HasFillNumber returns a boolean if a field has been set.
func (o *Order) HasFillNumber() bool {
	if o != nil && !IsNil(o.FillNumber) {
		return true
	}

	return false
}

// SetFillNumber gets a reference to the given float32 and assigns it to the FillNumber field.
func (o *Order) SetFillNumber(v float32) {
	o.FillNumber = &v
}

// GetQuantityDispensed returns the QuantityDispensed field value if set, zero value otherwise.
func (o *Order) GetQuantityDispensed() float32 {
	if o == nil || IsNil(o.QuantityDispensed) {
		var ret float32
		return ret
	}
	return *o.QuantityDispensed
}

// GetQuantityDispensedOk returns a tuple with the QuantityDispensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetQuantityDispensedOk() (*float32, bool) {
	if o == nil || IsNil(o.QuantityDispensed) {
		return nil, false
	}
	return o.QuantityDispensed, true
}

// HasQuantityDispensed returns a boolean if a field has been set.
func (o *Order) HasQuantityDispensed() bool {
	if o != nil && !IsNil(o.QuantityDispensed) {
		return true
	}

	return false
}

// SetQuantityDispensed gets a reference to the given float32 and assigns it to the QuantityDispensed field.
func (o *Order) SetQuantityDispensed(v float32) {
	o.QuantityDispensed = &v
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Order) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["partnerPrescriptionId"] = o.PartnerPrescriptionId
	if !IsNil(o.StatusReasonCode) {
		toSerialize["statusReasonCode"] = o.StatusReasonCode
	}
	if !IsNil(o.StatusDescription) {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	if !IsNil(o.AdditionalOrderDetails) {
		toSerialize["additionalOrderDetails"] = o.AdditionalOrderDetails
	}
	if !IsNil(o.FilledDate) {
		toSerialize["filledDate"] = o.FilledDate
	}
	if !IsNil(o.ShipDate) {
		toSerialize["shipDate"] = o.ShipDate
	}
	if !IsNil(o.FillNumber) {
		toSerialize["fillNumber"] = o.FillNumber
	}
	if !IsNil(o.QuantityDispensed) {
		toSerialize["quantityDispensed"] = o.QuantityDispensed
	}
	return toSerialize, nil
}

func (o *Order) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"partnerPrescriptionId",
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

	varOrder := _Order{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrder)

	if err != nil {
		return err
	}

	*o = Order(varOrder)

	return err
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


