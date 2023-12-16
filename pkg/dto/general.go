package dto

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type StandardResponse struct {
	Code       int                   `json:"code"`
	Message    string                `json:"message"`
	Data       interface{}           `json:"data"`
	Pagination *BasePaginationResult `json:"page_info"`
}

type StandardErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

type StandardDateResponse struct {
	CreatedDate time.Time  `json:"created_date"`
	CreatedBy   int        `json:"created_by"`
	UpdatedDate *time.Time `json:"updated_date"`
	UpdatedBy   *int       `json:"updated_by"`
	DeletedDate *time.Time `json:"deleted_date,omitempty"`
	DeletedBy   *int       `json:"deleted_by,omitempty"`
}

type BasePaginationResult struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Count    int `json:"count"`
}

func SuccessResponse(c *gin.Context, httpCode int, data interface{}, pagination *BasePaginationResult) {
	if statusText := http.StatusText(httpCode); len(statusText) == 0 {
		httpCode = http.StatusOK
	}

	c.JSON(httpCode, StandardResponse{
		Code:       httpCode,
		Message:    http.StatusText(httpCode),
		Data:       data,
		Pagination: pagination,
	})
}
