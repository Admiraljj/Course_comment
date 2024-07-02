package models

import "github.com/jinzhu/gorm"

// Course 模型
type Course struct {
	gorm.Model
	CourseName  string `gorm:"type:varchar(50);not null" json:"course_name" form:"course_name"`
	Credits     int    `gorm:"not null" json:"credits" form:"credits"`
	CourseType  string `gorm:"type:varchar(50);not null" json:"course_type" form:"course_type"`
	TeacherName string `gorm:"type:varchar(50);not null" json:"teacher_name" form:"teacher_name"`
}
