package gotelegrambot

import (
	"context"
)

// Invoice contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// LabeledPrice represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

// SendInvoice sends an invoice.
func (b *Bot) SendInvoice(ctx context.Context, chatID int64, title, description, payload, providerToken, currency string, prices []LabeledPrice, options ...SendInvoiceOption) (*Message, error) {
	params := map[string]interface{}{
		"chat_id":         chatID,
		"title":           title,
		"description":     description,
		"payload":         payload,
		"provider_token":  providerToken,
		"currency":        currency,
		"prices":          prices,
	}
	
	opts := defaultSendInvoiceOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	if opts.MaxTipAmount != 0 {
		params["max_tip_amount"] = opts.MaxTipAmount
	}
	
	if len(opts.SuggestedTipAmounts) > 0 {
		params["suggested_tip_amounts"] = opts.SuggestedTipAmounts
	}
	
	if opts.StartParameter != "" {
		params["start_parameter"] = opts.StartParameter
	}
	
	if opts.ProviderData != "" {
		params["provider_data"] = opts.ProviderData
	}
	
	if opts.PhotoURL != "" {
		params["photo_url"] = opts.PhotoURL
	}
	
	if opts.PhotoSize != 0 {
		params["photo_size"] = opts.PhotoSize
	}
	
	if opts.PhotoWidth != 0 {
		params["photo_width"] = opts.PhotoWidth
	}
	
	if opts.PhotoHeight != 0 {
		params["photo_height"] = opts.PhotoHeight
	}
	
	if opts.NeedName {
		params["need_name"] = true
	}
	
	if opts.NeedPhoneNumber {
		params["need_phone_number"] = true
	}
	
	if opts.NeedEmail {
		params["need_email"] = true
	}
	
	if opts.NeedShippingAddress {
		params["need_shipping_address"] = true
	}
	
	if opts.SendPhoneNumberToProvider {
		params["send_phone_number_to_provider"] = true
	}
	
	if opts.SendEmailToProvider {
		params["send_email_to_provider"] = true
	}
	
	if opts.IsFlexible {
		params["is_flexible"] = true
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
	err := b.makeRequest(ctx, "sendInvoice", params, &message)
	if err != nil {
		return nil, err
	}
	
	return &message, nil
}

// SendInvoiceOption is a function that configures SendInvoice options.
type SendInvoiceOption func(*sendInvoiceOptions)

// sendInvoiceOptions represents options for SendInvoice.
type sendInvoiceOptions struct {
	MaxTipAmount             int
	SuggestedTipAmounts      []int
	StartParameter           string
	ProviderData             string
	PhotoURL                 string
	PhotoSize                int
	PhotoWidth               int
	PhotoHeight              int
	NeedName                 bool
	NeedPhoneNumber          bool
	NeedEmail                bool
	NeedShippingAddress      bool
	SendPhoneNumberToProvider bool
	SendEmailToProvider      bool
	IsFlexible               bool
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

// Various option functions for SendInvoice would be defined here.

func defaultSendInvoiceOptions() sendInvoiceOptions {
	return sendInvoiceOptions{}
}

// AnswerShippingQuery answers a shipping query.
func (b *Bot) AnswerShippingQuery(ctx context.Context, shippingQueryID string, ok bool, options ...AnswerShippingQueryOption) error {
	params := map[string]interface{}{
		"shipping_query_id": shippingQueryID,
		"ok":                ok,
	}
	
	opts := defaultAnswerShippingQueryOptions()
	for _, opt := range options {
		opt(&opts)
	}
	
	if len(opts.ShippingOptions) > 0 {
		params["shipping_options"] = opts.ShippingOptions
	}
	
	if opts.ErrorMessage != "" {
		params["error_message"] = opts.ErrorMessage
	}
	
	return b.makeRequest(ctx, "answerShippingQuery", params, nil)
}

// AnswerShippingQueryOption is a function that configures AnswerShippingQuery options.
type AnswerShippingQueryOption func(*answerShippingQueryOptions)

// answerShippingQueryOptions represents options for AnswerShippingQuery.
type answerShippingQueryOptions struct {
	ShippingOptions []ShippingOption
	ErrorMessage    string
}

// ShippingOption represents one shipping option.
type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// Various option functions for AnswerShippingQuery would be defined here.

func defaultAnswerShippingQueryOptions() answerShippingQueryOptions {
	return answerShippingQueryOptions{}
}

// AnswerPreCheckoutQuery answers a pre-checkout query.
func (b *Bot) AnswerPreCheckoutQuery(ctx context.Context, preCheckoutQueryID string, ok bool, errorMessage string) error {
	params := map[string]interface{}{
		"pre_checkout_query_id": preCheckoutQueryID,
		"ok":                    ok,
	}
	
	if !ok && errorMessage != "" {
		params["error_message"] = errorMessage
	}
	
	return b.makeRequest(ctx, "answerPreCheckoutQuery", params, nil)
}