/*
Amazon Pharmacy APIs for ingesting Prescriptions, Patients, Orders and Payment Instruments

Amazon Pharmacy APIs provide the following functionalities for external partners:  - Patient creation, update, and view - Prescription creation, update, and view - Order creation, update, and view - Payment instrument creation and update - Prescription transfer request creation  ## Changelog for the Amazon Pharmacy APIs (from version 1.15.3):  ### 2025-08-15 v1.23.1 #### Added: - Added `DUR_REVIEW_IN_PROGRESS` to `OrderStatusReasonCode`. - This is a new status reason code to be returned for ORDER_FAILED and indicates the order is under DUR (Drug Utilization Review). Please retry placing the order using PutOrder API in 2-4 hours  ### 2025-08-05 v1.23.0 #### Changed: - Updated validation rules for prescriber partners:   - Allergies: Allow NOT_AVAILABLE in patient data when no allergy data is provided by the prescriber   - Medical conditions: Must always be set as NOT_AVAILABLE  ### 2025-08-01 v1.22.0 #### Changed: - Updated `MedicationPrescribed` structure in `PutPrescription API`:   - Modified `writtenDate`, `effectiveDate`, and `expirationDate` fields to accept both UTC DateTime (YYYY-MM-DDTHH:MM:SSZ) and simple Date (YYYY-MM-DD) formats   - Added recommendations to use DateTime over simple Date format - Updated API examples to demonstrate both supported formats  ### 2025-07-16 v1.21.0 #### Added: - Enhanced diagnosis code support in PutPrescription API:   - New optional `codeWithQualifier` field in `DiagnosisCode` structure to support additional code systems   - Made `icd10Code` field optional for prescriber partners when `codeWithQualifier` is provided - Added example payloads demonstrating usage of new fields  #### Changed: - Updated `Allergy` structure to make the `code` and `codeType` fields optional in `PutPatient API`, allowing prescriber partners to omit these fields when calling the PutPatient API - Updated documentation to reflect new code system options and field requirements - Improved documentation on `Name` field by making patient-specific suffix checks explicit  ### 2025-07-14 v1.20.1 #### Changed: - Updated the regex pattern for `patientDetails.name`, `prescriber.prescriberName`, `pharmacist.pharmacistName` field to accept special characters like `-`, `'`, `.` for `firstName`, `middleName`, `lastName` subfields.  ### 2025-07-07 v1.20.0 #### Added: - Added `INVALID_DELIVERY_ADDRESS` to `OrderStatusReasonCode`.  #### Changed: - Updated `MedicationPrescribed` structure to make the `daysSupply` field optional in `PutPrescription API`, allowing prescriber partners to omit this field when calling the PutPrescription API  ### 2025-07-01 v1.19.0 #### Added: - Asynchronous notifications for orders - Added documentation in `PutOrder` callback with `EventNotification` schema which signifies notification structure  #### Notes: - Order notifications are available for below order statuses for orders created through `PutOrder` API   - ORDER_CREATED   - ORDER_PROCESSING   - ORDER_ON_HOLD   - ORDER_SHIPPED   - ORDER_CANCELLED   - ORDER_FAILED  ### 2025-06-19 v1.18.2 #### Changed: - Updated `PatientDetails` structure to make the `sexAssignedAtBirth` field optional in `PutPatient API`, allowing partners to omit this field when calling the PutPatient API; however, this field will continue to be mandatory for hub partners who manage orders on behalf of customers  ### 2025-06-10 v1.18.1 #### Changed: - Removed mandatory requirement for the deprecated `unitOfMeasureCode` field in `Quantity`. - Updated documentation for quantity and refill fields. - Added comprehensive examples covering various prescription transfer scenarios. - Updated documentation to reflect optional status of `refillsTransferred` field.  ### 2025-06-03 v1.18.0 #### Added: - Added optional field `Observation` in PutPrescriptionRequest to support vital signs data. - Added new structures:   - `Observation` with measurement list and notes   - `Measurement` with vital sign, value, unit, and date fields   - `VitalSign` enum supporting HEIGHT and WEIGHT measurements - Updated `PutPrescription` API examples to demonstrate usage of Observation structure with height and weight measurements.   #### Changed:  - Updated `PatientDetails` structure to make the `smsConsent` field optional, allowing partners to omit this field when calling the PutPatient API. When not provided, `smsConsent` defaults to `false`.  ### 2025-03-24 v1.17.0 #### Added: - Asynchronous notifications for prescriptions. - Added documentation in `PutPrescription` callback with `EventNotification` schema which signifies notification structure  #### Changed: - `Insurance` now supports `cardholderId`. - `Insurance` and `OrderInsurance` field `encryptedCardholderId` is now deprecated. Please use `cardholderId` instead for `PutPatient` and `PutOrder` calls.  ### 2025-03-26 v1.16.4 #### Added: - Added prescriber-direct prescription example to the `PutPrescription` API documentation.  #### Changed: - Updated `PutPrescription` API examples to include NCI Codes for quantity UnitOfMeasureCode fields.  ### 2025-03-18 v1.16.3 #### Changed: - Modified the `PrescriptionTransferInDetails` field to be optional in the `PutPrescription` API schema. - For pharmacy-to-pharmacy transfers: This structure remains mandatory but is now enforced through server-side validation. - For prescriber-originated prescriptions: This structure is not required and should be omitted. - Modified `PutOrder` API to include examples for Insurance compliance exceptions and providing useful messaging when `InsuranceComplianceException` is returned.   ### 2025-01-06 v1.16.2 #### Added: - The changes below are applied to `PutPrescription` API. - Added new optional field `refillsRemaining` in `PrescriptionTransferred`. This field will be marked required in a future release. - Added a new required field `ndc` in `PreviousDispensedMedication`. - Added a new optional field `pharmacyRxNumber` in `PrescriptionTransferred`. This field will be marked as required in a future release. - The changes below are applied to `PutOrder` API. - Added `InsuranceComplianceException` for when setting OrderInsuranceDetails fails.  #### Changed: - The changes below are applied to `PutPrescription` API. - Renamed the `medicationDispensed` structure into `previousDispensedMedications` and updated documentation. - Updated deprecation date for `refillsTransferred`, `quantityTransferred`, `quantityRemaining`, `rxcui`, `unitOfMeasureCode`, `formCode` and  `diagnosis`. - Fixed documentation issue on enums `QuantityCodeListQualifier`, `ClinicalInformationQualifier`, `QuantityUnitOfMeasureCode`, `DrugDbCodeQualifier`, `StrengthFormCode` and `StrengthUnitOfMeasureCode`  ### 2025-01-06 v1.16.1 #### Added: - The changes below are applied to `PutPrescription` API. - Added new Enum `DrugDbCodeQualifier`, `QuantityCodeListQualifier`, `QuantityUnitOfMeasureCode`, `StrengthFormCode`, `StrengthUnitOfMeasureCode`. - Added new optional field `quantityRemaining` in `PrescriptionTransferred`. This field will be marked required in future release. - Added new structures `DrugDbCode`, `MedicationDispensed`, `MedicalDiagnosis` and `DiagnosisCode`. - Added new optional field `supervisingPhysician` in `PutPrescriptionRequest`. - Added new optional fields `labelerName`, `drugDbCode`, `note` and `medicalDiagnosis` in `MedicationPrescribed`. - Added new optional field `medicationDispensed` in `PrescriptionTransferInDetails`. #### Changed: - The changes below are applied to `PutPrescription` API. - Deprecated `rxcui`. use `DrugDbCode` structure instead. - Deprecated `unitOfMeasureCode` in `Quantity`. Use `quantityUnitOfMeasureCode` enum instead. - Deprecated `formCode` and `unitOfMeasureCode` in `Strength`. Use `strengthFormCode` and `strengthUnitOfMeasureCode` enums instead. - Deprecated `diagnosis`. Use `medicalDiagnosis` instead. - Updated documentation on field mandate for `refillsTransferred` and `quantityTransferred` in `PrescriptionTransferred`. - Updated `Pharmacy` sub-fields `ncpdpId`, `pharmacyName`, `pharmacistName`, `pharmacyAddress`, `primaryTelephone` to required to validate `sendingPharmacy` and `receivingPharmacy` in prescriptions.  ### 2024-12-13 v1.16.0 #### Changed: - Updated `BaseInsurance` field `personCode` to have regex validation `^[0-9]*$`: only numbers allowed. This applies to `PutPatient` and `PutOrder` insurances. - Updated `InsuranceList` max size to 27 from 50.  ### 2024-12-09 v1.15.10 #### Changed: - Updated `ICD10` pattern for existing medical condition and diagnosis for `PutPatient` and `PutPrescription` request model respectively to allow special ICD-10 codes.  ### 2024-11-04 v1.15.9 #### Added: - New optional fields `payerType`, `payerName`, `stateCode`, `startDate`, `expiryDate` to capture more information about the Insurance plan provided in `PutPatient` API request. #### Changed: - Updated `PutOrder` API to introduce capability to accept unencrypted primary & secondary insurances with the Order. This will be fully enabled in future releases.  ### 2024-10-15 v1.15.8 #### Added: - Updated `PutPrescription` API to make `prescriptionTransferInDetails` required structure in the `PutPrescriptionRequest` model. - Deprecated `PutPrescriptionRequest.sendingPharmacy`. Use `PrescriptionTransferInDetails.sendingPharmacy` structure instead.  ### 2024-09-23 v1.15.7 #### Added: - Added OrderStatusReasonCodes `COPAY_DISCREPANCY_DUE_TO_GOVERNMENT_INSURANCE_ON_FILE` and `COPAY_DISCREPANCY_DUE_TO_POTENTIAL_GOVERNMENT_INSURANCE_ON_FILE` to provide more information on copay failure reasons. This enhancement will help partners understand the reasons for order failures, making it easier to troubleshoot and resolve issues. - Added new `ClaimRejectionDetails` attribute to `AdditionalOrderDetails` in GetOrder API. This enhancement will help partners understand the reasons for order failures due to claim rejections, making it easier to troubleshoot and resolve issues. - Added new `PreconditionErrorCode` (`UNKNOWN_INSURANCE_PLAN_INPUT`) for unknown insurance plans in Amazon Pharmacy to provide more transparent feedback to partners during patient ingestion. This is exclusively for partners who place orders on behalf of customers.  ### 2024-09-03 v1.15.6 #### Added: - Updated examples for PutPatient to use allergiesDescriptor for allergyDetails, existingMedicalConditionsDescriptor for existingMedicalConditionDetails instead of noKnownAllergies and noKnownExistingMedicalConditions which are deprecated. Added existingMedicationsDescriptor for existingMedicationDetails, in locations where they were missing.   If you are still using noKnownAllergies and noKnownExistingMedicalConditions please plan to migrate to allergiesDescriptor and existingMedicalConditionsDescriptor. - The InvalidInputException structure has been enhanced with two new fields: `errorCode` and `errorFieldList`. These fields will provide additional information about the error, including the specific reason for the failure, the fields that caused the error, and a brief error message for each problematic field. - The PreconditionFailedException structure now includes the `errorFieldList` field. This field will contain a list of fields that caused the precondition failure. - A new enum called `InvalidInputErrorCode` has been introduced. This enum will encapsulate error codes related to invalid input scenarios, such as \"INVALID_BILLING_ADDRESS\". This enum will be included in InvalidInputException. - `PutPaymentInstrument` API will return InvalidInputException with \"INVALID_BILLING_ADDRESS\" errorCode when billing address fails validation with Amazon internal API. - Updated the API exception handling documentation with detailed descriptions and retry logic for common error status codes. This includes guidance on how to handle errors like `503`, `504`, `429`, and `403` for better integration and error management by external partners.  ### 2024-08-06 v1.15.5 #### Added: - Updated `PutPrescription` API to include new optional field `prescriptionTransferInDetails` in `PutPrescriptionRequest` model to capture details about the prescription transfer. This includes information if the prescription was previously dispensed or its a new prescription, the dates of the first and last fills, remaining refills and quantities. - Included examples in `PutPrescription` API for new prescription and previously dispensed prescriptions. - Deprecated `firstFillDate` and `lastFillDate` fields from `MedicationPrescribed`. This will be supported in `prescriptionTransferInDetails` structure instead. - Added documentation for enum fields for all APIs. - Included examples with expected exception in `PutPatient` API when the patient's insurance is invalid or not contracted with Amazon Pharmacy.  ### 2024-07-31 v1.15.4 #### Improved: - Updated the `directions` field to allow all strings for greater flexibility. - Enhanced `Address Line 1` and `Address Line 2` fields to allow all non-empty strings, providing more flexibility in address formatting.  ### 2024-07-14 v1.15.3 #### Added: - All APIs will accept `client-request-id` as header parameter. This can be set by Partners and it will be used to identify partner API requests. Please see updated request headers for all APIs.

API version: 1.23.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amazonpharmacy

import (
	"encoding/json"
	"fmt"
)

// StrengthUnitOfMeasureCode Strength Unit of Measure Code (Reference: https://evs.nci.nih.gov/ftp1/NCPDP/NCPDP.txt)  `C70518` - A unit of radioactivity defined as 3.7E-8 atomic disintegrations or other nuclear transformations per second or one quintillionth of curie (10E-18 curie).  `C42562` - A SI derived unit of activity of a radionuclide, equal to one nuclear disintegration or other nuclear transition from a particular energy state occurring in an amount of a radionuclide during one second-long time interval.  `C70515` - A unit of radioactivity defined as 3.7 E8 atomic disintegrations or other nuclear transformations per second. One centicurie is equal to 0.37 gigabecquerel.  `C48466` - A unit of radioactivity in the CGS system, defined as 3.7 E10 atomic disintegrations or other nuclear transformations per second. One curie is equal to 37 gigabecquerels.  `C25301` - The time for Earth to make a complete rotation on its axis; ordinarily divided into twenty-four hours, equal to 86 400 seconds. This also refers to a specific day.  `C70514` - A unit of radioactivity defined as 3.7 E9 atomic disintegrations or other nuclear transformations per second. One decicurie is equal to 3.7 gigabecquerels.  `C70517` - A unit of radioactivity defined as 3.7E-5 atomic disintegrations or other nuclear transformations per second or one quadrillionth of curie (10E-15 curie).  `C70513` - A unit of radioactivity equal to 10E9 nuclear disintegrations or other nuclear transformations per second, or to 10E9 becquerels.  `C48155` - The metric unit of mass equal to one thousandth of a kilogram. One gram equals approximately 15.432 grains or 0.035 273 966 ounce.  `C48579` - The unitage assigned by the WHO to International Biological Standards - substances, classed as biological according to the criteria provided by WHO Expert Committee on Biological Standardization (e.g. hormones, enzymes, and vaccines), to enable the results of biological and immunological assay procedures to be expressed in the same way throughout the world. The definition of an international unit is generally arbitrary and technical, and has to be officially approved by the International Conference for Unification of Formulae.  `C70511` - A unit of radioactivity equal to 10E3 nuclear disintegrations or other nuclear transformations per second, or to 10E3 becquerels.  `C28252` - A basic SI unit of mass. It is defined as the mass of an international prototype in the form of a platinum-iridium cylinder kept at Sevres in France. A kilogram is equal to 1,000 grams and 2.204 622 6 pounds.  `C48505` - The non-SI unit of volume accepted for use with the SI. One liter is equal to cubic decimeter, or one thousandth of cubic meter, or 1000 cubic centimeters, or approximately 61.023 744 cubic inches.  `C70512` - A unit of radioactivity equal to 10E6 nuclear disintegrations or other nuclear transformations per second, or to 10E6 becquerels.  `C42576` - A SI derived unit of mass concentration defined as the concentration of one kilogram of a substance per unit volume of the mixture equal to one cubic meter, or the concentration of one milligram of a substance per unit volume of the mixture equal to one milliliter, or one gram of a substance per one liter of the mixture. It is also a unit of mass density (volumic mass) defined as the density of substance which mass equal to one milligram occupies the volume one milliliter.  `C48507` - A unit of radioactivity equal to one millionth of a curie or 37 kilobecquerels, and corresponding to a radioactivity of 37 000 atomic disintegrations per second.  `C48152` - A metric unit of mass equal to one millionth of a gram or one thousandth of a milligram.  `C71205` - A unit of mass flow rate equivalent to the rate at which one millionth of a gram of matter crosses a given surface or is delivered to a given object or space over a period of time equal to twenty four hours. Microgram per 24 hours is also a dose administration rate unit equal to the rate at which one millionth of a gram of a product is administered per unit of time equal to twenty four hours.  `C91132` - The quantity of micrograms in a volume of fifteen milliliters of a substance.  `C67394` - A unit of mass flow rate equivalent to the rate at which one millionth of a gram of matter crosses a given surface or is delivered to a given object or space over a period of time equal to one hour.  `C64572` - A metric unit of mass concentration defined as the concentration of one gram of a substance per unit volume of the mixture equal to one cubic meter. The concept also refers to the metric unit of mass density (volumic mass) defined as the density of a substance which mass equal to one gram occupies the volume of one cubic meter.  `C91135` - The quantity of micrograms prescribed or delivered in three days.  `C48511` - A unit of radioactivity equal to one thousandth of a curie or 37 megabecquerels, and corresponding to a radioactivity of 37 millions of atomic disintegrations per second.  `C48512` - A unit of relative amount of a substance equal to one thousandth of an equivalent weight.  `C28253` - A metric unit of mass equal to one thousandth of a gram or 1000 micrograms. One milligram equals approximately 0.015 432 grain or 35.274 x 10E-6 ounce.  `C91131` - The quantity of milligrams in a volume of five milliliters of a substance.  `C28254` - A unit of volume equal to one millionth (10E-6) of a cubic meter, one thousandth of a liter, one cubic centimeter, or 0.061023 7 cubic inch. A cubic centimeter is the CGS unit of volume.  `C48513` - A unit of amount of substance equal to 0.001 mole.  `C67310` - A quantity equivalent to one million units (10E6 units).  `C67352` - A unit of radioactivity equal to one billionth of a curie or 37 becquerels, and corresponding to radioactivity of 37 atomic disintegrations per second.  `C25613` - A fraction or ratio with 100 understood as the denominator.  `C70516` - A unit of radioactivity defined as 0.037 atomic disintegrations or other nuclear transformations per second or one trillionth of curie (10E-12 curie). One picocurie is equal to 37 millibecquerels.  `C44278` - A single undivided thing occurring in the composition of something else.  `C38046` - Not stated explicitly or in detail.  `C70520` - A unit of radioactivity defined as 3.7E-14 atomic disintegrations or other nuclear transformations per second or one septillionth of curie (10E-24 curie).  `C70519` - A unit of radioactivity defined as 3.7E-11 atomic disintegrations or other nuclear transformations per second or one sextillionth of curie (10E-21 curie).
type StrengthUnitOfMeasureCode string

// List of StrengthUnitOfMeasureCode
const (
	C70518 StrengthUnitOfMeasureCode = "C70518"
	C42562 StrengthUnitOfMeasureCode = "C42562"
	C70515 StrengthUnitOfMeasureCode = "C70515"
	C48466 StrengthUnitOfMeasureCode = "C48466"
	C25301 StrengthUnitOfMeasureCode = "C25301"
	C70514 StrengthUnitOfMeasureCode = "C70514"
	C70517 StrengthUnitOfMeasureCode = "C70517"
	C70513 StrengthUnitOfMeasureCode = "C70513"
	C48155 StrengthUnitOfMeasureCode = "C48155"
	C48579 StrengthUnitOfMeasureCode = "C48579"
	C70511 StrengthUnitOfMeasureCode = "C70511"
	C28252 StrengthUnitOfMeasureCode = "C28252"
	C48505 StrengthUnitOfMeasureCode = "C48505"
	C70512 StrengthUnitOfMeasureCode = "C70512"
	C42576 StrengthUnitOfMeasureCode = "C42576"
	C48507 StrengthUnitOfMeasureCode = "C48507"
	C48152 StrengthUnitOfMeasureCode = "C48152"
	C71205 StrengthUnitOfMeasureCode = "C71205"
	C91132 StrengthUnitOfMeasureCode = "C91132"
	C67394 StrengthUnitOfMeasureCode = "C67394"
	C64572 StrengthUnitOfMeasureCode = "C64572"
	C91135 StrengthUnitOfMeasureCode = "C91135"
	C48511 StrengthUnitOfMeasureCode = "C48511"
	C48512 StrengthUnitOfMeasureCode = "C48512"
	C28253 StrengthUnitOfMeasureCode = "C28253"
	C91131 StrengthUnitOfMeasureCode = "C91131"
	C28254 StrengthUnitOfMeasureCode = "C28254"
	C48513 StrengthUnitOfMeasureCode = "C48513"
	C67310 StrengthUnitOfMeasureCode = "C67310"
	C67352 StrengthUnitOfMeasureCode = "C67352"
	C25613 StrengthUnitOfMeasureCode = "C25613"
	C70516 StrengthUnitOfMeasureCode = "C70516"
	C44278 StrengthUnitOfMeasureCode = "C44278"
	C38046 StrengthUnitOfMeasureCode = "C38046"
	C70520 StrengthUnitOfMeasureCode = "C70520"
	C70519 StrengthUnitOfMeasureCode = "C70519"
)

// All allowed values of StrengthUnitOfMeasureCode enum
var AllowedStrengthUnitOfMeasureCodeEnumValues = []StrengthUnitOfMeasureCode{
	"C70518",
	"C42562",
	"C70515",
	"C48466",
	"C25301",
	"C70514",
	"C70517",
	"C70513",
	"C48155",
	"C48579",
	"C70511",
	"C28252",
	"C48505",
	"C70512",
	"C42576",
	"C48507",
	"C48152",
	"C71205",
	"C91132",
	"C67394",
	"C64572",
	"C91135",
	"C48511",
	"C48512",
	"C28253",
	"C91131",
	"C28254",
	"C48513",
	"C67310",
	"C67352",
	"C25613",
	"C70516",
	"C44278",
	"C38046",
	"C70520",
	"C70519",
}

func (v *StrengthUnitOfMeasureCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StrengthUnitOfMeasureCode(value)
	for _, existing := range AllowedStrengthUnitOfMeasureCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StrengthUnitOfMeasureCode", value)
}

// NewStrengthUnitOfMeasureCodeFromValue returns a pointer to a valid StrengthUnitOfMeasureCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStrengthUnitOfMeasureCodeFromValue(v string) (*StrengthUnitOfMeasureCode, error) {
	ev := StrengthUnitOfMeasureCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StrengthUnitOfMeasureCode: valid values are %v", v, AllowedStrengthUnitOfMeasureCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StrengthUnitOfMeasureCode) IsValid() bool {
	for _, existing := range AllowedStrengthUnitOfMeasureCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StrengthUnitOfMeasureCode value
func (v StrengthUnitOfMeasureCode) Ptr() *StrengthUnitOfMeasureCode {
	return &v
}

type NullableStrengthUnitOfMeasureCode struct {
	value *StrengthUnitOfMeasureCode
	isSet bool
}

func (v NullableStrengthUnitOfMeasureCode) Get() *StrengthUnitOfMeasureCode {
	return v.value
}

func (v *NullableStrengthUnitOfMeasureCode) Set(val *StrengthUnitOfMeasureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableStrengthUnitOfMeasureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableStrengthUnitOfMeasureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrengthUnitOfMeasureCode(val *StrengthUnitOfMeasureCode) *NullableStrengthUnitOfMeasureCode {
	return &NullableStrengthUnitOfMeasureCode{value: val, isSet: true}
}

func (v NullableStrengthUnitOfMeasureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrengthUnitOfMeasureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
