package types

import (
	"encoding/json"
	"fmt"
)

// TenantType represent tenant type on catalog
type TenantType int8

// Tenant(*)Type represent tenant type enum on catalog
const (
	TenantEmptyType TenantType = iota
	TenantLoremType
	TenantIpsumType
)

var (
	_TenantTypeNameToValue = map[string]TenantType{
		"lorem": TenantLoremType,
		"ipsum": TenantIpsumType,
	}

	_TenantTypeValueToName = map[TenantType]string{
		TenantLoremType: "lorem",
		TenantIpsumType: "ipsum",
	}
)

// Scan is used for Scan
func (t *TenantType) Scan(value interface{}) error {
	val := TenantType(value.(int64))
	if val == 0 || val > 1 {
		return errInvalidEnum("tenant_type", fmt.Sprint(value.(int64)))
	}

	*t = val
	return nil
}

// MarshalJSON defined so that TenantType satisfies json.Marshaler
func (t TenantType) MarshalJSON() ([]byte, error) {
	s, ok := _TenantTypeValueToName[t]
	if !ok {
		return nil, errInvalidEnum("tenant_type", fmt.Sprint(t))
	}
	return json.Marshal(s)
}

// UnmarshalJSON defined so that TenantType satisfies json.Unmarshaler
func (r *TenantType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TenantType should be a string, got %s", data)
	}
	v, ok := _TenantTypeNameToValue[s]
	if !ok {
		return errInvalidEnum("tenant_type", s)
	}
	*r = v
	return nil
}
