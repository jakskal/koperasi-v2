package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, httpCode int, msg string, errors []string) {
	fakeData := []string{}
	if statusText := http.StatusText(httpCode); len(statusText) == 0 {
		httpCode = http.StatusBadRequest
	}

	c.JSON(httpCode, StandardResponse{
		Code:    httpCode,
		Message: msg,
		Data:    fakeData,
		Errors:  errors,
	})
}
