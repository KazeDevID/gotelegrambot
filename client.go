package gotelegrambot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

// makeRequest is the internal method that actually performs the API request.
func (b *Bot) makeRequest(ctx context.Context, method string, params interface{}, result interface{}) error {
	// Prepare URL
	url := b.APIEndpoint + "/" + method

	var req *http.Request
	var err error

	// Create the request
	if params != nil {
		jsonData, err := json.Marshal(params)
		if err != nil {
			return errors.Wrap(err, "failed to marshal request params")
		}

		req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
		if err != nil {
			return errors.Wrap(err, "failed to create request")
		}

		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return errors.Wrap(err, "failed to create request")
		}
	}

	// Make the request with retries
	var resp *http.Response
	var lastErr error

	for i := 0; i <= b.retryCount; i++ {
		if i > 0 {
			// Wait before retrying
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Duration(i) * time.Second):
				// Exponential backoff
			}
		}

		resp, err = b.Client.Do(req)
		if err == nil {
			break
		}

		lastErr = err
		b.debug("Request failed (attempt %d/%d): %v", i+1, b.retryCount+1, err)
	}

	if err != nil {
		return errors.Wrap(lastErr, "request failed after retries")
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	b.debug("Response: %s", string(body))

	// Parse response
	if err := ParseAPIResponse(body, result); err != nil {
		var apiErr *Error
		if errors.As(err, &apiErr) {
			apiErr.Response = resp
		}
		return err
	}

	return nil
}

// makeMultipartRequest makes a multipart request for uploading files.
func (b *Bot) makeMultipartRequest(ctx context.Context, method string, params map[string]interface{}, files map[string]string, result interface{}) error {
	// Prepare URL
	url := b.APIEndpoint + "/" + method

	// Create multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add files to the multipart writer
	for field, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			return errors.Wrapf(err, "failed to open file %s", filePath)
		}
		defer file.Close()

		part, err := writer.CreateFormFile(field, filepath.Base(filePath))
		if err != nil {
			return errors.Wrapf(err, "failed to create form file for %s", filePath)
		}

		if _, err := io.Copy(part, file); err != nil {
			return errors.Wrapf(err, "failed to copy file content for %s", filePath)
		}
	}

	// Add other parameters
	for key, value := range params {
		if value == nil {
			continue
		}

		switch v := value.(type) {
		case string:
			if err := writer.WriteField(key, v); err != nil {
				return errors.Wrapf(err, "failed to write field %s", key)
			}
		case []byte:
			if err := writer.WriteField(key, string(v)); err != nil {
				return errors.Wrapf(err, "failed to write field %s", key)
			}
		default:
			// For complex types, marshal to JSON
			jsonValue, err := json.Marshal(v)
			if err != nil {
				return errors.Wrapf(err, "failed to marshal value for field %s", key)
			}
			if err := writer.WriteField(key, string(jsonValue)); err != nil {
				return errors.Wrapf(err, "failed to write field %s", key)
			}
		}
	}

	// Close multipart writer
	if err := writer.Close(); err != nil {
		return errors.Wrap(err, "failed to close multipart writer")
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request with retries
	var resp *http.Response
	var lastErr error

	for i := 0; i <= b.retryCount; i++ {
		if i > 0 {
			// Wait before retrying
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Duration(i) * time.Second):
				// Exponential backoff
			}
		}

		resp, err = b.Client.Do(req)
		if err == nil {
			break
		}

		lastErr = err
		b.debug("Request failed (attempt %d/%d): %v", i+1, b.retryCount+1, err)
	}

	if err != nil {
		return errors.Wrap(lastErr, "request failed after retries")
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	b.debug("Response: %s", string(responseBody))

	// Parse response
	if err := ParseAPIResponse(responseBody, result); err != nil {
		var apiErr *Error
		if errors.As(err, &apiErr) {
			apiErr.Response = resp
		}
		return err
	}

	return nil
}

// IsInputFile checks if the interface is meant to be uploaded as a file.
func IsInputFile(file interface{}) bool {
	switch f := file.(type) {
	case string:
		// If it starts with file://, it's a local file
		return len(f) > 7 && f[:7] == "file://"
	case *os.File:
		return true
	case io.Reader:
		return true
	default:
		return false
	}
}

// GetFile gets information about a file.
func (b *Bot) GetFile(ctx context.Context, fileID string) (*File, error) {
	params := map[string]interface{}{
		"file_id": fileID,
	}

	var file File
	err := b.makeRequest(ctx, "getFile", params, &file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file")
	}

	// Set file URL
	file.URL = fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.Token, file.FilePath)

	return &file, nil
}

// File represents a file ready to be downloaded.
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
	URL          string `json:"-"` // URL to download the file
}

// DownloadFile downloads a file to the specified path.
func (b *Bot) DownloadFile(ctx context.Context, file *File, destPath string) error {
	if file.URL == "" {
		return errors.New("file URL is empty")
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, file.URL, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// Make request
	resp, err := b.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to download file")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	// Create destination file
	out, err := os.Create(destPath)
	if err != nil {
		return errors.Wrap(err, "failed to create destination file")
	}
	defer out.Close()

	// Copy response body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to write file")
	}

	return nil
}