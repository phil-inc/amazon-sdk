# MedicationPrescribed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrugDescription** | **string** | Name or description of the medication being prescribed. | 
**LabelerName** | Pointer to **string** | Drug labeler/manufacturer name. | [optional] 
**Quantity** | [**Quantity**](Quantity.md) |  | 
**Strength** | Pointer to [**Strength**](Strength.md) |  | [optional] 
**DaysSupply** | **float64** | Estimated number of days a prescription would last. | 
**WrittenDate** | **string** | The date on which the prescriber wrote the prescription, quoted in the format &#39;YYYY-MM-DD&#39;. | 
**NumberOfRefills** | **float64** | The number of refills allowed for this prescription. | 
**Directions** | **string** | Clinical directions meaningful to the patient as written by the prescriber, e.g., &#39;Take 1 tablet by mouth once weekly.&#39; | 
**Ndc** | **string** | The National Drug Code of the medication, without dashes or spaces. | 
**Rxcui** | Pointer to **string** | RxNorm concept unique identifier. This shape is deprecated: RxNorm codes are included in drugDbCode field. This field will no longer be supported starting from 2025-03-12. | [optional] 
**DrugDbCode** | Pointer to [**DrugDbCode**](DrugDbCode.md) |  | [optional] 
**DispenseAsWritten** | **string** | An indicator that the prescriber has instructed to dispense as written (DAW) code. This is typically used when a brand is chosen over a generic alternative based on the direction of the prescriber. 0 means no DAW is specified 1 means by prescriber 2 means a patient/user request 9 means other. | 
**EffectiveDate** | Pointer to **string** | The date on which the user is intended to start the prescription, quoted in the format &#39;YYYY-MM-DD&#39;. | [optional] 
**ExpirationDate** | Pointer to **string** | The date on which the prescription expires, quoted in the format &#39;YYYY-MM-DD&#39;. | [optional] 
**DeaSchedule** | Pointer to [**DeaSchedule**](DeaSchedule.md) |  | [optional] 
**Diagnosis** | Pointer to [**[]Diagnosis**](Diagnosis.md) | Medical diagnosis information for which the patient is prescribed. First item is the Primary Diagnosis, second is Secondary.  This shape is deprecated: This field will be replaced by &#x60;medicalDiagnosis&#x60;. This field will no longer be supported starting from 2025-03-12 | [optional] 
**MedicalDiagnosis** | Pointer to [**MedicalDiagnosis**](MedicalDiagnosis.md) |  | [optional] 
**FirstFillDate** | Pointer to **string** | The firstFillDate field is deprecated here and will be part of PrescriptionTransferred structure instead. The date on which the prescription was first filled by the pharmacy, in the format &#39;YYYY-MM-DD&#39;. This date helps track the initial dispensing of the medication to the patient.  This shape is deprecated: This field is moved to the &#x60;PrescriptionTransferred&#x60; data structure. It will no longer be available here. | [optional] 
**LastFillDate** | Pointer to **string** | The lastFillDate field is deprecated here and will be part of PrescriptionTransferred structure instead. The date on which the prescription was most recently filled by the pharmacy, in the format &#39;YYYY-MM-DD&#39;. This date helps track the latest dispensing of the medication and is useful for monitoring adherence and refill patterns.  This shape is deprecated: This field is moved to the &#x60;PrescriptionTransferred&#x60; data structure. It will no longer be available here. | [optional] 
**Note** | Pointer to **string** | Additional information on the Patient or Prescription provided by the Prescriber. | [optional] 

## Methods

### NewMedicationPrescribed

`func NewMedicationPrescribed(drugDescription string, quantity Quantity, daysSupply float64, writtenDate string, numberOfRefills float64, directions string, ndc string, dispenseAsWritten string, ) *MedicationPrescribed`

NewMedicationPrescribed instantiates a new MedicationPrescribed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMedicationPrescribedWithDefaults

`func NewMedicationPrescribedWithDefaults() *MedicationPrescribed`

NewMedicationPrescribedWithDefaults instantiates a new MedicationPrescribed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrugDescription

`func (o *MedicationPrescribed) GetDrugDescription() string`

GetDrugDescription returns the DrugDescription field if non-nil, zero value otherwise.

### GetDrugDescriptionOk

`func (o *MedicationPrescribed) GetDrugDescriptionOk() (*string, bool)`

GetDrugDescriptionOk returns a tuple with the DrugDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugDescription

`func (o *MedicationPrescribed) SetDrugDescription(v string)`

SetDrugDescription sets DrugDescription field to given value.


### GetLabelerName

`func (o *MedicationPrescribed) GetLabelerName() string`

