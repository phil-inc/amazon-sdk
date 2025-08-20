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

// StrengthFormCode Strength Form Code (Reference: https://evs.nci.nih.gov/ftp1/NCPDP/NCPDP.txt)  `C78746` - A drug pack that contains a 21 day supply of tablets, composed of a mixture of active and/or inert ingredient(s) that are pressed or compacted into a solid round, square or oval shaped form.  `C78747` - A drug pack that contains a 28 day supply of tablets, composed of a mixture of active and/or inert ingredient(s) that are pressed or compacted into a solid round, square or oval shaped form.  `C107670` - A syringe with a measured dose of medication that has an apparatus that enables the user to specify the amount of the medication to deliver in an injection.  `C64886` - A suppository intended for administration to adults.  `C42887` - Contents packaged under pressure that are released upon activation of an appropriate valve system.  `C42888` - A foam that is packaged under pressure and released upon activation of an appropriate valve system.  `C68935` - A mist that is packaged under pressure and released upon activation of an appropriate valve system.  `C69030` - A homogenous solution comingled with a propellant and packaged under pressure. Upon activation of an appropriate valve system, the liquid is expelled into the air, the propellant is vaporized and the ingredients are dispersed.  `C42889` - A spray that is packaged under pressure and released upon activation of an appropriate valve system.  `C60897` - A cream that is modified to enhance the release of active and/or inert ingredient(s).  `C91136` - A gel that is modified to enhance the release of active and/or inert ingredient(s).  `C60957` - A lotion that is modified to enhance the release of active and/or inert ingredient(s).  `C60984` - An ointment that is modified to enhance the release of active and/or inert ingredient(s).  `C91227` - Active and/or inert ingredient(s) in a device that is automated to deliver a designated dose.  `C69012` - A solid composed of single or multilayered strips of permeable material.  `C68950` - A solid or semi-solid in a singular unit that is rectangular or oval in shape.  `C91137` - A solid in the shape of a bar composed of active and/or inert ingredient(s) in an anionic surfactant, for cleansing.  `C42890` - A solid in the shape of a small sphere.  `C64875` - A capsule designed to be bitten to release the active and/or inert ingredient(s).  `C42891` - A solid or semi-solid in the shape of a square or rectangle.  `C64874` - A capsule intended for administration between the cheek and gum of the oral cavity. The active and/or inert ingredient(s) are usually absorbed through the oral mucosa.  `C91138` - A thin layer of water soluble or insoluble polymer intended to coat the inner surface of the cheek.  `C42755` - A tablet intended for administration between the cheek and gum of the oral cavity. The active and/or inert ingredient(s) are usually absorbed through the oral mucosa.  `C64878` - A solid composed of dry powdered active and/or inert ingredient(s) covered by an edible shell. The shell consists of two concave pieces of wafer made of flour and water.  `C68951` - A solid composed of active ingredient(s), excipients (usually electrolytes and bulking agents), and a solvent system which are lyophilized to yield a compressed solid with uniform ingredient distribution.  `C64880` - A solid or semi-solid with a sweet taste and intended for oral administration.  `C25158` - A solid contained within either a hard or soft soluble shell, usually prepared from gelatin.  `C68943` - A capsule designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a 12 hour time interval.  `C68944` - A capsule designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a 24 hour time interval.  `C45414` - A base composed of bonding agents intermingled with active and/or inert ingredient(s), which rapidly hardens into a firm mass upon application.  `C42892` - A bar designed to be chewed to release active and/or inert ingredient(s).  `C64876` - A capsule designed to be chewed to release the active and/or inert ingredient(s).  `C42893` - A tablet that must be chewed to release the active and/or inert ingredient(s).  `C42894` - A semi-solid composed of synthetic, polymerized polysaccharide and flavorings, intended to be chewed to release active and/or inert ingredient(s).  `C42678` - A solid composed of active and/or inert ingredient(s) wrapped in a paper cylinder whereby one end is burned. The ingredient is released in a smoke form, which is inhaled.  `C60884` - A solid composed active and/or inert ingredient(s) in a woven absorbent material.  `C42895` - A capsule that contains an additional outer layer over the capsule shell.  `C42896` - A capsule composed of small globular masses covered with an outer shell.  `C42897` - A tablet whose outer shell is covered with a substance.  `C69002` - A tablet that has been compacted into capsule shape and covered in sugar.  `C60891` - A solution or suspension composed of active and/or inert ingredient(s) with increased strength in a reduced volume, which is usually diluted prior to administration.  `C42899` - A concentrate which is usually diluted prior to administration by injection.  `C69001` - A concentrate that usually requires dilution prior to oral administration.  `C42898` - A concentrate which is usually diluted prior to administration and exists as a solution in its final form.  `C42900` - A solid with a shape bounded by a circular base converging upon a single vertex.  `C42731` - A solid, semi-solid, solution or suspension designed to release active and/or inert ingredient(s) at a controlled rate.  `C69024` - A capsule designed to release active and/or inert ingredient(s) at a controlled rate.  `C69026` - A solution or suspension designed to release active and/or inert ingredient(s) at a controlled rate.  `C69025` - A tablet designed to release active and/or inert ingredient(s) at a controlled rate.  `C28944` - A semi-solid composed of an emulsion of lipids, hydrocarbons, waxes, or polyols in water.  `C42901` - A solid that consists of naturally produced angular solids composed of singular, repeating units that are systematically arranged in an evenly spaced lattice.  `C64881` - A solid or semi-solid in the shape of a three dimensional square.  `C45415` - A solution, suspension or semi-solid composed of media conducive to microbial growth.  `C68949` - A capsule designed to dissolve after releasing the active and/or inert ingredient(s) at a controlled rate.  `C42730` - A solid, semi-solid, solution or suspension that has been coated with a substance that is designed to impede the immediate release of the active and/or inert ingredient(s) after administration.  `C42902` - A capsule covered with a substance that is designed to impede the immediate release of the active and/or inert ingredient(s) after administration.  `C42903` - A solid composed of small particles or grains and have been coated with a substance that is designed to impede the immediate release of the active and/or inert ingredient(s) after administration.  `C42997` - A tablet composed of individual fragments have been coated with a substance that is designed to impede the immediate release of the active and/or inert ingredient(s) after administration.  `C42904` - A capsule composed of active and/or inert ingredient(s) small, spherical structures that have been coated with a substance that is designed to impede the immediate release of the active and/or inert ingredient(s) after administration.  `C42905` - A tablet that has been coated with a substance that is designed to impede the immediate release of the active and/or inert ingredient(s) after administration.  `C69059` - An substance intended for use on the teeth or gingiva in the oral cavity.  `C68954` - A solid bounded by a circular base converging upon a single vertex and intended for administration to the teeth and gingiva in the oral cavity.  `C45413` - A resinous film-forming agent dissolved in a solvent or a calcium hydroxide suspension within a synthetic resin solution, which rapidly hardens into a firm mass upon administration.  `C42906` - A gel intended for use in the oral cavity to clean the teeth and maintain oral hygiene.  `C42907` - A paste of mild abrasives, detergents, flavoring agents, binders, fluoride, and other active and/or inert ingredient(s).  `C42908` - A powder composed of a mild abrasives, detergents, flavoring agents, binders, fluoride, and other active and/or inert ingredient(s).  `C42740` - A solid, semi-solid, solution or suspension covered with an outer permeable polymeric membrane designed to release active and/or inert ingredient(s) at a controlled, prolonged rate so as to reduce dosing frequency.  `C43525` - A flat, circular shaped form composed of a solid or powder that contains active and/or inert ingredient(s).  `C69071` - A tablet that dissolves or melts quickly (usually within a matter of seconds) when it comes into contact with a liquid.  `C42756` - A tablet composed of a large amount of active and/or inert ingredient(s) that can be used in compounding prescriptions to ensure accurate quantity of ingredient within the final, compounded form.  `C42741` - A solid, semi-solid, solution or suspension composed of active and/or inert ingredient(s) covered with or contained within a polymeric matrix such that dissolution is determined by the solubility rate of the given polymer.  `C42679` - A substance intended for use as a vaginal irrigant.  `C69033` - A powder which may be mixed with a liquid, intended for administration as a vaginal irrigant.  `C69032` - A solution intended for use as a vaginal irrigant.  `C42763` - A substance intended for administration in or on a wound.  `C29012` - A substance intended for administration in a drop-wise fashion.  `C60992` - A solution intended for administration in a drop-wise fashion.  `C60995` - A suspension intended for administration in a drop-wise fashion.  `C68997` - A suspension prepared immediately prior to dispensing or administration in a drop-wise fashion.  `C64883` - A powder intended to be dusted over the outer surface of the body.  `C42909` - A granule composed of active and/or inert ingredient(s), and a mixture of acids and sodium bicarbonate, which release carbon dioxide when dissolved in water.  `C64884` - A powder composed of active and/or inert ingredient(s), and a mixture of acids and sodium bicarbonate, which release carbon dioxide when dissolved in water.  `C42910` - A tablet composed of active and/or inert ingredient(s), and a mixture of acids and sodium bicarbonate, which release carbon dioxide when dissolved in water.  `C42911` - A patch that composed of active and/or inert ingredient(s) whose release is controlled by electronic impulses that release ingredients at a controlled, prolonged rate so as to reduce dosing frequency.  `C42912` - A solution composed of a clear, sweetened hydroalcoholic liquid in which active and/or inert ingredient(s) are dissolved.  `C42913` - A semi-solid composed of at least two immiscible liquids, one of which is dispersed as droplets within the other liquid, and stabilized with one or more emulsifying agents.  `C42914` - An emulsion intended for injection.  `C42915` - A solution or suspension intended for administration into the rectum.  `C64885` - A powder that is suspended in liquid prior to administration to the rectum.  `C64871` - A tablet that is suspended in liquid prior to administration to the rectum.  `C68945` - A capsule covered with a substance that is designed to delay the release of active and/or inert ingredient(s) until the capsule passes into the intestines.  `C42758` - A tablet covered with a substance that is designed to delay the release of active and/or inert ingredient(s) until the tablet passes into the intestines.  `C42742` - A solid, semi-solid, solution or suspension in which active and/or inert ingredient(s) release is controlled by the erosion rate of a carrier matrix.  `C42713` - A solid, semi-solid, solution or suspension designed to release active and/or inert ingredient(s) at a controlled, prolonged rate so as to reduce dosing frequency.  `C43451` - A solid in the shape of a small sphere intended to be embedded securely in the body and designed to release active and/or inert ingredient(s) at a controlled, prolonged rate.  `C42916` - A capsule designed to release active and/or inert ingredient(s) at a controlled, prolonged rate so as to reduce dosing frequency.  `C42917` - A capsule covered with a thin layer of substance and designed to release active and/or inert ingredient(s) at a controlled, prolonged rate.  `C42918` - A pellet covered with a thin layer of substance designed to release active and/or inert ingredient(s) at a controlled, prolonged rate.  `C42919` - A solid or liquid core composed of active and/or inert ingredient(s) enclosed by a polymer coated shell for controlled release.  `C91140` - An extended release capsule covered with a substance designed to delay release until the capsule passes into the intestines.  `C91141` - An extended release tablet covered with a substance designed to delay release until the tablet passes into the intestines.  `C60926` - A solid in the shape of a slender, elongated thread designed for controlled release.  `C42920` - A film designed to release active and/or inert ingredient(s) at a controlled, prolonged rate.  `C42928` - A capsule covered with a thin layer of water soluble or insoluble polymer for controlled release.  `C60929` - A substance that forms a suspension upon reconstitution for controlled release.  `C42935` - A solution that becomes gelatinous upon administration for controlled release.  `C69067` - A solid in the shape of a small particle or grain designed for controlled release.  `C42921` - A granule intended for administration as a suspension for controlled release.  `C42922` - An insert designed for controlled release.  `C60934` - A solution or suspension designed for controlled release.  `C42923` - A patch designed for controlled release.  `C42924` - A suppository designed for controlled release.  `C42925` - A suspension designed for controlled release.  `C42927` - A tablet designed for controlled release.  `C42929` - A distillate composed of active and/or inert ingredient(s) extracted from plant or animal matter.  `C42932` - A thin layer of water soluble or insoluble polymer for site administration.  `C42930` - A tablet covered with a thin layer of water soluble or insoluble polymer for controlled release.  `C42931` - A tablet covered with a thin layer of water soluble or insoluble polymer.  `C68982` - A solid composed of active and/or inert ingredient(s) resembling small, thin chips.  `C68991` - A solution or suspension distillate extracted from plant or animal matter with alcohol as solvent/preservative.  `C64898` - An aerated solution or suspension.  `C64899` - A solution or suspension that produces dense bubbles when combined with water.  `C60927` - A substance intended for administration as a solution.  `C60928` - A substance intended for administration as a suspension.  `C68966` - A solution reconstituted from frozen pre-mixed form for intravenous administration.  `C78748` - A solution or suspension for oral gargling.  `C42933` - An elastic aeriform fluid with separated molecules having freedom of movement.  `C42934` - A semi-solid composed of a cross-linked matrix of polymers within a liquid.  `C60994` - A solution for drop-wise administration that becomes gelatinous.  `C68973` - A solution that becomes gelatinous after administration.  `C42936` - A capsule covered with a thin layer of gelatin.  `C64872` - A tablet covered with a layer of gelatin.  `C48193` - An apparatus that converts substances into vapor or gas by heat or chemical reaction.  `C42937` - A solid composed of polysaccharide pellets impregnated with active/inert ingredient(s).  `C45416` - A solid composed of tissue for implantation.  `C42938` - A solid composed of small particles or grains.  `C69066` - A solid particle/grain mixed with liquid for solution/suspension.  `C42939` - A granule for solution administration.  `C42940` - A granule for suspension administration.  `C42941` - A semi-solid composed of sticky, mucilaginous polysaccharide.  `C64904` - A capsule covered with a rigid outer shell.  `C64882` - A solid composed of uniform homeopathic polysaccharide pellets.  `C42752` - A tablet for hypodermic injection.  `C42669` - A form designed for immediate release upon administration.  `C42942` - A solid for secure body insertion.  `C42943` - A pellet for secure body insertion.  `C42944` - Vaporized ingredients for inhalation.  `C91142` - Powdered medication for measured dose inhalation.  `C91143` - Liquid medication for inhalation.  `C64879` - A capsule of powders for oral inhalation.  `C42946` - A substance for injection.  `C42926` - A suspension for injection with controlled release.  `C42950` - A substance with lipid-complexed ingredients for injection.  `C42951` - A suspension with liposome-encapsulated ingredients for injection.  `C69037` - A freeze-dried powder for injection after reconstitution.  `C42945` - A solution for injection.  `C42988` - A sonicated suspension for injection.  `C42995` - A suspension for injection.  `C60933` - A solid or semi-solid intended for administration within the body.  `C78793` - A powder intended for use within the body.  `C68971` - A solution intended for administration into the intraperitoneal cavity.  `C47915` - A solid intended for placement within the uterine cavity where active and/or inert ingredient(s) are released from an apparatus at a controlled rate.  `C68967` - A solution intended for intravenous administration simultaneously with an established compatible maintenance infusion drug by a mutual intravenous access.  `C68965` - A solution composed of a homogenous liquid that contains active and/or inert ingredient(s) dissolved in a suitable solvent or mixture of mutually miscible solvents.  `C42947` - A solution or suspension used to bathe or flush open wounds or body cavities.  `C42948` - A semi-solid composed of suspensions of active and/or inert ingredient(s) in any type of highly viscous material, including gelatin or gel. The structural matrix of the jelly contains high levels of liquid.  `C47916` - A collection of unrelated materials that are used together to orchestrate dosage administration.  `C42949` - A semi-solid with viscosity similar to that of a lotion, which is intended for topical use and must be rubbed into the site of administration to release the active and/or inert ingredient(s).  `C60931` - A substance composed of active and/or inert ingredient(s) encapsulated in liposomes and intended for injection.  `C42952` - A solid or semi-solid with a waxy consistency in the shape of a stick and intended for use on the lips.  `C42953` - A substance in the fluid state of matter having no fixed shape but a fixed volume.  `C42954` - A capsule composed of liquid covered with an outer shell, which is usually prepared from gelatin.  `C68953` - A solution or suspension composed of an anionic surfactant used for cleansing.  `C69068` - A solid composed of a sugared, medicated candy mounted to a stick, which can be held while sucking or chewing on the candy.  `C29167` - A semi-solid composed of an oil in water emulsion, with lower viscosity than cream or gel.  `C60958` - A lotion composed of anionic surfactant intended for administration to the hair or scalp.  `C42955` - A solid composed of active and/or inert ingredient(s), sweeteners, other flavorings and mucilage, intended for oral administration.  `C42956` - A freeze dried powder intended for injection following reconstitution with a suitable solvent to form a suspension that is designed to release active and/or inert ingredient(s) at a controlled, prolonged rate so as to reduce dosing frequency.  `C42959` - A freeze dried powder composed of active and/or inert ingredient(s) encapsulated in liposomes and intended for injection following reconstitution with a suitable solvent to form a suspension.  `C42957` - A freeze dried powder intended for injection following reconstitution with a suitable solvent to form a solution.  `C42958` - A freeze dried powder intended for injection following reconstitution with a suitable solvent to form a suspension.  `C68988` - A solid composed of an impermeable occlusive backing, a formulation matrix in which the active and/or inert ingredient(s) are dissolved or dispersed and an adhesive layer. It is intended for external application.  `C91144` - A solid in the shape of a bar composed of active and inert ingredient(s) in an anionic surfactant, for cleansing.  `C68958` - A film composed of active and/or inert ingredient(s).  `C91145` - A liquid composed of active and inert ingredient(s) in an anionic surfactant, for cleansing.  `C68957` - A solid composed of active and/or inert ingredient(s) in a soft, non-adhesive and absorbent piece of fabric or other material.  `C91146` - A solution or suspension composed of active and inert ingredient(s) in an anionic surfactant and intended for administration to the hair and/or scalp.  `C68952` - A substance composed of active and/or inert ingredient(s) in an anionic surfactant, for cleansing.  `C64901` - A solid composed of a porous, interlacing, absorbent, usually shape retaining material that contains active and/or inert ingredient(s).  `C68955` - A solid composed of active and/or inert ingredient(s) on a small piece of absorbent material attached to one end of a stick.  `C91147` - A fabric or a narrow extruded synthetic material usually with an adhesive on one or both sides, impregnated with active and inert ingredient(s).  `C64877` - A capsule covered with a semipermeable film.  `C64873` - A tablet covered with a semipermeable film.  `C42960` - An aerosol contained in a device with valves that permits the release of a specified quantity of active and/or inert ingredient(s) upon activation.  `C91148` - A liquid medication delivered as a mist to be inhaled as a measured dose  `C60930` - A gel contained in a device with valves that permits the release of a uniform dose of gel that composed of active and/or inert ingredient(s) upon activation.  `C42961` - A powder contained in a device with valves that permits the dispensing of active and/or inert ingredient(s) upon activation.  `C42962` - A spray contained in a device with valves that permits the dispensing of a specified quantity of active and/or inert ingredient(s) upon activation.  `C64888` - A solution or suspension present in a small volume or mass, intended for administration to the rectum.  `C42712` - A solid, semi-solid, solution or suspension exhibiting an altered inherent rate of release of active and/or inert ingredient(s).  `C29269` - A solution or suspension intended for oral administration, for its deodorizing or antiseptic effects in the oral cavity. Usually, the mouthwash is swished around the mouth and expelled.  `C91149` - A spray intended for administration to the mucosa.  `C91150` - A solution intended for administration to the surface of the mucosa membranes.  `C42963` - A tablet composed of active and/or inert ingredient(s) that have been compressed into multiple layers and is designed to release active and/or inert ingredient(s) at a controlled, prolonged rate so as to reduce dosing frequency.  `C42964` - A tablet composed of active and/or inert ingredient(s) that have been compressed to form a tablet that contains multiple layers.  `C69064` - A substance intended for administration in the nares.  `C91151` - A cream intended for administration to the mucosa of the nose.  `C91152` - A gel intended for administration to the mucosa of the nose.  `C91153` - An inhalant intended for administration through the nose.  `C91154` - A medication delivered as a mist to be inhaled through the nose as a measured dose  `C91155` - An ointment intended for administration to the mucosa of the nose.  `C91156` - A solution intended for administration to the mucosa of the nose.  `C91157` - A spray intended for administration to the mucosa of the nose.  `C91158` - A suspension intended for administration to the mucosa of the nose.  `C68941` - A spray composed of active and/or inert ingredient(s) that are expelled from a delivery device without the use of a propellant.  `C69017` - A solid composed of a small piece of absorbent material that does not contain active ingredient(s), attached to one end of a small stick.  `C48624` - An indication that the dosage form is not relevant or appropriate.  `C69031` - A system that is placed in the eye, usually under the lower eyelid, from which the active and/or inert ingredient(s) diffuses through a membrane at a controlled rate.  `C42965` - A solution or suspension that is hydrophobic, soluble in organic solvents, and liquid at ambient temperature.  `C42966` - A semi-solid, viscous in texture, that may be composed of a variety of bases including hydrocarbons, emulsifiers or vegetable oils, and mixed with active and/or inert ingredient(s).  `C69039` - A substance intended for administration in or around the eye.  `C91159` - A cream intended for administration in or around the eye.  `C91160` - A gel intended for administration in or around the eye.  `C91161` - A solution intended for irrigation of the eye.  `C69038` - A solution or suspension intended for administration in the eye.  `C91162` - A ointment intended for administration in or around the eye.  `C91163` - A solution intended for administration in or around the eye.  `C91164` - A suspension intended for administration in or around the eye.  `C42744` - A substance intended for administration through the mouth.  `C91165` - A capsule intended for administration through the mouth.  `C91166` - A cream intended for administration through the mouth.  `C91167` - A foam intended for administration through the mouth.  `C91168` - A gel intended for administration through the mouth.  `C91169` - An ointment intended for administration through the mouth.  `C91170` - A paste intended for administration through the mouth.  `C91171` - A powder intended for administration through the mouth.  `C68981` - A suspension composed of active and/or inert ingredient(s) that requires reconstitution and is intended oral administration.  `C68996` - A solution intended for oral administration.  `C91172` - A spray intended for administration into the mouth.  `C91173` - A strip intended for administration through the mouth.  `C68992` - A suspension intended for oral administration.  `C68993` - A suspension composed of a suspension that has been prepared immediately prior to dispensing or oral administration.  `C43243` - A tablet intended for oral administration.  `C61006` - A tablet that disintegrates rapidly, usually within a matter of seconds, when placed upon the tongue but is also designed to impede the immediate release of the active and/or inert ingredient(s) after administration.  `C42999` - A tablet which disintegrates and dissolves quickly (usually within a matter of seconds) in the mouth. Additional water to complete dissolution is not necessary.  `C69008` - A tablet with a semi-permeable membrane composed of active and/or inert ingredient(s) and an osmotic agent with a laser drilled central core. Upon ingestion, water permeates the core and the ingredient is dissolved and released.  `C42760` - A solid designed to release active and/or inert ingredient(s) at a controlled, prolonged rate so as to reduce dosing frequency, based on osmotic pressure between the inside of the dosage form and the external environment.  `C69040` - A substance intended for administration either on the outer ear or into the auditory canal.  `C91174` - A cream intended for administration either on the outer ear or into the auditory canal.  `C91175` - An ointment intended for administration either on the outer ear or into the auditory canal.  `C91176` - A solution intended for administration either on the outer ear or into the auditory canal.  `C91177` - A suspension intended for administration either on the outer ear or into the auditory canal.  `C62653` - A number of individual items packaged as a unit.  `C47887` - A solid composed of a material that contains active and/or inert ingredient(s) which is intended to be placed in a cavity or depression within the body.  `C69016` - A solid composed of a soft, non-adhesive and absorbent piece of fabric or other material.  `C42746` - A substance intended for administration by injection.  `C42967` - A semi-solid composed of a large proportion of solids and finely dispersed active and/or inert ingredient(s) in a fat-based vehicle.  `C60985` - A solid or semi-solid composed of a lozenge, that contains active and/or inert ingredient(s) that dissolves when sucked in the mouth.  `C42968` - A solid composed of an impermeable occlusive backing and a formulation matrix in which the active and/or inert ingredient(s) are dissolved or dispersed; possibly includes an adhesive layer.  `C69042` - A solution or suspension intended for use in children.  `C64887` - A suppository intended for administration to children.  `C42969` - A solid composed of a small granular or compressed mass that contains highly purified active and/or inert ingredient(s).  `C42636` - The form in which active and/or inert ingredient(s) are physically presented.  `C25394` - A solid composed of a small, round object composed of active and/or inert ingredient(s).  `C42970` - A semi-solid composed of a soft, smooth mass with relatively thick consistency and strong adhesiveness that becomes solid when dry.  `C42736` - Polymeric microspheres composed of biocompatible polymer that contains active and/or inert ingredient(s) designed for release at a controlled rate.  `C47913` - A semi-solid composed of a soft, moist mass that is made by wetting an absorbent solid substance.  `C42972` - A solid composed of a mixture of dry, finely divided active and/or inert ingredient(s).  `C42971` - A powder composed of active and/or inert ingredient(s) and packaged under pressure. Ingredients are released upon activation of an appropriate valve system.  `C42977` - A powder intended for injection following reconstitution to form a suspension that is designed to release active and/or inert ingredient(s) at a controlled, prolonged rate so as to reduce dosing frequency.  `C42974` - A powder intended for injection following reconstitution to form a solution.  `C42976` - A powder intended for injection following reconstitution to form a suspension.  `C69069` - A powder intended for injection following reconstitution to form a solution or suspension.  `C64907` - A powder that yields a solution intended for oral administration following reconstitution.  `C64908` - A powder that yields a suspension intended for oral administration following reconstitution.  `C69070` - A powder mixed with a liquid to form a solution or suspension before administration.  `C42973` - A powder that yields a solution following reconstitution with an appropriate liquid.  `C42975` - A powder that forms a suspension when reconstituted with a liquid solution.  `C91139` - Powder intended for oral inhalation.  `C68983` - A granule composed of a mixture of small (powder like), dry particles or grains, which does not have the ability to release gas upon contact with water.  `C91179` - Active and/or inert ingredients in an applicator delivered as a single measured dose.  `C97717` - A syringe that lacks a conventional plunger, resembles a writing pen, and is designed to dispense a pre-loaded dose of a drug. It may be designed to deliver a single dose or be designed for repeated use.  `C91180` - A solution or suspension intended to be injected in a syringe, cartridge or pen.  `C68972` - A semi-solid, having a relatively thick consistency, intended for oral administration to facilitate ease of swallowing.  `C68984` - A substance that requires reconstitution with a suitable solvent and is intended for oral administration in a drop-wise fashion.  `C68985` - A solution that requires reconstitution with a suitable solvent, and is intended for oral administration.  `C69046` - A substance intended for administration in or around the rectum.  `C69045` - A cream intended for administration in or around the rectum.  `C91181` - A foam intended for administration to the rectum.  `C91182` - A gel intended for administration to the rectum.  `C69047` - An ointment intended for administration in or around the rectum.  `C91183` - A powder intended for administration to the rectum.  `C91184` - A spray intended for administration to the rectum.  `C68989` - A suppository intended for administration within the rectum.  `C42978` - A semi-solid composed of a mixture of gum and a plant based hydrocarbon.  `C60988` - A solid composed active and/or inert ingredient(s) of in the shape of a hoop.  `C42979` - A solution or suspension intended for irrigation purposes.  `C42980` - A semi-solid composed of a thick, fat or wax-based ointment or cerate with a consistency somewhere between an ointment and a plaster.  `C45244` - A substance with properties of both a solid and a liquid in terms of viscosity and rigidity. Semi-solids include, but are not limited to, ointments, creams, pastes and gels.  `C42981` - A solution or suspension composed of an anionic surfactant and intended for administration to the hair and/or scalp.  `C42982` - A suspension composed of anionic surfactant and intended for administration to the hair and/or scalp.  `C28276` - A patch intended for administration to the skin.  `C42983` - An anionic surfactant for cleansing.  `C64909` - A capsule composed of active and/or inert ingredient(s) covered with a soft outer shell, which is usually prepared from gelatin.  `C45235` - A substance having definite shape and volume manufactured for the administration of active and/or inert ingredient(s). Solids include tablets, capsules, powders, granules and certain suppositories.  `C42984` - A film that will dissolve in a liquid solvent to form a solution.  `C42985` - A tablet that will dissolve in a liquid solvent to form a solution.  `C42986` - A clear, homogeneous liquid composed of one or more chemical substances dissolved in a solvent or mixture of mutually miscible solvents.  `C69028` - A solution intended for reconstitution before administration.  `C42987` - A solution composed of an iced saline slurry with or without active and/or inert ingredient(s) intended to induce regional hypothermia from the point of administration.  `C47912` - A solid composed of a porous, interlacing, absorbent, usually shape retaining material, such as cellulose wood fibers or plastic polymer form.  `C42989` - A solution or suspension composed of active and/or inert ingredient(s) in oil or water, typically dispensed from an atomizer or nebulizer.  `C42990` - A spray composed of a suspension that contains active and/or inert ingredient(s).  `C68946` - A capsule composed active and/or inert ingredient(s) of as small pellet-like granules. The capsule is intended to be broken open, facilitating the scattering of the granules into soft foods, although it may also be swallowed whole.  `C42991` - A solid composed active and/or inert ingredient(s) in a long, slender cylindrical shape.  `C47914` - A solid composed of active and/or inert ingredient(s) in a long narrow piece of material.  `C42751` - A tablet intended for administration under the tongue.  `C42992` - A tablet covered with sugar.  `C42993` - A solid or semi-solid composed of active and/or inert ingredient(s) in a wax, fat or a glycerin gelatin jelly and that is conical or oval in shape. It is intended to be inserted into a body orifice.  `C42994` - Insoluble solid particles composed of active and/or inert ingredient(s) that are dispersed in a liquid.  `C68986` - A suspension designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a 12 hour time interval.  `C69027` - A suspension mixed with a liquid to form a solution or suspension before administration.  `C42672` - A dosage form designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a period of time.  `C69007` - A tablet intended for administration between the cheek and gum of the oral cavity. The active and/or inert ingredient(s) are usually absorbed through the oral mucosa and designed to be released slowly over a period of time.  `C69043` - A capsule designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a period of time.  `C69044` - A solution or suspension intended for oral administration and designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a period of time.  `C69036` - An injectable solution or suspension designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a period of time.  `C68947` - A capsule composed of active and/or inert ingredient(s) small, spherical structures that contain active and/or inert ingredient(s) encased within a shell. It is designed to release ingredients slowly so as to achieve a constant circulating concentration of the ingredient over a period of time.  `C78751` - A tablet designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a period of time.  `C47889` - A solid composed of fibrous strands of material that are used to bind wound edges together.  `C47898` - A solid composed of a small piece of absorbent material attached to one end of a small stick.  `C42996` - A solution or suspension composed of a viscid vehicle that contains a high concentration of sucrose or other sugars and active and/or inert ingredient(s).  `C42998` - A solid composed of a mixture of that active and/or inert ingredient(s) are pressed or compacted together, usually in the form of a relatively flat and round, square or oval shape.  `C69004` - A tablet designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a 12 hour time interval.  `C69003` - A tablet designed to release active and/or inert ingredient(s) slowly so as to achieve a constant circulating concentration of the ingredient over a 24 hour time interval.  `C60997` - A tablet composed of a conglomerate of particles that contains active and/or inert ingredient(s) that have been individually covered with a coating.  `C61004` - A tablet that changes into a solution when added to a liquid solvent.  `C61005` - A tablet that changes into a suspension when added to a liquid solution.  `C69006` - A tablet composed of small fragments of singular, repeating units of that are systematically arranged over an evenly spaced lattice.  `C42759` - A compressed cylindrical tablet composed of medicated powder and dispersed with a mixture of lactose, powdered sucrose and a moistening agent.  `C47892` - A solid composed of a plug made from absorbent material.  `C47897` - A solid composed of narrow woven fabric or a narrow extruded synthetic material, usually with an adhesive on one or both sides.  `C78749` - A semi-solid that is oily and viscous and composed of a hydrocarbon based resin formed from the breakdown and distillation of organic substances.  `C69063` - A spray intended for administration to the mucosa of the throat.  `C43000` - A solution composed of an alcoholic extract or solution of nonvolatile active and/or inert ingredient(s).  `C91186` - A paste intended for administration to the teeth.  `C42747` - A substance intended for administration to a body surface.  `C91187` - A cream intended for administration to a body surface.  `C91188` - A foam intended for administration to a body surface.  `C91189` - A gel intended for administration to a body surface.  `C91190` - A lotion intended for administration to a body surface.  `C91191` - A oil intended for administration to a body surface.  `C91192` - A ointment intended for administration to a body surface.  `C91193` - A powder intended for administration to a body surface.  `C91178` - A powder spray intended for administration to the skin.  `C64905` - A solution intended for administration to a body surface.  `C91194` - A spray intended for topical administration.  `C68998` - A suspension intended for administration to a body surface.  `C43001` - A solid discoid composed of active and/or inert ingredient(s) in a suitably flavored base, that dissolves when sucked in the mouth.  `C43002` - An indication that the dosage form has yet to be assigned.  `C64902` - A solid composed of a porous, interlacing, absorbent, usually shape retaining material that does not contain active ingredient(s).  `C38046` - Not stated explicitly or in detail.  `C69052` - A substance intended for administration into the urethra.  `C91195` - A gel intended for administration into the urethra.  `C69053` - A suppository intended for administration within the urethra.  `C69048` - A substance intended for administration in or around the vagina.  `C69050` - A cream intended for administration in or around the vagina.  `C47890` - A dome-shaped insert with a springy, flexible rim intended to be inserted into the vagina and positioned behind the pelvic bone to completely cover the cervix.  `C69051` - A foam composed of active and/or inert ingredient(s), intended for administration in or around the vagina.  `C91197` - A gel composed of active and/or inert ingredient(s), intended for administration in or around the vagina.  `C69054` - An ointment intended for administration in or around the vagina.  `C91198` - A powder composed of active and/or inert ingredient(s), intended for administration in or around the vagina.  `C91199` - A ring composed of active and/or inert ingredient(s), intended for administration in or around the vagina.  `C91200` - A spray composed of active and/or inert ingredient(s), intended for administration in or around the vagina.  `C68990` - A suppository intended for administration within the vagina.  `C69049` - A tablet intended for vaginal administration.  `C43003` - A solid composed of a thin slice of material that contain active and/or inert ingredient(s).  `C64903` - A solution or suspension composed of active and/or inert ingredient(s) and intended for administration as an irrigant and cleanser.  `C78750` - A semi-solid composed of lipids made up of hydrocarbons or esters of fatty acids; active and/or inert ingredient(s) are commonly suspended within the wax matrix layer.  `C64910` - A relatively long and often cylindrical solid composed of active and/or inert ingredient(s) intended for wound application.
type StrengthFormCode string

