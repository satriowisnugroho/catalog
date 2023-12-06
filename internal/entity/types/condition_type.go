package types

import (
	"encoding/json"
	"fmt"
)

// ConditionType represent product condition type
type ConditionType int8

// Condition(*)Type represent product condition type enum
const (
	ConditionEmptyType ConditionType = iota
	ConditionNewType
	ConditionPrelovedType
)

var (
	ConditionTypeNameToValue = map[string]ConditionType{
		"new":      ConditionNewType,
		"preloved": ConditionPrelovedType,
	}

	_ConditionTypeValueToName = map[ConditionType]string{
		ConditionNewType:      "new",
		ConditionPrelovedType: "preloved",
	}
)

// Scan is used for Scan
func (t *ConditionType) Scan(value interface{}) error {
	val := ConditionType(value.(int64))
	if val == 0 || int(value.(int64)) > len(ConditionTypeNameToValue) {
		return errInvalidEnum("condition_type", fmt.Sprint(value.(int64)))
	}

	*t = val
	return nil
}

// MarshalJSON defined so that ConditionType satisfies json.Marshaler
func (t ConditionType) MarshalJSON() ([]byte, error) {
	s, ok := _ConditionTypeValueToName[t]
	if !ok {
		return nil, errInvalidEnum("condition_type", fmt.Sprint(t))
	}
	return json.Marshal(s)
}

// UnmarshalJSON defined so that ConditionType satisfies json.Unmarshaler
func (r *ConditionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ConditionType should be a string, got %s", data)
	}
	v, ok := ConditionTypeNameToValue[s]
	if !ok {
		return errInvalidValue("condition_type", s)
	}
	*r = v
	return nil
}
