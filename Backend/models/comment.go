package models

import "github.com/jinzhu/gorm"

// Comment 模型
type Comment struct {
	gorm.Model
	CourseId    int    `gorm:"not null" json:"course_id" form:"course_id" binding:"required"`
	UserId      int    `gorm:"not null" json:"user_id" form:"user_id" binding:"required"`
	CommentText string `gorm:"type:text;not null" json:"comment_text" form:"comment_text" binding:"required"`
	CommentDate string `gorm:"type:datetime;not null" json:"comment_date" form:"comment_date" binding:"required"`
}
