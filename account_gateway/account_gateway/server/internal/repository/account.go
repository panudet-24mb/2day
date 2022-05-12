package repository

import (
	uuid "github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountID      int `gorm:"primaryKey;autoIncrement:true"`
	UserID         int
	PublicID       uuid.UUID `gorm:"type:uuid"`
	ProfilePicture string
	AvatarIcon     string
	AvatarColor    string
	AvatarUsage    bool
	LocalLanguage  string
	NightMode      bool
	AccountCode    string `gorm:"type:varchar(20)"`
	AccountGroupID int
	NamePrefix     int
	MobilePrefix   string `gorm:"type:varchar(10)"`
	MobilePhone    string `gorm:"type:varchar(15)"`
	FirstName      string
	LastName       string
	DateOfBirth    string
	Timezone       string
	GenderID       int
	CountryID      int
	UserNote       string
}

type AccountRepository interface {
	GetAllAccount() ([]Account, error)
	GetByAccountId(int) (*Account, error)
	CreateAccount(Account) (*Account, error)
	UpdateAccount(Account) (*Account, error)
	UpdateOneAccount(i interface{}, userId int) error
	DeleteOneAccount(int) error
	FindAccount(Account) (*Account, error)
}
