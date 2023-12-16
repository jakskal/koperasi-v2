package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/pkg/dto"
)

func ErrorHandler(c *gin.Context, err any) {
	code := http.StatusInternalServerError
	message := http.StatusText(code)
	errors := ""

	if _err, ok := err.(error); ok {
		message = _err.Error()
	}
	if _str, ok := err.(string); ok {
		message = _str
	}

	c.AbortWithStatusJSON(code, dto.StandardErrorResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
	})
}
