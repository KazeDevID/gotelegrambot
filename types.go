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

// MessageEntity represents one special entity in a text message.
type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	URL           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

// Animation represents an animation file.
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Audio represents an audio file.
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

// Document represents a general file.
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// VideoNote represents a video message.
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Voice represents a voice note.
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	Longitude             float64 `json:"longitude"`
	Latitude              float64 `json:"latitude"`
	HorizontalAccuracy    float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod            int     `json:"live_period,omitempty"`
	Heading               int     `json:"heading,omitempty"`
	ProximityAlertRadius  int     `json:"proximity_alert_radius,omitempty"`
}

// Venue represents a venue.
type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id,omitempty"`
	FoursquareType  string    `json:"foursquare_type,omitempty"`
	GooglePlaceID   string    `json:"google_place_id,omitempty"`
	GooglePlaceType string    `json:"google_place_type,omitempty"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// More types omitted for brevity...
// In a real implementation, all types from the Telegram Bot API would be defined here.