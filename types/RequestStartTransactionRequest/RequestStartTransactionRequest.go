// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package RequestStartTransactionRequest

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType, v)
	}
	*j = CostKindEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId in CustomDataType: required")
	}
	type Plain CustomDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType(plain)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken" yaml:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type" yaml:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken in AdditionalInfoType: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in AdditionalInfoType: required")
	}
	type Plain AdditionalInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType(plain)
	return nil
}

type ChargingProfileKindEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType_1, v)
	}
	*j = RecurrencyKindEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType, v)
	}
	*j = ChargingProfileKindEnumType(v)
	return nil
}

const ChargingProfileKindEnumTypeAbsolute ChargingProfileKindEnumType = "Absolute"
const ChargingProfileKindEnumTypeRecurring ChargingProfileKindEnumType = "Recurring"
const ChargingProfileKindEnumTypeRelative ChargingProfileKindEnumType = "Relative"

type ChargingProfilePurposeEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken in IdTokenType: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in IdTokenType: required")
	}
	type Plain IdTokenType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.AdditionalInfo) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "additionalInfo", 1)
	}
	*j = IdTokenType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType, v)
	}
	*j = ChargingProfilePurposeEnumType(v)
	return nil
}

const ChargingProfilePurposeEnumTypeChargingStationExternalConstraints ChargingProfilePurposeEnumType = "ChargingStationExternalConstraints"
const ChargingProfilePurposeEnumTypeChargingStationMaxProfile ChargingProfilePurposeEnumType = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumTypeTxDefaultProfile ChargingProfilePurposeEnumType = "TxDefaultProfile"
const ChargingProfilePurposeEnumTypeTxProfile ChargingProfilePurposeEnumType = "TxProfile"

type ChargingProfileKindEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_1, v)
	}
	*j = IdTokenEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType_1, v)
	}
	*j = ChargingProfileKindEnumType_1(v)
	return nil
}

const ChargingProfileKindEnumType_1_Absolute ChargingProfileKindEnumType_1 = "Absolute"
const ChargingProfileKindEnumType_1_Recurring ChargingProfileKindEnumType_1 = "Recurring"
const ChargingProfileKindEnumType_1_Relative ChargingProfileKindEnumType_1 = "Relative"

type ChargingProfilePurposeEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType, v)
	}
	*j = IdTokenEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_1, v)
	}
	*j = ChargingProfilePurposeEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_1, v)
	}
	*j = CostKindEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_1, v)
	}
	*j = ChargingRateUnitEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingProfileKind"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfileKind in ChargingProfileType: required")
	}
	if v, ok := raw["chargingProfilePurpose"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfilePurpose in ChargingProfileType: required")
	}
	if v, ok := raw["chargingSchedule"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedule in ChargingProfileType: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in ChargingProfileType: required")
	}
	if v, ok := raw["stackLevel"]; !ok || v == nil {
		return fmt.Errorf("field stackLevel in ChargingProfileType: required")
	}
	type Plain ChargingProfileType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.ChargingSchedule) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "chargingSchedule", 1)
	}
	if len(plain.ChargingSchedule) > 3 {
		return fmt.Errorf("field %s length: must be <= %d", "chargingSchedule", 3)
	}
	*j = ChargingProfileType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType, v)
	}
	*j = RecurrencyKindEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingScheduleType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingRateUnit"]; !ok || v == nil {
		return fmt.Errorf("field chargingRateUnit in ChargingScheduleType: required")
	}
	if v, ok := raw["chargingSchedulePeriod"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedulePeriod in ChargingScheduleType: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in ChargingScheduleType: required")
	}
	type Plain ChargingScheduleType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.ChargingSchedulePeriod) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "chargingSchedulePeriod", 1)
	}
	if len(plain.ChargingSchedulePeriod) > 1024 {
		return fmt.Errorf("field %s length: must be <= %d", "chargingSchedulePeriod", 1024)
	}
	*j = ChargingScheduleType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in SalesTariffType: required")
	}
	if v, ok := raw["salesTariffEntry"]; !ok || v == nil {
		return fmt.Errorf("field salesTariffEntry in SalesTariffType: required")
	}
	type Plain SalesTariffType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.SalesTariffEntry) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "salesTariffEntry", 1)
	}
	if len(plain.SalesTariffEntry) > 1024 {
		return fmt.Errorf("field %s length: must be <= %d", "salesTariffEntry", 1024)
	}
	*j = SalesTariffType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType, v)
	}
	*j = ChargingRateUnitEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffEntryType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["relativeTimeInterval"]; !ok || v == nil {
		return fmt.Errorf("field relativeTimeInterval in SalesTariffEntryType: required")
	}
	type Plain SalesTariffEntryType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.ConsumptionCost) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "consumptionCost", 1)
	}
	if len(plain.ConsumptionCost) > 3 {
		return fmt.Errorf("field %s length: must be <= %d", "consumptionCost", 3)
	}
	*j = SalesTariffEntryType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RelativeTimeIntervalType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["start"]; !ok || v == nil {
		return fmt.Errorf("field start in RelativeTimeIntervalType: required")
	}
	type Plain RelativeTimeIntervalType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RelativeTimeIntervalType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsumptionCostType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cost"]; !ok || v == nil {
		return fmt.Errorf("field cost in ConsumptionCostType: required")
	}
	if v, ok := raw["startValue"]; !ok || v == nil {
		return fmt.Errorf("field startValue in ConsumptionCostType: required")
	}
	type Plain ConsumptionCostType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.Cost) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "cost", 1)
	}
	if len(plain.Cost) > 3 {
		return fmt.Errorf("field %s length: must be <= %d", "cost", 3)
	}
	*j = ConsumptionCostType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit in ChargingSchedulePeriodType: required")
	}
	if v, ok := raw["startPeriod"]; !ok || v == nil {
		return fmt.Errorf("field startPeriod in ChargingSchedulePeriodType: required")
	}
	type Plain ChargingSchedulePeriodType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["amount"]; !ok || v == nil {
		return fmt.Errorf("field amount in CostType: required")
	}
	if v, ok := raw["costKind"]; !ok || v == nil {
		return fmt.Errorf("field costKind in CostType: required")
	}
	type Plain CostType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostType(plain)
	return nil
}

const ChargingProfilePurposeEnumType_1_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_1 = "ChargingStationExternalConstraints"
const ChargingProfilePurposeEnumType_1_ChargingStationMaxProfile ChargingProfilePurposeEnumType_1 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_1_TxDefaultProfile ChargingProfilePurposeEnumType_1 = "TxDefaultProfile"
const ChargingProfilePurposeEnumType_1_TxProfile ChargingProfilePurposeEnumType_1 = "TxProfile"

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
//
type ChargingProfileType struct {
	// ChargingProfileKind corresponds to the JSON schema field "chargingProfileKind".
	ChargingProfileKind ChargingProfileKindEnumType_1 `json:"chargingProfileKind" yaml:"chargingProfileKind"`

	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose ChargingProfilePurposeEnumType_1 `json:"chargingProfilePurpose" yaml:"chargingProfilePurpose"`

	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule []ChargingScheduleType `json:"chargingSchedule" yaml:"chargingSchedule"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Id of ChargingProfile.
	//
	Id int `json:"id" yaml:"id"`

	// RecurrencyKind corresponds to the JSON schema field "recurrencyKind".
	RecurrencyKind *RecurrencyKindEnumType `json:"recurrencyKind,omitempty" yaml:"recurrencyKind,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Value determining level in hierarchy stack of profiles. Higher values have
	// precedence over lower values. Lowest level is 0.
	//
	StackLevel int `json:"stackLevel" yaml:"stackLevel"`

	// SHALL only be included if ChargingProfilePurpose is set to TxProfile. The
	// transactionId is used to match the profile to a specific transaction.
	//
	TransactionId *string `json:"transactionId,omitempty" yaml:"transactionId,omitempty"`

	// Charging_ Profile. Valid_ From. Date_ Time
	// urn:x-oca:ocpp:uid:1:569234
	// Point in time at which the profile starts to be valid. If absent, the profile
	// is valid as soon as it is received by the Charging Station.
	//
	ValidFrom *string `json:"validFrom,omitempty" yaml:"validFrom,omitempty"`

	// Charging_ Profile. Valid_ To. Date_ Time
	// urn:x-oca:ocpp:uid:1:569235
	// Point in time at which the profile stops to be valid. If absent, the profile is
	// valid until it is replaced by another profile.
	//
	ValidTo *string `json:"validTo,omitempty" yaml:"validTo,omitempty"`
}

