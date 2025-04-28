/*
Package gotelegrambot provides a comprehensive Go library for the Telegram Bot API.

# Overview

This library provides a powerful interface to interact with the Telegram Bot API,
supporting both polling and webhook modes, context-based operations, and a complete
implementation of all API methods.

# Getting Started

To create a new bot:

	bot, err := gotelegrambot.New("YOUR_BOT_TOKEN", gotelegrambot.WithDebug(true))
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

# Long Polling

To receive updates using long polling:

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.StartPolling(ctx, func(ctx context.Context, update *gotelegrambot.Update) error {
		// Handle the update
		return nil
	}, gotelegrambot.WithTimeout(60))

# Webhook

To set up a webhook:

	err := bot.SetWebhook(ctx, gotelegrambot.WebhookConfig{
		URL: "https://example.com/webhook",
	})

To handle webhook updates:

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		bot.WebhookHandler(handleUpdate).ServeHTTP(w, r)
	})

# Sending Messages

To send a simple message:

	_, err := bot.SendMessage(ctx, chatID, "Hello, world!")

With formatting and reply markup:

	_, err := bot.SendMessage(ctx, chatID, "Hello, *world*!",
		gotelegrambot.WithParseMode("Markdown"),
		gotelegrambot.WithReplyMarkup(keyboard))

# Handling Media

To send a photo:

	_, err := bot.SendPhoto(ctx, chatID, "https://example.com/image.jpg",
		gotelegrambot.WithCaption("An example image"))

# Error Handling

The library provides detailed error information from the Telegram API:

	_, err := bot.SendMessage(ctx, chatID, "Hello")
	if err != nil {
		var apiErr *gotelegrambot.Error
		if errors.As(err, &apiErr) {
			log.Printf("API Error %d: %s", apiErr.Code, apiErr.Message)
		} else {
			log.Printf("Other error: %v", err)
		}
	}

# Cancellation

All API methods accept a context.Context, allowing for proper cancellation:

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := bot.SendMessage(ctx, chatID, "This request will timeout after 5 seconds")
*/
package gotelegrambot