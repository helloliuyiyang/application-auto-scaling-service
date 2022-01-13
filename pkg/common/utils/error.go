package utils

import (
	"fmt"
)

// constants for error type
const (
	ErrGeneral int = 0

	ErrService int = 5
)

//constants for prefix type
const (
	aassPrefix = "aass"
)

// APIError holds API errors for function manager
type APIError struct {
	ErrorPrefix string
	ErrorCode   int
	ErrorType   int
	ErrorMsg    string
}

// InternalServerError error
func InternalServerError() APIError {
	return APIError{aassPrefix, 1, ErrGeneral, "Internal Server Error!"}
}

// InternalServerWithError error with error msg
func InternalServerWithError(errMsg string) APIError {
	return APIError{aassPrefix, 1, ErrGeneral,
		fmt.Sprintf("Internal Server Error: %s", errMsg)}
}

// UserHasNoAccessError Role without permission
func UserHasNoAccessError() APIError {
	return APIError{aassPrefix, 5, ErrGeneral,
		"the current user have no access to service"}
}

// RequestQueryParamInvalid check request URI error
func RequestQueryParamInvalid() APIError {
	return APIError{aassPrefix, 6, ErrGeneral,
		fmt.Sprintf("the query parameter is illegal")}
}

// RequestBodyParamInvalid check request body error
func RequestBodyParamInvalid(errMsg string) APIError {
	return APIError{aassPrefix, 7, ErrGeneral,
		fmt.Sprintf("The request body parameter is illegal: %s.", errMsg)}
}

// RequestURIInvalidError error
func RequestURIInvalidError() APIError {
	return APIError{aassPrefix, 8, ErrGeneral, "request uri is illegal!"}
}

// RequestBodyInvalid error
func RequestBodyInvalid() APIError {
	return APIError{aassPrefix, 10, ErrGeneral, "the request body is illegal!"}
}

// RequestURIParamEmptyError error
func RequestURIParamEmptyError(pathParamName string) APIError {
	return APIError{aassPrefix, 31, ErrGeneral,
		fmt.Sprintf("request uri path param (%s) can't be empty!Please input it.", pathParamName)}
}

func InternalError(err error) APIError {
	return APIError{aassPrefix, 500, ErrService,
		fmt.Sprintf("internal err: %v", err)}
}
