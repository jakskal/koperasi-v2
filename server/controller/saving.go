package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/internal/service/saving"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/paginator"
	"gorm.io/gorm"
)

type SavingController struct {
	savingService saving.SavingService
}

func InitSavingController(db *gorm.DB) *SavingController {
	return &SavingController{
		savingService: saving.NewSavingService(db),
	}
}

func (s SavingController) Create(c *gin.Context) {
	var request dto.CreateSavingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	err := s.savingService.Create(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to create saving",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s SavingController) List(c *gin.Context) {
	var request dto.GetSavingListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	page, pagesize := paginator.GetPaginationQuery(c)
	request.Page = page
	request.PageSize = pagesize

	resp, err := s.savingService.List(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to get saving list",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, resp.Savings, &resp.Pagination)
}

func (s SavingController) Get(c *gin.Context) {
	savingIDParam := c.Param("id")
	savingID, err := strconv.Atoi(savingIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert saving id to int", err.Error())
		return
	}

	ctx := c.Request.Context()

	response, err := s.savingService.Get(ctx, savingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to get saving by id",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s SavingController) Update(c *gin.Context) {
	savingIDParam := c.Param("id")
	savingID, err := strconv.Atoi(savingIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert saving id to int", err.Error())
		return
	}

	var request dto.UpdateSavingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	request.ID = savingID
	err = s.savingService.Update(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to update saving",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s SavingController) Delete(c *gin.Context) {

	savingIDParam := c.Param("id")
	savingID, err := strconv.Atoi(savingIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert saving id to int", err.Error())
		return
	}
	ctx := c.Request.Context()

	err = s.savingService.Delete(ctx, savingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to delete saving",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s SavingController) CreateSavingType(c *gin.Context) {
	var request dto.CreateSavingTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	err := s.savingService.CreateSavingType(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to create saving type",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s SavingController) UpdateSavingType(c *gin.Context) {

	savingTypeIDParam := c.Param("id")
	savingTypeID, err := strconv.Atoi(savingTypeIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert saving id to int", err.Error())
		return
	}

	var request dto.UpdateSavingTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	request.ID = savingTypeID
	err = s.savingService.UpdateSavingType(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to update saving type",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s SavingController) ListSavingType(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := s.savingService.ListSavingType(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to update saving type",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, res, nil)
}
