package users

import (
	// "strings"
	// "time"

	// "github.com/bnkamalesh/errors"
)

type User struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Mobile	string `json:"mobile,omitempty"`
	// CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// func (u *User) setDefaults() {
// 	now := time.Now()
// 	if u.CreatedAt == nil {
// 		u.CreatedAt = &now
// 	}

// 	if u.UpdatedAt == nil {
// 		u.UpdatedAt = &now
// 	}
// }

// func (u *User) Sanitize() {
// 	u.FirstName = strings.TrimSpace(u.FirstName)
// 	u.LastName = strings.TrimSpace(u.LastName)
// 	u.Email = strings.TrimSpace(u.Email)
// 	u.Mobile = strings.TrimSpace(u.Mobile)
// }

// func (u *User) Validate() error {

// 	if u.FirstName == "" {
// 		return errors.Validation("First name is required")
// 	}

// 	if u.LastName == "" {
// 		return errors.Validation("Last name is required")
// 	}

// 	err := validateEmail(u.Email)
// 	if err != nil {
// 		return err
// 	}

// 	err = validateMobile(u.Mobile)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func validateEmail(email string) error {
// 	if email == "" {
// 		return errors.Validation("Email is required")
// 	}
// 	parts := strings.Split(email, "@")
// 	if len(parts) != 2 {
// 		return errors.New("Invalid email")
// 	}

// 	if parts[0] == "" {
// 		return errors.New("Invalid email")
// 	}

// 	if parts[1] == "" {
// 		return errors.New("Invalid email")
// 	}

// 	return nil
// }

// func validateMobile(mobile string) error {
// 	if mobile == "" {
// 		return errors.Validation("Mobile is required")
// 	}

// 	if len(mobile) != 10 {
// 		return errors.New("Invalid mobile number")
// 	}

// 	return nil
// }