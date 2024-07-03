package models

import "github.com/jinzhu/gorm"

// Rating 模型
type Rating struct {
	gorm.Model
	CourseId int `gorm:"not null" json:"course_id" form:"course_id" binding:"required"`
	UserId   int `gorm:"not null" json:"user_id" form:"user_id" binding:"required"`
	Rating   int `gorm:"not null" json:"rating" form:"rating" binding:"required"`
}
