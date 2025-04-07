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

// checks if the TransferToPharmacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferToPharmacy{}

// TransferToPharmacy struct for TransferToPharmacy
type TransferToPharmacy struct {
	// An NCPDP-assigned national provider identification number that assists pharmacies in their interactions with federal agencies and third party providers.
	NcpdpId string `json:"ncpdpId" validate:"regexp=^\\\\s*\\\\S.*$"`
	// A National Provider System (NPS) assigned identifier for health care providers.
	Npi string `json:"npi" validate:"regexp=^\\\\d{10}$"`
	// The pharmacy this prescription needs transfer to.
	PharmacyName string `json:"pharmacyName" validate:"regexp=^\\\\s*\\\\S.*$"`
	// The Drug Enforcement Administration (DEA) assigned number to all businesses that manufacture or distribute controlled pharmaceuticals,         all health professionals who dispense, administer, or prescribe controlled pharmaceuticals and all pharmacies that dispense prescriptions.
	DeaNumber string `json:"deaNumber" validate:"regexp=^\\\\s*\\\\S.*$"`
	PharmacistName Name `json:"pharmacistName"`
	PharmacyAddress Address `json:"pharmacyAddress"`
	// Pharmacy telephone number.
	PrimaryTelephone string `json:"primaryTelephone" validate:"regexp=^1?[2-9]\\\\d{2}[2-9]\\\\d{6}$"`
	// Pharmacy fax number.
	Fax string `json:"fax" validate:"regexp=^1?[2-9]\\\\d{2}[2-9]\\\\d{6}$"`
}

type _TransferToPharmacy TransferToPharmacy

// NewTransferToPharmacy instantiates a new TransferToPharmacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferToPharmacy(ncpdpId string, npi string, pharmacyName string, deaNumber string, pharmacistName Name, pharmacyAddress Address, primaryTelephone string, fax string) *TransferToPharmacy {
	this := TransferToPharmacy{}
	this.NcpdpId = ncpdpId
	this.Npi = npi
	this.PharmacyName = pharmacyName
	this.DeaNumber = deaNumber
	this.PharmacistName = pharmacistName
	this.PharmacyAddress = pharmacyAddress
	this.PrimaryTelephone = primaryTelephone
	this.Fax = fax
	return &this
}

// NewTransferToPharmacyWithDefaults instantiates a new TransferToPharmacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferToPharmacyWithDefaults() *TransferToPharmacy {
	this := TransferToPharmacy{}
	return &this
}

// GetNcpdpId returns the NcpdpId field value
func (o *TransferToPharmacy) GetNcpdpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NcpdpId
}

// GetNcpdpIdOk returns a tuple with the NcpdpId field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetNcpdpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NcpdpId, true
}

// SetNcpdpId sets field value
func (o *TransferToPharmacy) SetNcpdpId(v string) {
	o.NcpdpId = v
}

// GetNpi returns the Npi field value
func (o *TransferToPharmacy) GetNpi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Npi
}

// GetNpiOk returns a tuple with the Npi field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetNpiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Npi, true
}

// SetNpi sets field value
func (o *TransferToPharmacy) SetNpi(v string) {
	o.Npi = v
}

// GetPharmacyName returns the PharmacyName field value
func (o *TransferToPharmacy) GetPharmacyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PharmacyName
}

// GetPharmacyNameOk returns a tuple with the PharmacyName field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetPharmacyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PharmacyName, true
}

// SetPharmacyName sets field value
func (o *TransferToPharmacy) SetPharmacyName(v string) {
	o.PharmacyName = v
}

// GetDeaNumber returns the DeaNumber field value
func (o *TransferToPharmacy) GetDeaNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeaNumber
}

// GetDeaNumberOk returns a tuple with the DeaNumber field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetDeaNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeaNumber, true
}

// SetDeaNumber sets field value
func (o *TransferToPharmacy) SetDeaNumber(v string) {
	o.DeaNumber = v
}

// GetPharmacistName returns the PharmacistName field value
func (o *TransferToPharmacy) GetPharmacistName() Name {
	if o == nil {
		var ret Name
		return ret
	}

	return o.PharmacistName
}

// GetPharmacistNameOk returns a tuple with the PharmacistName field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetPharmacistNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PharmacistName, true
}

// SetPharmacistName sets field value
func (o *TransferToPharmacy) SetPharmacistName(v Name) {
	o.PharmacistName = v
}

// GetPharmacyAddress returns the PharmacyAddress field value
func (o *TransferToPharmacy) GetPharmacyAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.PharmacyAddress
}

// GetPharmacyAddressOk returns a tuple with the PharmacyAddress field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetPharmacyAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PharmacyAddress, true
}

// SetPharmacyAddress sets field value
func (o *TransferToPharmacy) SetPharmacyAddress(v Address) {
	o.PharmacyAddress = v
}

// GetPrimaryTelephone returns the PrimaryTelephone field value
func (o *TransferToPharmacy) GetPrimaryTelephone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryTelephone
}

// GetPrimaryTelephoneOk returns a tuple with the PrimaryTelephone field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetPrimaryTelephoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryTelephone, true
}

// SetPrimaryTelephone sets field value
func (o *TransferToPharmacy) SetPrimaryTelephone(v string) {
	o.PrimaryTelephone = v
}

// GetFax returns the Fax field value
func (o *TransferToPharmacy) GetFax() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fax
}

// GetFaxOk returns a tuple with the Fax field value
// and a boolean to check if the value has been set.
func (o *TransferToPharmacy) GetFaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fax, true
}

// SetFax sets field value
func (o *TransferToPharmacy) SetFax(v string) {
	o.Fax = v
}

func (o TransferToPharmacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferToPharmacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ncpdpId"] = o.NcpdpId
	toSerialize["npi"] = o.Npi
	toSerialize["pharmacyName"] = o.PharmacyName
	toSerialize["deaNumber"] = o.DeaNumber
	toSerialize["pharmacistName"] = o.PharmacistName
	toSerialize["pharmacyAddress"] = o.PharmacyAddress
	toSerialize["primaryTelephone"] = o.PrimaryTelephone
	toSerialize["fax"] = o.Fax
	return toSerialize, nil
}

func (o *TransferToPharmacy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ncpdpId",
		"npi",
		"pharmacyName",
		"deaNumber",
		"pharmacistName",
		"pharmacyAddress",
		"primaryTelephone",
		"fax",
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

	varTransferToPharmacy := _TransferToPharmacy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferToPharmacy)

	if err != nil {
		return err
	}

	*o = TransferToPharmacy(varTransferToPharmacy)

	return err
}

type NullableTransferToPharmacy struct {
	value *TransferToPharmacy
	isSet bool
}

func (v NullableTransferToPharmacy) Get() *TransferToPharmacy {
	return v.value
}

func (v *NullableTransferToPharmacy) Set(val *TransferToPharmacy) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferToPharmacy) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferToPharmacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferToPharmacy(val *TransferToPharmacy) *NullableTransferToPharmacy {
	return &NullableTransferToPharmacy{value: val, isSet: true}
}

func (v NullableTransferToPharmacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferToPharmacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


