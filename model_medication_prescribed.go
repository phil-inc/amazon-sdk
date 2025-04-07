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

// checks if the MedicationPrescribed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicationPrescribed{}

// MedicationPrescribed Structure representing the details of the medication prescribed.
type MedicationPrescribed struct {
	// Name or description of the medication being prescribed.
	DrugDescription string `json:"drugDescription" validate:"regexp=^\\\\s*\\\\S.*$"`
	// Drug labeler/manufacturer name.
	LabelerName *string `json:"labelerName,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
	Quantity Quantity `json:"quantity"`
	Strength *Strength `json:"strength,omitempty"`
	// Estimated number of days a prescription would last.
	DaysSupply float64 `json:"daysSupply"`
	// The date on which the prescriber wrote the prescription, quoted in the format 'YYYY-MM-DD'.
	WrittenDate string `json:"writtenDate" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// The number of refills allowed for this prescription.
	NumberOfRefills float64 `json:"numberOfRefills"`
	// Clinical directions meaningful to the patient as written by the prescriber, e.g., 'Take 1 tablet by mouth once weekly.'
	Directions string `json:"directions"`
	// The National Drug Code of the medication, without dashes or spaces.
	Ndc string `json:"ndc" validate:"regexp=^\\\\d{11}$"`
	// RxNorm concept unique identifier. This shape is deprecated: RxNorm codes are included in drugDbCode field. This field will no longer be supported starting from 2025-03-12.
	// Deprecated
	Rxcui *string `json:"rxcui,omitempty" validate:"regexp=^\\\\d{1,10}$"`
	DrugDbCode *DrugDbCode `json:"drugDbCode,omitempty"`
	// An indicator that the prescriber has instructed to dispense as written (DAW) code. This is typically used when a brand is chosen over a generic alternative based on the direction of the prescriber. 0 means no DAW is specified 1 means by prescriber 2 means a patient/user request 9 means other.
	DispenseAsWritten string `json:"dispenseAsWritten" validate:"regexp=^[0-9]$"`
	// The date on which the user is intended to start the prescription, quoted in the format 'YYYY-MM-DD'.
	EffectiveDate *string `json:"effectiveDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// The date on which the prescription expires, quoted in the format 'YYYY-MM-DD'.
	ExpirationDate *string `json:"expirationDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	DeaSchedule *DeaSchedule `json:"deaSchedule,omitempty"`
	// Medical diagnosis information for which the patient is prescribed. First item is the Primary Diagnosis, second is Secondary.  This shape is deprecated: This field will be replaced by `medicalDiagnosis`. This field will no longer be supported starting from 2025-03-12
	// Deprecated
	Diagnosis []Diagnosis `json:"diagnosis,omitempty"`
	MedicalDiagnosis *MedicalDiagnosis `json:"medicalDiagnosis,omitempty"`
	// The firstFillDate field is deprecated here and will be part of PrescriptionTransferred structure instead. The date on which the prescription was first filled by the pharmacy, in the format 'YYYY-MM-DD'. This date helps track the initial dispensing of the medication to the patient.  This shape is deprecated: This field is moved to the `PrescriptionTransferred` data structure. It will no longer be available here.
	// Deprecated
	FirstFillDate *string `json:"firstFillDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// The lastFillDate field is deprecated here and will be part of PrescriptionTransferred structure instead. The date on which the prescription was most recently filled by the pharmacy, in the format 'YYYY-MM-DD'. This date helps track the latest dispensing of the medication and is useful for monitoring adherence and refill patterns.  This shape is deprecated: This field is moved to the `PrescriptionTransferred` data structure. It will no longer be available here.
	// Deprecated
	LastFillDate *string `json:"lastFillDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	// Additional information on the Patient or Prescription provided by the Prescriber.
	Note *string `json:"note,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
}

type _MedicationPrescribed MedicationPrescribed

// NewMedicationPrescribed instantiates a new MedicationPrescribed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicationPrescribed(drugDescription string, quantity Quantity, daysSupply float64, writtenDate string, numberOfRefills float64, directions string, ndc string, dispenseAsWritten string) *MedicationPrescribed {
	this := MedicationPrescribed{}
	this.DrugDescription = drugDescription
	this.Quantity = quantity
	this.DaysSupply = daysSupply
	this.WrittenDate = writtenDate
	this.NumberOfRefills = numberOfRefills
	this.Directions = directions
	this.Ndc = ndc
	this.DispenseAsWritten = dispenseAsWritten
	return &this
}

// NewMedicationPrescribedWithDefaults instantiates a new MedicationPrescribed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicationPrescribedWithDefaults() *MedicationPrescribed {
	this := MedicationPrescribed{}
	return &this
}

// GetDrugDescription returns the DrugDescription field value
func (o *MedicationPrescribed) GetDrugDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DrugDescription
}

// GetDrugDescriptionOk returns a tuple with the DrugDescription field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetDrugDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrugDescription, true
}

// SetDrugDescription sets field value
func (o *MedicationPrescribed) SetDrugDescription(v string) {
	o.DrugDescription = v
}

// GetLabelerName returns the LabelerName field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetLabelerName() string {
	if o == nil || IsNil(o.LabelerName) {
		var ret string
		return ret
	}
	return *o.LabelerName
}

// GetLabelerNameOk returns a tuple with the LabelerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetLabelerNameOk() (*string, bool) {
	if o == nil || IsNil(o.LabelerName) {
		return nil, false
	}
	return o.LabelerName, true
}

// HasLabelerName returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasLabelerName() bool {
	if o != nil && !IsNil(o.LabelerName) {
		return true
	}

	return false
}

// SetLabelerName gets a reference to the given string and assigns it to the LabelerName field.
func (o *MedicationPrescribed) SetLabelerName(v string) {
	o.LabelerName = &v
}

// GetQuantity returns the Quantity field value
func (o *MedicationPrescribed) GetQuantity() Quantity {
	if o == nil {
		var ret Quantity
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetQuantityOk() (*Quantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *MedicationPrescribed) SetQuantity(v Quantity) {
	o.Quantity = v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetStrength() Strength {
	if o == nil || IsNil(o.Strength) {
		var ret Strength
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetStrengthOk() (*Strength, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given Strength and assigns it to the Strength field.
func (o *MedicationPrescribed) SetStrength(v Strength) {
	o.Strength = &v
}

// GetDaysSupply returns the DaysSupply field value
func (o *MedicationPrescribed) GetDaysSupply() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DaysSupply
}

// GetDaysSupplyOk returns a tuple with the DaysSupply field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetDaysSupplyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysSupply, true
}

// SetDaysSupply sets field value
func (o *MedicationPrescribed) SetDaysSupply(v float64) {
	o.DaysSupply = v
}

// GetWrittenDate returns the WrittenDate field value
func (o *MedicationPrescribed) GetWrittenDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WrittenDate
}

// GetWrittenDateOk returns a tuple with the WrittenDate field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetWrittenDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WrittenDate, true
}

// SetWrittenDate sets field value
func (o *MedicationPrescribed) SetWrittenDate(v string) {
	o.WrittenDate = v
}

// GetNumberOfRefills returns the NumberOfRefills field value
func (o *MedicationPrescribed) GetNumberOfRefills() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NumberOfRefills
}

// GetNumberOfRefillsOk returns a tuple with the NumberOfRefills field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetNumberOfRefillsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfRefills, true
}

// SetNumberOfRefills sets field value
func (o *MedicationPrescribed) SetNumberOfRefills(v float64) {
	o.NumberOfRefills = v
}

// GetDirections returns the Directions field value
func (o *MedicationPrescribed) GetDirections() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Directions
}

// GetDirectionsOk returns a tuple with the Directions field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetDirectionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Directions, true
}

// SetDirections sets field value
func (o *MedicationPrescribed) SetDirections(v string) {
	o.Directions = v
}

// GetNdc returns the Ndc field value
func (o *MedicationPrescribed) GetNdc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ndc
}

// GetNdcOk returns a tuple with the Ndc field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetNdcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ndc, true
}

// SetNdc sets field value
func (o *MedicationPrescribed) SetNdc(v string) {
	o.Ndc = v
}

// GetRxcui returns the Rxcui field value if set, zero value otherwise.
// Deprecated
func (o *MedicationPrescribed) GetRxcui() string {
	if o == nil || IsNil(o.Rxcui) {
		var ret string
		return ret
	}
	return *o.Rxcui
}

// GetRxcuiOk returns a tuple with the Rxcui field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MedicationPrescribed) GetRxcuiOk() (*string, bool) {
	if o == nil || IsNil(o.Rxcui) {
		return nil, false
	}
	return o.Rxcui, true
}

// HasRxcui returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasRxcui() bool {
	if o != nil && !IsNil(o.Rxcui) {
		return true
	}

	return false
}

// SetRxcui gets a reference to the given string and assigns it to the Rxcui field.
// Deprecated
func (o *MedicationPrescribed) SetRxcui(v string) {
	o.Rxcui = &v
}

// GetDrugDbCode returns the DrugDbCode field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetDrugDbCode() DrugDbCode {
	if o == nil || IsNil(o.DrugDbCode) {
		var ret DrugDbCode
		return ret
	}
	return *o.DrugDbCode
}

// GetDrugDbCodeOk returns a tuple with the DrugDbCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetDrugDbCodeOk() (*DrugDbCode, bool) {
	if o == nil || IsNil(o.DrugDbCode) {
		return nil, false
	}
	return o.DrugDbCode, true
}

// HasDrugDbCode returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasDrugDbCode() bool {
	if o != nil && !IsNil(o.DrugDbCode) {
		return true
	}

	return false
}

// SetDrugDbCode gets a reference to the given DrugDbCode and assigns it to the DrugDbCode field.
func (o *MedicationPrescribed) SetDrugDbCode(v DrugDbCode) {
	o.DrugDbCode = &v
}

// GetDispenseAsWritten returns the DispenseAsWritten field value
func (o *MedicationPrescribed) GetDispenseAsWritten() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DispenseAsWritten
}

// GetDispenseAsWrittenOk returns a tuple with the DispenseAsWritten field value
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetDispenseAsWrittenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DispenseAsWritten, true
}

// SetDispenseAsWritten sets field value
func (o *MedicationPrescribed) SetDispenseAsWritten(v string) {
	o.DispenseAsWritten = v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetEffectiveDate() string {
	if o == nil || IsNil(o.EffectiveDate) {
		var ret string
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetEffectiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveDate) {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasEffectiveDate() bool {
	if o != nil && !IsNil(o.EffectiveDate) {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given string and assigns it to the EffectiveDate field.
func (o *MedicationPrescribed) SetEffectiveDate(v string) {
	o.EffectiveDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *MedicationPrescribed) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetDeaSchedule returns the DeaSchedule field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetDeaSchedule() DeaSchedule {
	if o == nil || IsNil(o.DeaSchedule) {
		var ret DeaSchedule
		return ret
	}
	return *o.DeaSchedule
}

// GetDeaScheduleOk returns a tuple with the DeaSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetDeaScheduleOk() (*DeaSchedule, bool) {
	if o == nil || IsNil(o.DeaSchedule) {
		return nil, false
	}
	return o.DeaSchedule, true
}

// HasDeaSchedule returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasDeaSchedule() bool {
	if o != nil && !IsNil(o.DeaSchedule) {
		return true
	}

	return false
}

// SetDeaSchedule gets a reference to the given DeaSchedule and assigns it to the DeaSchedule field.
func (o *MedicationPrescribed) SetDeaSchedule(v DeaSchedule) {
	o.DeaSchedule = &v
}

// GetDiagnosis returns the Diagnosis field value if set, zero value otherwise.
// Deprecated
func (o *MedicationPrescribed) GetDiagnosis() []Diagnosis {
	if o == nil || IsNil(o.Diagnosis) {
		var ret []Diagnosis
		return ret
	}
	return o.Diagnosis
}

// GetDiagnosisOk returns a tuple with the Diagnosis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MedicationPrescribed) GetDiagnosisOk() ([]Diagnosis, bool) {
	if o == nil || IsNil(o.Diagnosis) {
		return nil, false
	}
	return o.Diagnosis, true
}

// HasDiagnosis returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasDiagnosis() bool {
	if o != nil && !IsNil(o.Diagnosis) {
		return true
	}

	return false
}

// SetDiagnosis gets a reference to the given []Diagnosis and assigns it to the Diagnosis field.
// Deprecated
func (o *MedicationPrescribed) SetDiagnosis(v []Diagnosis) {
	o.Diagnosis = v
}

// GetMedicalDiagnosis returns the MedicalDiagnosis field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetMedicalDiagnosis() MedicalDiagnosis {
	if o == nil || IsNil(o.MedicalDiagnosis) {
		var ret MedicalDiagnosis
		return ret
	}
	return *o.MedicalDiagnosis
}

// GetMedicalDiagnosisOk returns a tuple with the MedicalDiagnosis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetMedicalDiagnosisOk() (*MedicalDiagnosis, bool) {
	if o == nil || IsNil(o.MedicalDiagnosis) {
		return nil, false
	}
	return o.MedicalDiagnosis, true
}

// HasMedicalDiagnosis returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasMedicalDiagnosis() bool {
	if o != nil && !IsNil(o.MedicalDiagnosis) {
		return true
	}

	return false
}

// SetMedicalDiagnosis gets a reference to the given MedicalDiagnosis and assigns it to the MedicalDiagnosis field.
func (o *MedicationPrescribed) SetMedicalDiagnosis(v MedicalDiagnosis) {
	o.MedicalDiagnosis = &v
}

// GetFirstFillDate returns the FirstFillDate field value if set, zero value otherwise.
// Deprecated
func (o *MedicationPrescribed) GetFirstFillDate() string {
	if o == nil || IsNil(o.FirstFillDate) {
		var ret string
		return ret
	}
	return *o.FirstFillDate
}

// GetFirstFillDateOk returns a tuple with the FirstFillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MedicationPrescribed) GetFirstFillDateOk() (*string, bool) {
	if o == nil || IsNil(o.FirstFillDate) {
		return nil, false
	}
	return o.FirstFillDate, true
}

// HasFirstFillDate returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasFirstFillDate() bool {
	if o != nil && !IsNil(o.FirstFillDate) {
		return true
	}

	return false
}

// SetFirstFillDate gets a reference to the given string and assigns it to the FirstFillDate field.
// Deprecated
func (o *MedicationPrescribed) SetFirstFillDate(v string) {
	o.FirstFillDate = &v
}

// GetLastFillDate returns the LastFillDate field value if set, zero value otherwise.
// Deprecated
func (o *MedicationPrescribed) GetLastFillDate() string {
	if o == nil || IsNil(o.LastFillDate) {
		var ret string
		return ret
	}
	return *o.LastFillDate
}

// GetLastFillDateOk returns a tuple with the LastFillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MedicationPrescribed) GetLastFillDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastFillDate) {
		return nil, false
	}
	return o.LastFillDate, true
}

// HasLastFillDate returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasLastFillDate() bool {
	if o != nil && !IsNil(o.LastFillDate) {
		return true
	}

	return false
}

// SetLastFillDate gets a reference to the given string and assigns it to the LastFillDate field.
// Deprecated
func (o *MedicationPrescribed) SetLastFillDate(v string) {
	o.LastFillDate = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *MedicationPrescribed) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationPrescribed) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *MedicationPrescribed) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *MedicationPrescribed) SetNote(v string) {
	o.Note = &v
}

func (o MedicationPrescribed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicationPrescribed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["drugDescription"] = o.DrugDescription
	if !IsNil(o.LabelerName) {
		toSerialize["labelerName"] = o.LabelerName
	}
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	toSerialize["daysSupply"] = o.DaysSupply
	toSerialize["writtenDate"] = o.WrittenDate
	toSerialize["numberOfRefills"] = o.NumberOfRefills
	toSerialize["directions"] = o.Directions
	toSerialize["ndc"] = o.Ndc
	if !IsNil(o.Rxcui) {
		toSerialize["rxcui"] = o.Rxcui
	}
	if !IsNil(o.DrugDbCode) {
		toSerialize["drugDbCode"] = o.DrugDbCode
	}
	toSerialize["dispenseAsWritten"] = o.DispenseAsWritten
	if !IsNil(o.EffectiveDate) {
		toSerialize["effectiveDate"] = o.EffectiveDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.DeaSchedule) {
		toSerialize["deaSchedule"] = o.DeaSchedule
	}
	if !IsNil(o.Diagnosis) {
		toSerialize["diagnosis"] = o.Diagnosis
	}
	if !IsNil(o.MedicalDiagnosis) {
		toSerialize["medicalDiagnosis"] = o.MedicalDiagnosis
	}
	if !IsNil(o.FirstFillDate) {
		toSerialize["firstFillDate"] = o.FirstFillDate
	}
	if !IsNil(o.LastFillDate) {
		toSerialize["lastFillDate"] = o.LastFillDate
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

func (o *MedicationPrescribed) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"drugDescription",
		"quantity",
		"daysSupply",
		"writtenDate",
		"numberOfRefills",
		"directions",
		"ndc",
		"dispenseAsWritten",
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

	varMedicationPrescribed := _MedicationPrescribed{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicationPrescribed)

	if err != nil {
		return err
	}

	*o = MedicationPrescribed(varMedicationPrescribed)

	return err
}

type NullableMedicationPrescribed struct {
	value *MedicationPrescribed
	isSet bool
}

func (v NullableMedicationPrescribed) Get() *MedicationPrescribed {
	return v.value
}

func (v *NullableMedicationPrescribed) Set(val *MedicationPrescribed) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicationPrescribed) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicationPrescribed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicationPrescribed(val *MedicationPrescribed) *NullableMedicationPrescribed {
	return &NullableMedicationPrescribed{value: val, isSet: true}
}

func (v NullableMedicationPrescribed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicationPrescribed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


