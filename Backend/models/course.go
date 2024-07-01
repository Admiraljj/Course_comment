package models

import "github.com/jinzhu/gorm"

// Course 模型
type Course struct {
	gorm.Model
	CourseName string `gorm:"type:varchar(50);not null"`
	Credits    int    `gorm:"not null"`
	CourseType string `gorm:"type:varchar(50);not null"`
	TeacherId  int    `gorm:"not null"`
}