type ChargingRateUnitEnumType string

const ChargingRateUnitEnumTypeA ChargingRateUnitEnumType = "A"
const ChargingRateUnitEnumTypeW ChargingRateUnitEnumType = "W"

type ChargingRateUnitEnumType_1 string

const ChargingRateUnitEnumType_1_A ChargingRateUnitEnumType_1 = "A"
const ChargingRateUnitEnumType_1_W ChargingRateUnitEnumType_1 = "W"

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
//
type ChargingSchedulePeriodType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit" yaml:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty" yaml:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty" yaml:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod" yaml:"startPeriod"`
}

// Charging_ Schedule
// urn:x-oca:ocpp:uid:2:233256
// Charging schedule structure defines a list of charging periods, as used in:
// GetCompositeSchedule.conf and ChargingProfile.
//
type ChargingScheduleType struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType `json:"chargingRateUnit" yaml:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType `json:"chargingSchedulePeriod" yaml:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Charging_ Schedule. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569236
	// Duration of the charging schedule in seconds. If the duration is left empty,
	// the last period will continue indefinitely or until end of the transaction if
	// chargingProfilePurpose = TxProfile.
	//
	Duration *int `json:"duration,omitempty" yaml:"duration,omitempty"`

	// Identifies the ChargingSchedule.
	//
	Id int `json:"id" yaml:"id"`

	// Charging_ Schedule. Min_ Charging_ Rate. Numeric
	// urn:x-oca:ocpp:uid:1:569239
	// Minimum charging rate supported by the EV. The unit of measure is defined by
	// the chargingRateUnit. This parameter is intended to be used by a local smart
	// charging algorithm to optimize the power allocation for in the case a charging
	// process is inefficient at lower charging rates. Accepts at most one digit
	// fraction (e.g. 8.1)
	//
	MinChargingRate *float64 `json:"minChargingRate,omitempty" yaml:"minChargingRate,omitempty"`

	// SalesTariff corresponds to the JSON schema field "salesTariff".
	SalesTariff *SalesTariffType `json:"salesTariff,omitempty" yaml:"salesTariff,omitempty"`

	// Charging_ Schedule. Start_ Schedule. Date_ Time
	// urn:x-oca:ocpp:uid:1:569237
	// Starting point of an absolute schedule. If absent the schedule will be relative
	// to start of charging.
	//
	StartSchedule *string `json:"startSchedule,omitempty" yaml:"startSchedule,omitempty"`
}

// Consumption_ Cost
// urn:x-oca:ocpp:uid:2:233259
//
type ConsumptionCostType struct {
	// Cost corresponds to the JSON schema field "cost".
	Cost []CostType `json:"cost" yaml:"cost"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Consumption_ Cost. Start_ Value. Numeric
	// urn:x-oca:ocpp:uid:1:569246
	// The lowest level of consumption that defines the starting point of this
	// consumption block. The block interval extends to the start of the next
	// interval.
	//
	StartValue float64 `json:"startValue" yaml:"startValue"`
}

type CostKindEnumType string

const CostKindEnumTypeCarbonDioxideEmission CostKindEnumType = "CarbonDioxideEmission"
const CostKindEnumTypeRelativePricePercentage CostKindEnumType = "RelativePricePercentage"
const CostKindEnumTypeRenewableGenerationPercentage CostKindEnumType = "RenewableGenerationPercentage"

