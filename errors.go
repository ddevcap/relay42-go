package relay42

import "net/http"

// ErrorResponse holds error response data
type ErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message,omitempty"`
}

// Error returns the error message
func (e ErrorResponse) Error() string {
	return e.Message
}

// APIError holds the API error request, response and error
type APIError struct {
	req *http.Request
	res *http.Response
	err *ErrorResponse
}

// BadRequestError holds the API error
type BadRequestError struct {
	APIError
}

// UnauthorizedError holds the API error
type UnauthorizedError struct {
	APIError
}

// ForbiddenError holds the API error
type ForbiddenError struct {
	APIError
}

// InternalServerError holds the API error
type InternalServerError struct {
	APIError
}

// Error returns the error message
func (e BadRequestError) Error() string {
	return e.err.Message
}

// Error returns the error message
func (e UnauthorizedError) Error() string {
	return e.err.Message
}

// Error returns the error message
func (e ForbiddenError) Error() string {
	return e.err.Message
}

// Error returns the error message
func (e InternalServerError) Error() string {
	return e.err.Message
}
