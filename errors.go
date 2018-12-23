package robinhood

import (
	"fmt"
	"strings"
)

type APIError struct {
	StatusCode int
	Errors     map[string][]string
}

func (e APIError) Error() string {
	var errors []string
	for k, vs := range e.Errors {
		err := fmt.Sprintf("%v: %v", k, vs)
		errors = append(errors, err)
	}

	errMsg := strings.Join(errors, "; ")
	return fmt.Sprintf("%v: %v", e.StatusCode, errMsg)
}
