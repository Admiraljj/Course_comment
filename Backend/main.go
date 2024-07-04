package main

import (
	"Backend/conf"
	"Backend/controllers"
	"Backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 导入 MySQL 驱动
)

func main() {
	r := gin.Default()
	db := conf.InitDB()
	if db != nil {
		defer db.Close()
	} else {
		fmt.Println("Database connection failed, exiting...")
		return
	}
	db.AutoMigrate(&models.User{}, &models.Course{}, &models.Comment{}, &models.Rating{})
	controllers.UserControllers(r, db)
	controllers.CommentControllers(r, db)
	controllers.CourseControllers(r, db)
	controllers.RatingControllers(r, db)

	err := r.Run()
	if err != nil {
		return
	}
}
