package dto

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, httpCode int, msg string, errors string) {
	if statusText := http.StatusText(httpCode); len(statusText) == 0 {
		httpCode = http.StatusBadRequest
	}

	c.JSON(httpCode, StandardErrorResponse{
		Code:    httpCode,
		Message: msg,
		Errors:  errors,
	})
}

func ErrorWrap(msg string, err error) error {
	return errors.New(fmt.Sprintf("%s, err: %s", msg, err.Error()))
}
