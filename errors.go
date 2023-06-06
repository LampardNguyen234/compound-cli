package main

import "fmt"

const (
	errUnexpected = iota
	errClient
	errInvalidDuration
	errCreateTx

	errInvalidPrivateKey
	errInvalidAddress
	errInvalidValidatorAddress
)

var errCodeMessages = map[int]struct {
	Code    int
	Message string
}{
	errUnexpected:              {-1000, "unexpected error"},
	errClient:                  {-1001, "client error"},
	errInvalidDuration:         {-1002, "invalid duration"},
	errCreateTx:                {-1003, "failed to create transaction"},
	errInvalidPrivateKey:       {-2000, "invalid private key"},
	errInvalidAddress:          {-2001, "invalid address"},
	errInvalidValidatorAddress: {-2002, "invalid validator address"},
}

type appError struct {
	Code    int
	Message string
	Err     error
}

// Error satisfies the error interface and prints human-readable errors.
func (e appError) Error() error {
	if e.Err != nil {
		return fmt.Errorf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Errorf("[%d] %s", e.Code, e.Message)
}

func newAppError(key int, err ...error) error {
	res := appError{
		Code:    errCodeMessages[key].Code,
		Message: errCodeMessages[key].Message,
	}

	if len(err) > 0 {
		res.Err = err[0]
	}

	return res.Error()
}
