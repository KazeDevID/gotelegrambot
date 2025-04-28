package gotelegrambot

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WebhookConfig represents webhook configuration options.
type WebhookConfig struct {
	URL                string
	Certificate        interface{}
	IPAddress          string
	MaxConnections     int
	AllowedUpdates     []string
	DropPendingUpdates bool
	SecretToken        string
}

// SetWebhook sets a webhook for receiving updates.
func (b *Bot) SetWebhook(ctx context.Context, config WebhookConfig) error {
	params := map[string]interface{}{
		"url": config.URL,
	}
	
	if config.Certificate != nil {
		params["certificate"] = config.Certificate
	}
	
	if config.IPAddress != "" {
		params["ip_address"] = config.IPAddress
	}
	
	if config.MaxConnections != 0 {
		params["max_connections"] = config.MaxConnections
	}
	
	if len(config.AllowedUpdates) > 0 {
		params["allowed_updates"] = config.AllowedUpdates
	}
	
	if config.DropPendingUpdates {
		params["drop_pending_updates"] = true
	}
	
	if config.SecretToken != "" {
		params["secret_token"] = config.SecretToken
	}
	
	// TODO: Make the actual API request
	// This is a placeholder
	return nil
}

// DeleteWebhook deletes the webhook.
func (b *Bot) DeleteWebhook(ctx context.Context, dropPendingUpdates bool) error {
	params := map[string]interface{}{}
	
	if dropPendingUpdates {
		params["drop_pending_updates"] = true
	}
	
	// TODO: Make the actual API request
	// This is a placeholder
	return nil
}

// GetWebhookInfo gets current webhook status.
func (b *Bot) GetWebhookInfo(ctx context.Context) (*WebhookInfo, error) {
	// TODO: Make the actual API request
	// This is a placeholder
	return nil, nil
}

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	IPAddress            string   `json:"ip_address,omitempty"`
	LastErrorDate        int      `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date,omitempty"`
	MaxConnections       int      `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

// WebhookHandler creates an http.Handler for processing webhook requests.
func (b *Bot) WebhookHandler(handler UpdateHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		body, err := io.ReadAll(r.Body)
		if err != nil {
			b.debug("Error reading request body: %v", err)
			http.Error(w, "Error reading request", http.StatusInternalServerError)
			return
		}
		
		var update Update
		if err := json.Unmarshal(body, &update); err != nil {
			b.debug("Error unmarshaling update: %v", err)
			http.Error(w, "Error parsing update", http.StatusBadRequest)
			return
		}
		
		if err := handler(r.Context(), &update); err != nil {
			b.debug("Error handling update: %v", err)
			http.Error(w, "Error processing update", http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusOK)
	})
}

// StartWebhookServer starts an HTTPS server for the webhook.
func (b *Bot) StartWebhookServer(addr string, certFile, keyFile string, handler UpdateHandler) error {
	server := &http.Server{
		Addr:    addr,
		Handler: b.WebhookHandler(handler),
	}
	
	return server.ListenAndServeTLS(certFile, keyFile)
}

// StartWebhookServerTLS starts an HTTPS server for the webhook with a custom TLS config.
func (b *Bot) StartWebhookServerTLS(addr string, tlsConfig *tls.Config, handler UpdateHandler) error {
	server := &http.Server{
		Addr:      addr,
		Handler:   b.WebhookHandler(handler),
		TLSConfig: tlsConfig,
	}
	
	return server.ListenAndServeTLS("", "")
}