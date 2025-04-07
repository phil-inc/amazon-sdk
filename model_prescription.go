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

// checks if the Prescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prescription{}

// Prescription Structure representing a prescription.
type Prescription struct {
	Status PrescriptionStatus `json:"status"`
	// A more detailed explanation of the prescription's status.
	StatusDescription string `json:"statusDescription" validate:"regexp=^\\\\s*\\\\S.*$"`
	AmazonPharmacy Pharmacy `json:"amazonPharmacy"`
	// Unit of the drug quantity, e.g., Caps/Tabs/ML/Dev.
	DrugUnits *string `json:"drugUnits,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
	// The National Drug Code of the medication, without dashes or spaces.
	Ndc *string `json:"ndc,omitempty" validate:"regexp=^\\\\d{11}$"`
	// Number of refills authorized on the Prescription.
	RefillsAuthorized *float32 `json:"refillsAuthorized,omitempty"`
	// Number of refills remaining on the Prescription.
	RefillsRemaining *float32 `json:"refillsRemaining,omitempty"`
	// Number of total fills (first fill + refill) remaining on the Prescription.
	FillsRemaining *float32 `json:"fillsRemaining,omitempty"`
	// The number of days the authorizedQuantityPerFill will supply, mapped to Prescription's daysSupply.
	DaysSupply *float32 `json:"daysSupply,omitempty"`
	// The pharmacy-assigned number for this prescription.
	RxNumber *string `json:"rxNumber,omitempty" validate:"regexp=^\\\\s*\\\\S.*$"`
	TransferRequestStatus *TransferRequestStatus `json:"transferRequestStatus,omitempty"`
	TotalQuantityRemaining *Quantity `json:"totalQuantityRemaining,omitempty"`
	// A list of fills for this prescription
	Fills []PrescriptionFill `json:"fills,omitempty"`
	Prescriber *Prescriber `json:"prescriber,omitempty"`
}

type _Prescription Prescription

// NewPrescription instantiates a new Prescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrescription(status PrescriptionStatus, statusDescription string, amazonPharmacy Pharmacy) *Prescription {
	this := Prescription{}
	this.Status = status
	this.StatusDescription = statusDescription
	this.AmazonPharmacy = amazonPharmacy
	return &this
}

// NewPrescriptionWithDefaults instantiates a new Prescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrescriptionWithDefaults() *Prescription {
	this := Prescription{}
	return &this
}

// GetStatus returns the Status field value
func (o *Prescription) GetStatus() PrescriptionStatus {
	if o == nil {
		var ret PrescriptionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Prescription) GetStatusOk() (*PrescriptionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Prescription) SetStatus(v PrescriptionStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *Prescription) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *Prescription) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *Prescription) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetAmazonPharmacy returns the AmazonPharmacy field value
func (o *Prescription) GetAmazonPharmacy() Pharmacy {
	if o == nil {
		var ret Pharmacy
		return ret
	}

	return o.AmazonPharmacy
}

// GetAmazonPharmacyOk returns a tuple with the AmazonPharmacy field value
// and a boolean to check if the value has been set.
func (o *Prescription) GetAmazonPharmacyOk() (*Pharmacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonPharmacy, true
}

// SetAmazonPharmacy sets field value
func (o *Prescription) SetAmazonPharmacy(v Pharmacy) {
	o.AmazonPharmacy = v
}

// GetDrugUnits returns the DrugUnits field value if set, zero value otherwise.
func (o *Prescription) GetDrugUnits() string {
	if o == nil || IsNil(o.DrugUnits) {
		var ret string
		return ret
	}
	return *o.DrugUnits
}

// GetDrugUnitsOk returns a tuple with the DrugUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetDrugUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.DrugUnits) {
		return nil, false
	}
	return o.DrugUnits, true
}

// HasDrugUnits returns a boolean if a field has been set.
func (o *Prescription) HasDrugUnits() bool {
	if o != nil && !IsNil(o.DrugUnits) {
		return true
	}

	return false
}

// SetDrugUnits gets a reference to the given string and assigns it to the DrugUnits field.
func (o *Prescription) SetDrugUnits(v string) {
	o.DrugUnits = &v
}

// GetNdc returns the Ndc field value if set, zero value otherwise.
func (o *Prescription) GetNdc() string {
	if o == nil || IsNil(o.Ndc) {
		var ret string
		return ret
	}
	return *o.Ndc
}

// GetNdcOk returns a tuple with the Ndc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetNdcOk() (*string, bool) {
	if o == nil || IsNil(o.Ndc) {
		return nil, false
	}
	return o.Ndc, true
}

// HasNdc returns a boolean if a field has been set.
func (o *Prescription) HasNdc() bool {
	if o != nil && !IsNil(o.Ndc) {
		return true
	}

	return false
}

// SetNdc gets a reference to the given string and assigns it to the Ndc field.
func (o *Prescription) SetNdc(v string) {
	o.Ndc = &v
}

// GetRefillsAuthorized returns the RefillsAuthorized field value if set, zero value otherwise.
func (o *Prescription) GetRefillsAuthorized() float32 {
	if o == nil || IsNil(o.RefillsAuthorized) {
		var ret float32
		return ret
	}
	return *o.RefillsAuthorized
}

// GetRefillsAuthorizedOk returns a tuple with the RefillsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetRefillsAuthorizedOk() (*float32, bool) {
	if o == nil || IsNil(o.RefillsAuthorized) {
		return nil, false
	}
	return o.RefillsAuthorized, true
}

// HasRefillsAuthorized returns a boolean if a field has been set.
func (o *Prescription) HasRefillsAuthorized() bool {
	if o != nil && !IsNil(o.RefillsAuthorized) {
		return true
	}

	return false
}

// SetRefillsAuthorized gets a reference to the given float32 and assigns it to the RefillsAuthorized field.
func (o *Prescription) SetRefillsAuthorized(v float32) {
	o.RefillsAuthorized = &v
}

// GetRefillsRemaining returns the RefillsRemaining field value if set, zero value otherwise.
func (o *Prescription) GetRefillsRemaining() float32 {
	if o == nil || IsNil(o.RefillsRemaining) {
		var ret float32
		return ret
	}
	return *o.RefillsRemaining
}

// GetRefillsRemainingOk returns a tuple with the RefillsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetRefillsRemainingOk() (*float32, bool) {
	if o == nil || IsNil(o.RefillsRemaining) {
		return nil, false
	}
	return o.RefillsRemaining, true
}

// HasRefillsRemaining returns a boolean if a field has been set.
func (o *Prescription) HasRefillsRemaining() bool {
	if o != nil && !IsNil(o.RefillsRemaining) {
		return true
	}

	return false
}

// SetRefillsRemaining gets a reference to the given float32 and assigns it to the RefillsRemaining field.
func (o *Prescription) SetRefillsRemaining(v float32) {
	o.RefillsRemaining = &v
}

// GetFillsRemaining returns the FillsRemaining field value if set, zero value otherwise.
func (o *Prescription) GetFillsRemaining() float32 {
	if o == nil || IsNil(o.FillsRemaining) {
		var ret float32
		return ret
	}
	return *o.FillsRemaining
}

// GetFillsRemainingOk returns a tuple with the FillsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetFillsRemainingOk() (*float32, bool) {
	if o == nil || IsNil(o.FillsRemaining) {
		return nil, false
	}
	return o.FillsRemaining, true
}

// HasFillsRemaining returns a boolean if a field has been set.
func (o *Prescription) HasFillsRemaining() bool {
	if o != nil && !IsNil(o.FillsRemaining) {
		return true
	}

	return false
}

// SetFillsRemaining gets a reference to the given float32 and assigns it to the FillsRemaining field.
func (o *Prescription) SetFillsRemaining(v float32) {
	o.FillsRemaining = &v
}

// GetDaysSupply returns the DaysSupply field value if set, zero value otherwise.
func (o *Prescription) GetDaysSupply() float32 {
	if o == nil || IsNil(o.DaysSupply) {
		var ret float32
		return ret
	}
	return *o.DaysSupply
}

// GetDaysSupplyOk returns a tuple with the DaysSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetDaysSupplyOk() (*float32, bool) {
	if o == nil || IsNil(o.DaysSupply) {
		return nil, false
	}
	return o.DaysSupply, true
}

// HasDaysSupply returns a boolean if a field has been set.
func (o *Prescription) HasDaysSupply() bool {
	if o != nil && !IsNil(o.DaysSupply) {
		return true
	}

	return false
}

// SetDaysSupply gets a reference to the given float32 and assigns it to the DaysSupply field.
func (o *Prescription) SetDaysSupply(v float32) {
	o.DaysSupply = &v
}

// GetRxNumber returns the RxNumber field value if set, zero value otherwise.
func (o *Prescription) GetRxNumber() string {
	if o == nil || IsNil(o.RxNumber) {
		var ret string
		return ret
	}
	return *o.RxNumber
}

// GetRxNumberOk returns a tuple with the RxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetRxNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RxNumber) {
		return nil, false
	}
	return o.RxNumber, true
}

// HasRxNumber returns a boolean if a field has been set.
func (o *Prescription) HasRxNumber() bool {
	if o != nil && !IsNil(o.RxNumber) {
		return true
	}

	return false
}

// SetRxNumber gets a reference to the given string and assigns it to the RxNumber field.
func (o *Prescription) SetRxNumber(v string) {
	o.RxNumber = &v
}

// GetTransferRequestStatus returns the TransferRequestStatus field value if set, zero value otherwise.
func (o *Prescription) GetTransferRequestStatus() TransferRequestStatus {
	if o == nil || IsNil(o.TransferRequestStatus) {
		var ret TransferRequestStatus
		return ret
	}
	return *o.TransferRequestStatus
}

// GetTransferRequestStatusOk returns a tuple with the TransferRequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetTransferRequestStatusOk() (*TransferRequestStatus, bool) {
	if o == nil || IsNil(o.TransferRequestStatus) {
		return nil, false
	}
	return o.TransferRequestStatus, true
}

// HasTransferRequestStatus returns a boolean if a field has been set.
func (o *Prescription) HasTransferRequestStatus() bool {
	if o != nil && !IsNil(o.TransferRequestStatus) {
		return true
	}

	return false
}

// SetTransferRequestStatus gets a reference to the given TransferRequestStatus and assigns it to the TransferRequestStatus field.
func (o *Prescription) SetTransferRequestStatus(v TransferRequestStatus) {
	o.TransferRequestStatus = &v
}

// GetTotalQuantityRemaining returns the TotalQuantityRemaining field value if set, zero value otherwise.
func (o *Prescription) GetTotalQuantityRemaining() Quantity {
	if o == nil || IsNil(o.TotalQuantityRemaining) {
		var ret Quantity
		return ret
	}
	return *o.TotalQuantityRemaining
}

// GetTotalQuantityRemainingOk returns a tuple with the TotalQuantityRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetTotalQuantityRemainingOk() (*Quantity, bool) {
	if o == nil || IsNil(o.TotalQuantityRemaining) {
		return nil, false
	}
	return o.TotalQuantityRemaining, true
}

// HasTotalQuantityRemaining returns a boolean if a field has been set.
func (o *Prescription) HasTotalQuantityRemaining() bool {
	if o != nil && !IsNil(o.TotalQuantityRemaining) {
		return true
	}

	return false
}

// SetTotalQuantityRemaining gets a reference to the given Quantity and assigns it to the TotalQuantityRemaining field.
func (o *Prescription) SetTotalQuantityRemaining(v Quantity) {
	o.TotalQuantityRemaining = &v
}

// GetFills returns the Fills field value if set, zero value otherwise.
func (o *Prescription) GetFills() []PrescriptionFill {
	if o == nil || IsNil(o.Fills) {
		var ret []PrescriptionFill
		return ret
	}
	return o.Fills
}

// GetFillsOk returns a tuple with the Fills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetFillsOk() ([]PrescriptionFill, bool) {
	if o == nil || IsNil(o.Fills) {
		return nil, false
	}
	return o.Fills, true
}

// HasFills returns a boolean if a field has been set.
func (o *Prescription) HasFills() bool {
	if o != nil && !IsNil(o.Fills) {
		return true
	}

	return false
}

// SetFills gets a reference to the given []PrescriptionFill and assigns it to the Fills field.
func (o *Prescription) SetFills(v []PrescriptionFill) {
	o.Fills = v
}

// GetPrescriber returns the Prescriber field value if set, zero value otherwise.
func (o *Prescription) GetPrescriber() Prescriber {
	if o == nil || IsNil(o.Prescriber) {
		var ret Prescriber
		return ret
	}
	return *o.Prescriber
}

// GetPrescriberOk returns a tuple with the Prescriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prescription) GetPrescriberOk() (*Prescriber, bool) {
	if o == nil || IsNil(o.Prescriber) {
		return nil, false
	}
	return o.Prescriber, true
}

// HasPrescriber returns a boolean if a field has been set.
func (o *Prescription) HasPrescriber() bool {
	if o != nil && !IsNil(o.Prescriber) {
		return true
	}

	return false
}

// SetPrescriber gets a reference to the given Prescriber and assigns it to the Prescriber field.
func (o *Prescription) SetPrescriber(v Prescriber) {
	o.Prescriber = &v
}

func (o Prescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Prescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	toSerialize["amazonPharmacy"] = o.AmazonPharmacy
	if !IsNil(o.DrugUnits) {
		toSerialize["drugUnits"] = o.DrugUnits
	}
	if !IsNil(o.Ndc) {
		toSerialize["ndc"] = o.Ndc
	}
	if !IsNil(o.RefillsAuthorized) {
		toSerialize["refillsAuthorized"] = o.RefillsAuthorized
	}
	if !IsNil(o.RefillsRemaining) {
		toSerialize["refillsRemaining"] = o.RefillsRemaining
	}
	if !IsNil(o.FillsRemaining) {
		toSerialize["fillsRemaining"] = o.FillsRemaining
	}
	if !IsNil(o.DaysSupply) {
		toSerialize["daysSupply"] = o.DaysSupply
	}
	if !IsNil(o.RxNumber) {
		toSerialize["rxNumber"] = o.RxNumber
	}
	if !IsNil(o.TransferRequestStatus) {
		toSerialize["transferRequestStatus"] = o.TransferRequestStatus
	}
	if !IsNil(o.TotalQuantityRemaining) {
		toSerialize["totalQuantityRemaining"] = o.TotalQuantityRemaining
	}
	if !IsNil(o.Fills) {
		toSerialize["fills"] = o.Fills
	}
	if !IsNil(o.Prescriber) {
		toSerialize["prescriber"] = o.Prescriber
	}
	return toSerialize, nil
}

func (o *Prescription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"statusDescription",
		"amazonPharmacy",
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

	varPrescription := _Prescription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrescription)

	if err != nil {
		return err
	}

	*o = Prescription(varPrescription)

	return err
}

type NullablePrescription struct {
	value *Prescription
	isSet bool
}

func (v NullablePrescription) Get() *Prescription {
	return v.value
}

func (v *NullablePrescription) Set(val *Prescription) {
	v.value = val
	v.isSet = true
}

func (v NullablePrescription) IsSet() bool {
	return v.isSet
}

func (v *NullablePrescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrescription(val *Prescription) *NullablePrescription {
	return &NullablePrescription{value: val, isSet: true}
}

func (v NullablePrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


