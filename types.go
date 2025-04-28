package gotelegrambot

import (
	"encoding/json"
	"time"
)

// Update represents a Telegram update.
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         *ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    *ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

// User represents a Telegram user or bot.
type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls         bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo        bool `json:"can_change_info,omitempty"`
	CanInviteUsers       bool `json:"can_invite_users,omitempty"`
	CanPinMessages       bool `json:"can_pin_messages,omitempty"`
}

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

// Chat represents a chat.
type Chat struct {
	ID                    int64       `json:"id"`
	Type                  string      `json:"type"`
	Title                 string      `json:"title,omitempty"`
	Username              string      `json:"username,omitempty"`
	FirstName             string      `json:"first_name,omitempty"`
	LastName              string      `json:"last_name,omitempty"`
	IsForum               bool        `json:"is_forum,omitempty"`
	Photo                 *ChatPhoto  `json:"photo,omitempty"`
	ActiveUsernames       []string    `json:"active_usernames,omitempty"`
	EmojiStatusCustomEmojiID string   `json:"emoji_status_custom_emoji_id,omitempty"`
	Bio                   string      `json:"bio,omitempty"`
	HasPrivateForwards    bool        `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages    bool        `json:"join_to_send_messages,omitempty"`
	JoinByRequest         bool        `json:"join_by_request,omitempty"`
	Description           string      `json:"description,omitempty"`
	InviteLink            string      `json:"invite_link,omitempty"`
	PinnedMessage         *Message    `json:"pinned_message,omitempty"`
	Permissions           *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay         int         `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime int         `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent   bool        `json:"has_protected_content,omitempty"`
	StickerSetName        string      `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      bool        `json:"can_set_sticker_set,omitempty"`
	LinkedChatID          int64       `json:"linked_chat_id,omitempty"`
	Location              *ChatLocation `json:"location,omitempty"`
}

// Sticker represents a sticker.
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Type         string        `json:"type"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	IsVideo      bool          `json:"is_video"`
	Thumbnail    *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji        string        `json:"emoji,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// Game represents a game.
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

// MessageAutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// PassportData contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

// EncryptedPassportElement contains information about documents or other Telegram Passport elements shared with the bot.
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

// PassportFile represents a file uploaded to Telegram Passport.
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

// EncryptedCredentials contains data required for decrypting and authenticating PassportData.
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// ProximityAlertTriggered represents a service message about a user in the chat triggering a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VideoChatStarted represents a service message about a video chat started in the chat.
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// WebAppData contains data sent from a Web App to the bot.
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// Message represents a message.
type Message struct {
	MessageID              int                `json:"message_id"`
	From                   *User              `json:"from,omitempty"`
	SenderChat             *Chat              `json:"sender_chat,omitempty"`
	Date                   int                `json:"date"`
	Chat                   *Chat              `json:"chat"`
	ForwardFrom            *User              `json:"forward_from,omitempty"`
	ForwardFromChat        *Chat              `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID   int                `json:"forward_from_message_id,omitempty"`
	ForwardSignature       string             `json:"forward_signature,omitempty"`
	ForwardSenderName      string             `json:"forward_sender_name,omitempty"`
	ForwardDate            int                `json:"forward_date,omitempty"`
	IsAutomaticForward     bool               `json:"is_automatic_forward,omitempty"`
	ReplyToMessage         *Message           `json:"reply_to_message,omitempty"`
	ViaBot                 *User              `json:"via_bot,omitempty"`
	EditDate               int                `json:"edit_date,omitempty"`
	HasProtectedContent    bool               `json:"has_protected_content,omitempty"`
	MediaGroupID           string             `json:"media_group_id,omitempty"`
	AuthorSignature        string             `json:"author_signature,omitempty"`
	Text                   string             `json:"text,omitempty"`
	Entities               []MessageEntity    `json:"entities,omitempty"`
	Animation              *Animation         `json:"animation,omitempty"`
	Audio                  *Audio             `json:"audio,omitempty"`
	Document               *Document          `json:"document,omitempty"`
	Photo                  []PhotoSize        `json:"photo,omitempty"`
	Sticker                *Sticker           `json:"sticker,omitempty"`
	Video                  *Video             `json:"video,omitempty"`
	VideoNote              *VideoNote         `json:"video_note,omitempty"`
	Voice                  *Voice             `json:"voice,omitempty"`
	Caption                string             `json:"caption,omitempty"`
	CaptionEntities        []MessageEntity    `json:"caption_entities,omitempty"`
	HasMediaSpoiler        bool               `json:"has_media_spoiler,omitempty"`
	Contact                *Contact           `json:"contact,omitempty"`
	Dice                   *Dice              `json:"dice,omitempty"`
	Game                   *Game              `json:"game,omitempty"`
	Poll                   *Poll              `json:"poll,omitempty"`
	Venue                  *Venue             `json:"venue,omitempty"`
	Location               *Location          `json:"location,omitempty"`
	NewChatMembers         []User             `json:"new_chat_members,omitempty"`
	LeftChatMember         *User              `json:"left_chat_member,omitempty"`
	NewChatTitle           string             `json:"new_chat_title,omitempty"`
	NewChatPhoto           []PhotoSize        `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto        bool               `json:"delete_chat_photo,omitempty"`
	GroupChatCreated       bool               `json:"group_chat_created,omitempty"`
	SupergroupChatCreated  bool               `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated     bool               `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID        int64              `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID      int64              `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage          *Message           `json:"pinned_message,omitempty"`
	Invoice                *Invoice           `json:"invoice,omitempty"`
	SuccessfulPayment      *SuccessfulPayment `json:"successful_payment,omitempty"`
	ConnectedWebsite       string             `json:"connected_website,omitempty"`
	PassportData           *PassportData      `json:"passport_data,omitempty"`
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	VideoChatScheduled     *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted       *VideoChatStarted  `json:"video_chat_started,omitempty"`
	VideoChatEnded         *VideoChatEnded    `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData             *WebAppData        `json:"web_app_data,omitempty"`
	ReplyMarkup            *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}