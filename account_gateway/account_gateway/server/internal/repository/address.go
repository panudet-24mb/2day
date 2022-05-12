package repository

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	AddressID          int `gorm:"primaryKey;autoIncrement:true"`
	AccountID          int
	AddressNumber      string `gorm:"type:varchar(100)"`
	AddressName        string `gorm:"type:varchar(100)"`
	Address1           string `gorm:"type:varchar(100)"`
	Address2           string `gorm:"type:varchar(100)"`
	Province           string `gorm:"type:varchar(100)"`
	District           string `gorm:"type:varchar(100)"`
	SubDistrict        string `gorm:"type:varchar(100)"`
	ZipCode            string `gorm:"type:varchar(100)"`
	Latitude           string `gorm:"type:varchar(30)"`
	Longtitude         string `gorm:"type:varchar(30)"`
	Status             bool
	AddressGroupID     int
	AddressTypeID      int
	RefPartnerID       int
	RefPartnerBranchID int
}
type AddressGroup struct {
	ID   int
	Name string
}
type AddressType struct {
	ID   int
	Name string
}
