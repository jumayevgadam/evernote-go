package httpError

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jumayevgadam/evernote-go/internal/config"
	"github.com/jumayevgadam/evernote-go/pkg/httpError/errlist"
	"github.com/jumayevgadam/evernote-go/pkg/logger"
)

// RestErr interface is used for handling errors.
type RestErr interface {
	Status() int
	Error() string
	Message() string
}

// RestError struct keeps error details.
type RestError struct {
	ErrStatus  int    `json:"status"`
	ErrKind    string `json:"kind"`
	ErrMessage string `json:"message"`
}

// Status returns error status.
func (e RestError) Status() int {
	return e.ErrStatus
}

// Error returns error kind.
func (e RestError) Error() string {
	return e.ErrMessage
}

// Message returns error message.
func (e RestError) Message() string {
	return fmt.Sprintf("%s: %s", e.ErrKind, e.ErrMessage)
}

// NewRestError returns new RestError.
func NewRestError(status int, kind, message string) RestErr {
	return &RestError{
		ErrStatus:  status,
		ErrKind:    kind,
		ErrMessage: message,
	}
}

// Predefined error constructors.
// NewBadRequest returns new bad request error.
func NewBadRequestError(message string) RestErr {
	return NewRestError(
		http.StatusBadRequest,
		errlist.ErrBadRequest.Error(),
		message,
	)
}

// NewInternalServer returns new internal server error.
func NewInternalServerError(message string) RestErr {
	return NewRestError(
		http.StatusInternalServerError,
		errlist.ErrInternalServer.Error(),
		message,
	)
}

// NewNotFound returns new not found error.
func NewNotFoundError(message string) RestErr {
	return NewRestError(
		http.StatusNotFound,
		errlist.ErrNotFound.Error(),
		message,
	)
}

// NewBadQueryParams returns new bad query params error.
func NewBadQueryParamsError(message string) RestErr {
	return NewRestError(
		http.StatusBadRequest,
		errlist.ErrBadQueryParams.Error(),
		message,
	)
}

// NewUnauthorized returns new unauthorized error.
func NewUnauthorizedError(message string) RestErr {
	return NewRestError(
		http.StatusUnauthorized,
		errlist.ErrUnauthorized.Error(),
		message,
	)
}

// NewForbidden returns new forbidden error.
func NewForbiddenError(message string) RestErr {
	return NewRestError(
		http.StatusForbidden,
		errlist.ErrForbidden.Error(),
		message,
	)
}

// NewConflict returns new conflict error.
func NewConflictError(message string) RestErr {
	return NewRestError(
		http.StatusConflict,
		errlist.ErrConflict.Error(),
		message,
	)
}

// ParseError returns error based on error kind.
func ParseError(err error) RestErr {
	switch {
	// pgx errors.
	case errors.Is(err, pgx.ErrNoRows):
		return NewNotFoundError(err.Error())
	case errors.Is(err, pgx.ErrTooManyRows):
		return NewConflictError(err.Error())
	case errors.Is(err, pgx.ErrTxClosed),
		errors.Is(err, pgx.ErrTxCommitRollback):
		return NewInternalServerError(err.Error())
	// SQLSTATE errors.
	case strings.Contains(err.Error(), "SQLSTATE"):
		return ParseSQLError(err)

	// validation errors.
	case errors.As(err, &validator.ValidationErrors{}):
		return ParseValidationError(err)

	default:
		// If already a RestErr, return as-is.
		var restErr RestErr
		if errors.As(err, &restErr) {
			return restErr
		}

		return NewInternalServerError(err.Error())
	}
}

// ParseSQLError returns error based on SQLSTATE.
func ParseSQLError(err error) RestErr {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// CLASS 22
		case "22001": // numeric value out of range
			return NewBadRequestError("Numeric" + errlist.ErrRange.Error() + pgErr.Message + "\n")

		// CLASS 23
		case "23505": // Unique violation
			return NewConflictError("Unique constraint violation: " + errlist.ErrConflict.Error() + pgErr.Message + "\n")
		case "23503": // Foreign key violation
			return NewBadRequestError("Foreign key violation: " + pgErr.Message + "\n")
		case "23502": // Not null violation
			return NewBadRequestError("Not null violation: " + pgErr.Message + "\n")

		// CLASS 40
		case "40001": // serialization failure
			return NewConflictError("Serialization error: " + pgErr.Message + "\n")
		// CLASS 42
		case "42601": // syntax error
			return NewBadRequestError("Syntax error in sql statements: " + pgErr.Message + "\n")

		}
	}

	if strings.Contains(err.Error(), "no corresponding field found") {
		return NewBadRequestError(err.Error() + "\n")
	}

	return NewBadRequestError(err.Error() + "\n")
}

// ParseValidationError returns error based on validation error.
func ParseValidationError(err error) RestErr {
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		return NewBadRequestError(err.Error())
	}

	errorMsgs := make([]string, 0, len(validationErrs))

	for _, fieldErr := range validationErrs {
		errorMsgs = append(errorMsgs, fmt.Sprintf("Field: %s, Error: %s", fieldErr.Field(), fieldErr.Tag()))
	}

	return NewBadRequestError(strings.Join(errorMsgs, "\n"))
}

// Response returns ErrorResponse, for clean syntax I took function name Response.
func Response(c *gin.Context, err error) {
	logger := logger.NewAPILogger(&config.Config{})
	logger.InitLogger()

	errStatus, errResponse := ParseError(err).Status(), ParseError(err).Message()
	c.JSON(errStatus, errResponse)
}
