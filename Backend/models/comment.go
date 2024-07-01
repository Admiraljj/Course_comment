package models

import "github.com/jinzhu/gorm"

// Comment 模型
type Comment struct {
	gorm.Model
	CourseId    int    `gorm:"not null"`
	UserId      int    `gorm:"not null"`
	CommentText string `gorm:"type:text;not null"`
	CommentDate string `gorm:"type:datetime;not null"`
}
