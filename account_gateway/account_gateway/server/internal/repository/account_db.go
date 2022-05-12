package repository

import (
	"errors"

	"gorm.io/gorm"
)

type AccountRepositoryDB struct {
	db *gorm.DB
}

func NewAccountRepositoryDB(db *gorm.DB) AccountRepository {
	return AccountRepositoryDB{db: db}
}

func (r AccountRepositoryDB) GetAllAccount() ([]Account, error) {
	accounts := []Account{}
	result := r.db.Find(&accounts)

	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

func (r AccountRepositoryDB) GetByAccountId(id int) (*Account, error) {
	account := Account{}
	result := r.db.First(&account, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (r AccountRepositoryDB) CreateAccount(u Account) (*Account, error) {
	result := r.db.Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &u, result.Error

}

func (r AccountRepositoryDB) UpdateAccount(u Account) (*Account, error) {

	result := r.db.Model(&Account{}).Where("user_id = ?", u.UserID).Updates(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, result.Error

}
func (r AccountRepositoryDB) UpdateOneAccount(updateterm interface{}, userId int) error {

	// result := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Where("user_id = ?", u.UserID).Updates(&u)
	result := r.db.Model(&Account{}).Where("user_id = ?", userId).Updates(updateterm)
	if result.Error != nil {
		return result.Error
	}

	// acc, err := r.GetByAccountId(updateterm.(Account).UserID)
	// if err != nil {
	// 	return nil, err
	// }

	return nil

}

func (r AccountRepositoryDB) DeleteOneAccount(id int) error {
	result := r.db.Delete(&Account{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (r AccountRepositoryDB) FindAccount(u Account) (*Account, error) {
	account := Account{}
	err := r.db.Where("account_id = ? OR user_id = ?", u.AccountID, u.UserID).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &account, nil
}