// List of StrengthFormCode
const (
	C78746   StrengthFormCode = "C78746"
	C78747   StrengthFormCode = "C78747"
	C107670  StrengthFormCode = "C107670"
	C64886   StrengthFormCode = "C64886"
	C42887   StrengthFormCode = "C42887"
	C42888   StrengthFormCode = "C42888"
	C68935   StrengthFormCode = "C68935"
	C69030   StrengthFormCode = "C69030"
	C42889   StrengthFormCode = "C42889"
	C60897   StrengthFormCode = "C60897"
	C91136   StrengthFormCode = "C91136"
	C60957   StrengthFormCode = "C60957"
	C60984   StrengthFormCode = "C60984"
	C91227   StrengthFormCode = "C91227"
	C69012   StrengthFormCode = "C69012"
	C68950   StrengthFormCode = "C68950"
	C91137   StrengthFormCode = "C91137"
	C42890   StrengthFormCode = "C42890"
	C64875   StrengthFormCode = "C64875"
	C42891   StrengthFormCode = "C42891"
	C64874   StrengthFormCode = "C64874"
	C91138   StrengthFormCode = "C91138"
	C42755   StrengthFormCode = "C42755"
	C64878   StrengthFormCode = "C64878"
	C68951   StrengthFormCode = "C68951"
	C64880   StrengthFormCode = "C64880"
	C25158   StrengthFormCode = "C25158"
	C68943   StrengthFormCode = "C68943"
	C68944   StrengthFormCode = "C68944"
	C45414   StrengthFormCode = "C45414"
	C42892   StrengthFormCode = "C42892"
	C64876   StrengthFormCode = "C64876"
	C42893   StrengthFormCode = "C42893"
	C42894   StrengthFormCode = "C42894"
	C42678   StrengthFormCode = "C42678"
	C60884   StrengthFormCode = "C60884"
	C42895   StrengthFormCode = "C42895"
	C42896   StrengthFormCode = "C42896"
	C42897   StrengthFormCode = "C42897"
	C69002   StrengthFormCode = "C69002"
	C60891   StrengthFormCode = "C60891"
	C42899   StrengthFormCode = "C42899"
	C69001   StrengthFormCode = "C69001"
	C42898   StrengthFormCode = "C42898"
	C42900   StrengthFormCode = "C42900"
	C42731   StrengthFormCode = "C42731"
	C69024   StrengthFormCode = "C69024"
	C69026   StrengthFormCode = "C69026"
	C69025   StrengthFormCode = "C69025"
	C28944   StrengthFormCode = "C28944"
	C42901   StrengthFormCode = "C42901"
	C64881   StrengthFormCode = "C64881"
	C45415   StrengthFormCode = "C45415"
	C68949   StrengthFormCode = "C68949"
	C42730   StrengthFormCode = "C42730"
	C42902   StrengthFormCode = "C42902"
	C42903   StrengthFormCode = "C42903"
	C42997   StrengthFormCode = "C42997"
	C42904   StrengthFormCode = "C42904"
	C42905   StrengthFormCode = "C42905"
	C69059   StrengthFormCode = "C69059"
	C68954   StrengthFormCode = "C68954"
	C45413   StrengthFormCode = "C45413"
	C42906   StrengthFormCode = "C42906"
	C42907   StrengthFormCode = "C42907"
	C42908   StrengthFormCode = "C42908"
	C42740   StrengthFormCode = "C42740"
	C43525   StrengthFormCode = "C43525"
	C69071   StrengthFormCode = "C69071"
	C42756   StrengthFormCode = "C42756"
	C42741   StrengthFormCode = "C42741"
	C42679   StrengthFormCode = "C42679"
	C69033   StrengthFormCode = "C69033"
	C69032   StrengthFormCode = "C69032"
	C42763   StrengthFormCode = "C42763"
	C29012   StrengthFormCode = "C29012"
	C60992   StrengthFormCode = "C60992"
	C60995   StrengthFormCode = "C60995"
	C68997   StrengthFormCode = "C68997"
	C64883   StrengthFormCode = "C64883"
	C42909   StrengthFormCode = "C42909"
	C64884   StrengthFormCode = "C64884"
	C42910   StrengthFormCode = "C42910"
	C42911   StrengthFormCode = "C42911"
	C42912   StrengthFormCode = "C42912"
	C42913   StrengthFormCode = "C42913"
	C42914   StrengthFormCode = "C42914"
	C42915   StrengthFormCode = "C42915"
	C64885   StrengthFormCode = "C64885"
	C64871   StrengthFormCode = "C64871"
	C68945   StrengthFormCode = "C68945"
	C42758   StrengthFormCode = "C42758"
	C42742   StrengthFormCode = "C42742"
	C42713   StrengthFormCode = "C42713"
	C43451   StrengthFormCode = "C43451"
	C42916   StrengthFormCode = "C42916"
	C42917   StrengthFormCode = "C42917"
	C42918   StrengthFormCode = "C42918"
	C42919   StrengthFormCode = "C42919"
	C91140   StrengthFormCode = "C91140"
	C91141   StrengthFormCode = "C91141"
	C60926   StrengthFormCode = "C60926"
	C42920   StrengthFormCode = "C42920"
	C42928   StrengthFormCode = "C42928"
	C60929   StrengthFormCode = "C60929"
	C42935   StrengthFormCode = "C42935"
	C69067   StrengthFormCode = "C69067"
	C42921   StrengthFormCode = "C42921"
	C42922   StrengthFormCode = "C42922"
	C60934   StrengthFormCode = "C60934"
	C42923   StrengthFormCode = "C42923"
	C42924   StrengthFormCode = "C42924"
	C42925   StrengthFormCode = "C42925"
	C42927   StrengthFormCode = "C42927"
	C42929   StrengthFormCode = "C42929"
	C42932   StrengthFormCode = "C42932"
	C42930   StrengthFormCode = "C42930"
	C42931   StrengthFormCode = "C42931"
	C68982   StrengthFormCode = "C68982"
	C68991   StrengthFormCode = "C68991"
	C64898   StrengthFormCode = "C64898"
	C64899   StrengthFormCode = "C64899"
	C60927   StrengthFormCode = "C60927"
	C60928   StrengthFormCode = "C60928"
	C68966   StrengthFormCode = "C68966"
	C78748   StrengthFormCode = "C78748"
	C42933   StrengthFormCode = "C42933"
	C42934   StrengthFormCode = "C42934"
	C60994   StrengthFormCode = "C60994"
	C68973   StrengthFormCode = "C68973"
	C42936   StrengthFormCode = "C42936"
	C64872   StrengthFormCode = "C64872"
	C48193   StrengthFormCode = "C48193"
	C42937   StrengthFormCode = "C42937"
	C45416   StrengthFormCode = "C45416"
	C42938   StrengthFormCode = "C42938"
	C69066   StrengthFormCode = "C69066"
	C42939   StrengthFormCode = "C42939"
	C42940   StrengthFormCode = "C42940"
	C42941   StrengthFormCode = "C42941"
	C64904   StrengthFormCode = "C64904"
	C64882   StrengthFormCode = "C64882"
	C42752   StrengthFormCode = "C42752"
	C42669   StrengthFormCode = "C42669"
	C42942   StrengthFormCode = "C42942"
	C42943   StrengthFormCode = "C42943"
	C42944   StrengthFormCode = "C42944"
	C91142   StrengthFormCode = "C91142"
	C91143   StrengthFormCode = "C91143"
	C64879   StrengthFormCode = "C64879"
	C42946   StrengthFormCode = "C42946"
	C42926   StrengthFormCode = "C42926"
	C42950   StrengthFormCode = "C42950"
	C42951   StrengthFormCode = "C42951"
	C69037   StrengthFormCode = "C69037"
	C42945   StrengthFormCode = "C42945"
	C42988   StrengthFormCode = "C42988"
	C42995   StrengthFormCode = "C42995"
	C60933   StrengthFormCode = "C60933"
	C78793   StrengthFormCode = "C78793"
	C68971   StrengthFormCode = "C68971"
	C47915   StrengthFormCode = "C47915"
	C68967   StrengthFormCode = "C68967"
	C68965   StrengthFormCode = "C68965"
	C42947   StrengthFormCode = "C42947"
	C42948   StrengthFormCode = "C42948"
	C47916   StrengthFormCode = "C47916"
	C42949   StrengthFormCode = "C42949"
	C60931   StrengthFormCode = "C60931"
	C42952   StrengthFormCode = "C42952"
	C42953   StrengthFormCode = "C42953"
	C42954   StrengthFormCode = "C42954"
	C68953   StrengthFormCode = "C68953"
	C69068   StrengthFormCode = "C69068"
	C29167   StrengthFormCode = "C29167"
	C60958   StrengthFormCode = "C60958"
	C42955   StrengthFormCode = "C42955"
	C42956   StrengthFormCode = "C42956"
	C42959   StrengthFormCode = "C42959"
	C42957   StrengthFormCode = "C42957"
	C42958   StrengthFormCode = "C42958"
	C68988   StrengthFormCode = "C68988"
	C91144   StrengthFormCode = "C91144"
	C68958   StrengthFormCode = "C68958"
	C91145   StrengthFormCode = "C91145"
	C68957   StrengthFormCode = "C68957"
	C91146   StrengthFormCode = "C91146"
	C68952   StrengthFormCode = "C68952"
	C64901   StrengthFormCode = "C64901"
	C68955   StrengthFormCode = "C68955"
	C91147   StrengthFormCode = "C91147"
	C64877   StrengthFormCode = "C64877"
	C64873   StrengthFormCode = "C64873"
	C42960   StrengthFormCode = "C42960"
	C91148   StrengthFormCode = "C91148"
	C60930   StrengthFormCode = "C60930"
	C42961   StrengthFormCode = "C42961"
	C42962   StrengthFormCode = "C42962"
	C64888   StrengthFormCode = "C64888"
	C42712   StrengthFormCode = "C42712"
	C29269   StrengthFormCode = "C29269"
	C91149   StrengthFormCode = "C91149"
	C91150   StrengthFormCode = "C91150"
	C42963   StrengthFormCode = "C42963"
	C42964   StrengthFormCode = "C42964"
	C69064   StrengthFormCode = "C69064"
	C91151   StrengthFormCode = "C91151"
	C91152   StrengthFormCode = "C91152"
	C91153   StrengthFormCode = "C91153"
	C91154   StrengthFormCode = "C91154"
	C91155   StrengthFormCode = "C91155"
	C91156   StrengthFormCode = "C91156"
	C91157   StrengthFormCode = "C91157"
	C91158   StrengthFormCode = "C91158"
	C68941   StrengthFormCode = "C68941"
	C69017   StrengthFormCode = "C69017"
	C48624   StrengthFormCode = "C48624"
	C69031   StrengthFormCode = "C69031"
	C42965   StrengthFormCode = "C42965"
	C42966   StrengthFormCode = "C42966"
	C69039   StrengthFormCode = "C69039"
	C91159   StrengthFormCode = "C91159"
	C91160   StrengthFormCode = "C91160"
	C91161   StrengthFormCode = "C91161"
	C69038   StrengthFormCode = "C69038"
	C91162   StrengthFormCode = "C91162"
	C91163   StrengthFormCode = "C91163"
	C91164   StrengthFormCode = "C91164"
	C42744   StrengthFormCode = "C42744"
	C91165   StrengthFormCode = "C91165"
	C91166   StrengthFormCode = "C91166"
	C91167   StrengthFormCode = "C91167"
	C91168   StrengthFormCode = "C91168"
	C91169   StrengthFormCode = "C91169"
	C91170   StrengthFormCode = "C91170"
	C91171   StrengthFormCode = "C91171"
	C68981   StrengthFormCode = "C68981"
	C68996   StrengthFormCode = "C68996"
	C91172   StrengthFormCode = "C91172"
	C91173   StrengthFormCode = "C91173"
	C68992   StrengthFormCode = "C68992"
	C68993   StrengthFormCode = "C68993"
	C43243   StrengthFormCode = "C43243"
	C61006   StrengthFormCode = "C61006"
	C42999   StrengthFormCode = "C42999"
	C69008   StrengthFormCode = "C69008"
	C42760   StrengthFormCode = "C42760"
	C69040   StrengthFormCode = "C69040"
	C91174   StrengthFormCode = "C91174"
	C91175   StrengthFormCode = "C91175"
	C91176   StrengthFormCode = "C91176"
	C91177   StrengthFormCode = "C91177"
	C62653   StrengthFormCode = "C62653"
	C47887   StrengthFormCode = "C47887"
	C69016   StrengthFormCode = "C69016"
	C42746   StrengthFormCode = "C42746"
	C42967   StrengthFormCode = "C42967"
	C60985   StrengthFormCode = "C60985"
	C42968   StrengthFormCode = "C42968"
	C69042   StrengthFormCode = "C69042"
	C64887   StrengthFormCode = "C64887"
	C42969   StrengthFormCode = "C42969"
	C42636   StrengthFormCode = "C42636"
	C25394   StrengthFormCode = "C25394"
	C42970   StrengthFormCode = "C42970"
	C42736   StrengthFormCode = "C42736"
	C47913   StrengthFormCode = "C47913"
	C42972   StrengthFormCode = "C42972"
	C42971   StrengthFormCode = "C42971"
	C42977   StrengthFormCode = "C42977"
	C42974   StrengthFormCode = "C42974"
	C42976   StrengthFormCode = "C42976"
	C69069   StrengthFormCode = "C69069"
	C64907   StrengthFormCode = "C64907"
	C64908   StrengthFormCode = "C64908"
	C69070   StrengthFormCode = "C69070"
	C42973   StrengthFormCode = "C42973"
	C42975   StrengthFormCode = "C42975"
	C91139   StrengthFormCode = "C91139"
	C68983   StrengthFormCode = "C68983"
	C91179   StrengthFormCode = "C91179"
	C97717   StrengthFormCode = "C97717"
	C91180   StrengthFormCode = "C91180"
	C68972   StrengthFormCode = "C68972"
	C68984   StrengthFormCode = "C68984"
	C68985   StrengthFormCode = "C68985"
	C69046   StrengthFormCode = "C69046"
	C69045   StrengthFormCode = "C69045"
	C91181   StrengthFormCode = "C91181"
	C91182   StrengthFormCode = "C91182"
	C69047   StrengthFormCode = "C69047"
	C91183   StrengthFormCode = "C91183"
	C91184   StrengthFormCode = "C91184"
	C68989   StrengthFormCode = "C68989"
	C42978   StrengthFormCode = "C42978"
	C60988   StrengthFormCode = "C60988"
	C42979   StrengthFormCode = "C42979"
	C42980   StrengthFormCode = "C42980"
	C45244   StrengthFormCode = "C45244"
	C42981   StrengthFormCode = "C42981"
	C42982   StrengthFormCode = "C42982"
	C28276   StrengthFormCode = "C28276"
	C42983   StrengthFormCode = "C42983"
	C64909   StrengthFormCode = "C64909"
	C45235   StrengthFormCode = "C45235"
	C42984   StrengthFormCode = "C42984"
	C42985   StrengthFormCode = "C42985"
	C42986   StrengthFormCode = "C42986"
	C69028   StrengthFormCode = "C69028"
	C42987   StrengthFormCode = "C42987"
	C47912   StrengthFormCode = "C47912"
	C42989   StrengthFormCode = "C42989"
	C42990   StrengthFormCode = "C42990"
	C68946   StrengthFormCode = "C68946"
	C42991   StrengthFormCode = "C42991"
	C47914   StrengthFormCode = "C47914"
	C42751   StrengthFormCode = "C42751"
	C42992   StrengthFormCode = "C42992"
	C42993   StrengthFormCode = "C42993"
	C42994   StrengthFormCode = "C42994"
	C68986   StrengthFormCode = "C68986"
	C69027   StrengthFormCode = "C69027"
	C42672   StrengthFormCode = "C42672"
	C69007   StrengthFormCode = "C69007"
	C69043   StrengthFormCode = "C69043"
	C69044   StrengthFormCode = "C69044"
	C69036   StrengthFormCode = "C69036"
	C68947   StrengthFormCode = "C68947"
	C78751   StrengthFormCode = "C78751"
	C47889   StrengthFormCode = "C47889"
	C47898   StrengthFormCode = "C47898"
	C42996   StrengthFormCode = "C42996"
	C42998   StrengthFormCode = "C42998"
	C69004   StrengthFormCode = "C69004"
	C69003   StrengthFormCode = "C69003"
	C60997   StrengthFormCode = "C60997"
	C61004   StrengthFormCode = "C61004"
	C61005   StrengthFormCode = "C61005"
	C69006   StrengthFormCode = "C69006"
	C42759   StrengthFormCode = "C42759"
	C47892   StrengthFormCode = "C47892"
	C47897   StrengthFormCode = "C47897"
	C78749   StrengthFormCode = "C78749"
	C69063   StrengthFormCode = "C69063"
	C43000   StrengthFormCode = "C43000"
	C91186   StrengthFormCode = "C91186"
	C42747   StrengthFormCode = "C42747"
	C91187   StrengthFormCode = "C91187"
	C91188   StrengthFormCode = "C91188"
	C91189   StrengthFormCode = "C91189"
	C91190   StrengthFormCode = "C91190"
	C91191   StrengthFormCode = "C91191"
	C91192   StrengthFormCode = "C91192"
	C91193   StrengthFormCode = "C91193"
	C91178   StrengthFormCode = "C91178"
	C64905   StrengthFormCode = "C64905"
	C91194   StrengthFormCode = "C91194"
	C68998   StrengthFormCode = "C68998"
	C43001   StrengthFormCode = "C43001"
	C43002   StrengthFormCode = "C43002"
	C64902   StrengthFormCode = "C64902"
	C38046_2 StrengthFormCode = "C38046"
	C69052   StrengthFormCode = "C69052"
	C91195   StrengthFormCode = "C91195"
	C69053   StrengthFormCode = "C69053"
	C69048   StrengthFormCode = "C69048"
	C69050   StrengthFormCode = "C69050"
	C47890   StrengthFormCode = "C47890"
	C69051   StrengthFormCode = "C69051"
	C91197   StrengthFormCode = "C91197"
	C69054   StrengthFormCode = "C69054"
	C91198   StrengthFormCode = "C91198"
	C91199   StrengthFormCode = "C91199"
	C91200   StrengthFormCode = "C91200"
	C68990   StrengthFormCode = "C68990"
	C69049   StrengthFormCode = "C69049"
	C43003   StrengthFormCode = "C43003"
	C64903   StrengthFormCode = "C64903"
	C78750   StrengthFormCode = "C78750"
	C64910   StrengthFormCode = "C64910"
)

