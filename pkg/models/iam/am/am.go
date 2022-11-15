package am

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	sql "io.github/devopssphere/pkg/storage/iam/am"
)

type AccessManagementInterface interface {
	AddUser(email, password string) error
	ResetPassword(email, oldPass, newPass string) error
	DelUser(email string) error
}

type amOperator struct {
	db *gorm.DB
}

func NewAMOperator(db *gorm.DB) AccessManagementInterface {
	return amOperator{
		db: db,
	}
}

func (a amOperator) AddUser(email, password string) error {
	return a.db.Transaction(func(tx *gorm.DB) error {
		saltpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := sql.User{
			Email:    email,
			Password: string(saltpass),
		}

		return sql.SaveUser(tx, user)
	})
}

func (a amOperator) ResetPassword(email, oldPass, newPass string) error {
	return a.db.Transaction(func(tx *gorm.DB) error {
		// lookup user from email
		user, err := sql.FindUser(tx, email)
		if err != nil {
			return err
		}

		// password checker
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPass))
		if err != nil {
			return err
		}

		saltpass, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(saltpass)

		return sql.UpdateUser(tx, user)
	})
}

func (a amOperator) DelUser(email string) error {
	return a.db.Transaction(func(tx *gorm.DB) error {
		// lookup user from email
		existed, err := sql.ExistUser(tx, email)
		if err != nil {
			return err
		}
		if !existed {
			return errors.New("Account not found.")
		}

		return sql.DelUser(tx, email)
	})
}
