package gotelegrambot

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
}

// KeyboardButton represents one button of the reply keyboard.
type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestContact  bool                        `json:"request_contact,omitempty"`
	RequestLocation bool                        `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
}

// KeyboardButtonPollType represents type of a poll which is allowed to be created.
type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

// WebAppInfo contains information about a Web App.
type WebAppInfo struct {
	URL string `json:"url"`
}

// ReplyKeyboardRemove represents a message to remove a custom keyboard.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard attached to a message.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo   `json:"web_app,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// CallbackGame is a placeholder, currently holds no information.
type CallbackGame struct{}

// ForceReply represents a placeholder to force a reply from the user.
type ForceReply struct {
	ForceReply             bool   `json:"force_reply"`
	InputFieldPlaceholder  string `json:"input_field_placeholder,omitempty"`
	Selective              bool   `json:"selective,omitempty"`
}

// NewInlineKeyboardMarkup creates a new inline keyboard markup.
func NewInlineKeyboardMarkup(buttons ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	return &InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}

// NewInlineKeyboardButtonRow creates a new row of inline keyboard buttons.
func NewInlineKeyboardButtonRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	return buttons
}

// NewInlineKeyboardButtonURL creates a URL button.
func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

// NewInlineKeyboardButtonCallback creates a callback button.
func NewInlineKeyboardButtonCallback(text, callbackData string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	}
}

// NewReplyKeyboardMarkup creates a new reply keyboard markup.
func NewReplyKeyboardMarkup(keyboard [][]KeyboardButton, opts ...ReplyKeyboardOption) *ReplyKeyboardMarkup {
	markup := &ReplyKeyboardMarkup{
		Keyboard: keyboard,
	}
	
	for _, opt := range opts {
		opt(markup)
	}
	
	return markup
}

// ReplyKeyboardOption configures a ReplyKeyboardMarkup.
type ReplyKeyboardOption func(*ReplyKeyboardMarkup)

// WithIsPersistent sets the IsPersistent flag.
func WithIsPersistent(isPersistent bool) ReplyKeyboardOption {
	return func(markup *ReplyKeyboardMarkup) {
		markup.IsPersistent = isPersistent
	}
}

// WithResizeKeyboard sets the ResizeKeyboard flag.
func WithResizeKeyboard(resize bool) ReplyKeyboardOption {
	return func(markup *ReplyKeyboardMarkup) {
		markup.ResizeKeyboard = resize
	}
}

// WithOneTimeKeyboard sets the OneTimeKeyboard flag.
func WithOneTimeKeyboard(oneTime bool) ReplyKeyboardOption {
	return func(markup *ReplyKeyboardMarkup) {
		markup.OneTimeKeyboard = oneTime
	}
}

// WithInputFieldPlaceholder sets the input field placeholder.
func WithInputFieldPlaceholder(placeholder string) ReplyKeyboardOption {
	return func(markup *ReplyKeyboardMarkup) {
		markup.InputFieldPlaceholder = placeholder
	}
}

// WithSelective sets the Selective flag.
func WithSelective(selective bool) ReplyKeyboardOption {
	return func(markup *ReplyKeyboardMarkup) {
		markup.Selective = selective
	}
}

// NewReplyKeyboardRemove creates a new reply keyboard remove.
func NewReplyKeyboardRemove(selective bool) *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// NewForceReply creates a new force reply.
func NewForceReply(selective bool, placeholder string) *ForceReply {
	return &ForceReply{
		ForceReply:            true,
		Selective:             selective,
		InputFieldPlaceholder: placeholder,
	}
}