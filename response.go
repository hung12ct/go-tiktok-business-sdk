package tiktokbiz

import (
	"encoding/json"
	"fmt"
)

// BaseResponse contains shared API response data fields
type BaseResponse struct {
	// Code is the return code
	Code int `json:"code"`
	// Message is the return message
	Message string `json:"message"`
	// RequestID is the request log ID, uniquely identifies a request
	RequestID string `json:"request_id,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Message)
}

type BaseResponseWithData struct {
	BaseResponse
	Data json.RawMessage `json:"data,omitempty"`
}

// PageInfo contains general pagination data.
type PageInfo struct {
	// Page is the current page number.
	Page int `json:"page,omitempty"`
	// PageSize is the number of items per page.
	PageSize int `json:"page_size,omitempty"`
	// TotalNumber is the total number of items.
	TotalNumber int64 `json:"total_number,omitempty"`
	// TotalPage is the total number of pages.
	TotalPage int `json:"total_page,omitempty"`
}
