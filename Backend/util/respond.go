package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Respond(c *gin.Context, httpCode int, code int, message string, data interface{}) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func RespondSuccess(c *gin.Context, data interface{}) {
	Respond(c, http.StatusOK, 0, "success", data)
}

func RespondError(c *gin.Context, httpCode int, code int, message string) {
	Respond(c, httpCode, code, message, nil)
}
