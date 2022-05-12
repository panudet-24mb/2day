package service

import (
	"github.com/gofrs/uuid"
)

type CreateAccountForm struct {
	UserID         int
	ProfilePicture string `json:"profile_picture"`
	AvatarIcon     string `json:"avatar_icon"`
	AvatarColor    string `json:"avatar_color"`
	AvatarUsage    bool   `json:"avatar_usage"`
	NamePrefix     int    `json:"name_prefix"`
	MobilePrefix   string `json:"mobile_prefix"`
	MobilePhone    string `json:"mobile_phone"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DateOfBirth    string `json:"date_of_birth"`
	GenderID       int    `json:"gender_id"`
	CountryID      int    `json:"country_id"`
}
type RequstFindAccount struct {
	UserID int
}
type AccountResponse struct {
	AccountID      int       `json:"account_id"`
	UserID         int       `json:"user_id"`
	PublicID       uuid.UUID `json:"public_id"`
	ProfilePicture string    `json:"profile_picture"`
	AvatarIcon     string    `json:"avatar_icon"`
	AvatarColor    string    `json:"avatar_color"`
	AvatarUsage    bool      `json:"avatar_usage"`
	LocalLanguage  string    `json:"local_language"`
	NightMode      bool      `json:"night_mode"`
	AccountCode    string    `json:"account_code"`
	AccountGroupID int       `json:"account_group_id"`
	NamePrefix     int       `json:"name_prefix"`
	MobilePrefix   string    `json:"mobile_prefix"`
	MobilePhone    string    `json:"mobile_phone"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DateOfBirth    string    `json:"date_of_birth"`
	Timezone       string    `json:"timezone"`
	GenderID       int       `json:"gender_id"`
	CountryID      int       `json:"country_id"`
}
type UpdateAccountForm struct {
	UserID         int
	ProfilePicture string `json:"profile_picture,omitempty"`
	AvatarIcon     string `json:"avatar_icon,omitempty"`
	AvatarColor    string `json:"avatar_color,omitempty"`
	AvatarUsage    bool   `json:"avatar_usage,omitempty"`
	LocalLanguage  string `json:"local_language,omitempty"`
	NightMode      bool   `json:"night_mode,omitempty"`
	AccountCode    string `json:"account_code,omitempty"`
	AccountGroupID int    `json:"account_group_id,omitempty"`
	NamePrefix     int    `json:"name_prefix,omitempty"`
	MobilePrefix   string `json:"mobile_prefix,omitempty"`
	MobilePhone    string `json:"mobile_phone,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	DateOfBirth    string `json:"date_of_birth,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
	GenderID       int    `json:"gender_id,omitempty"`
	CountryID      int    `json:"country_id,omitempty"`
}

type AccountService interface {
	CreateAccount(CreateAccountForm *CreateAccountForm) (*AccountResponse, error)
	FindAccount(RequstFindAccount *RequstFindAccount) (*AccountResponse, error)
	UpdateAccount(UpdateAccountForm *UpdateAccountForm) (*AccountResponse, error)
}
