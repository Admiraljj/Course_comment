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

func CommentControllers(r *gin.Engine, db *gorm.DB) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 替换为你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	getCommentsByCourseId(r, db)
	addComment(r, db)
}

func getCommentsByCourseId(r *gin.Engine, db *gorm.DB) {
	r.GET("/comment/:course_id", func(c *gin.Context) {
		courseId := c.Param("course_id")
		var comments []models.Comment
		if err := db.Where("course_id = ?", courseId).Find(&comments).Error; err != nil {
			util.RespondError(c, http.StatusBadRequest, 1, err.Error())
			return
		}
		util.RespondSuccess(c, comments)
	})
}

func addComment(r *gin.Engine, db *gorm.DB) {
	r.POST("/comment/add", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		user, err := util.ParseToken(token)
		if err != nil {
			util.RespondError(c, http.StatusBadRequest, 1, "请先登录")
			return
		}
		var comment models.Comment
		if err := c.ShouldBind(&comment); err != nil {
			util.RespondError(c, http.StatusBadRequest, 2, err.Error())
			return
		}
		if comment.CommentText == "" {
			util.RespondError(c, http.StatusBadRequest, 3, "评论内容不能为空")
			return
		}
		comment.UserName = user.Username
		comment.UserId = int(user.ID)
		comment.CommentDate = time.Now().Format("2006-01-02 15:04:05")
		if err := db.Create(&comment).Error; err != nil {
			util.RespondError(c, http.StatusInternalServerError, 3, err.Error())
			return
		}
		util.RespondSuccess(c, gin.H{"评论": comment.CommentText + "创建成功"})
	})
}

//func getAllCourses(r *gin.Engine, db *gorm.DB) {
//	r.GET("/courses", func(c *gin.Context) {
//		var courses []models.Course
//		if err := db.Find(&courses).Error; err != nil {
//			util.RespondError(c, http.StatusInternalServerError, 1, err.Error())
//			return
//		}
//		util.RespondSuccess(c, courses)
//	})
//}
