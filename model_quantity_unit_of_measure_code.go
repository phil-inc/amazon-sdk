/*
Amazon Pharmacy APIs for ingesting Prescriptions, Patients, Orders and Payment Instruments

Amazon Pharmacy APIs provide the following functionalities for external partners:  - Patient creation, update, and view - Prescription creation, update, and view - Order creation, update, and view - Payment instrument creation and update - Prescription transfer request creation  ## Changelog for the Amazon Pharmacy APIs (from version 1.15.3): ### 2025-03-26 v1.16.4 #### Added: - Added prescriber-direct prescription example to the `PutPrescription` API documentation.  #### Changed: - Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.  ### 2025-03-18 v1.16.3 #### Changed: - Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema. - For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation. - For prescriber-originated prescriptions: This structure is not required and should be omitted. - Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.  ### 2025-01-06 v1.16.2 #### Added: - The changes below are applied to `PutPrescription` API. - Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release. - Added a new required field `ndc` in `PreviousDispensedMedication`. - Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release. - The changes below are applied to `PutOrder` API. - Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.  #### Changed: - The changes below are applied to `PutPrescription` API. - Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation. - Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`. - Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`  ### 2025-01-06 v1.16.1 #### Added: - The changes below are applied to `PutPrescription` API. - Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`. - Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release. - Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`. - Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`. - Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`. - Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`. #### Changed: - The changes below are applied to `PutPrescription` API. - Deprecated `rxcui`. use `DrugDbCode` structure instead. - Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead. - Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead. - Deprecated `diagnosis`. Use `medicalDiagnosis` instead. - Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`. - Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.  ### 2024-12-13 v1.16.0 #### Changed: - Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances. - Updated `InsuranceList` max size to 27 from 50.  ### 2024-12-09 v1.15.10 #### Changed: - Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.  ### 2024-11-04 v1.15.9 #### Added: - New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request. #### Changed: - Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.  ### 2024-10-15 v1.15.8 #### Added: - Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model. - Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.  ### 2024-09-23 v1.15.7 #### Added: - Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues. - Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues. - Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.  ### 2024-09-03 v1.15.6 #### Added: - Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.   If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor. - The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field. - The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure. - A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException. - `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API. - Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.  ### 2024-08-06 v1.15.5 #### Added: - Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities. - Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions. - Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead. - Added documentation for enum fields for all APIs. - Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.  ### 2024-07-31 v1.15.4 #### Improved: - Updated the `directions` field to allow all strings for greater flexibility. - Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.  ### 2024-07-14 v1.15.3 #### Added: - All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs. 

API version: 1.16.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amazonpharmacy

import (
	"encoding/json"
	"fmt"
)

// QuantityUnitOfMeasureCode Quantity Unit of Measure Code (Reference: https://evs.nci.nih.gov/ftp1/NCPDP/NCPDP.txt)  `C62412` - A dosing unit equal to the amount of active ingredient(s) contained in a single applicator.  `C54564` - A dosing unit equal to the amount of active ingredient(s) contained in a blister.  `C64696` - A dosing unit equal to the amount of active ingredient(s) contained in a caplet.  `C48480` - A dosing unit equal to the amount of active ingredient(s) contained in a capsule.  `C64933` - Used to refer to every member of a group of people or things, considered individually.  `C53499` - A dosing unit equal to the amount of active ingredient(s) contained in a film.  `C48155` - The metric unit of mass equal to one thousandth of a kilogram. One gram equals approximately 15.432 grains or 0.035 273 966 ounce.  `C69124` - A dosing unit equal to the amount of active ingredient(s) contained in a gum.  `C48499` - A dosing unit equal to the amount of active ingredient(s) contained in an implant.  `C62276` - A dosing unit equal to the amount of active ingredient(s) contained in an insert.  `C48504` - A dosing unit equal to the amount of active ingredient(s) contained in a kit.  `C120263` - A small, sharp, needle-like instrument that is used to puncture the skin.  `C48506` - A dosing unit equal to the amount of active ingredient(s) contained in a lozenge.  `C28254` - A unit of volume equal to one millionth (10E-6) of a cubic meter, one thousandth of a liter, one cubic centimeter, or 0.061023 7 cubic inch. A cubic centimeter is the CGS unit of volume.  `C48521` - A dosing unit equal to the amount of active ingredient(s) contained in a packet.  `C65032` - A dosing unit equal to the amount of active ingredient(s) contained in a pad.  `C48524` - A dosing unit equal to the amount of active ingredient(s) contained in a patch.  `C120216` - A single use, hollow needle embedded in a plastic hub, which is then attached to a preloaded syringe (pen) to facilitate the delivery of medication.  `C62609` - A dosing unit equal to the amount of active ingredient(s) contained in a ring.  `C53502` - A dosing unit equal to the amount of active ingredient(s) contained in a sponge.  `C53503` - A dosing unit equal to the amount of active ingredient(s) contained in a stick.  `C48538` - A dosing unit equal to the amount of active ingredient(s) contained in a strip.  `C48539` - A dosing unit equal to the amount of active ingredient(s) contained in a suppository.  `C53504` - A dosing unit equal to the amount of active ingredient(s) contained in a swab.  `C48542` - A dosing unit equal to the amount of active ingredient(s) contained in a tablet.  `C48548` - A dosing unit equal to the amount of active ingredient(s) contained in a troche.  `C38046` - Not stated explicitly or in detail.  `C48552` - A dosing unit equal to the amount of active ingredient(s) contained in a wafer.
type QuantityUnitOfMeasureCode string

// List of QuantityUnitOfMeasureCode
const (
	C62412 QuantityUnitOfMeasureCode = "C62412"
	C54564 QuantityUnitOfMeasureCode = "C54564"
	C64696 QuantityUnitOfMeasureCode = "C64696"
	C48480 QuantityUnitOfMeasureCode = "C48480"
	C64933 QuantityUnitOfMeasureCode = "C64933"
	C53499 QuantityUnitOfMeasureCode = "C53499"
	C48155_1 QuantityUnitOfMeasureCode = "C48155"
	C69124 QuantityUnitOfMeasureCode = "C69124"
	C48499 QuantityUnitOfMeasureCode = "C48499"
	C62276 QuantityUnitOfMeasureCode = "C62276"
	C48504 QuantityUnitOfMeasureCode = "C48504"
	C120263 QuantityUnitOfMeasureCode = "C120263"
	C48506 QuantityUnitOfMeasureCode = "C48506"
	C28254_1 QuantityUnitOfMeasureCode = "C28254"
	C48521 QuantityUnitOfMeasureCode = "C48521"
	C65032 QuantityUnitOfMeasureCode = "C65032"
	C48524 QuantityUnitOfMeasureCode = "C48524"
	C120216 QuantityUnitOfMeasureCode = "C120216"
	C62609 QuantityUnitOfMeasureCode = "C62609"
	C53502 QuantityUnitOfMeasureCode = "C53502"
	C53503 QuantityUnitOfMeasureCode = "C53503"
	C48538 QuantityUnitOfMeasureCode = "C48538"
	C48539 QuantityUnitOfMeasureCode = "C48539"
	C53504 QuantityUnitOfMeasureCode = "C53504"
	C48542 QuantityUnitOfMeasureCode = "C48542"
	C48548 QuantityUnitOfMeasureCode = "C48548"
	C38046_1 QuantityUnitOfMeasureCode = "C38046"
	C48552 QuantityUnitOfMeasureCode = "C48552"
)

// All allowed values of QuantityUnitOfMeasureCode enum
var AllowedQuantityUnitOfMeasureCodeEnumValues = []QuantityUnitOfMeasureCode{
	"C62412",
	"C54564",
	"C64696",
	"C48480",
	"C64933",
	"C53499",
	"C48155",
	"C69124",
	"C48499",
	"C62276",
	"C48504",
	"C120263",
	"C48506",
	"C28254",
	"C48521",
	"C65032",
	"C48524",
	"C120216",
	"C62609",
	"C53502",
	"C53503",
	"C48538",
	"C48539",
	"C53504",
	"C48542",
	"C48548",
	"C38046",
	"C48552",
}

func (v *QuantityUnitOfMeasureCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QuantityUnitOfMeasureCode(value)
	for _, existing := range AllowedQuantityUnitOfMeasureCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QuantityUnitOfMeasureCode", value)
}

// NewQuantityUnitOfMeasureCodeFromValue returns a pointer to a valid QuantityUnitOfMeasureCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuantityUnitOfMeasureCodeFromValue(v string) (*QuantityUnitOfMeasureCode, error) {
	ev := QuantityUnitOfMeasureCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuantityUnitOfMeasureCode: valid values are %v", v, AllowedQuantityUnitOfMeasureCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuantityUnitOfMeasureCode) IsValid() bool {
	for _, existing := range AllowedQuantityUnitOfMeasureCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuantityUnitOfMeasureCode value
func (v QuantityUnitOfMeasureCode) Ptr() *QuantityUnitOfMeasureCode {
	return &v
}

type NullableQuantityUnitOfMeasureCode struct {
	value *QuantityUnitOfMeasureCode
	isSet bool
}

func (v NullableQuantityUnitOfMeasureCode) Get() *QuantityUnitOfMeasureCode {
	return v.value
}

func (v *NullableQuantityUnitOfMeasureCode) Set(val *QuantityUnitOfMeasureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableQuantityUnitOfMeasureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableQuantityUnitOfMeasureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuantityUnitOfMeasureCode(val *QuantityUnitOfMeasureCode) *NullableQuantityUnitOfMeasureCode {
	return &NullableQuantityUnitOfMeasureCode{value: val, isSet: true}
}

func (v NullableQuantityUnitOfMeasureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuantityUnitOfMeasureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