GetLabelerName returns the LabelerName field if non-nil, zero value otherwise.

### GetLabelerNameOk

`func (o *MedicationPrescribed) GetLabelerNameOk() (*string, bool)`

GetLabelerNameOk returns a tuple with the LabelerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelerName

`func (o *MedicationPrescribed) SetLabelerName(v string)`

SetLabelerName sets LabelerName field to given value.

### HasLabelerName

`func (o *MedicationPrescribed) HasLabelerName() bool`

HasLabelerName returns a boolean if a field has been set.

### GetQuantity

`func (o *MedicationPrescribed) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MedicationPrescribed) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MedicationPrescribed) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.


### GetStrength

`func (o *MedicationPrescribed) GetStrength() Strength`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *MedicationPrescribed) GetStrengthOk() (*Strength, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *MedicationPrescribed) SetStrength(v Strength)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *MedicationPrescribed) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetDaysSupply

`func (o *MedicationPrescribed) GetDaysSupply() float64`

GetDaysSupply returns the DaysSupply field if non-nil, zero value otherwise.

### GetDaysSupplyOk

`func (o *MedicationPrescribed) GetDaysSupplyOk() (*float64, bool)`

GetDaysSupplyOk returns a tuple with the DaysSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSupply

`func (o *MedicationPrescribed) SetDaysSupply(v float64)`

SetDaysSupply sets DaysSupply field to given value.


### GetWrittenDate

`func (o *MedicationPrescribed) GetWrittenDate() string`

GetWrittenDate returns the WrittenDate field if non-nil, zero value otherwise.

### GetWrittenDateOk

`func (o *MedicationPrescribed) GetWrittenDateOk() (*string, bool)`

GetWrittenDateOk returns a tuple with the WrittenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenDate

`func (o *MedicationPrescribed) SetWrittenDate(v string)`

SetWrittenDate sets WrittenDate field to given value.


### GetNumberOfRefills

`func (o *MedicationPrescribed) GetNumberOfRefills() float64`

GetNumberOfRefills returns the NumberOfRefills field if non-nil, zero value otherwise.

### GetNumberOfRefillsOk

`func (o *MedicationPrescribed) GetNumberOfRefillsOk() (*float64, bool)`

GetNumberOfRefillsOk returns a tuple with the NumberOfRefills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRefills

`func (o *MedicationPrescribed) SetNumberOfRefills(v float64)`

SetNumberOfRefills sets NumberOfRefills field to given value.


### GetDirections

`func (o *MedicationPrescribed) GetDirections() string`

GetDirections returns the Directions field if non-nil, zero value otherwise.

### GetDirectionsOk

`func (o *MedicationPrescribed) GetDirectionsOk() (*string, bool)`

GetDirectionsOk returns a tuple with the Directions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirections

`func (o *MedicationPrescribed) SetDirections(v string)`

SetDirections sets Directions field to given value.


### GetNdc

`func (o *MedicationPrescribed) GetNdc() string`

GetNdc returns the Ndc field if non-nil, zero value otherwise.

### GetNdcOk

`func (o *MedicationPrescribed) GetNdcOk() (*string, bool)`

GetNdcOk returns a tuple with the Ndc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdc

`func (o *MedicationPrescribed) SetNdc(v string)`

SetNdc sets Ndc field to given value.


### GetRxcui

`func (o *MedicationPrescribed) GetRxcui() string`

GetRxcui returns the Rxcui field if non-nil, zero value otherwise.

### GetRxcuiOk

`func (o *MedicationPrescribed) GetRxcuiOk() (*string, bool)`

GetRxcuiOk returns a tuple with the Rxcui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxcui

`func (o *MedicationPrescribed) SetRxcui(v string)`

SetRxcui sets Rxcui field to given value.

### HasRxcui

`func (o *MedicationPrescribed) HasRxcui() bool`

HasRxcui returns a boolean if a field has been set.

### GetDrugDbCode

`func (o *MedicationPrescribed) GetDrugDbCode() DrugDbCode`

GetDrugDbCode returns the DrugDbCode field if non-nil, zero value otherwise.

### GetDrugDbCodeOk

`func (o *MedicationPrescribed) GetDrugDbCodeOk() (*DrugDbCode, bool)`

GetDrugDbCodeOk returns a tuple with the DrugDbCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugDbCode

`func (o *MedicationPrescribed) SetDrugDbCode(v DrugDbCode)`

SetDrugDbCode sets DrugDbCode field to given value.

### HasDrugDbCode

`func (o *MedicationPrescribed) HasDrugDbCode() bool`

HasDrugDbCode returns a boolean if a field has been set.

### GetDispenseAsWritten

`func (o *MedicationPrescribed) GetDispenseAsWritten() string`

GetDispenseAsWritten returns the DispenseAsWritten field if non-nil, zero value otherwise.

### GetDispenseAsWrittenOk

`func (o *MedicationPrescribed) GetDispenseAsWrittenOk() (*string, bool)`

GetDispenseAsWrittenOk returns a tuple with the DispenseAsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispenseAsWritten

`func (o *MedicationPrescribed) SetDispenseAsWritten(v string)`

SetDispenseAsWritten sets DispenseAsWritten field to given value.


### GetEffectiveDate

`func (o *MedicationPrescribed) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *MedicationPrescribed) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *MedicationPrescribed) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *MedicationPrescribed) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *MedicationPrescribed) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *MedicationPrescribed) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *MedicationPrescribed) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *MedicationPrescribed) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetDeaSchedule

