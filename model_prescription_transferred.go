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

// checks if the PrescriptionTransferred type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrescriptionTransferred{}

// PrescriptionTransferred Structure capturing details about fills and quantity for the prescription being transferred.
type PrescriptionTransferred struct {
	PrescriptionFillTransferInType PrescriptionFillTransferInType `json:"prescriptionFillTransferInType"`
	// The date on which the prescription was first filled by the pharmacy, in the format 'YYYY-MM-DD'. This date helps track the initial dispensing of the medication to the patient. This field is required if prescription is previously dispensed (PrescriptionFillTransferInType.DISPENSED_BEFORE).
	FirstFillDate *string `json:"firstFillDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// The date on which the prescription was most recently filled by the pharmacy, in the format 'YYYY-MM-DD'. This date helps track the latest dispensing of the medication and is useful for monitoring adherence and refill patterns. This field is required if prescription is previously dispensed (PrescriptionFillTransferInType.DISPENSED_BEFORE).
	LastFillDate *string `json:"lastFillDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// The number of refills left on the prescription. This field will be marked as required for both NewRx and previously dispensed Rx starting from 2025-03-12.
	RefillsRemaining *float64 `json:"refillsRemaining,omitempty"`
	// The number of refills being transferred to Amazon Pharmacy. This field will be marked as required for both NewRx and previously dispensed Rx starting from 2025-03-12.
	RefillsTransferred *float64 `json:"refillsTransferred,omitempty"`
	QuantityTransferred *Quantity `json:"quantityTransferred,omitempty"`
	QuantityRemaining *Quantity `json:"quantityRemaining,omitempty"`
	// Contains information about previous medication dispenses for the prescription. Each entry represents a medication dispensed by the sending pharmacy. This provides details about actual dispensing when the filled medication differs from what was prescribed. When PrescriptionFillTransferInType.DISPENSED_BEFORE is true, this object must contain either at least one aggregate dispense detail or a list of unique previous dispenses. This field will be marked as required for previously dispensed Rx starting from 2025-03-12.
	PreviousDispensedMedications []PreviousDispensedMedication `json:"previousDispensedMedications,omitempty"`
	// Prescription number assigned by the sending pharmacy. This field will be marked as required starting from 2025-03-12.
	PharmacyRxNumber *string `json:"pharmacyRxNumber,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
}

type _PrescriptionTransferred PrescriptionTransferred

// NewPrescriptionTransferred instantiates a new PrescriptionTransferred object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrescriptionTransferred(prescriptionFillTransferInType PrescriptionFillTransferInType) *PrescriptionTransferred {
	this := PrescriptionTransferred{}
	this.PrescriptionFillTransferInType = prescriptionFillTransferInType
	return &this
}

// NewPrescriptionTransferredWithDefaults instantiates a new PrescriptionTransferred object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrescriptionTransferredWithDefaults() *PrescriptionTransferred {
	this := PrescriptionTransferred{}
	return &this
}

// GetPrescriptionFillTransferInType returns the PrescriptionFillTransferInType field value
func (o *PrescriptionTransferred) GetPrescriptionFillTransferInType() PrescriptionFillTransferInType {
	if o == nil {
		var ret PrescriptionFillTransferInType
		return ret
	}

	return o.PrescriptionFillTransferInType
}

// GetPrescriptionFillTransferInTypeOk returns a tuple with the PrescriptionFillTransferInType field value
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetPrescriptionFillTransferInTypeOk() (*PrescriptionFillTransferInType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrescriptionFillTransferInType, true
}

// SetPrescriptionFillTransferInType sets field value
func (o *PrescriptionTransferred) SetPrescriptionFillTransferInType(v PrescriptionFillTransferInType) {
	o.PrescriptionFillTransferInType = v
}

// GetFirstFillDate returns the FirstFillDate field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetFirstFillDate() string {
	if o == nil || IsNil(o.FirstFillDate) {
		var ret string
		return ret
	}
	return *o.FirstFillDate
}

// GetFirstFillDateOk returns a tuple with the FirstFillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetFirstFillDateOk() (*string, bool) {
	if o == nil || IsNil(o.FirstFillDate) {
		return nil, false
	}
	return o.FirstFillDate, true
}

// HasFirstFillDate returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasFirstFillDate() bool {
	if o != nil && !IsNil(o.FirstFillDate) {
		return true
	}

	return false
}

// SetFirstFillDate gets a reference to the given string and assigns it to the FirstFillDate field.
func (o *PrescriptionTransferred) SetFirstFillDate(v string) {
	o.FirstFillDate = &v
}

// GetLastFillDate returns the LastFillDate field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetLastFillDate() string {
	if o == nil || IsNil(o.LastFillDate) {
		var ret string
		return ret
	}
	return *o.LastFillDate
}

// GetLastFillDateOk returns a tuple with the LastFillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetLastFillDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastFillDate) {
		return nil, false
	}
	return o.LastFillDate, true
}

// HasLastFillDate returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasLastFillDate() bool {
	if o != nil && !IsNil(o.LastFillDate) {
		return true
	}

	return false
}

// SetLastFillDate gets a reference to the given string and assigns it to the LastFillDate field.
func (o *PrescriptionTransferred) SetLastFillDate(v string) {
	o.LastFillDate = &v
}

// GetRefillsRemaining returns the RefillsRemaining field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetRefillsRemaining() float64 {
	if o == nil || IsNil(o.RefillsRemaining) {
		var ret float64
		return ret
	}
	return *o.RefillsRemaining
}

// GetRefillsRemainingOk returns a tuple with the RefillsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetRefillsRemainingOk() (*float64, bool) {
	if o == nil || IsNil(o.RefillsRemaining) {
		return nil, false
	}
	return o.RefillsRemaining, true
}

// HasRefillsRemaining returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasRefillsRemaining() bool {
	if o != nil && !IsNil(o.RefillsRemaining) {
		return true
	}

	return false
}

// SetRefillsRemaining gets a reference to the given float64 and assigns it to the RefillsRemaining field.
func (o *PrescriptionTransferred) SetRefillsRemaining(v float64) {
	o.RefillsRemaining = &v
}

// GetRefillsTransferred returns the RefillsTransferred field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetRefillsTransferred() float64 {
	if o == nil || IsNil(o.RefillsTransferred) {
		var ret float64
		return ret
	}
	return *o.RefillsTransferred
}

// GetRefillsTransferredOk returns a tuple with the RefillsTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetRefillsTransferredOk() (*float64, bool) {
	if o == nil || IsNil(o.RefillsTransferred) {
		return nil, false
	}
	return o.RefillsTransferred, true
}

// HasRefillsTransferred returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasRefillsTransferred() bool {
	if o != nil && !IsNil(o.RefillsTransferred) {
		return true
	}

	return false
}

// SetRefillsTransferred gets a reference to the given float64 and assigns it to the RefillsTransferred field.
func (o *PrescriptionTransferred) SetRefillsTransferred(v float64) {
	o.RefillsTransferred = &v
}

// GetQuantityTransferred returns the QuantityTransferred field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetQuantityTransferred() Quantity {
	if o == nil || IsNil(o.QuantityTransferred) {
		var ret Quantity
		return ret
	}
	return *o.QuantityTransferred
}

// GetQuantityTransferredOk returns a tuple with the QuantityTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetQuantityTransferredOk() (*Quantity, bool) {
	if o == nil || IsNil(o.QuantityTransferred) {
		return nil, false
	}
	return o.QuantityTransferred, true
}

// HasQuantityTransferred returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasQuantityTransferred() bool {
	if o != nil && !IsNil(o.QuantityTransferred) {
		return true
	}

	return false
}

// SetQuantityTransferred gets a reference to the given Quantity and assigns it to the QuantityTransferred field.
func (o *PrescriptionTransferred) SetQuantityTransferred(v Quantity) {
	o.QuantityTransferred = &v
}

// GetQuantityRemaining returns the QuantityRemaining field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetQuantityRemaining() Quantity {
	if o == nil || IsNil(o.QuantityRemaining) {
		var ret Quantity
		return ret
	}
	return *o.QuantityRemaining
}

// GetQuantityRemainingOk returns a tuple with the QuantityRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetQuantityRemainingOk() (*Quantity, bool) {
	if o == nil || IsNil(o.QuantityRemaining) {
		return nil, false
	}
	return o.QuantityRemaining, true
}

// HasQuantityRemaining returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasQuantityRemaining() bool {
	if o != nil && !IsNil(o.QuantityRemaining) {
		return true
	}

	return false
}

// SetQuantityRemaining gets a reference to the given Quantity and assigns it to the QuantityRemaining field.
func (o *PrescriptionTransferred) SetQuantityRemaining(v Quantity) {
	o.QuantityRemaining = &v
}

// GetPreviousDispensedMedications returns the PreviousDispensedMedications field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetPreviousDispensedMedications() []PreviousDispensedMedication {
	if o == nil || IsNil(o.PreviousDispensedMedications) {
		var ret []PreviousDispensedMedication
		return ret
	}
	return o.PreviousDispensedMedications
}

// GetPreviousDispensedMedicationsOk returns a tuple with the PreviousDispensedMedications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetPreviousDispensedMedicationsOk() ([]PreviousDispensedMedication, bool) {
	if o == nil || IsNil(o.PreviousDispensedMedications) {
		return nil, false
	}
	return o.PreviousDispensedMedications, true
}

// HasPreviousDispensedMedications returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasPreviousDispensedMedications() bool {
	if o != nil && !IsNil(o.PreviousDispensedMedications) {
		return true
	}

	return false
}

// SetPreviousDispensedMedications gets a reference to the given []PreviousDispensedMedication and assigns it to the PreviousDispensedMedications field.
func (o *PrescriptionTransferred) SetPreviousDispensedMedications(v []PreviousDispensedMedication) {
	o.PreviousDispensedMedications = v
}

// GetPharmacyRxNumber returns the PharmacyRxNumber field value if set, zero value otherwise.
func (o *PrescriptionTransferred) GetPharmacyRxNumber() string {
	if o == nil || IsNil(o.PharmacyRxNumber) {
		var ret string
		return ret
	}
	return *o.PharmacyRxNumber
}

// GetPharmacyRxNumberOk returns a tuple with the PharmacyRxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrescriptionTransferred) GetPharmacyRxNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PharmacyRxNumber) {
		return nil, false
	}
	return o.PharmacyRxNumber, true
}

// HasPharmacyRxNumber returns a boolean if a field has been set.
func (o *PrescriptionTransferred) HasPharmacyRxNumber() bool {
	if o != nil && !IsNil(o.PharmacyRxNumber) {
		return true
	}

	return false
}

// SetPharmacyRxNumber gets a reference to the given string and assigns it to the PharmacyRxNumber field.
func (o *PrescriptionTransferred) SetPharmacyRxNumber(v string) {
	o.PharmacyRxNumber = &v
}

func (o PrescriptionTransferred) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrescriptionTransferred) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prescriptionFillTransferInType"] = o.PrescriptionFillTransferInType
	if !IsNil(o.FirstFillDate) {
		toSerialize["firstFillDate"] = o.FirstFillDate
	}
	if !IsNil(o.LastFillDate) {
		toSerialize["lastFillDate"] = o.LastFillDate
	}
	if !IsNil(o.RefillsRemaining) {
		toSerialize["refillsRemaining"] = o.RefillsRemaining
	}
	if !IsNil(o.RefillsTransferred) {
		toSerialize["refillsTransferred"] = o.RefillsTransferred
	}
	if !IsNil(o.QuantityTransferred) {
		toSerialize["quantityTransferred"] = o.QuantityTransferred
	}
	if !IsNil(o.QuantityRemaining) {
		toSerialize["quantityRemaining"] = o.QuantityRemaining
	}
	if !IsNil(o.PreviousDispensedMedications) {
		toSerialize["previousDispensedMedications"] = o.PreviousDispensedMedications
	}
	if !IsNil(o.PharmacyRxNumber) {
		toSerialize["pharmacyRxNumber"] = o.PharmacyRxNumber
	}
	return toSerialize, nil
}

func (o *PrescriptionTransferred) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prescriptionFillTransferInType",
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

	varPrescriptionTransferred := _PrescriptionTransferred{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrescriptionTransferred)

	if err != nil {
		return err
	}

	*o = PrescriptionTransferred(varPrescriptionTransferred)

	return err
}

type NullablePrescriptionTransferred struct {
	value *PrescriptionTransferred
	isSet bool
}

func (v NullablePrescriptionTransferred) Get() *PrescriptionTransferred {
	return v.value
}

func (v *NullablePrescriptionTransferred) Set(val *PrescriptionTransferred) {
	v.value = val
	v.isSet = true
}

func (v NullablePrescriptionTransferred) IsSet() bool {
	return v.isSet
}

func (v *NullablePrescriptionTransferred) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrescriptionTransferred(val *PrescriptionTransferred) *NullablePrescriptionTransferred {
	return &NullablePrescriptionTransferred{value: val, isSet: true}
}

func (v NullablePrescriptionTransferred) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrescriptionTransferred) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


