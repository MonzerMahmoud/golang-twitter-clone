package interfaces

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"fullName,omitempty"`
	Email     string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}

type Tweet struct {
	gorm.Model
	Body string `json:"body,omitempty"`
	UserID uint `json:"userId,omitempty"`
	//User User `json:"user,omitempty"`
	// CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type ResponseUser struct {
	ID uint
	FullName string
	Email string
	Username string
}
