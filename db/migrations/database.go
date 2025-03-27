package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
