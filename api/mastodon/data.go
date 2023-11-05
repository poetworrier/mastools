// The api package contains http json bindings for Mastodon
package mastodon

import (
	"fmt"
	"log/slog"
)

// Status represents a Mastodon status update.
// See: https://docs.joinmastodon.org/entities/Status
type Status struct {
	ID string `json:"id"` // ID of the status in the database
	// TODO: implement UnmarshallJSON for this type to be time.Time
	CreatedAt          string             `json:"created_at,omitempty"`             // The date when this status was created
	InReplyToID        string             `json:"in_reply_to_id,omitempty"`         // ID of the status being replied to
	InReplyToAccountID string             `json:"in_reply_to_account_id,omitempty"` // ID of the account that authored the status being replied to
	Sensitive          bool               `json:"sensitive"`                        // Is this status marked as sensitive content?
	SpoilerText        string             `json:"spoiler_text"`                     // Subject or summary line, below which status content is collapsed until expanded
	Visibility         string             `json:"visibility"`                       // Visibility of this status
	Language           string             `json:"language,omitempty"`               // Primary language of this status
	URI                string             `json:"uri"`                              // URI of the status used for federation
	URL                string             `json:"url,omitempty"`                    // A link to the status's HTML representation
	RepliesCount       int                `json:"replies_count"`                    // How many replies this status has received
	ReblogsCount       int                `json:"reblogs_count"`                    // How many boosts this status has received
	FavouritesCount    int                `json:"favourites_count"`                 // How many favourites this status has received
	Favourited         bool               `json:"favourited,omitempty"`             // Have you favourited this status?
	Reblogged          bool               `json:"reblogged,omitempty"`              // Have you boosted this status?
	Muted              bool               `json:"muted,omitempty"`                  // Have you muted notifications for this status's conversation?
	Bookmarked         bool               `json:"bookmarked,omitempty"`             // Have you bookmarked this status?
	Content            string             `json:"content"`                          // HTML-encoded status content
	Reblog             *Status            `json:"reblog,omitempty"`                 // The status being reblogged
	Application        *Application       `json:"application,omitempty"`            // The application used to post this status
	Account            *Account           `json:"account,omitempty"`                // The account that authored this status
	MediaAttachments   []*MediaAttachment `json:"media_attachments,omitempty"`      // Media that is attached to this status
	Mentions           []*Mention         `json:"mentions,omitempty"`               // Mentions of users within the status content
	Tags               []*Tag             `json:"tags,omitempty"`                   // Hashtags used within the status content
	Emojis             []*CustomEmoji     `json:"emojis,omitempty"`                 // Custom emoji to be used when rendering status content
	Card               *PreviewCard       `json:"card,omitempty"`                   // Preview card for links included within status content
	Poll               *Poll              `json:"poll,omitempty"`                   // The poll attached to the status
}

func (s Status) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("ID", s.ID),
		slog.String("CreatedAt", s.CreatedAt),
		slog.String("InReplyToID", fmt.Sprintf("%v", s.InReplyToID)),
		slog.String("InReplyToAccountID", fmt.Sprintf("%v", s.InReplyToAccountID)),
		slog.Bool("Sensitive", s.Sensitive),
		slog.String("SpoilerText", s.SpoilerText),
		slog.String("Visibility", s.Visibility),
		slog.String("Language", s.Language),
		slog.String("URI", s.URI),
		slog.String("URL", s.URL),
		slog.Int("RepliesCount", s.RepliesCount),
		slog.Int("ReblogsCount", s.ReblogsCount),
		slog.Int("FavouritesCount", s.FavouritesCount),
		slog.Bool("Favourited", s.Favourited),
		slog.Bool("Reblogged", s.Reblogged),
		slog.Bool("Muted", s.Muted),
		slog.Bool("Bookmarked", s.Bookmarked),
		slog.String("Content", s.Content),
		slog.String("Reblog", fmt.Sprintf("%v", s.Reblog)),
		slog.String("Application", fmt.Sprintf("%v", s.Application)),
		slog.String("Account", fmt.Sprintf("%v", s.Account)),
		slog.String("MediaAttachments", fmt.Sprintf("%v", s.MediaAttachments)),
		slog.String("Mentions", fmt.Sprintf("%v", s.Mentions)),
		slog.String("Tags", fmt.Sprintf("%v", s.Tags)),
		slog.String("Emojis", fmt.Sprintf("%v", s.Emojis)),
		slog.String("Card", fmt.Sprintf("%v", s.Card)),
		slog.String("Poll", fmt.Sprintf("%v", s.Poll)),
	)
}

// Application represents the application used to post a status.
type Application struct {
	Name    string `json:"name"`    // The name of the application that posted the status
	Website string `json:"website"` // The website associated with the application that posted the status
}

// Account represents a Mastodon user account.
// TODO: support fields as needed
type Account struct {
	ID          string `json:"id"`           // ID of the account in the database
	Username    string `json:"username"`     // The username of the account
	Acct        string `json:"acct"`         // The webfinger acct: URI of the account
	DisplayName string `json:"display_name"` // The display name of the account
	Locked      bool   `json:"locked"`       // Is the account locked?
	Bot         bool   `json:"bot"`          // Is the account a bot?
	// Discoverable   bool      `json:"discoverable"`     // Is the account discoverable?
	// Group          bool      `json:"group"`            // Is the account a group?
	CreatedAt string `json:"created_at"` // The date when the account was created
	// Note           string    `json:"note"`             // HTML-encoded account biography
	URL    string `json:"url"`    // The location of the account's profile
	Avatar string `json:"avatar"` // The URL of the account's avatar
	// AvatarStatic   string    `json:"avatar_static"`    // The static URL of the account's avatar
	// Header         string    `json:"header"`           // The URL of the account's header image
	// HeaderStatic   string    `json:"header_static"`    // The static URL of the account's header image
	// FollowersCount int       `json:"followers_count"`  // The number of followers of the account
	// FollowingCount int       `json:"following_count"`  // The number of accounts the account is following
	// StatusesCount  int       `json:"statuses_count"`   // The number of status updates the account has made
	LastStatusAt string `json:"last_status_at"` // The date when the account's last status was posted
	// Emojis         []*CustomEmoji  `json:"emojis,omitempty"` // Custom emojis associated with the account
	// Fields         []*Field  `json:"fields,omitempty"` // Custom profile fields of the account
}

// Mention represents a mention of a user within a status.
type Mention struct {
	ID       string `json:"id"`       // The account ID of the mentioned user
	Username string `json:"username"` // The username of the mentioned user
	URL      string `json:"url"`      // The location of the mentioned user's profile
	Acct     string `json:"acct"`     // The webfinger acct: URI of the mentioned user
}

// Tag represents a hashtag used within a status.
type Tag struct {
	Name string `json:"name"` // The value of the hashtag after the # sign
	URL  string `json:"url"`  // A link to the hashtag on the instance
}

// TODO: fill out the stubs below as needed

// MediaAttachment represents a media attachment to a status.
type MediaAttachment struct {
	// Attributes of media attachment go here
}

// CustomEmoji represents a custom emoji used in a status.
type CustomEmoji struct {
	// Attributes of custom emoji go here
}

// PreviewCard represents a preview card for links in a status.
type PreviewCard struct {
	// Attributes of preview card go here
}

// Poll represents a poll attached to a status.
type Poll struct {
	// Attributes of poll go here
}
