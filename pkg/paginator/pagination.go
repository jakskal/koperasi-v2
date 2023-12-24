package paginator

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PaginateGin add pagination gorm with http server gin.
func PaginateGin(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// // GetPaginationQuery get page and pageSize for pagination purpose
func GetPaginationQuery(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "0"))
	if page == 0 {
		page = 1
	}

	switch {
	case pageSize > 20:
		pageSize = 20
	case pageSize <= 0:
		pageSize = 15
	}

	return page, pageSize
}
