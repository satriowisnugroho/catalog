package types

import (
	"encoding/json"
	"fmt"
)

// CategoryType represent product category type
type CategoryType int8

// Category(*)Type represent product category type enum
const (
	CategoryEmptyType CategoryType = iota
	CategoryBookType
	CategoryComputerType
	CategoryBagType
)

var (
	CategoryTypeNameToValue = map[string]CategoryType{
		"book":     CategoryBookType,
		"computer": CategoryComputerType,
		"bag":      CategoryBagType,
	}

	_CategoryTypeValueToName = map[CategoryType]string{
		CategoryBookType:     "book",
		CategoryComputerType: "computer",
		CategoryBagType:      "bag",
	}
)

// Scan is used for Scan
func (t *CategoryType) Scan(value interface{}) error {
	val := CategoryType(value.(int64))
	if val == 0 || int(value.(int64)) > len(CategoryTypeNameToValue) {
		return errInvalidEnum("category_type", fmt.Sprint(value.(int64)))
	}

	*t = val
	return nil
}

// MarshalJSON defined so that CategoryType satisfies json.Marshaler
func (t CategoryType) MarshalJSON() ([]byte, error) {
	s, ok := _CategoryTypeValueToName[t]
	if !ok {
		return nil, errInvalidEnum("category_type", fmt.Sprint(t))
	}
	return json.Marshal(s)
}

// UnmarshalJSON defined so that CategoryType satisfies json.Unmarshaler
func (r *CategoryType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CategoryType should be a string, got %s", data)
	}
	v, ok := CategoryTypeNameToValue[s]
	if !ok {
		return errInvalidValue("category_type", s)
	}
	*r = v
	return nil
}
