package errors

import (
	"fmt"
)

type HttpError struct {
	Status  int
	Message string
}

func (err HttpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", err.Status, err.Message)
}
