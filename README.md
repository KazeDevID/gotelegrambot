# Go Telegram Bot API

A comprehensive Go library for the Telegram Bot API that follows modern Go idioms and provides a clean, expressive interface for building Telegram bots.

## Features

- Complete implementation of the Telegram Bot API
- Support for both polling and webhook methods
- Context-based API for proper cancellation
- Error handling with retry mechanisms
- Type-safe interfaces
- Fluent builder-style API for complex parameters
- Comprehensive documentation and examples

## Installation

```bash
go get github.com/KazeDevID/gotelegrambot
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KazeDevID/gotelegrambot"
)

func main() {
	// Create a new bot with your token
	bot, err := gotelegrambot.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	// Create a context with cancel function for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start polling for updates
	err = bot.StartPolling(ctx, func(ctx context.Context, update *gotelegrambot.Update) error {
		// Handle text messages
		if update.Message != nil && update.Message.Text != "" {
			_, err := bot.SendMessage(ctx, update.Message.Chat.ID,
				fmt.Sprintf("You said: %s", update.Message.Text))
			return err
		}
		return nil
	})
	
	if err != nil {
		log.Fatalf("Error starting polling: %v", err)
	}

	// Wait for termination signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// Cancel the context to stop polling and clean up
	cancel()
}
```

## Message Sending

### Text Messages

```go
// Basic text message
_, err := bot.SendMessage(ctx, chatID, "Hello, world!")

// With Markdown formatting
_, err := bot.SendMessage(ctx, chatID, "Hello, *world*!",
	gotelegrambot.WithParseMode("Markdown"))

// With HTML formatting
_, err := bot.SendMessage(ctx, chatID, "<b>Hello</b>, world!",
	gotelegrambot.WithParseMode("HTML"))
```

### Media Messages

```go
// Send a photo
_, err := bot.SendPhoto(ctx, chatID, "https://example.com/image.jpg",
	gotelegrambot.WithCaption("Check out this photo!"))

// Send a document
_, err := bot.SendDocument(ctx, chatID, "/path/to/file.pdf")
```

### Keyboards

```go
// Inline keyboard
markup := gotelegrambot.NewInlineKeyboardMarkup(
	gotelegrambot.NewInlineKeyboardButtonRow(
		gotelegrambot.NewInlineKeyboardButtonURL("Visit Website", "https://example.com"),
		gotelegrambot.NewInlineKeyboardButtonCallback("Click Me", "callback_data"),
	),
)

_, err := bot.SendMessage(ctx, chatID, "Choose an option:",
	gotelegrambot.WithReplyMarkup(markup))

// Reply keyboard
keyboard := gotelegrambot.NewReplyKeyboardMarkup(
	[][]gotelegrambot.KeyboardButton{
		{{Text: "Option 1"}, {Text: "Option 2"}},
		{{Text: "Option 3"}, {Text: "Option 4"}},
	},
	gotelegrambot.WithResizeKeyboard(true),
)

_, err := bot.SendMessage(ctx, chatID, "Select an option:",
	gotelegrambot.WithReplyMarkup(keyboard))
```

## Webhook Setup

```go
// Set webhook
err := bot.SetWebhook(ctx, gotelegrambot.WebhookConfig{
	URL: "https://example.com/webhook",
})

// Create webhook handler
http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
	bot.WebhookHandler(handleUpdate).ServeHTTP(w, r)
})

// Start HTTP server
http.ListenAndServe(":8080", nil)
```

## Error Handling

```go
_, err := bot.SendMessage(ctx, chatID, "Hello")
if err != nil {
	var apiErr *gotelegrambot.Error
	if errors.As(err, &apiErr) {
		log.Printf("API Error %d: %s", apiErr.Code, apiErr.Message)
		
		// Check if we need to retry
		if gotelegrambot.RetryableError(err) {
			// Handle retry logic
		}
	} else {
		log.Printf("Other error: %v", err)
	}
}
```

## Documentation

For full documentation and examples, see the [Go package documentation](https://kazedevid.github.io/gotelegrambot/).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
