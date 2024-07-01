package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 导入 MySQL 驱动
)

// User 模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
	Role     string `gorm:"type:varchar(10);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
}

// Course 模型
type Course struct {
	gorm.Model
	CourseName string `gorm:"type:varchar(50);not null"`
	Credits    int    `gorm:"not null"`
	CourseType string `gorm:"type:varchar(50);not null"`
	TeacherId  int    `gorm:"not null"`
}

// Comment 模型
type Comment struct {
	gorm.Model
	CourseId    int    `gorm:"not null"`
	UserId      int    `gorm:"not null"`
	CommentText string `gorm:"type:text;not null"`
	CommentDate string `gorm:"type:datetime;not null"`
}

// Rating 模型
type Rating struct {
	gorm.Model
	CourseId   int    `gorm:"not null"`
	UserId     int    `gorm:"not null"`
	Rating     int    `gorm:"not null"`
	RatingDate string `gorm:"type:datetime;not null"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.Run() // listen and serve on
	fmt.Println("Hello, World!")
	db := InitDB()
	if db != nil {
		defer db.Close()
	} else {
		fmt.Println("Database connection failed, exiting...")
		return
	}
	db.AutoMigrate(&User{}, &Course{}, &Comment{}, &Rating{})

}

func InitDB() *gorm.DB {
	drive := "mysql"
	username := "root"
	password := "1029admiral"
	host := "127.0.0.1"
	port := "3306"
	database := "course_comment"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open(drive, args)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil
	}
	return db
}
