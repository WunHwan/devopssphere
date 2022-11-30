package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(o *Options) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(o.DSN), &gorm.Config{})
}
