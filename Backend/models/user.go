package models

import "github.com/jinzhu/gorm"

// User 模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
	Role     string `gorm:"type:varchar(10);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
}
