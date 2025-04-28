package gotelegrambot

import (
	"context"
)

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// AnswerCallbackQuery sends an answer to a callback query.
func (b *Bot) AnswerCallbackQuery(ctx context.Context, callbackQueryID string, options ...AnswerCallbackQueryOption) error {
	params := map[string]interface{}{
		"callback_query_id": callbackQueryID,
	}
	
	opts := defaultAnswerCallbackQueryOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	if opts.Text != "" {
		params["text"] = opts.Text
	}
	
	if opts.ShowAlert {
		params["show_alert"] = true
	}
	
	if opts.URL != "" {
		params["url"] = opts.URL
	}
	
	if opts.CacheTime != 0 {
		params["cache_time"] = opts.CacheTime
	}
	
	return b.makeRequest(ctx, "answerCallbackQuery", params, nil)
}

// AnswerCallbackQueryOption is a function that configures AnswerCallbackQuery options.
type AnswerCallbackQueryOption func(*answerCallbackQueryOptions)

// answerCallbackQueryOptions represents options for AnswerCallbackQuery.
type answerCallbackQueryOptions struct {
	Text      string
	ShowAlert bool
	URL       string
	CacheTime int
}

// WithCallbackText sets the text for the answer.
func WithCallbackText(text string) AnswerCallbackQueryOption {
	return func(o *answerCallbackQueryOptions) {
		o.Text = text
	}
}

// WithShowAlert sets whether to show an alert instead of a notification.
func WithShowAlert(showAlert bool) AnswerCallbackQueryOption {
	return func(o *answerCallbackQueryOptions) {
		o.ShowAlert = showAlert
	}
}

// WithCallbackURL sets the URL to open.
func WithCallbackURL(url string) AnswerCallbackQueryOption {
	return func(o *answerCallbackQueryOptions) {
		o.URL = url
	}
}

// WithCallbackCacheTime sets the cache time for the answer.
func WithCallbackCacheTime(cacheTime int) AnswerCallbackQueryOption {
	return func(o *answerCallbackQueryOptions) {
		o.CacheTime = cacheTime
	}
}

func defaultAnswerCallbackQueryOptions() answerCallbackQueryOptions {
	return answerCallbackQueryOptions{}
}

// EditMessageText edits a text message.
func (b *Bot) EditMessageText(ctx context.Context, options ...EditMessageTextOption) (*Message, error) {
	params := map[string]interface{}{}
	
	opts := defaultEditMessageTextOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	// Fill in params from options
	if opts.ChatID != 0 {
		params["chat_id"] = opts.ChatID
	}
	
	if opts.MessageID != 0 {
		params["message_id"] = opts.MessageID
	}
	
	if opts.InlineMessageID != "" {
		params["inline_message_id"] = opts.InlineMessageID
	}
	
	params["text"] = opts.Text
	
	if opts.ParseMode != "" {
		params["parse_mode"] = opts.ParseMode
	}
	
	if len(opts.Entities) > 0 {
		params["entities"] = opts.Entities
	}
	
	if opts.DisableWebPagePreview {
		params["disable_web_page_preview"] = true
	}
	
	if opts.ReplyMarkup != nil {
		params["reply_markup"] = opts.ReplyMarkup
	}
	
	var message Message
	err := b.makeRequest(ctx, "editMessageText", params, &message)
	if err != nil {
		return nil, err
	}
	
	return &message, nil
}

// EditMessageTextOption is a function that configures EditMessageText options.
type EditMessageTextOption func(*editMessageTextOptions)

// editMessageTextOptions represents options for EditMessageText.
type editMessageTextOptions struct {
	ChatID                int64
	MessageID             int
	InlineMessageID       string
	Text                  string
	ParseMode             string
	Entities              []MessageEntity
	DisableWebPagePreview bool
	ReplyMarkup           interface{}
}

// WithChatID sets the chat ID for editing a message.
func WithChatID(chatID int64) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.ChatID = chatID
	}
}

// WithMessageID sets the message ID for editing a message.
func WithMessageID(messageID int) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.MessageID = messageID
	}
}

// WithInlineMessageID sets the inline message ID for editing a message.
func WithInlineMessageID(inlineMessageID string) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.InlineMessageID = inlineMessageID
	}
}

// WithText sets the text for editing a message.
func WithText(text string) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.Text = text
	}
}

// WithEditParseMode sets the parse mode for editing a message.
func WithEditParseMode(parseMode string) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.ParseMode = parseMode
	}
}

// WithEditEntities sets the entities for editing a message.
func WithEditEntities(entities []MessageEntity) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.Entities = entities
	}
}

// WithEditDisableWebPagePreview disables link previews for editing a message.
func WithEditDisableWebPagePreview(disable bool) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.DisableWebPagePreview = disable
	}
}

