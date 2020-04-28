package application

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	NewDB() DB
	NewTX() TX
}

type DB interface {
	DB() *gorm.DB
}

type TX interface {
	Transaction(fn func(db *gorm.DB) error) error
}
