package am

import "gorm.io/gorm"

type User struct {
	Email    string
	Password string
}

func SaveUser(tx *gorm.DB, user User) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