`func (o *MedicationPrescribed) GetDeaSchedule() DeaSchedule`

GetDeaSchedule returns the DeaSchedule field if non-nil, zero value otherwise.

### GetDeaScheduleOk

`func (o *MedicationPrescribed) GetDeaScheduleOk() (*DeaSchedule, bool)`

GetDeaScheduleOk returns a tuple with the DeaSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaSchedule

`func (o *MedicationPrescribed) SetDeaSchedule(v DeaSchedule)`

SetDeaSchedule sets DeaSchedule field to given value.

### HasDeaSchedule

`func (o *MedicationPrescribed) HasDeaSchedule() bool`

HasDeaSchedule returns a boolean if a field has been set.

### GetDiagnosis

`func (o *MedicationPrescribed) GetDiagnosis() []Diagnosis`

GetDiagnosis returns the Diagnosis field if non-nil, zero value otherwise.

### GetDiagnosisOk

`func (o *MedicationPrescribed) GetDiagnosisOk() (*[]Diagnosis, bool)`

GetDiagnosisOk returns a tuple with the Diagnosis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosis

`func (o *MedicationPrescribed) SetDiagnosis(v []Diagnosis)`

SetDiagnosis sets Diagnosis field to given value.

### HasDiagnosis

`func (o *MedicationPrescribed) HasDiagnosis() bool`

HasDiagnosis returns a boolean if a field has been set.

### GetMedicalDiagnosis

`func (o *MedicationPrescribed) GetMedicalDiagnosis() MedicalDiagnosis`

GetMedicalDiagnosis returns the MedicalDiagnosis field if non-nil, zero value otherwise.

### GetMedicalDiagnosisOk

`func (o *MedicationPrescribed) GetMedicalDiagnosisOk() (*MedicalDiagnosis, bool)`

GetMedicalDiagnosisOk returns a tuple with the MedicalDiagnosis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalDiagnosis

`func (o *MedicationPrescribed) SetMedicalDiagnosis(v MedicalDiagnosis)`

SetMedicalDiagnosis sets MedicalDiagnosis field to given value.

### HasMedicalDiagnosis

`func (o *MedicationPrescribed) HasMedicalDiagnosis() bool`

HasMedicalDiagnosis returns a boolean if a field has been set.

### GetFirstFillDate

`func (o *MedicationPrescribed) GetFirstFillDate() string`

GetFirstFillDate returns the FirstFillDate field if non-nil, zero value otherwise.

### GetFirstFillDateOk

`func (o *MedicationPrescribed) GetFirstFillDateOk() (*string, bool)`

GetFirstFillDateOk returns a tuple with the FirstFillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFillDate

`func (o *MedicationPrescribed) SetFirstFillDate(v string)`

SetFirstFillDate sets FirstFillDate field to given value.

### HasFirstFillDate

`func (o *MedicationPrescribed) HasFirstFillDate() bool`

HasFirstFillDate returns a boolean if a field has been set.

### GetLastFillDate

`func (o *MedicationPrescribed) GetLastFillDate() string`

GetLastFillDate returns the LastFillDate field if non-nil, zero value otherwise.

### GetLastFillDateOk

`func (o *MedicationPrescribed) GetLastFillDateOk() (*string, bool)`

GetLastFillDateOk returns a tuple with the LastFillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFillDate

`func (o *MedicationPrescribed) SetLastFillDate(v string)`

SetLastFillDate sets LastFillDate field to given value.

### HasLastFillDate

`func (o *MedicationPrescribed) HasLastFillDate() bool`

HasLastFillDate returns a boolean if a field has been set.

### GetNote

`func (o *MedicationPrescribed) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *MedicationPrescribed) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *MedicationPrescribed) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *MedicationPrescribed) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


