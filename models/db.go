package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var error error
	db, error = gorm.Open(sqlite.Open("diary.db"), &gorm.Config{})
	if error != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Diary{})
}
