package gotelegrambot

import (
	"context"
)

// Poll represents a poll.
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       int             `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int             `json:"open_period,omitempty"`
	CloseDate             int             `json:"close_date,omitempty"`
}

// PollOption represents an option in a poll.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// PollAnswer represents an answer in a poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// SendPoll sends a poll.
func (b *Bot) SendPoll(ctx context.Context, chatID interface{}, question string, options []string, pollOptions ...SendPollOption) (*Message, error) {
	params := map[string]interface{}{
		"chat_id":  chatID,
		"question": question,
		"options":  options,
	}
	
	opts := defaultSendPollOptions()
	for _, opt := range pollOptions {
		opt(&opts)
	}
	
	if opts.IsAnonymous != nil {
		params["is_anonymous"] = *opts.IsAnonymous
	}
	
	if opts.Type != "" {
		params["type"] = opts.Type
	}
	
	if opts.AllowsMultipleAnswers != nil {
		params["allows_multiple_answers"] = *opts.AllowsMultipleAnswers
	}
	
	if opts.CorrectOptionID != nil {
		params["correct_option_id"] = *opts.CorrectOptionID
	}
	
	if opts.Explanation != "" {
		params["explanation"] = opts.Explanation
	}
	
	if opts.ExplanationParseMode != "" {
		params["explanation_parse_mode"] = opts.ExplanationParseMode
	}
	
	if len(opts.ExplanationEntities) > 0 {
		params["explanation_entities"] = opts.ExplanationEntities
	}
	
	if opts.OpenPeriod != nil {
		params["open_period"] = *opts.OpenPeriod
	}
	
	if opts.CloseDate != nil {
		params["close_date"] = *opts.CloseDate
	}
	
	if opts.IsClosed != nil {
		params["is_closed"] = *opts.IsClosed
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
	err := b.makeRequest(ctx, "sendPoll", params, &message)
	if err != nil {
		return nil, err
	}
	
	return &message, nil
}

// SendPollOption is a function that configures SendPoll options.
type SendPollOption func(*sendPollOptions)

// sendPollOptions represents options for SendPoll.
type sendPollOptions struct {
	IsAnonymous             *bool
	Type                    string
	AllowsMultipleAnswers   *bool
	CorrectOptionID         *int
	Explanation             string
	ExplanationParseMode    string
	ExplanationEntities     []MessageEntity
	OpenPeriod              *int
	CloseDate               *int
	IsClosed                *bool
	DisableNotification     bool
	ProtectContent          bool
	ReplyToMessageID        int
	AllowSendingWithoutReply bool
	ReplyMarkup             interface{}
}

// WithIsAnonymous sets whether the poll is anonymous.
func WithIsAnonymous(isAnonymous bool) SendPollOption {
	return func(o *sendPollOptions) {
		o.IsAnonymous = &isAnonymous
	}
}

// WithPollType sets the type of the poll.
func WithPollType(pollType string) SendPollOption {
	return func(o *sendPollOptions) {
		o.Type = pollType
	}
}

// WithAllowsMultipleAnswers sets whether the poll allows multiple answers.
func WithAllowsMultipleAnswers(allowsMultipleAnswers bool) SendPollOption {
	return func(o *sendPollOptions) {
		o.AllowsMultipleAnswers = &allowsMultipleAnswers
	}
}

// WithCorrectOptionID sets the correct option ID for the poll.
func WithCorrectOptionID(correctOptionID int) SendPollOption {
	return func(o *sendPollOptions) {
		o.CorrectOptionID = &correctOptionID
	}
}

// WithExplanation sets the explanation for the poll.
func WithExplanation(explanation string) SendPollOption {
	return func(o *sendPollOptions) {
		o.Explanation = explanation
	}
}

// WithExplanationParseMode sets the parse mode for the explanation.
func WithExplanationParseMode(parseMode string) SendPollOption {
	return func(o *sendPollOptions) {
		o.ExplanationParseMode = parseMode
	}
}

// WithExplanationEntities sets the explanation entities for the poll.
func WithExplanationEntities(entities []MessageEntity) SendPollOption {
	return func(o *sendPollOptions) {
		o.ExplanationEntities = entities
	}
}

// WithOpenPeriod sets the open period for the poll.
func WithOpenPeriod(openPeriod int) SendPollOption {
	return func(o *sendPollOptions) {
		o.OpenPeriod = &openPeriod
	}
}

// WithCloseDate sets the close date for the poll.
func WithCloseDate(closeDate int) SendPollOption {
	return func(o *sendPollOptions) {
		o.CloseDate = &closeDate
	}
}

// WithIsClosed sets whether the poll is closed.
func WithIsClosed(isClosed bool) SendPollOption {
	return func(o *sendPollOptions) {
		o.IsClosed = &isClosed
	}
}

// Various other option functions would be defined here,
// similar to the message options.

func defaultSendPollOptions() sendPollOptions {
	return sendPollOptions{}
}

// StopPoll stops a poll.
func (b *Bot) StopPoll(ctx context.Context, chatID interface{}, messageID int, replyMarkup interface{}) (*Poll, error) {
	params := map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	}
	
	if replyMarkup != nil {
		params["reply_markup"] = replyMarkup
	}
	
	var poll Poll
	err := b.makeRequest(ctx, "stopPoll", params, &poll)
	if err != nil {
		return nil, err
	}
	
	return &poll, nil
}