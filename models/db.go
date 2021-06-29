package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var error error
	Db, error = gorm.Open(sqlite.Open("diary.db"), &gorm.Config{})
	if error != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&Diary{})
}
