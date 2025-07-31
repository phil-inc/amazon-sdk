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

// checks if the PutOrderRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutOrderRequestContent{}

// PutOrderRequestContent Structure representing the request for placing an order.
type PutOrderRequestContent struct {
	// A unique Order identifier provided by the partner.
	PartnerOrderId string `json:"partnerOrderId" validate:"regexp=^[0-9a-zA-Z-]{1,40}$"`
	// A unique Prescription Identifier provided by the partner.
	PartnerPrescriptionId string  `json:"partnerPrescriptionId" validate:"regexp=^[0-9a-zA-Z-]{1,40}$"`
	ShippingAddress       Address `json:"shippingAddress"`
	// Customer's payment instrument ID, previously created using the payment-instrument API. This field is required when the quoted copay amount is greater than zero and placeOrderOnDate is null.
	PartnerPaymentInstrumentId *string `json:"partnerPaymentInstrumentId,omitempty" validate:"regexp=^[0-9a-zA-Z-]{1,40}$"`
	// An estimated copay for the customer. If the actual copay is higher than the estimated copay by a certain amount($XX)     or percentage (XX%), the shipment will be put on hold.
	QuotedCopay *string `json:"quotedCopay,omitempty" validate:"regexp=^([0-9]*[.])?[0-9]+$"`
	// Order will be held until this date and sent for order placement only when the date is passed.     It is quoted in the format 'YYYY-MM-DD'.
	PlaceOrderOnDate      *string                `json:"placeOrderOnDate,omitempty" validate:"regexp=^\\\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"`
	OrderInsuranceDetails *OrderInsuranceDetails `json:"orderInsuranceDetails,omitempty"`
}

type _PutOrderRequestContent PutOrderRequestContent

// NewPutOrderRequestContent instantiates a new PutOrderRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutOrderRequestContent(partnerOrderId string, partnerPrescriptionId string, shippingAddress Address) *PutOrderRequestContent {
	this := PutOrderRequestContent{}
	this.PartnerOrderId = partnerOrderId
	this.PartnerPrescriptionId = partnerPrescriptionId
	this.ShippingAddress = shippingAddress
	return &this
}

// NewPutOrderRequestContentWithDefaults instantiates a new PutOrderRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutOrderRequestContentWithDefaults() *PutOrderRequestContent {
	this := PutOrderRequestContent{}
	return &this
}

// GetPartnerOrderId returns the PartnerOrderId field value
func (o *PutOrderRequestContent) GetPartnerOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerOrderId
}

// GetPartnerOrderIdOk returns a tuple with the PartnerOrderId field value
// and a boolean to check if the value has been set.
func (o *PutOrderRequestContent) GetPartnerOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerOrderId, true
}

// SetPartnerOrderId sets field value
func (o *PutOrderRequestContent) SetPartnerOrderId(v string) {
	o.PartnerOrderId = v
}

// GetPartnerPrescriptionId returns the PartnerPrescriptionId field value
func (o *PutOrderRequestContent) GetPartnerPrescriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerPrescriptionId
}

// GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field value
// and a boolean to check if the value has been set.
func (o *PutOrderRequestContent) GetPartnerPrescriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerPrescriptionId, true
}

// SetPartnerPrescriptionId sets field value
func (o *PutOrderRequestContent) SetPartnerPrescriptionId(v string) {
	o.PartnerPrescriptionId = v
}

// GetShippingAddress returns the ShippingAddress field value
func (o *PutOrderRequestContent) GetShippingAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value
// and a boolean to check if the value has been set.
func (o *PutOrderRequestContent) GetShippingAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingAddress, true
}

// SetShippingAddress sets field value
func (o *PutOrderRequestContent) SetShippingAddress(v Address) {
	o.ShippingAddress = v
}

// GetPartnerPaymentInstrumentId returns the PartnerPaymentInstrumentId field value if set, zero value otherwise.
func (o *PutOrderRequestContent) GetPartnerPaymentInstrumentId() string {
	if o == nil || IsNil(o.PartnerPaymentInstrumentId) {
		var ret string
		return ret
	}
	return *o.PartnerPaymentInstrumentId
}

// GetPartnerPaymentInstrumentIdOk returns a tuple with the PartnerPaymentInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrderRequestContent) GetPartnerPaymentInstrumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerPaymentInstrumentId) {
		return nil, false
	}
	return o.PartnerPaymentInstrumentId, true
}

// HasPartnerPaymentInstrumentId returns a boolean if a field has been set.
func (o *PutOrderRequestContent) HasPartnerPaymentInstrumentId() bool {
	if o != nil && !IsNil(o.PartnerPaymentInstrumentId) {
		return true
	}

	return false
}

// SetPartnerPaymentInstrumentId gets a reference to the given string and assigns it to the PartnerPaymentInstrumentId field.
func (o *PutOrderRequestContent) SetPartnerPaymentInstrumentId(v string) {
	o.PartnerPaymentInstrumentId = &v
}