// All allowed values of StrengthFormCode enum
var AllowedStrengthFormCodeEnumValues = []StrengthFormCode{
	"C78746",
	"C78747",
	"C107670",
	"C64886",
	"C42887",
	"C42888",
	"C68935",
	"C69030",
	"C42889",
	"C60897",
	"C91136",
	"C60957",
	"C60984",
	"C91227",
	"C69012",
	"C68950",
	"C91137",
	"C42890",
	"C64875",
	"C42891",
	"C64874",
	"C91138",
	"C42755",
	"C64878",
	"C68951",
	"C64880",
	"C25158",
	"C68943",
	"C68944",
	"C45414",
	"C42892",
	"C64876",
	"C42893",
	"C42894",
	"C42678",
	"C60884",
	"C42895",
	"C42896",
	"C42897",
	"C69002",
	"C60891",
	"C42899",
	"C69001",
	"C42898",
	"C42900",
	"C42731",
	"C69024",
	"C69026",
	"C69025",
	"C28944",
	"C42901",
	"C64881",
	"C45415",
	"C68949",
	"C42730",
	"C42902",
	"C42903",
	"C42997",
	"C42904",
	"C42905",
	"C69059",
	"C68954",
	"C45413",
	"C42906",
	"C42907",
	"C42908",
	"C42740",
	"C43525",
	"C69071",
	"C42756",
	"C42741",
	"C42679",
	"C69033",
	"C69032",
	"C42763",
	"C29012",
	"C60992",
	"C60995",
	"C68997",
	"C64883",
	"C42909",
	"C64884",
	"C42910",
	"C42911",
	"C42912",
	"C42913",
	"C42914",
	"C42915",
	"C64885",
	"C64871",
	"C68945",
	"C42758",
	"C42742",
	"C42713",
	"C43451",
	"C42916",
	"C42917",
	"C42918",
	"C42919",
	"C91140",
	"C91141",
	"C60926",
	"C42920",
	"C42928",
	"C60929",
	"C42935",
	"C69067",
	"C42921",
	"C42922",
	"C60934",
	"C42923",
	"C42924",
	"C42925",
	"C42927",
	"C42929",
	"C42932",
	"C42930",
	"C42931",
	"C68982",
	"C68991",
	"C64898",
	"C64899",
	"C60927",
	"C60928",
	"C68966",
	"C78748",
	"C42933",
	"C42934",
	"C60994",
	"C68973",
	"C42936",
	"C64872",
	"C48193",
	"C42937",
	"C45416",
	"C42938",
	"C69066",
	"C42939",
	"C42940",
	"C42941",
	"C64904",
	"C64882",
	"C42752",
	"C42669",
	"C42942",
	"C42943",
	"C42944",
	"C91142",
	"C91143",
	"C64879",
	"C42946",
	"C42926",
	"C42950",
	"C42951",
	"C69037",
	"C42945",
	"C42988",
	"C42995",
	"C60933",
	"C78793",
	"C68971",
	"C47915",
	"C68967",
	"C68965",
	"C42947",
	"C42948",
	"C47916",
	"C42949",
	"C60931",
	"C42952",
	"C42953",
	"C42954",
	"C68953",
	"C69068",
	"C29167",
	"C60958",
	"C42955",
	"C42956",
	"C42959",
	"C42957",
	"C42958",
	"C68988",
	"C91144",
	"C68958",
	"C91145",
	"C68957",
	"C91146",
	"C68952",
	"C64901",
	"C68955",
	"C91147",
	"C64877",
	"C64873",
	"C42960",
	"C91148",
	"C60930",
	"C42961",
	"C42962",
	"C64888",
	"C42712",
	"C29269",
	"C91149",
	"C91150",
	"C42963",
	"C42964",
	"C69064",
	"C91151",
	"C91152",
	"C91153",
	"C91154",
	"C91155",
	"C91156",
	"C91157",
	"C91158",
	"C68941",
	"C69017",
	"C48624",
	"C69031",
	"C42965",
	"C42966",
	"C69039",
	"C91159",
	"C91160",
	"C91161",
	"C69038",
	"C91162",
	"C91163",
	"C91164",
	"C42744",
	"C91165",
	"C91166",
	"C91167",
	"C91168",
	"C91169",
	"C91170",
	"C91171",
	"C68981",
	"C68996",
	"C91172",
	"C91173",
	"C68992",
	"C68993",
	"C43243",
	"C61006",
	"C42999",
	"C69008",
	"C42760",
	"C69040",
	"C91174",
	"C91175",
	"C91176",
	"C91177",
	"C62653",
	"C47887",
	"C69016",
	"C42746",
	"C42967",
	"C60985",
	"C42968",
	"C69042",
	"C64887",
	"C42969",
	"C42636",
	"C25394",
	"C42970",
	"C42736",
	"C47913",
	"C42972",
	"C42971",
	"C42977",
	"C42974",
	"C42976",
	"C69069",
	"C64907",
	"C64908",
	"C69070",
	"C42973",
	"C42975",
	"C91139",
	"C68983",
	"C91179",
	"C97717",
	"C91180",
	"C68972",
	"C68984",
	"C68985",
	"C69046",
	"C69045",
	"C91181",
	"C91182",
	"C69047",
	"C91183",
	"C91184",
	"C68989",
	"C42978",
	"C60988",
	"C42979",
	"C42980",
	"C45244",
	"C42981",
	"C42982",
	"C28276",
	"C42983",
	"C64909",
	"C45235",
	"C42984",
	"C42985",
	"C42986",
	"C69028",
	"C42987",
	"C47912",
	"C42989",
	"C42990",
	"C68946",
	"C42991",
	"C47914",
	"C42751",
	"C42992",
	"C42993",
	"C42994",
	"C68986",
	"C69027",
	"C42672",
	"C69007",
	"C69043",
	"C69044",
	"C69036",
	"C68947",
	"C78751",
	"C47889",
	"C47898",
	"C42996",
	"C42998",
	"C69004",
	"C69003",
	"C60997",
	"C61004",
	"C61005",
	"C69006",
	"C42759",
	"C47892",
	"C47897",
	"C78749",
	"C69063",
	"C43000",
	"C91186",
	"C42747",
	"C91187",
	"C91188",
	"C91189",
	"C91190",
	"C91191",
	"C91192",
	"C91193",
	"C91178",
	"C64905",
	"C91194",
	"C68998",
	"C43001",
	"C43002",
	"C64902",
	"C38046",
	"C69052",
	"C91195",
	"C69053",
	"C69048",
	"C69050",
	"C47890",
	"C69051",
	"C91197",
	"C69054",
	"C91198",
	"C91199",
	"C91200",
	"C68990",
	"C69049",
	"C43003",
	"C64903",
	"C78750",
	"C64910",
}

func (v *StrengthFormCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StrengthFormCode(value)
	for _, existing := range AllowedStrengthFormCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StrengthFormCode", value)
}

// NewStrengthFormCodeFromValue returns a pointer to a valid StrengthFormCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStrengthFormCodeFromValue(v string) (*StrengthFormCode, error) {
	ev := StrengthFormCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StrengthFormCode: valid values are %v", v, AllowedStrengthFormCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StrengthFormCode) IsValid() bool {
	for _, existing := range AllowedStrengthFormCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StrengthFormCode value
func (v StrengthFormCode) Ptr() *StrengthFormCode {
	return &v
}

type NullableStrengthFormCode struct {
	value *StrengthFormCode
	isSet bool
}

func (v NullableStrengthFormCode) Get() *StrengthFormCode {
	return v.value
}

func (v *NullableStrengthFormCode) Set(val *StrengthFormCode) {
	v.value = val
	v.isSet = true
}

func (v NullableStrengthFormCode) IsSet() bool {
	return v.isSet
}

func (v *NullableStrengthFormCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrengthFormCode(val *StrengthFormCode) *NullableStrengthFormCode {
	return &NullableStrengthFormCode{value: val, isSet: true}
}

func (v NullableStrengthFormCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrengthFormCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
