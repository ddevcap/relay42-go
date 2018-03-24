package relay42

import "net/http"

// ErrorResponse model
type ErrorResponse struct {
	ErrorCode int		`json:"errorCode"`
	Message   string    `json:"message,omitempty"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

// APIError model
type APIError struct {
	req *http.Request
	res *http.Response
	err *ErrorResponse
}

type BadRequestError struct {
	APIError
}

type UnauthorizedError struct {
	APIError
}

type ForbiddenError struct {
	APIError
}

type InternalServerError struct {
	APIError
}

func (e BadRequestError) Error() string {
	return e.err.Message
}

func (e UnauthorizedError) Error() string {
	return e.err.Message
}

func (e ForbiddenError) Error() string {
	return e.err.Message
}

func (e InternalServerError) Error() string {
	return e.err.Message
}

