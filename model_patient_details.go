/*
Amazon Pharmacy APIs for ingesting Prescriptions, Patients, Orders and Payment Instruments

Amazon Pharmacy APIs provide the following functionalities for external partners:  - Patient creation, update, and view - Prescription creation, update, and view - Order creation, update, and view - Payment instrument creation and update - Prescription transfer request creation  ## Changelog for the Amazon Pharmacy APIs (from version 1.15.3):  ### 2025-07-14 v1.20.1 #### Changed: - Updated the regex pattern for `patientDetails.name`, `prescriber.prescriberName`, `pharmacist.pharmacistName` field to accept special characters like `-`, `'`, `.` for `firstName`, `middleName`, `lastName` subfields.  ### 2025-07-07 v1.20.0 #### Added: - `INVALID_DELIVERY_ADDRESS` to `OrderStatusReasonCode`.  #### Changed: - Updated `MedicationPrescribed` structure to make the `daysSupply` field optional in `PutPrescription API`, allowing prescriber partners to omit this field when calling the PutPrescription API  ### 2025-07-01 v1.19.0 #### Added: - Asynchronous notifications for orders - Added documentation in `PutOrder` callback with `EventNotification` schema which signifies notification structure  #### Notes: - Order notifications are available for below order statuses for orders created through `PutOrder` API   - ORDER_CREATED   - ORDER_PROCESSING   - ORDER_ON_HOLD   - ORDER_SHIPPED   - ORDER_CANCELLED   - ORDER_FAILED  ### 2025-06-19 v1.18.2 #### Changed: - Updated `PatientDetails` structure to make the `sexAssignedAtBirth` field optional in `PutPatient API`, allowing partners to omit this field when calling the PutPatient API; however, this field will continue to be mandatory for hub partners who manage orders on behalf of customers  ### 2025-06-10 v1.18.1 #### Changed: - Removed mandatory requirement for the deprecated `unitOfMeasureCode` field in `Quantity`. - Updated documentation for quantity and refill fields. - Added comprehensive examples covering various prescription transfer scenarios. - Updated documentation to reflect optional status of `refillsTransferred` field.  ### 2025-06-03 v1.18.0 #### Added: - Added optional field `Observation` in PutPrescriptionRequest to support vital signs data. - Added new structures:   - `Observation` with measurement list and notes   - `Measurement` with vital sign, value, unit, and date fields   - `VitalSign` enum supporting HEIGHT and WEIGHT measurements - Updated `PutPrescription` API examples to demonstrate usage of Observation structure with height and weight measurements.   #### Changed:  - Updated `PatientDetails` structure to make the `smsConsent` field optional, allowing partners to omit this field when calling the PutPatient API. When not provided, `smsConsent` defaults to `false`.  ### 2025-03-24 v1.17.0 #### Added: - Asynchronous notifications for prescriptions. - Added documentation in `PutPrescription` callback with `EventNotification` schema which signifies notification structure  #### Changed: - `Insurance` now supports `cardholderId`. - `Insurance` and `OrderInsurance` field `encryptedCardholderId` is now deprecated. Please use `cardholderId` instead for `PutPatient` and `PutOrder` calls.  ### 2025-03-26 v1.16.4 #### Added: - Added prescriber-direct prescription example to the `PutPrescription` API documentation.  #### Changed: - Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.  ### 2025-03-18 v1.16.3 #### Changed: - Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema. - For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation. - For prescriber-originated prescriptions: This structure is not required and should be omitted. - Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.   ### 2025-01-06 v1.16.2 #### Added: - The changes below are applied to `PutPrescription` API. - Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release. - Added a new required field `ndc` in `PreviousDispensedMedication`. - Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release. - The changes below are applied to `PutOrder` API. - Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.  #### Changed: - The changes below are applied to `PutPrescription` API. - Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation. - Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`. - Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`  ### 2025-01-06 v1.16.1 #### Added: - The changes below are applied to `PutPrescription` API. - Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`. - Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release. - Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`. - Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`. - Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`. - Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`. #### Changed: - The changes below are applied to `PutPrescription` API. - Deprecated `rxcui`. use `DrugDbCode` structure instead. - Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead. - Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead. - Deprecated `diagnosis`. Use `medicalDiagnosis` instead. - Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`. - Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.  ### 2024-12-13 v1.16.0 #### Changed: - Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances. - Updated `InsuranceList` max size to 27 from 50.  ### 2024-12-09 v1.15.10 #### Changed: - Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.  ### 2024-11-04 v1.15.9 #### Added: - New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request. #### Changed: - Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.  ### 2024-10-15 v1.15.8 #### Added: - Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model. - Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.  ### 2024-09-23 v1.15.7 #### Added: - Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues. - Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues. - Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.  ### 2024-09-03 v1.15.6 #### Added: - Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.   If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor. - The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field. - The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure. - A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException. - `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API. - Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.  ### 2024-08-06 v1.15.5 #### Added: - Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities. - Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions. - Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead. - Added documentation for enum fields for all APIs. - Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.  ### 2024-07-31 v1.15.4 #### Improved: - Updated the `directions` field to allow all strings for greater flexibility. - Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.  ### 2024-07-14 v1.15.3 #### Added: - All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs.

API version: 1.20.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amazonpharmacy

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PatientDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatientDetails{}

// PatientDetails Structure representing Patient details obtained via the PutPatient API request.
type PatientDetails struct {
	Name Name `json:"name"`
	// Patient's date of birth in the format 'yyyy-mm-dd'.
	DateOfBirth        string              `json:"dateOfBirth" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	SexAssignedAtBirth *SexAssignedAtBirth `json:"sexAssignedAtBirth,omitempty"`
	// List of contact information of the patient.
	ContactInfo []ContactInfo `json:"contactInfo"`
	// Gender of the patient.
	Gender *string `json:"gender,omitempty"`
	// Email address of the patient.
	Email *string `json:"email,omitempty" validate:"regexp=^([a-zA-Z0-9_\\\\.\\\\-\\\\+])+\\\\@(([a-zA-Z0-9\\\\-])+\\\\.)+([a-zA-Z0-9]{2,4})+$"`
	// Patient's consent to receive SMS notifications for signup and other communication.
	SmsConsent                      *bool                            `json:"smsConsent,omitempty"`
	AllergyDetails                  AllergyDetails                   `json:"allergyDetails"`
	ExistingMedicalConditionDetails ExistingMedicalConditionsDetails `json:"existingMedicalConditionDetails"`
	PregnancyStatus                 PregnancyStatus                  `json:"pregnancyStatus"`
	ExistingMedicationDetails       *ExistingMedicationDetails       `json:"existingMedicationDetails,omitempty"`
}

type _PatientDetails PatientDetails

// NewPatientDetails instantiates a new PatientDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatientDetails(name Name, dateOfBirth string, contactInfo []ContactInfo, allergyDetails AllergyDetails, existingMedicalConditionDetails ExistingMedicalConditionsDetails, pregnancyStatus PregnancyStatus) *PatientDetails {
	this := PatientDetails{}
	this.Name = name
	this.DateOfBirth = dateOfBirth
	this.ContactInfo = contactInfo
	this.AllergyDetails = allergyDetails
	this.ExistingMedicalConditionDetails = existingMedicalConditionDetails
	this.PregnancyStatus = pregnancyStatus
	return &this
}

// NewPatientDetailsWithDefaults instantiates a new PatientDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatientDetailsWithDefaults() *PatientDetails {
	this := PatientDetails{}
	return &this
}

// GetName returns the Name field value
func (o *PatientDetails) GetName() Name {
	if o == nil {
		var ret Name
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PatientDetails) SetName(v Name) {
	o.Name = v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *PatientDetails) GetDateOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetDateOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *PatientDetails) SetDateOfBirth(v string) {
	o.DateOfBirth = v
}

// GetSexAssignedAtBirth returns the SexAssignedAtBirth field value if set, zero value otherwise.
func (o *PatientDetails) GetSexAssignedAtBirth() SexAssignedAtBirth {
	if o == nil || IsNil(o.SexAssignedAtBirth) {
		var ret SexAssignedAtBirth
		return ret
	}
	return *o.SexAssignedAtBirth
}

// GetSexAssignedAtBirthOk returns a tuple with the SexAssignedAtBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetSexAssignedAtBirthOk() (*SexAssignedAtBirth, bool) {
	if o == nil || IsNil(o.SexAssignedAtBirth) {
		return nil, false
	}
	return o.SexAssignedAtBirth, true
}

// HasSexAssignedAtBirth returns a boolean if a field has been set.
func (o *PatientDetails) HasSexAssignedAtBirth() bool {
	if o != nil && !IsNil(o.SexAssignedAtBirth) {
		return true
	}

	return false
}

// SetSexAssignedAtBirth gets a reference to the given SexAssignedAtBirth and assigns it to the SexAssignedAtBirth field.
func (o *PatientDetails) SetSexAssignedAtBirth(v SexAssignedAtBirth) {
	o.SexAssignedAtBirth = &v
}

// GetContactInfo returns the ContactInfo field value
func (o *PatientDetails) GetContactInfo() []ContactInfo {
	if o == nil {
		var ret []ContactInfo
		return ret
	}

	return o.ContactInfo
}

// GetContactInfoOk returns a tuple with the ContactInfo field value
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetContactInfoOk() ([]ContactInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactInfo, true
}

// SetContactInfo sets field value
func (o *PatientDetails) SetContactInfo(v []ContactInfo) {
	o.ContactInfo = v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *PatientDetails) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *PatientDetails) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *PatientDetails) SetGender(v string) {
	o.Gender = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatientDetails) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatientDetails) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatientDetails) SetEmail(v string) {
	o.Email = &v
}

// GetSmsConsent returns the SmsConsent field value if set, zero value otherwise.
func (o *PatientDetails) GetSmsConsent() bool {
	if o == nil || IsNil(o.SmsConsent) {
		var ret bool
		return ret
	}
	return *o.SmsConsent
}

// GetSmsConsentOk returns a tuple with the SmsConsent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetSmsConsentOk() (*bool, bool) {
	if o == nil || IsNil(o.SmsConsent) {
		return nil, false
	}
	return o.SmsConsent, true
}

// HasSmsConsent returns a boolean if a field has been set.
func (o *PatientDetails) HasSmsConsent() bool {
	if o != nil && !IsNil(o.SmsConsent) {
		return true
	}

	return false
}

// SetSmsConsent gets a reference to the given bool and assigns it to the SmsConsent field.
func (o *PatientDetails) SetSmsConsent(v bool) {
	o.SmsConsent = &v
}

// GetAllergyDetails returns the AllergyDetails field value
func (o *PatientDetails) GetAllergyDetails() AllergyDetails {
	if o == nil {
		var ret AllergyDetails
		return ret
	}

	return o.AllergyDetails
}

// GetAllergyDetailsOk returns a tuple with the AllergyDetails field value
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetAllergyDetailsOk() (*AllergyDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllergyDetails, true
}

// SetAllergyDetails sets field value
func (o *PatientDetails) SetAllergyDetails(v AllergyDetails) {
	o.AllergyDetails = v
}

// GetExistingMedicalConditionDetails returns the ExistingMedicalConditionDetails field value
func (o *PatientDetails) GetExistingMedicalConditionDetails() ExistingMedicalConditionsDetails {
	if o == nil {
		var ret ExistingMedicalConditionsDetails
		return ret
	}

	return o.ExistingMedicalConditionDetails
}

// GetExistingMedicalConditionDetailsOk returns a tuple with the ExistingMedicalConditionDetails field value
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetExistingMedicalConditionDetailsOk() (*ExistingMedicalConditionsDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExistingMedicalConditionDetails, true
}

// SetExistingMedicalConditionDetails sets field value
func (o *PatientDetails) SetExistingMedicalConditionDetails(v ExistingMedicalConditionsDetails) {
	o.ExistingMedicalConditionDetails = v
}

// GetPregnancyStatus returns the PregnancyStatus field value
func (o *PatientDetails) GetPregnancyStatus() PregnancyStatus {
	if o == nil {
		var ret PregnancyStatus
		return ret
	}

	return o.PregnancyStatus
}

// GetPregnancyStatusOk returns a tuple with the PregnancyStatus field value
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetPregnancyStatusOk() (*PregnancyStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PregnancyStatus, true
}

// SetPregnancyStatus sets field value
func (o *PatientDetails) SetPregnancyStatus(v PregnancyStatus) {
	o.PregnancyStatus = v
}

// GetExistingMedicationDetails returns the ExistingMedicationDetails field value if set, zero value otherwise.
func (o *PatientDetails) GetExistingMedicationDetails() ExistingMedicationDetails {
	if o == nil || IsNil(o.ExistingMedicationDetails) {
		var ret ExistingMedicationDetails
		return ret
	}
	return *o.ExistingMedicationDetails
}

// GetExistingMedicationDetailsOk returns a tuple with the ExistingMedicationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientDetails) GetExistingMedicationDetailsOk() (*ExistingMedicationDetails, bool) {
	if o == nil || IsNil(o.ExistingMedicationDetails) {
		return nil, false
	}
	return o.ExistingMedicationDetails, true
}

// HasExistingMedicationDetails returns a boolean if a field has been set.
func (o *PatientDetails) HasExistingMedicationDetails() bool {
	if o != nil && !IsNil(o.ExistingMedicationDetails) {
		return true
	}

	return false
}

// SetExistingMedicationDetails gets a reference to the given ExistingMedicationDetails and assigns it to the ExistingMedicationDetails field.
func (o *PatientDetails) SetExistingMedicationDetails(v ExistingMedicationDetails) {
	o.ExistingMedicationDetails = &v
}

func (o PatientDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatientDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["dateOfBirth"] = o.DateOfBirth
	if !IsNil(o.SexAssignedAtBirth) {
		toSerialize["sexAssignedAtBirth"] = o.SexAssignedAtBirth
	}
	toSerialize["contactInfo"] = o.ContactInfo
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.SmsConsent) {
		toSerialize["smsConsent"] = o.SmsConsent
	}
	toSerialize["allergyDetails"] = o.AllergyDetails
	toSerialize["existingMedicalConditionDetails"] = o.ExistingMedicalConditionDetails
	toSerialize["pregnancyStatus"] = o.PregnancyStatus
	if !IsNil(o.ExistingMedicationDetails) {
		toSerialize["existingMedicationDetails"] = o.ExistingMedicationDetails
	}
	return toSerialize, nil
}

func (o *PatientDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"dateOfBirth",
		"contactInfo",
		"allergyDetails",
		"existingMedicalConditionDetails",
		"pregnancyStatus",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPatientDetails := _PatientDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatientDetails)

	if err != nil {
		return err
	}

	*o = PatientDetails(varPatientDetails)

	return err
}

type NullablePatientDetails struct {
	value *PatientDetails
	isSet bool
}

func (v NullablePatientDetails) Get() *PatientDetails {
	return v.value
}

func (v *NullablePatientDetails) Set(val *PatientDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePatientDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePatientDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatientDetails(val *PatientDetails) *NullablePatientDetails {
	return &NullablePatientDetails{value: val, isSet: true}
}

func (v NullablePatientDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatientDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
