package models

import "time"

type Diary struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `json:"diary_title"`
	Content   string `json:"diary_content"`
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
