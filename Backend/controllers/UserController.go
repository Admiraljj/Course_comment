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

func UserControllers(r *gin.Engine, db *gorm.DB) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 替换为你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	UserRegister(r, db)
	UserLogin(r, db)
	GetUserInfoByToken(r)
}

func UserRegister(r *gin.Engine, db *gorm.DB) {
	r.POST("/user/register", func(c *gin.Context) {
		var user models.User

		// 使用 ShouldBind 解析请求数据，不论是 JSON 还是 form-data
		if err := c.ShouldBind(&user); err != nil {
			util.RespondError(c, http.StatusBadRequest, 1, err.Error())
			return
		}

		// 检查用户名是否已经存在
		var existingUserByUsername models.User
		if err := db.Where("username = ?", user.Username).First(&existingUserByUsername).Error; err == nil {
			util.RespondError(c, http.StatusBadRequest, 2, "用户名已存在")
			return
		}

		// 检查电子邮件是否已经存在
		var existingUserByEmail models.User
		if err := db.Where("email = ?", user.Email).First(&existingUserByEmail).Error; err == nil {
			util.RespondError(c, http.StatusBadRequest, 3, "邮箱已存在")
			return
		}

		// 检查电子邮件是否为空
		if user.Email == "" {
			util.RespondError(c, http.StatusBadRequest, 4, "邮箱不能为空")
			return
		}

		// 保存原有密码
		originalPassword := user.Password
		// 对密码进行sha256加密
		user.Password = util.Encryption(user.Password)
		// 设置用户角色
		user.Role = "user" // 默认角色为 user
		// 保存用户到数据库
		if err := db.Create(&user).Error; err != nil {
			util.RespondError(c, http.StatusInternalServerError, 5, err.Error())
			return
		}

		util.RespondSuccess(c, gin.H{
			"username": user.Username,
			"password": originalPassword,
			"email":    user.Email,
		})
	})
}

func UserLogin(r *gin.Engine, db *gorm.DB) {
	r.POST("/user/login", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			util.RespondError(c, http.StatusBadRequest, 1, err.Error())
			return
		}

		// 对密码进行sha256加密
		user.Password = util.Encryption(user.Password)

		var existingUser models.User
		if err := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser).Error; err != nil {
			util.RespondError(c, http.StatusBadRequest, 2, "用户名或密码错误")
			return
		}

		util.RespondSuccess(c, gin.H{
			"token": util.GenerateToken(existingUser),
		})
	})
}

func GetUserInfoByToken(r *gin.Engine) {
	r.GET("/user/info", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			util.RespondError(c, http.StatusBadRequest, 1, "无效的token")
			return
		}

		user, err := util.ParseToken(token)
		if err != nil {
			util.RespondError(c, http.StatusUnauthorized, 2, "无效的token")
			return
		}

		if user == nil {
			util.RespondError(c, http.StatusUnauthorized, 3, "用户未找到")
			return
		}

		util.RespondSuccess(c, gin.H{
			"username": user.Username,
			"role":     user.Role,
			"email":    user.Email,
		})
	})
}
