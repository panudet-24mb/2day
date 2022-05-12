package repository

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return UserRepositoryDB{db: db}
}

func (r UserRepositoryDB) GetAll() ([]User, error) {
	users := []User{}
	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r UserRepositoryDB) GetById(id int) (*User, error) {
	user := User{}
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r UserRepositoryDB) Create(u User) (*User, error) {
	result := r.db.Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &u, result.Error

}

func (r UserRepositoryDB) Update(u User) (*User, error) {

	result := r.db.Model(&User{}).Where("user_id = ?", u.UserID).Updates(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, result.Error

}
func (r UserRepositoryDB) UpdateOne(u User) (*User, error) {

	result := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &u, result.Error

}

func (r UserRepositoryDB) DeleteOne(id int) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (r UserRepositoryDB) FindUser(u User) (*User, error) {
	user := User{}
	err := r.db.Where("email = ? OR user_name = ?", u.Email, u.UserName).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}
