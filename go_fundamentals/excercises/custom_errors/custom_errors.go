package custom_errors

import (
	"reflect"
)

type InternalError struct {
	description string
}

type ThirdPartyError struct {
	description string
}

type OtherError struct {
	description string
}

func (e InternalError) Error() string {
	return "Internal Error:" + e.description
}

func (e ThirdPartyError) Error() string {
	return "Third Party Error" + e.description
}

func (e OtherError) Error() string {
	return "Other Error" + e.description
}

func IsErrorType(e error, errorType string) bool {

	if reflect.TypeOf(e).Name() == errorType {
		return true
	}

	return false
}
