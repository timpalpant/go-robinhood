package robinhood

import (
	"fmt"
	"strings"
)

type APIError struct {
	StatusCode int
	Errors     map[string]interface{}
}

func (e APIError) Error() string {
	var errors []string
	for k, v := range e.Errors {
		err := fmt.Sprintf("%v: %v", k, v)
		errors = append(errors, err)
	}

	errMsg := strings.Join(errors, "; ")
	return fmt.Sprintf("%v: %v", e.StatusCode, errMsg)
}
