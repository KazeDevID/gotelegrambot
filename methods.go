package gotelegrambot

import (
	"context"
	"fmt"
)

// SendMessage sends a text message.
func (b *Bot) SendMessage(ctx context.Context, chatID interface{}, text string, options ...SendMessageOption) (*Message, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}
	
	opts := defaultSendMessageOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	if opts.ParseMode != "" {
		params["parse_mode"] = opts.ParseMode
	}
	
	if len(opts.Entities) > 0 {
		params["entities"] = opts.Entities
	}
	
	if opts.DisableWebPagePreview {
		params["disable_web_page_preview"] = true
	}
	
	if opts.DisableNotification {
		params["disable_notification"] = true
	}
	
	if opts.ProtectContent {
		params["protect_content"] = true
	}
	
	if opts.ReplyToMessageID != 0 {
		params["reply_to_message_id"] = opts.ReplyToMessageID
	}
	
	if opts.AllowSendingWithoutReply {
		params["allow_sending_without_reply"] = true
	}
	
	if opts.ReplyMarkup != nil {
		params["reply_markup"] = opts.ReplyMarkup
	}
	
	var message Message
	// TODO: Make the actual API request and parse response
	// This is a placeholder
	return &message, nil
}

// sendMessageOptions represents options for SendMessage.
type sendMessageOptions struct {
	ParseMode                string
	Entities                 []MessageEntity
	DisableWebPagePreview    bool
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

// SendMessageOption is a function that configures SendMessage options.
type SendMessageOption func(*sendMessageOptions)

// WithParseMode sets the parse mode for the message.
func WithParseMode(parseMode string) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.ParseMode = parseMode
	}
}

// WithEntities sets the entities for the message.
func WithEntities(entities []MessageEntity) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.Entities = entities
	}
}

// WithDisableWebPagePreview disables link previews.
func WithDisableWebPagePreview(disable bool) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.DisableWebPagePreview = disable
	}
}

// WithDisableNotification disables notifications for the message.
func WithDisableNotification(disable bool) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.DisableNotification = disable
	}
}

// WithProtectContent protects the content of the message.
func WithProtectContent(protect bool) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.ProtectContent = protect
	}
}

// WithReplyToMessageID sets the message to reply to.
func WithReplyToMessageID(messageID int) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.ReplyToMessageID = messageID
	}
}

// WithAllowSendingWithoutReply allows sending messages that are replies without the original message.
func WithAllowSendingWithoutReply(allow bool) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.AllowSendingWithoutReply = allow
	}
}

// WithReplyMarkup sets the reply markup for the message.
func WithReplyMarkup(markup interface{}) SendMessageOption {
	return func(o *sendMessageOptions) {
		o.ReplyMarkup = markup
	}
}

func defaultSendMessageOptions() sendMessageOptions {
	return sendMessageOptions{}
}

// SendPhoto sends a photo.
func (b *Bot) SendPhoto(ctx context.Context, chatID interface{}, photo interface{}, options ...SendPhotoOption) (*Message, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
		"photo":   photo,
	}
	
	opts := defaultSendPhotoOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	// Fill in the params based on options
	// ...
	
	var message Message
	// TODO: Make the actual API request and parse response
	return &message, nil
}

// SendPhotoOption is a function that configures SendPhoto options.
type SendPhotoOption func(*sendPhotoOptions)

// sendPhotoOptions represents options for SendPhoto.
type sendPhotoOptions struct {
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	HasSpoiler               bool
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func defaultSendPhotoOptions() sendPhotoOptions {
	return sendPhotoOptions{}
}

// Additional message sending methods would be implemented here,
// following the same pattern as SendMessage and SendPhoto.
// This includes methods for sending videos, audio, documents,
// locations, venues, contacts, etc.