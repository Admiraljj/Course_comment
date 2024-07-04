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

func CourseControllers(r *gin.Engine, db *gorm.DB) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 替换为你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	addCourse(r, db)
	getAllCourses(r, db)
	deleteCourse(r, db)
	getCourseInfoById(r, db)
}

func addCourse(r *gin.Engine, db *gorm.DB) {
	r.POST("/course/add", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			util.RespondError(c, http.StatusBadRequest, 1, "token不能为空")
			return
		}
		var course models.Course
		if err := c.ShouldBind(&course); err != nil {
			util.RespondError(c, http.StatusBadRequest, 2, err.Error())
			return
		}
		user, _ := util.ParseToken(token)
		// 从数据库中检查用户是否存在，且角色为管理员
		var existingUser models.User
		if err := db.Where("username = ? AND role = ?", user.Username, "admin").First(&existingUser).Error; err != nil {
			util.RespondError(c, http.StatusBadRequest, 3, "用户不存在或无权限")
			return
		}
		// 检查课程名是否为空
		if course.CourseName == "" {
			util.RespondError(c, http.StatusBadRequest, 4, "课程名不能为空")
			return
		}
		// 检查学分是否为空
		if course.Credits == 0 {
			util.RespondError(c, http.StatusBadRequest, 5, "学分不能为空")
			return
		}
		// 检查教师ID是否为空
		if course.TeacherName == "" {
			util.RespondError(c, http.StatusBadRequest, 6, "教师名称不能为空")
			return
		}
		// 检查课程类型是否选择
		if course.CourseType == "" {
			util.RespondError(c, http.StatusBadRequest, 7, "课程类型不能为空")
			return
		}
		// 检查课程是否已存在，需要课程名和教师名一致
		var existingCourse models.Course
		if err := db.Where("course_name = ? AND teacher_name = ?", course.CourseName, course.TeacherName).First(&existingCourse).Error; err == nil {
			util.RespondError(c, http.StatusBadRequest, 8, "课程已存在")
			return
		}
		// 保存课程到数据库
		if err := db.Create(&course).Error; err != nil {
			util.RespondError(c, http.StatusInternalServerError, 9, err.Error())
			return
		}

		util.RespondSuccess(c, gin.H{
			"course": course.CourseName + "添加成功",
		})
	})
}

func getAllCourses(r *gin.Engine, db *gorm.DB) {
	r.GET("/courses", func(c *gin.Context) {
		var courses []models.Course
		if err := db.Find(&courses).Error; err != nil {
			util.RespondError(c, http.StatusInternalServerError, 1, err.Error())
			return
		}
		util.RespondSuccess(c, courses)
	})
}

func deleteCourse(r *gin.Engine, db *gorm.DB) {
	r.GET("/course/delete/:id", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			util.RespondError(c, http.StatusBadRequest, 1, "token不能为空")
			return
		}

		user, _ := util.ParseToken(token)
		// 从数据库中检查用户是否存在，且角色为管理员
		var existingUser models.User
		if err := db.Where("username = ? AND role = ?", user.Username, "admin").First(&existingUser).Error; err != nil {
			util.RespondError(c, http.StatusBadRequest, 3, "用户不存在或无权限")
			return
		}

		// 获取路径参数
		courseID := c.Param("id")

		// 根据 CourseID 删除课程, 同时删除 ratings 和 comments
		if err := db.Where("id = ?", courseID).Delete(&models.Course{}).Error; err != nil {
			util.RespondError(c, http.StatusInternalServerError, 4, err.Error())
			return
		}
		if err := db.Where("course_id = ?", courseID).Delete(&models.Rating{}).Error; err != nil {
			util.RespondError(c, http.StatusInternalServerError, 5, err.Error())
			return
		}
		if err := db.Where("course_id = ?", courseID).Delete(&models.Comment{}).Error; err != nil {
			util.RespondError(c, http.StatusInternalServerError, 6, err.Error())
			return
		}

		util.RespondSuccess(c, gin.H{
			"message": "课程删除成功",
		})
	})
}

func getCourseInfoById(r *gin.Engine, db *gorm.DB) {
	r.GET("/course/:course_id", func(c *gin.Context) {
		courseId := c.Param("course_id")
		var course models.Course
		if err := db.Where("id = ?", courseId).First(&course).Error; err != nil {
			util.RespondError(c, http.StatusBadRequest, 1, err.Error())
			return
		}
		util.RespondSuccess(c, course)
	})
}
