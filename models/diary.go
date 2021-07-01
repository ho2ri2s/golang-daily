package models

import "time"

type Diary struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time `gorm:"index"`
}

/** This equal To Below
type Diary struct {
	gorm.Model
	Title     string
	Content   string
}
*/
