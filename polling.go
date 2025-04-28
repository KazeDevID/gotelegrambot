package gotelegrambot

import (
	"context"
	"fmt"
	"time"
)

// UpdateHandler is a function that handles an update.
type UpdateHandler func(ctx context.Context, update *Update) error

// StartPolling starts long polling for updates.
func (b *Bot) StartPolling(ctx context.Context, handler UpdateHandler, options ...PollingOption) error {
	if handler == nil {
		return fmt.Errorf("update handler is required")
	}

	b.mu.Lock()
	b.updateHandler = handler
	b.mu.Unlock()

	opts := defaultPollingOptions()
	for _, opt := range options {
		opt(&opts)
	}

	go b.startPollingLoop(ctx, opts)
	return nil
}

// PollingOptions represents options for polling.
type pollingOptions struct {
	Timeout  int
	Limit    int
	Offset   int
	AllowedUpdates []string
	PollInterval time.Duration
}

// PollingOption is a function that configures polling options.
type PollingOption func(*pollingOptions)

// WithTimeout sets the timeout for long polling.
func WithTimeout(timeout int) PollingOption {
	return func(o *pollingOptions) {
		o.Timeout = timeout
	}
}

// WithLimit sets the limit of updates to fetch per request.
func WithLimit(limit int) PollingOption {
	return func(o *pollingOptions) {
		o.Limit = limit
	}
}

// WithOffset sets the initial offset for updates.
func WithOffset(offset int) PollingOption {
	return func(o *pollingOptions) {
		o.Offset = offset
	}
}

// WithAllowedUpdates sets the types of updates to receive.
func WithAllowedUpdates(updates []string) PollingOption {
	return func(o *pollingOptions) {
		o.AllowedUpdates = updates
	}
}

// WithPollInterval sets the interval between polling requests.
func WithPollInterval(interval time.Duration) PollingOption {
	return func(o *pollingOptions) {
		o.PollInterval = interval
	}
}

func defaultPollingOptions() pollingOptions {
	return pollingOptions{
		Timeout:  60,
		Limit:    100,
		Offset:   0,
		AllowedUpdates: []string{},
		PollInterval: 100 * time.Millisecond,
	}
}

// startPollingLoop starts a loop that polls for updates.
func (b *Bot) startPollingLoop(ctx context.Context, opts pollingOptions) {
	b.debug("Starting polling loop")
	
	offset := opts.Offset
	
	for {
		select {
		case <-ctx.Done():
			b.debug("Polling loop canceled: %v", ctx.Err())
			return
		case <-b.shutdownChan:
			b.debug("Polling loop stopped by shutdown")
			return
		default:
			updates, err := b.getUpdates(ctx, offset, opts.Limit, opts.Timeout, opts.AllowedUpdates)
			if err != nil {
				b.debug("Error getting updates: %v", err)
				time.Sleep(opts.PollInterval)
				continue
			}
			
			for _, update := range updates {
				if update.UpdateID >= offset {
					offset = update.UpdateID + 1
				}
				
				go func(update Update) {
					err := b.processUpdate(ctx, &update)
					if err != nil {
						b.debug("Error processing update: %v", err)
					}
				}(update)
			}
			
			if len(updates) == 0 {
				time.Sleep(opts.PollInterval)
			}
		}
	}
}

// getUpdates gets updates from the Telegram API.
func (b *Bot) getUpdates(ctx context.Context, offset, limit, timeout int, allowedUpdates []string) ([]Update, error) {
	params := map[string]interface{}{
		"offset":  offset,
		"limit":   limit,
		"timeout": timeout,
	}
	
	if len(allowedUpdates) > 0 {
		params["allowed_updates"] = allowedUpdates
	}
	
	// TODO: Implement the actual API call
	// This is a placeholder
	return []Update{}, nil
}

// processUpdate processes a single update.
func (b *Bot) processUpdate(ctx context.Context, update *Update) error {
	b.mu.RLock()
	handler := b.updateHandler
	b.mu.RUnlock()
	
	if handler != nil {
		return handler(ctx, update)
	}
	
	return nil
}