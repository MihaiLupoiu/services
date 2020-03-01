package models

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

// User structure for a user
type User struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName   string    `gorm:"size:255;not null;" json:"firstname"`
	LastName    string    `gorm:"size:255;not null;" json:"lastname"`
	Email       string    `gorm:"size:255;not null;unique" json:"email"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	PhoneNumber int       `gorm:"size:15;" json:"phone_number"`
	Country     string    `gorm:"size:100;" json:"country"`
	PostalCode  int       `gorm:"size:5;" json:"postal_code"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Create user and save in database
func (u *User) Create(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Validate User is valid
func (u *User) Validate() error {
	if u.FirstName == "" {
		return errors.New("Required First Name")
	}
	if u.LastName == "" {
		return errors.New("Required LastName")
	}
	if u.Password == "" {
		return errors.New("Required Password")
	}
	if u.Email == "" {
		return errors.New("Required Email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}

	return nil
}

// ValidateLogin basic authentication provided
func (u *User) ValidateLogin() error {
	if u.Password == "" {
		return errors.New("Required Password")
	}
	if u.Email == "" {
		return errors.New("Required Email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

// FindByEmail user by using email in database
func (u *User) FindByEmail(db *gorm.DB) (*User, error) {
	return nil, nil
}

// Update user in database
func (u *User) Update(db *gorm.DB) (*User, error) {
	return nil, nil
}

// Delete user in database
func (u *User) Delete(db *gorm.DB) (*User, error) {
	return nil, nil
}
