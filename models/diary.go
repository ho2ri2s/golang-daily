package models

import "time"

// https://echo.labstack.com/guide/binding/
type Diary struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `json:"title" form:"diary_title"`
	Content   string `json:"content" form:"diary_content"`
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

type Diaries []Diary