type CostKindEnumType_1 string

const CostKindEnumType_1_CarbonDioxideEmission CostKindEnumType_1 = "CarbonDioxideEmission"
const CostKindEnumType_1_RelativePricePercentage CostKindEnumType_1 = "RelativePricePercentage"
const CostKindEnumType_1_RenewableGenerationPercentage CostKindEnumType_1 = "RenewableGenerationPercentage"

// Cost
// urn:x-oca:ocpp:uid:2:233258
//
type CostType struct {
	// Cost. Amount. Amount
	// urn:x-oca:ocpp:uid:1:569244
	// The estimated or actual cost per kWh
	//
	Amount int `json:"amount" yaml:"amount"`

	// Cost. Amount_ Multiplier. Integer
	// urn:x-oca:ocpp:uid:1:569245
	// Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The
	// final value is determined by: amount * 10 ^ amountMultiplier
	//
	AmountMultiplier *int `json:"amountMultiplier,omitempty" yaml:"amountMultiplier,omitempty"`

	// CostKind corresponds to the JSON schema field "costKind".
	CostKind CostKindEnumType `json:"costKind" yaml:"costKind"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

type IdTokenEnumType string

const IdTokenEnumTypeCentral IdTokenEnumType = "Central"
const IdTokenEnumTypeEMAID IdTokenEnumType = "eMAID"
const IdTokenEnumTypeISO14443 IdTokenEnumType = "ISO14443"
const IdTokenEnumTypeISO15693 IdTokenEnumType = "ISO15693"
const IdTokenEnumTypeKeyCode IdTokenEnumType = "KeyCode"
const IdTokenEnumTypeLocal IdTokenEnumType = "Local"
const IdTokenEnumTypeMacAddress IdTokenEnumType = "MacAddress"
const IdTokenEnumTypeNoAuthorization IdTokenEnumType = "NoAuthorization"

type IdTokenEnumType_1 string

const IdTokenEnumType_1_Central IdTokenEnumType_1 = "Central"
const IdTokenEnumType_1_EMAID IdTokenEnumType_1 = "eMAID"
const IdTokenEnumType_1_ISO14443 IdTokenEnumType_1 = "ISO14443"
const IdTokenEnumType_1_ISO15693 IdTokenEnumType_1 = "ISO15693"
const IdTokenEnumType_1_KeyCode IdTokenEnumType_1 = "KeyCode"
const IdTokenEnumType_1_Local IdTokenEnumType_1 = "Local"
const IdTokenEnumType_1_MacAddress IdTokenEnumType_1 = "MacAddress"
const IdTokenEnumType_1_NoAuthorization IdTokenEnumType_1 = "NoAuthorization"

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType `json:"additionalInfo,omitempty" yaml:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken" yaml:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_1 `json:"type" yaml:"type"`
}

type RecurrencyKindEnumType string

const RecurrencyKindEnumTypeDaily RecurrencyKindEnumType = "Daily"
const RecurrencyKindEnumTypeWeekly RecurrencyKindEnumType = "Weekly"

type RecurrencyKindEnumType_1 string

const RecurrencyKindEnumType_1_Daily RecurrencyKindEnumType_1 = "Daily"
const RecurrencyKindEnumType_1_Weekly RecurrencyKindEnumType_1 = "Weekly"

// Relative_ Timer_ Interval
// urn:x-oca:ocpp:uid:2:233270
//
type RelativeTimeIntervalType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Relative_ Timer_ Interval. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569280
	// Duration of the interval, in seconds.
	//
	Duration *int `json:"duration,omitempty" yaml:"duration,omitempty"`

	// Relative_ Timer_ Interval. Start. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569279
	// Start of the interval, in seconds from NOW.
	//
	Start int `json:"start" yaml:"start"`
}

