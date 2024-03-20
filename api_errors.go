package main

import (
	"errors"
	"fmt"
)

type APIError struct {
	StatusCode int
	Err        error
}

func NewAPIError(statusCode int, errMsg string) error {
	return &APIError{
		StatusCode: statusCode,
		Err:        errors.New(errMsg),
	}
}
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: Status Code: [%d], Error Message: %s", e.StatusCode, e.Err)
}