// GetQuotedCopay returns the QuotedCopay field value if set, zero value otherwise.
func (o *PutOrderRequestContent) GetQuotedCopay() string {
	if o == nil || IsNil(o.QuotedCopay) {
		var ret string
		return ret
	}
	return *o.QuotedCopay
}

// GetQuotedCopayOk returns a tuple with the QuotedCopay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrderRequestContent) GetQuotedCopayOk() (*string, bool) {
	if o == nil || IsNil(o.QuotedCopay) {
		return nil, false
	}
	return o.QuotedCopay, true
}

// HasQuotedCopay returns a boolean if a field has been set.
func (o *PutOrderRequestContent) HasQuotedCopay() bool {
	if o != nil && !IsNil(o.QuotedCopay) {
		return true
	}

	return false
}

// SetQuotedCopay gets a reference to the given string and assigns it to the QuotedCopay field.
func (o *PutOrderRequestContent) SetQuotedCopay(v string) {
	o.QuotedCopay = &v
}

// GetPlaceOrderOnDate returns the PlaceOrderOnDate field value if set, zero value otherwise.
func (o *PutOrderRequestContent) GetPlaceOrderOnDate() string {
	if o == nil || IsNil(o.PlaceOrderOnDate) {
		var ret string
		return ret
	}
	return *o.PlaceOrderOnDate
}

// GetPlaceOrderOnDateOk returns a tuple with the PlaceOrderOnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrderRequestContent) GetPlaceOrderOnDateOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceOrderOnDate) {
		return nil, false
	}
	return o.PlaceOrderOnDate, true
}

// HasPlaceOrderOnDate returns a boolean if a field has been set.
func (o *PutOrderRequestContent) HasPlaceOrderOnDate() bool {
	if o != nil && !IsNil(o.PlaceOrderOnDate) {
		return true
	}

	return false
}

// SetPlaceOrderOnDate gets a reference to the given string and assigns it to the PlaceOrderOnDate field.
func (o *PutOrderRequestContent) SetPlaceOrderOnDate(v string) {
	o.PlaceOrderOnDate = &v
}

// GetOrderInsuranceDetails returns the OrderInsuranceDetails field value if set, zero value otherwise.
func (o *PutOrderRequestContent) GetOrderInsuranceDetails() OrderInsuranceDetails {
	if o == nil || IsNil(o.OrderInsuranceDetails) {
		var ret OrderInsuranceDetails
		return ret
	}
	return *o.OrderInsuranceDetails
}

// GetOrderInsuranceDetailsOk returns a tuple with the OrderInsuranceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrderRequestContent) GetOrderInsuranceDetailsOk() (*OrderInsuranceDetails, bool) {
	if o == nil || IsNil(o.OrderInsuranceDetails) {
		return nil, false
	}
	return o.OrderInsuranceDetails, true
}

// HasOrderInsuranceDetails returns a boolean if a field has been set.
func (o *PutOrderRequestContent) HasOrderInsuranceDetails() bool {
	if o != nil && !IsNil(o.OrderInsuranceDetails) {
		return true
	}

	return false
}

// SetOrderInsuranceDetails gets a reference to the given OrderInsuranceDetails and assigns it to the OrderInsuranceDetails field.
func (o *PutOrderRequestContent) SetOrderInsuranceDetails(v OrderInsuranceDetails) {
	o.OrderInsuranceDetails = &v
}

func (o PutOrderRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutOrderRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["partnerOrderId"] = o.PartnerOrderId
	toSerialize["partnerPrescriptionId"] = o.PartnerPrescriptionId
	toSerialize["shippingAddress"] = o.ShippingAddress
	if !IsNil(o.PartnerPaymentInstrumentId) {
		toSerialize["partnerPaymentInstrumentId"] = o.PartnerPaymentInstrumentId
	}
	if !IsNil(o.QuotedCopay) {
		toSerialize["quotedCopay"] = o.QuotedCopay
	}
	if !IsNil(o.PlaceOrderOnDate) {
		toSerialize["placeOrderOnDate"] = o.PlaceOrderOnDate
	}
	if !IsNil(o.OrderInsuranceDetails) {
		toSerialize["orderInsuranceDetails"] = o.OrderInsuranceDetails
	}
	return toSerialize, nil
}

func (o *PutOrderRequestContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"partnerOrderId",
		"partnerPrescriptionId",
		"shippingAddress",
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

	varPutOrderRequestContent := _PutOrderRequestContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPutOrderRequestContent)

	if err != nil {
		return err
	}

	*o = PutOrderRequestContent(varPutOrderRequestContent)

	return err
}

type NullablePutOrderRequestContent struct {
	value *PutOrderRequestContent
	isSet bool
}

func (v NullablePutOrderRequestContent) Get() *PutOrderRequestContent {
	return v.value
}

func (v *NullablePutOrderRequestContent) Set(val *PutOrderRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePutOrderRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePutOrderRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutOrderRequestContent(val *PutOrderRequestContent) *NullablePutOrderRequestContent {
	return &NullablePutOrderRequestContent{value: val, isSet: true}
}

func (v NullablePutOrderRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutOrderRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