type RequestStartTransactionRequestJson struct {
	// ChargingProfile corresponds to the JSON schema field "chargingProfile".
	ChargingProfile *ChargingProfileType `json:"chargingProfile,omitempty" yaml:"chargingProfile,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Number of the EVSE on which to start the transaction. EvseId SHALL be &gt; 0
	//
	EvseId *int `json:"evseId,omitempty" yaml:"evseId,omitempty"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType `json:"groupIdToken,omitempty" yaml:"groupIdToken,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken IdTokenType `json:"idToken" yaml:"idToken"`

	// Id given by the server to this start request. The Charging Station might return
	// this in the &lt;&lt;transactioneventrequest, TransactionEventRequest&gt;&gt;,
	// letting the server know which transaction was started for this request. Use to
	// start a transaction.
	//
	RemoteStartId int `json:"remoteStartId" yaml:"remoteStartId"`
}

// Sales_ Tariff_ Entry
// urn:x-oca:ocpp:uid:2:233271
//
type SalesTariffEntryType struct {
	// ConsumptionCost corresponds to the JSON schema field "consumptionCost".
	ConsumptionCost []ConsumptionCostType `json:"consumptionCost,omitempty" yaml:"consumptionCost,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Sales_ Tariff_ Entry. E_ Price_ Level. Unsigned_ Integer
	// urn:x-oca:ocpp:uid:1:569281
	// Defines the price level of this SalesTariffEntry (referring to
	// NumEPriceLevels). Small values for the EPriceLevel represent a cheaper
	// TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	//
	EPriceLevel *int `json:"ePriceLevel,omitempty" yaml:"ePriceLevel,omitempty"`

	// RelativeTimeInterval corresponds to the JSON schema field
	// "relativeTimeInterval".
	RelativeTimeInterval RelativeTimeIntervalType `json:"relativeTimeInterval" yaml:"relativeTimeInterval"`
}

// Sales_ Tariff
// urn:x-oca:ocpp:uid:2:233272
// NOTE: This dataType is based on dataTypes from &lt;&lt;ref-ISOIEC15118-2,ISO
// 15118-2&gt;&gt;.
//
type SalesTariffType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// SalesTariff identifier used to identify one sales tariff. An SAID remains a
	// unique identifier for one schedule throughout a charging session.
	//
	Id int `json:"id" yaml:"id"`

	// Sales_ Tariff. Num_ E_ Price_ Levels. Counter
	// urn:x-oca:ocpp:uid:1:569284
	// Defines the overall number of distinct price levels used across all provided
	// SalesTariff elements.
	//
	NumEPriceLevels *int `json:"numEPriceLevels,omitempty" yaml:"numEPriceLevels,omitempty"`

	// Sales_ Tariff. Sales. Tariff_ Description
	// urn:x-oca:ocpp:uid:1:569283
	// A human readable title/short description of the sales tariff e.g. for HMI
	// display purposes.
	//
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty" yaml:"salesTariffDescription,omitempty"`

	// SalesTariffEntry corresponds to the JSON schema field "salesTariffEntry".
	SalesTariffEntry []SalesTariffEntryType `json:"salesTariffEntry" yaml:"salesTariffEntry"`
}

var enumValues_ChargingProfileKindEnumType = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfileKindEnumType_1 = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfilePurposeEnumType = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_1 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingRateUnitEnumType = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_1 = []interface{}{
	"W",
	"A",
}
var enumValues_CostKindEnumType = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_1 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_IdTokenEnumType = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_1 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_RecurrencyKindEnumType = []interface{}{
	"Daily",
	"Weekly",
}
var enumValues_RecurrencyKindEnumType_1 = []interface{}{
	"Daily",
	"Weekly",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStartTransactionRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken in RequestStartTransactionRequestJson: required")
	}
	if v, ok := raw["remoteStartId"]; !ok || v == nil {
		return fmt.Errorf("field remoteStartId in RequestStartTransactionRequestJson: required")
	}
	type Plain RequestStartTransactionRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequestStartTransactionRequestJson(plain)
	return nil
}
