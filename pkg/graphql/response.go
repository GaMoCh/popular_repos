package graphql

import (
	"strings"
)

type response struct {
	Data   interface{}   `json:"data,omitempty"`
	Errors []ErrResponse `json:"errors,omitempty"`
}

type ErrResponse struct {
	Message string
}

func (e *ErrResponse) Error() string {
	return "graphql query: " + e.Message
}

func joinResponseErrors(errs []ErrResponse) error {
	messages := make([]string, len(errs))
	for i, err := range errs {
		messages[i] = err.Message
	}

	return &ErrResponse{
		Message: strings.Join(messages, "; "),
	}
}
