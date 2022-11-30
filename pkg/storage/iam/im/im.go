package im

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Email     string    `gorm:"email"`
	Password  string    `gorm:"pass"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func SaveUser(tx *gorm.DB, user User) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func FindUser(tx *gorm.DB, email string) (*User, error) {
	user := new(User)

	if err := tx.Where(&User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func ExistUser(tx *gorm.DB, email string) (bool, error) {
	var count int64
	var err error

	err = tx.Where(&User{Email: email}).Limit(1).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UpdateUser(tx *gorm.DB, user *User) error {
	return tx.Save(user).Error
}

func DelUser(tx *gorm.DB, email string) error {
	return tx.Delete(&User{Email: email}).Error
}
