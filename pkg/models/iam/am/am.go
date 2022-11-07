package am

import (
	"gorm.io/gorm"
	sql "io.github/devopssphere/pkg/storage/iam/am"
)

type AccessManagementInterface interface {
	AddUser(email, password string) error
	ResetPassword(email, oldPass, newPass string) error
	DelUser(name string) error
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
		user := sql.User{
			Email:    email,
			Password: password,
		}

		return sql.SaveUser(tx, user)
	})
}

func (a amOperator) ResetPassword(email, oldPass, newPass string) error {
	//TODO implement me
	panic("implement me")
}

func (a amOperator) DelUser(name string) error {
	//TODO implement me
	panic("implement me")
}
