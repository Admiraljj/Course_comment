package models

import "github.com/jinzhu/gorm"

// Rating 模型
type Rating struct {
	gorm.Model
	CourseId   int    `gorm:"not null"`
	UserId     int    `gorm:"not null"`
	Rating     int    `gorm:"not null"`
	RatingDate string `gorm:"type:datetime;not null"`
}
