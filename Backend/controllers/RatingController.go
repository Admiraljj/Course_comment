package controllers

import (
	"Backend/models"
	"Backend/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func RatingControllers(r *gin.Engine, db *gorm.DB) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 替换为你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	addOrAlterRating(r, db)
	getRatingsByCourseId(r, db)
}

func addOrAlterRating(r *gin.Engine, db *gorm.DB) {
	r.POST("/rating/add", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			util.RespondError(c, http.StatusBadRequest, 1, "token不能为空")
			return
		}
		var rating models.Rating
		if err := c.ShouldBind(&rating); err != nil {
			util.RespondError(c, http.StatusBadRequest, 2, err.Error())
			return
		}
		user, err := util.ParseToken(token)
		if err != nil {
			util.RespondError(c, http.StatusUnauthorized, 3, "无效的 token")
			return
		}
		rating.UserId = int(user.ID)

		var existingRating models.Rating
		if err := db.Where("course_id = ? AND user_id = ?", rating.CourseId, rating.UserId).First(&existingRating).Error; err == nil {
			existingRating.Rating = rating.Rating
			db.Save(&existingRating)
			util.RespondSuccess(c, existingRating)
			return
		} else {
			db.Create(&rating)
			util.RespondSuccess(c, rating)
			return
		}
	})
}

func getRatingsByCourseId(r *gin.Engine, db *gorm.DB) {
	r.GET("/rating/:course_id", func(c *gin.Context) {
		courseId := c.Param("course_id")
		if courseId == "" {
			util.RespondError(c, http.StatusBadRequest, 1, "course_id不能为空")
			return
		}
		var ratings []models.Rating
		db.Where("course_id = ?", courseId).Find(&ratings)
		util.RespondSuccess(c, ratings)
	})
}
