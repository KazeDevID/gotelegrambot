package gotelegrambot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// APIResponse represents a response from the Telegram API.
type APIResponse struct {
	OK          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// ResponseParameters contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

// Error represents an error from the Telegram API.
type Error struct {
	Code        int
	Message     string
	Parameters  *ResponseParameters
	Response    *http.Response
}

// Error returns a string representation of the error.
func (e *Error) Error() string {
	return fmt.Sprintf("telegram: %d %s", e.Code, e.Message)
}

// ParseAPIResponse parses an API response into a generic struct.
func ParseAPIResponse(data []byte, target interface{}) error {
	var resp APIResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return errors.Wrap(err, "failed to parse API response")
	}
	
	if !resp.OK {
		return &Error{
			Code:       resp.ErrorCode,
			Message:    resp.Description,
			Parameters: resp.Parameters,
		}
	}
	
	if target == nil {
		return nil
	}
	
	if err := json.Unmarshal(resp.Result, target); err != nil {
		return errors.Wrap(err, "failed to parse result")
	}
	
	return nil
}

// RetryableError checks if an error is retryable.
func RetryableError(err error) bool {
	if err == nil {
		return false
	}
	
	var apiErr *Error
	if errors.As(err, &apiErr) {
		// Retry on rate limit errors or internal server errors
		return apiErr.Code == 429 || (apiErr.Code >= 500 && apiErr.Code < 600)
	}
	
	// Retry on network errors, etc.
	return true
}

// WaitTime returns the time to wait before retrying if there's a retry-after header.
func (e *Error) WaitTime() int {
	if e.Parameters != nil && e.Parameters.RetryAfter > 0 {
		return e.Parameters.RetryAfter
	}
	return 0
}