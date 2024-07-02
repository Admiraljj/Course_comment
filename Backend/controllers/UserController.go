package controllers

import (
	"Backend/models"
	"Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func UserControllers(r *gin.Engine, db *gorm.DB) {
	UserRegister(r, db)
	UserLogin(r, db)
	GetUserInfoByToken(r)
}

func UserRegister(r *gin.Engine, db *gorm.DB) {
	r.POST("/user/register", func(c *gin.Context) {
		var user models.User

		// 使用 ShouldBind 解析请求数据，不论是 JSON 还是 form-data
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 检查用户名是否已经存在
		var existingUserByUsername models.User
		if err := db.Where("username = ?", user.Username).First(&existingUserByUsername).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
			return
		}

		// 检查电子邮件是否已经存在
		var existingUserByEmail models.User
		if err := db.Where("email = ?", user.Email).First(&existingUserByEmail).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已存在"})
			return
		}
		// 检查电子邮件是否为空
		if user.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱不能为空"})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 对密码进行sha256加密
		user.Password = util.Encryption(user.Password)

		var existingUser models.User
		if err := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": util.GenerateToken(existingUser),
		})
	})
}

// 通过token获取username的接口
func GetUserInfoByToken(r *gin.Engine) {
	r.GET("/user/info", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		user, _ := util.ParseToken(token)
		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"role":     user.Role,
			"email":    user.Email,
		})
	})
}
