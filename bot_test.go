package gotelegrambot

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Test with valid token
	bot, err := New("valid_token")
	assert.NoError(t, err)
	assert.NotNil(t, bot)
	assert.Equal(t, "valid_token", bot.Token)
	assert.Equal(t, APIBaseURL+"valid_token", bot.APIEndpoint)
	assert.NotNil(t, bot.Client)
	assert.Equal(t, DefaultRetryCount, bot.retryCount)
	
	// Test with empty token
	bot, err = New("")
	assert.Error(t, err)
	assert.Nil(t, bot)
	
	// Test with options
	customClient := &http.Client{Timeout: 30 * time.Second}
	bot, err = New("valid_token",
		WithHTTPClient(customClient),
		WithDebug(true),
		WithBuffer(200),
		WithRetryCount(5),
	)
	assert.NoError(t, err)
	assert.NotNil(t, bot)
	assert.Equal(t, customClient, bot.Client)
	assert.True(t, bot.Debug)
	assert.Equal(t, 200, bot.Buffer)
	assert.Equal(t, 5, bot.retryCount)
}

func TestParseAPIResponse(t *testing.T) {
	// Test successful response
	successJSON := `{"ok":true,"result":{"id":123,"first_name":"Test","username":"test_bot"}}`
	var user User
	err := ParseAPIResponse([]byte(successJSON), &user)
	assert.NoError(t, err)
	assert.Equal(t, int64(123), user.ID)
	assert.Equal(t, "Test", user.FirstName)
	assert.Equal(t, "test_bot", user.Username)
	
	// Test error response
	errorJSON := `{"ok":false,"error_code":400,"description":"Bad Request: chat not found"}`
	err = ParseAPIResponse([]byte(errorJSON), &user)
	assert.Error(t, err)
	var apiErr *Error
	assert.ErrorAs(t, err, &apiErr)
	assert.Equal(t, 400, apiErr.Code)
	assert.Equal(t, "Bad Request: chat not found", apiErr.Message)
	
	// Test invalid JSON
	invalidJSON := `{"ok":true,"result":`
	err = ParseAPIResponse([]byte(invalidJSON), &user)
	assert.Error(t, err)
}

func TestRetryableError(t *testing.T) {
	// Test nil error
	assert.False(t, RetryableError(nil))
	
	// Test API errors
	rateLimit := &Error{Code: 429, Message: "Too Many Requests"}
	assert.True(t, RetryableError(rateLimit))
	
	serverError := &Error{Code: 500, Message: "Internal Server Error"}
	assert.True(t, RetryableError(serverError))
	
	clientError := &Error{Code: 400, Message: "Bad Request"}
	assert.False(t, RetryableError(clientError))
}

func TestWebhookHandler(t *testing.T) {
	bot, _ := New("test_token")
	
	// Create a test handler
	var receivedUpdate *Update
	handler := func(ctx context.Context, update *Update) error {
		receivedUpdate = update
		return nil
	}
	
	// Create a test server
	server := httptest.NewServer(bot.WebhookHandler(handler))
	defer server.Close()
	
	// Test with valid update
	updateJSON := `{"update_id":123456,"message":{"message_id":1,"from":{"id":123,"first_name":"Test","is_bot":false},"chat":{"id":123,"first_name":"Test","type":"private"},"date":1600000000,"text":"Hello, world!"}}`
	
	resp, err := http.Post(server.URL, "application/json", []byte(updateJSON))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotNil(t, receivedUpdate)
	assert.Equal(t, 123456, receivedUpdate.UpdateID)
	assert.NotNil(t, receivedUpdate.Message)
	assert.Equal(t, "Hello, world!", receivedUpdate.Message.Text)
	
	// Test with invalid method
	resp, err = http.Get(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
	
	// Test with invalid JSON
	resp, err = http.Post(server.URL, "application/json", []byte(`{"update_id":123456`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// More tests would be defined here for various methods and functionality.