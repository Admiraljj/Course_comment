package models

import "github.com/jinzhu/gorm"

// User 模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" form:"username" binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password" form:"password" binding:"required"`
	Role     string `gorm:"type:varchar(10);not null" json:"role" form:"role"`
	Email    string `gorm:"type:varchar(255);not null" json:"email" form:"email"`
}
