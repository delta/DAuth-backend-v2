package impl

import (
	"errors"
	"net/http"
)

// Response error response
type Response struct {
	Error       error
	ErrorCode   int
	Description string
	URI         string
	StatusCode  int
	Header      http.Header
}

// NewResponse create the response pointer
func NewResponse(err error, statusCode int) *Response {
	return &Response{
		Error:      err,
		StatusCode: statusCode,
	}
}

// SetHeader sets the header entries associated with key to
// the single element value.
func (r *Response) SetHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)
}

var (
	ErrInvalidRequest      = errors.New("Invalid request")
	ErrUserNotFound        = errors.New("User not found")
	ErrInternalServerError = errors.New("Internal server error")
)

// Descriptions error description
var Descriptions = map[error]string{
	ErrInvalidRequest:      "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed",
	ErrUserNotFound:        "The user is not is not found",
	ErrInternalServerError: "Internal server error",
}

// StatusCodes response error HTTP status code
var StatusCodes = map[error]int{
	ErrInvalidRequest:      400,
	ErrUserNotFound:        401,
	ErrInternalServerError: 500,
}
