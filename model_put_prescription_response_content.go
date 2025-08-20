/*
Amazon Pharmacy APIs for ingesting Prescriptions, Patients, Orders and Payment Instruments

Amazon Pharmacy APIs provide the following functionalities for external partners:  - Patient creation, update, and view - Prescription creation, update, and view - Order creation, update, and view - Payment instrument creation and update - Prescription transfer request creation  ## Changelog for the Amazon Pharmacy APIs (from version 1.15.3):  ### 2025-08-15 v1.23.1 #### Added: - Added `DUR_REVIEW_IN_PROGRESS` to `OrderStatusReasonCode`. - This is a new status reason code to be returned for ORDER_FAILED and indicates the order is under DUR (Drug Utilization Review). Please retry placing the order using PutOrder API in 2-4 hours  ### 2025-08-05 v1.23.0 #### Changed: - Updated validation rules for prescriber partners:   - Allergies: Allow NOT_AVAILABLE in patient data when no allergy data is provided by the prescriber   - Medical conditions: Must always be set as NOT_AVAILABLE  ### 2025-08-01 v1.22.0 #### Changed: - Updated `MedicationPrescribed` structure in `PutPrescription API`:   - Modified `writtenDate`, `effectiveDate`, and `expirationDate` fields to accept both UTC DateTime (YYYY-MM-DDTHH:MM:SSZ) and simple Date (YYYY-MM-DD) formats   - Added recommendations to use DateTime over simple Date format - Updated API examples to demonstrate both supported formats  ### 2025-07-16 v1.21.0 #### Added: - Enhanced diagnosis code support in PutPrescription API:   - New optional `codeWithQualifier` field in `DiagnosisCode` structure to support additional code systems   - Made `icd10Code` field optional for prescriber partners when `codeWithQualifier` is provided - Added example payloads demonstrating usage of new fields  #### Changed: - Updated `Allergy` structure to make the `code` and `codeType` fields optional in `PutPatient API`, allowing prescriber partners to omit these fields when calling the PutPatient API - Updated documentation to reflect new code system options and field requirements - Improved documentation on `Name` field by making patient-specific suffix checks explicit  ### 2025-07-14 v1.20.1 #### Changed: - Updated the regex pattern for `patientDetails.name`, `prescriber.prescriberName`, `pharmacist.pharmacistName` field to accept special characters like `-`, `'`, `.` for `firstName`, `middleName`, `lastName` subfields.  ### 2025-07-07 v1.20.0 #### Added: - Added `INVALID_DELIVERY_ADDRESS` to `OrderStatusReasonCode`.  #### Changed: - Updated `MedicationPrescribed` structure to make the `daysSupply` field optional in `PutPrescription API`, allowing prescriber partners to omit this field when calling the PutPrescription API  ### 2025-07-01 v1.19.0 #### Added: - Asynchronous notifications for orders - Added documentation in `PutOrder` callback with `EventNotification` schema which signifies notification structure  #### Notes: - Order notifications are available for below order statuses for orders created through `PutOrder` API   - ORDER_CREATED   - ORDER_PROCESSING   - ORDER_ON_HOLD   - ORDER_SHIPPED   - ORDER_CANCELLED   - ORDER_FAILED  ### 2025-06-19 v1.18.2 #### Changed: - Updated `PatientDetails` structure to make the `sexAssignedAtBirth` field optional in `PutPatient API`, allowing partners to omit this field when calling the PutPatient API; however, this field will continue to be mandatory for hub partners who manage orders on behalf of customers  ### 2025-06-10 v1.18.1 #### Changed: - Removed mandatory requirement for the deprecated `unitOfMeasureCode` field in `Quantity`. - Updated documentation for quantity and refill fields. - Added comprehensive examples covering various prescription transfer scenarios. - Updated documentation to reflect optional status of `refillsTransferred` field.  ### 2025-06-03 v1.18.0 #### Added: - Added optional field `Observation` in PutPrescriptionRequest to support vital signs data. - Added new structures:   - `Observation` with measurement list and notes   - `Measurement` with vital sign, value, unit, and date fields   - `VitalSign` enum supporting HEIGHT and WEIGHT measurements - Updated `PutPrescription` API examples to demonstrate usage of Observation structure with height and weight measurements.   #### Changed:  - Updated `PatientDetails` structure to make the `smsConsent` field optional, allowing partners to omit this field when calling the PutPatient API. When not provided, `smsConsent` defaults to `false`.  ### 2025-03-24 v1.17.0 #### Added: - Asynchronous notifications for prescriptions. - Added documentation in `PutPrescription` callback with `EventNotification` schema which signifies notification structure  #### Changed: - `Insurance` now supports `cardholderId`. - `Insurance` and `OrderInsurance` field `encryptedCardholderId` is now deprecated. Please use `cardholderId` instead for `PutPatient` and `PutOrder` calls.  ### 2025-03-26 v1.16.4 #### Added: - Added prescriber-direct prescription example to the `PutPrescription` API documentation.  #### Changed: - Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.  ### 2025-03-18 v1.16.3 #### Changed: - Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema. - For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation. - For prescriber-originated prescriptions: This structure is not required and should be omitted. - Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.   ### 2025-01-06 v1.16.2 #### Added: - The changes below are applied to `PutPrescription` API. - Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release. - Added a new required field `ndc` in `PreviousDispensedMedication`. - Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release. - The changes below are applied to `PutOrder` API. - Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.  #### Changed: - The changes below are applied to `PutPrescription` API. - Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation. - Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`. - Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`  ### 2025-01-06 v1.16.1 #### Added: - The changes below are applied to `PutPrescription` API. - Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`. - Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release. - Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`. - Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`. - Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`. - Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`. #### Changed: - The changes below are applied to `PutPrescription` API. - Deprecated `rxcui`. use `DrugDbCode` structure instead. - Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead. - Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead. - Deprecated `diagnosis`. Use `medicalDiagnosis` instead. - Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`. - Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.  ### 2024-12-13 v1.16.0 #### Changed: - Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances. - Updated `InsuranceList` max size to 27 from 50.  ### 2024-12-09 v1.15.10 #### Changed: - Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.  ### 2024-11-04 v1.15.9 #### Added: - New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request. #### Changed: - Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.  ### 2024-10-15 v1.15.8 #### Added: - Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model. - Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.  ### 2024-09-23 v1.15.7 #### Added: - Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues. - Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues. - Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.  ### 2024-09-03 v1.15.6 #### Added: - Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.   If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor. - The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field. - The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure. - A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException. - `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API. - Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.  ### 2024-08-06 v1.15.5 #### Added: - Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities. - Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions. - Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead. - Added documentation for enum fields for all APIs. - Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.  ### 2024-07-31 v1.15.4 #### Improved: - Updated the `directions` field to allow all strings for greater flexibility. - Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.  ### 2024-07-14 v1.15.3 #### Added: - All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs.

API version: 1.23.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amazonpharmacy

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PutPrescriptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutPrescriptionResponseContent{}

// PutPrescriptionResponseContent Structure representing the response for adding a prescription.
type PutPrescriptionResponseContent struct {
	ReceivingPharmacy Pharmacy `json:"receivingPharmacy"`
	// The partner Prescription ID of the added prescription.
	PartnerPrescriptionId string `json:"partnerPrescriptionId" validate:"regexp=^[0-9a-zA-Z-]{1,40}$"`
}

type _PutPrescriptionResponseContent PutPrescriptionResponseContent

// NewPutPrescriptionResponseContent instantiates a new PutPrescriptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutPrescriptionResponseContent(receivingPharmacy Pharmacy, partnerPrescriptionId string) *PutPrescriptionResponseContent {
	this := PutPrescriptionResponseContent{}
	this.ReceivingPharmacy = receivingPharmacy
	this.PartnerPrescriptionId = partnerPrescriptionId
	return &this
}

// NewPutPrescriptionResponseContentWithDefaults instantiates a new PutPrescriptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutPrescriptionResponseContentWithDefaults() *PutPrescriptionResponseContent {
	this := PutPrescriptionResponseContent{}
	return &this
}

// GetReceivingPharmacy returns the ReceivingPharmacy field value
func (o *PutPrescriptionResponseContent) GetReceivingPharmacy() Pharmacy {
	if o == nil {
		var ret Pharmacy
		return ret
	}

	return o.ReceivingPharmacy
}

// GetReceivingPharmacyOk returns a tuple with the ReceivingPharmacy field value
// and a boolean to check if the value has been set.
func (o *PutPrescriptionResponseContent) GetReceivingPharmacyOk() (*Pharmacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivingPharmacy, true
}

// SetReceivingPharmacy sets field value
func (o *PutPrescriptionResponseContent) SetReceivingPharmacy(v Pharmacy) {
	o.ReceivingPharmacy = v
}

// GetPartnerPrescriptionId returns the PartnerPrescriptionId field value
func (o *PutPrescriptionResponseContent) GetPartnerPrescriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerPrescriptionId
}

// GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field value
// and a boolean to check if the value has been set.
func (o *PutPrescriptionResponseContent) GetPartnerPrescriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerPrescriptionId, true
}

// SetPartnerPrescriptionId sets field value
func (o *PutPrescriptionResponseContent) SetPartnerPrescriptionId(v string) {
	o.PartnerPrescriptionId = v
}

func (o PutPrescriptionResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutPrescriptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["receivingPharmacy"] = o.ReceivingPharmacy
	toSerialize["partnerPrescriptionId"] = o.PartnerPrescriptionId
	return toSerialize, nil
}

func (o *PutPrescriptionResponseContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"receivingPharmacy",
		"partnerPrescriptionId",
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

	varPutPrescriptionResponseContent := _PutPrescriptionResponseContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPutPrescriptionResponseContent)

	if err != nil {
		return err
	}

	*o = PutPrescriptionResponseContent(varPutPrescriptionResponseContent)

	return err
}

type NullablePutPrescriptionResponseContent struct {
	value *PutPrescriptionResponseContent
	isSet bool
}

func (v NullablePutPrescriptionResponseContent) Get() *PutPrescriptionResponseContent {
	return v.value
}

func (v *NullablePutPrescriptionResponseContent) Set(val *PutPrescriptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePutPrescriptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePutPrescriptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutPrescriptionResponseContent(val *PutPrescriptionResponseContent) *NullablePutPrescriptionResponseContent {
	return &NullablePutPrescriptionResponseContent{value: val, isSet: true}
}

func (v NullablePutPrescriptionResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutPrescriptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
