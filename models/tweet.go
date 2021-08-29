package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Tweet struct {
	Id string `gorm:"notNull;primaryKey" json:"id"`

	UserId string `gorm:"notNull"`

	// References the original top-level tweet
	ConversationId string `gorm:"notNull" json:"conversationId"`

	// References the parent tweet
	InReplyToTweetId null.String `json:"inReplyToTweetId"`

	Text      string    `gorm:"notNull" json:"text"`
	CreatedAt time.Time `gorm:"notNull" json:"createdAt"`
	UpdatedAt time.Time `gorm:"notNull" json:"updatedAt"`
}

type EntityUrl struct {
	// Url with http/https
	DisplayUrl string
	// Full url
	ExpandedUrl string
	// Defines the position in the tweet text. Example: [5, 17]
	Indices [2]uint
}

type EntityUserMention struct {
	// Mentionned user id
	Id string
	// Mentionned user full name
	FullName string
	// Mentionned user username
	Username string
	// Defines the position in the tweet text. Example: [5, 17]
	Indices [2]uint
}

type EntityHashtag struct {
	Text string
	// Defines the position in the tweet text. Example: [5, 17]
	Indices [2]uint
}

type Entities struct {
	Hashtags     []EntityHashtag
	Urls         []EntityUrl
	UserMentions []EntityUserMention
}

type TweetFull struct {
	// Same as above, plus other generated properties
	Entities Entities

	FavoriteCount uint
	Favorited     bool

	RetweetCount uint
	Retweeted    bool

	ReplyCount uint
}
