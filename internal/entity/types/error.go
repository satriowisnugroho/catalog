package types

import (
	"fmt"
	"net/http"

	"github.com/satriowisnugroho/catalog/internal/response"
)

func errInvalidValue(enumType string, value string) response.CustomError {
	return errInvalid("value", enumType, value)
}

func errInvalidEnum(enumType string, value string) response.CustomError {
	return errInvalid("enum", enumType, value)
}

func errInvalid(errType string, enumType string, value string) response.CustomError {
	return response.CustomError{
		Message:  fmt.Sprintf("%s is invalid %s for %s", value, errType, enumType),
		Field:    enumType,
		HTTPCode: http.StatusUnprocessableEntity,
	}
}
