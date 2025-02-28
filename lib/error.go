package lib

import "errors"

type ApiError struct {
	Err       error                   `json:"-"`
	RawErrors *[]string               `json:"rawErrors,omitempty"`
	Relogin   *bool                   `json:"relogin,omitempty"`
	Fields    *map[string]interface{} `json:"fields,omitempty"`
}

func NewApiError(message string, rawErrors []string, relogin bool, fields map[string]interface{}) *ApiError {
	return &ApiError{
		Err:       errors.New(message),
		RawErrors: &rawErrors,
		Relogin:   &relogin,
		Fields:    &fields,
	}
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}
