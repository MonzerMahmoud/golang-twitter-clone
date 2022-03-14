package interfaces

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"fullName,omitempty"`
	Email     string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
	Following int `json:"following,omitempty"`
	Followers int `json:"followers,omitempty"`
}

type Tweet struct {
	gorm.Model
	Body string `json:"body,omitempty"`
	UserID uint `json:"userId,omitempty"`
	//User User `json:"user,omitempty"`
	// CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type Follow struct {
	gorm.Model
	FollowerID uint `json:"followerId,omitempty"`
	FolloweeID uint `json:"followeeId,omitempty"`
}

type ResponseUser struct {
	ID uint
	FullName string
	Email string
	Username string
}

type ResponseTweet struct {
	ID uint
	Body string
	UserID uint
	User User
	CreatedAt time.Time
}

type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string `json:"message"`
}
