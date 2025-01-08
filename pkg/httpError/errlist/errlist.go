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

	// ErrNotFound is used when requested data is not found.
	ErrNotFound = errors.New("not found")

	// ErrUnauthorized is used when user is not authorized.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden is used when user is forbidden to access.
	ErrForbidden = errors.New("forbidden")

	// ErrConflict is used when conflict occurs.
	ErrConflict = errors.New("conflict")

	// ErrRange is used when range is not valid.
	ErrRange = errors.New("range is not valid")
)
