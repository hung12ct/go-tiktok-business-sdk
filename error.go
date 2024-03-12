package tiktokbiz

import "fmt"

// APIError provides error information returned by the TikTok Business API.
type APIError struct {
	Code           any     `json:"code,omitempty"`
	Message        string  `json:"message"`
	Param          *string `json:"param,omitempty"`
	Type           string  `json:"type"`
	HTTPStatusCode int     `json:"-"`
}

// RequestError provides informations about generic request errors.
type RequestError struct {
	HTTPStatusCode int
	Err            error
}

type ErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	if e.HTTPStatusCode > 0 {
		return fmt.Sprintf("error, status code: %d, message: %s", e.HTTPStatusCode, e.Message)
	}

	return e.Message
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("error, status code: %d, message: %s", e.HTTPStatusCode, e.Err)
}
