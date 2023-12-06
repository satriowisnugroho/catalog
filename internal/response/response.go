package response

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// SuccessBody holds data for success response
type SuccessBody struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    interface{} `json:"meta"`
}

// CustomError holds data for customized error
type CustomError struct {
	Message  string
	Field    string
	Code     int
	HTTPCode int
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (c CustomError) Error() string {
	return c.Message
}

// ErrorInfo holds error detail
type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Field   string `json:"field,omitempty"`
}

// Error implement error
func (e ErrorInfo) Error() string {
	return fmt.Sprintf(
		"error - msg: %s, code: %d, field: %s",
		e.Message,
		e.Code,
		e.Field,
	)
}

// ErrorBody holds data for error response
type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (e ErrorBody) Error() string {
	errMsg := "response - errors"
	for _, err := range e.Errors {
		errMsg += fmt.Sprintf("\n\t%s", err.Error())
	}

	return errMsg
}

// ErrorResponse error from http.Response
type ErrorResponse struct {
	ErrorBody  ErrorBody
	HTTPStatus int
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (e ErrorResponse) Error() string {
	errMsg := e.ErrorBody.Error()
	errMsg += fmt.Sprintf("\nhttp_status: %d", e.HTTPStatus)

	return errMsg
}

// MetaInfo holds meta data
type MetaInfo struct {
	HTTPStatus int `json:"http_status,omitempty"`
	Offset     int `json:"offset,omitempty"`
	Limit      int `json:"limit,omitempty"`
	Total      int `json:"total,omitempty"`
}

// // PaginationMetaInfo is used to define meta info for pagination meta
// type PaginationMetaInfo struct {
// 	HTTPStatus uint `json:"http_status"`
// 	Offset     int  `json:"offset"`
// 	Limit      int  `json:"limit"`
// 	Total      int  `json:"total"`
// }

const (
	// ErrorCodeNotFound Error code for path not found
	ErrorCodeNotFound = 10000
	// ErrorCodeInsufficientStock Error code for insufficient quantity
	ErrorCodeInsufficientStock = 10001
	// ErrorCodeForbidden Error code for forbidden
	ErrorCodeForbidden = 10002

	// ErrorCodeNoSQLTransactionFound Error code for no sql transaction found
	ErrorCodeNoSQLTransactionFound = 9000
	// ErrorCodeUnexpectedError Error code for unexpected error
	ErrorCodeUnexpectedError = 9999
)

var (
	// ErrInternalServerError custom error on unexpected error
	ErrInternalServerError = CustomError{
		Message:  "Internal Server Error",
		Code:     ErrorCodeUnexpectedError,
		HTTPCode: http.StatusInternalServerError,
	}
	// ErrNotFound define error if record is not found in database
	ErrNotFound = CustomError{
		Message:  "Record not found",
		Code:     ErrorCodeNotFound,
		HTTPCode: http.StatusNotFound,
	}
	// ErrInsufficientStock define error when insufficient quantity
	ErrInsufficientStock = CustomError{
		Message:  "Insufficient quantity",
		Code:     ErrorCodeInsufficientStock,
		HTTPCode: http.StatusUnprocessableEntity,
	}
	// ErrForbidden define error when forbidden
	ErrForbidden = CustomError{
		Message:  "Forbidden",
		Code:     ErrorCodeForbidden,
		HTTPCode: http.StatusForbidden,
	}

	// ErrNoSQLTransactionFound defines no sql transaction when do the db transaction
	ErrNoSQLTransactionFound = CustomError{
		Message:  "No SQL transaction found",
		Code:     ErrorCodeNoSQLTransactionFound,
		HTTPCode: http.StatusInternalServerError,
	}
)

// BuildSuccess is a function to create SuccessBody
func BuildSuccess(data interface{}, message string, meta interface{}) SuccessBody {
	return SuccessBody{
		Data:    data,
		Message: message,
		Meta:    meta,
	}
}

// OK wrap success response
func OK(c *gin.Context, data interface{}, message string) {
	successResponse := BuildSuccess(data, message, MetaInfo{HTTPStatus: http.StatusOK})
	c.JSON(http.StatusOK, successResponse)
}

// OKWithPagination wrap success response with pagination meta
func OKWithPagination(c *gin.Context, data interface{}, message string, total, offset, limit int) {
	successResponse := BuildSuccess(data, message, MetaInfo{
		HTTPStatus: http.StatusOK,
		Total:      total,
		Offset:     offset,
		Limit:      limit,
	})
	c.JSON(http.StatusOK, successResponse)
}

// InternalServerErrorBody for default internal server error
func InternalServerErrorBody() ErrorBody {
	return ErrorBody{
		Errors: []ErrorInfo{
			{
				Message: ErrInternalServerError.Message,
				Code:    ErrInternalServerError.Code,
				Field:   ErrInternalServerError.Field,
			},
		},
		Meta: MetaInfo{
			HTTPStatus: ErrInternalServerError.HTTPCode,
		},
	}
}

// BuildError is a function to create ErrorBody
func BuildError(errors ...error) ErrorBody {
	if len(errors) == 0 {
		return InternalServerErrorBody()
	}

	errInfos := []ErrorInfo{}

	for _, err := range errors {
		switch errOrig := err.(type) {
		case CustomError:
			return ErrorBody{
				Errors: []ErrorInfo{
					{
						Message: errOrig.Message,
						Code:    errOrig.Code,
						Field:   errOrig.Field,
					},
				},
				Meta: MetaInfo{
					HTTPStatus: errOrig.HTTPCode,
				},
			}
		case ErrorInfo:
			errInfos = append(errInfos, errOrig)
		case ErrorBody:
			return errOrig
		case ErrorResponse:
			return errOrig.ErrorBody
		default:
			return InternalServerErrorBody()
		}
	}

	return ErrorBody{
		Errors: errInfos,
	}
}

// Error wrap error response
func Error(c *gin.Context, err error) {
	if err == context.Canceled || err == context.DeadlineExceeded {
		return
	}

	if ce, ok := err.(CustomError); ok {
		errorResponse := BuildError(ce)
		c.AbortWithStatusJSON(ce.HTTPCode, errorResponse)
		return
	}

	causer := errors.Cause(err)
	customError := causer

	var meta MetaInfo
	errorResponse := BuildError(customError)

	// case for error info with no http status
	if errorResponse.Meta == nil {
		errorResponse.Meta = MetaInfo{
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	meta = errorResponse.Meta.(MetaInfo)
	c.AbortWithStatusJSON(meta.HTTPStatus, errorResponse)
}
