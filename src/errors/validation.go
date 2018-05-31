package errors

import (
	"fmt"

	"github.com/AlexSugak/skycoin-promo/src/util"
	validator "gopkg.in/go-playground/validator.v9"
)

// ValidationErrorResponse is entity that contains property name and its error message
type ValidationErrorResponse struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

// ValidationError represents validation error
type ValidationError struct {
	Errors []ValidationErrorResponse
}

func (e ValidationError) Error() string {
	return fmt.Sprint(e.Errors)
}

// CreateSingleValidationError creates ValidationError for single error
func CreateSingleValidationError(key string, message string) ValidationError {
	errorResponse := make([]ValidationErrorResponse, 0)

	ve := ValidationErrorResponse{
		Key:     key,
		Message: message,
	}
	errorResponse = append(errorResponse, ve)

	return ValidationError{Errors: errorResponse}
}

// ValidatorErrorsResponse convert validator.ValidationErrors to ValidationError
func ValidatorErrorsResponse(errors validator.ValidationErrors) ValidationError {
	errorResponse := make([]ValidationErrorResponse, 0)

	for i := 0; i < len(errors); i++ {
		ve := ValidationErrorResponse{
			Key:     util.FirstToLower(errors[i].Field()),
			Message: mapFieldErrorToErrorMessage(errors[i]),
		}
		errorResponse = append(errorResponse, ve)
	}

	return ValidationError{Errors: errorResponse}
}

func mapFieldErrorToErrorMessage(fe validator.FieldError) string {
	tag := fe.Tag()

	switch tag {
	case "required":
		return "is required"
	case "len":
		return fmt.Sprintf("length should be %s", fe.Param())
	case "max":
		return fmt.Sprintf("can't be more then %s", fe.Param())
	case "min":
		return fmt.Sprintf("can't be less then %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("should be one of: %s", fe.Param())
	case "email":
		return fmt.Sprintf("string is not email address")
	default:
		return "unknown error"
	}
}
