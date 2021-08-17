package github

type Err struct {
	err error
}

func (e *Err) Error() string {
	return "github: " + e.err.Error()
}

func (e *Err) Unwrap() error {
	return e.err
}

type ErrInvalidToken struct {
	err error
}

func (e *ErrInvalidToken) Error() string {
	return "invalid token; " + e.err.Error()
}

func (e *ErrInvalidToken) Unwrap() error {
	return e.err
}
