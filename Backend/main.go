package main

import (
	"Backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 导入 MySQL 驱动
)

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
	db.AutoMigrate(&models.User{}, &models.Course{}, &models.Comment{})

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