// WithEditReplyMarkup sets the reply markup for editing a message.
func WithEditReplyMarkup(markup interface{}) EditMessageTextOption {
	return func(o *editMessageTextOptions) {
		o.ReplyMarkup = markup
	}
}

func defaultEditMessageTextOptions() editMessageTextOptions {
	return editMessageTextOptions{}
}

// DeleteMessage deletes a message.
func (b *Bot) DeleteMessage(ctx context.Context, chatID interface{}, messageID int) error {
	params := map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	}
	
	return b.makeRequest(ctx, "deleteMessage", params, nil)
}

// SendChatAction sends a chat action.
func (b *Bot) SendChatAction(ctx context.Context, chatID interface{}, action string) error {
	params := map[string]interface{}{
		"chat_id": chatID,
		"action":  action,
	}
	
	return b.makeRequest(ctx, "sendChatAction", params, nil)
}

// ChatActionType represents the type of chat action.
const (
	ChatActionTyping          = "typing"
	ChatActionUploadPhoto     = "upload_photo"
	ChatActionRecordVideo     = "record_video"
	ChatActionUploadVideo     = "upload_video"
	ChatActionRecordVoice     = "record_voice"
	ChatActionUploadVoice     = "upload_voice"
	ChatActionUploadDocument  = "upload_document"
	ChatActionChooseSticker   = "choose_sticker"
	ChatActionFindLocation    = "find_location"
	ChatActionRecordVideoNote = "record_video_note"
	ChatActionUploadVideoNote = "upload_video_note"
)

// ForwardMessage forwards a message.
func (b *Bot) ForwardMessage(ctx context.Context, chatID interface{}, fromChatID interface{}, messageID int, options ...ForwardMessageOption) (*Message, error) {
	params := map[string]interface{}{
		"chat_id":      chatID,
		"from_chat_id": fromChatID,
		"message_id":   messageID,
	}
	
	opts := defaultForwardMessageOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	if opts.DisableNotification {
		params["disable_notification"] = true
	}
	
	if opts.ProtectContent {
		params["protect_content"] = true
	}
	
	var message Message
	err := b.makeRequest(ctx, "forwardMessage", params, &message)
	if err != nil {
		return nil, err
	}
	
	return &message, nil
}

// ForwardMessageOption is a function that configures ForwardMessage options.
type ForwardMessageOption func(*forwardMessageOptions)

// forwardMessageOptions represents options for ForwardMessage.
type forwardMessageOptions struct {
	DisableNotification bool
	ProtectContent      bool
}

// WithForwardDisableNotification disables notifications for forwarding a message.
func WithForwardDisableNotification(disable bool) ForwardMessageOption {
	return func(o *forwardMessageOptions) {
		o.DisableNotification = disable
	}
}

// WithForwardProtectContent protects the content of the forwarded message.
func WithForwardProtectContent(protect bool) ForwardMessageOption {
	return func(o *forwardMessageOptions) {
		o.ProtectContent = protect
	}
}

func defaultForwardMessageOptions() forwardMessageOptions {
	return forwardMessageOptions{}
}

// CopyMessage copies a message.
func (b *Bot) CopyMessage(ctx context.Context, chatID interface{}, fromChatID interface{}, messageID int, options ...CopyMessageOption) (*MessageID, error) {
	params := map[string]interface{}{
		"chat_id":      chatID,
		"from_chat_id": fromChatID,
		"message_id":   messageID,
	}
	
	opts := defaultCopyMessageOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	// Fill in params from options
	if opts.Caption != "" {
		params["caption"] = opts.Caption
	}
	
	if opts.ParseMode != "" {
		params["parse_mode"] = opts.ParseMode
	}
	
	if len(opts.CaptionEntities) > 0 {
		params["caption_entities"] = opts.CaptionEntities
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
	
	var messageID MessageID
	err := b.makeRequest(ctx, "copyMessage", params, &messageID)
	if err != nil {
		return nil, err
	}
	
	return &messageID, nil
}

// MessageID represents a unique message identifier.
type MessageID struct {
	MessageID int `json:"message_id"`
}

// CopyMessageOption is a function that configures CopyMessage options.
type CopyMessageOption func(*copyMessageOptions)

// copyMessageOptions represents options for CopyMessage.
type copyMessageOptions struct {
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

// Various option functions for CopyMessage would be defined here,
// similar to the other option functions above.

func defaultCopyMessageOptions() copyMessageOptions {
	return copyMessageOptions{}
}

// Additional methods for advanced features would be implemented here,
// such as:
// - User and chat management
// - Poll creation
// - Payment methods
// - Game methods
// - etc.