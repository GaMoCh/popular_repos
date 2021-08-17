package graphql

import "errors"

type Err struct {
	err error
}

func (e *Err) Error() string {
	return "graphql: " + e.err.Error()
}

func (e *Err) Unwrap() error {
	return e.err
}

type ErrStatus struct {
	StatusCode int
	Status     string
}

func (e *ErrStatus) Error() string {
	return "wrong HTTP status: " + e.Status
}

var ErrNilRequest = errors.New("graphQL request not provided")
