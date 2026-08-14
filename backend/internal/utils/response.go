package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weaver/api/internal/models"
)

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, models.Response{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Timestamp:  time.Now(),
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, models.Response{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
		Timestamp:  time.Now(),
	})
}

func CreatedResponse(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, 201, message, data)
}

func OKResponse(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, 200, message, data)
}

func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, 400, message)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, 401, message)
}

func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, 404, message)
}

func InternalErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, 500, message)
}
