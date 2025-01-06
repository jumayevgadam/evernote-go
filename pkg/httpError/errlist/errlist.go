package errlist

import "errors"

// We will use this package to keep all error messages in one place.
var (
	// ErrInternalServer is used when internal server error occurs.
	ErrInternalServer = errors.New("internal server error")

	// ErrBadRequest is used when bad request occurs.
	ErrBadRequest = errors.New("bad request")

	// ErrBadQueryParams is used when bad query params occurs.
	ErrBadQueryParams = errors.New("bad query params")
)
