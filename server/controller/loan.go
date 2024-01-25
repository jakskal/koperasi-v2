package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/internal/service/loan"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/paginator"
	"gorm.io/gorm"
)

type Loanontroller struct {
	loanService loan.LoanService
}

func InitLoanController(db *gorm.DB) *Loanontroller {
	return &Loanontroller{
		loanService: loan.NewLoanService(db),
	}
}

func (s Loanontroller) Create(c *gin.Context) {
	var request dto.CreateLoanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	err := s.loanService.Create(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to create loan",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s Loanontroller) List(c *gin.Context) {
	var request dto.GetLoanListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	page, pagesize := paginator.GetPaginationQuery(c)
	request.Page = page
	request.PageSize = pagesize

	resp, err := s.loanService.List(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to get loan list",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, resp.Loans, &resp.Pagination)
}

func (s Loanontroller) Get(c *gin.Context) {
	loanIDParam := c.Param("id")
	loanID, err := strconv.Atoi(loanIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert loan id to int", err.Error())
		return
	}

	ctx := c.Request.Context()

	response, err := s.loanService.Get(ctx, loanID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to get loan by id",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s Loanontroller) Update(c *gin.Context) {
	loanIDParam := c.Param("id")
	loanID, err := strconv.Atoi(loanIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert loan id to int", err.Error())
		return
	}

	var request dto.UpdateLoanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	request.ID = loanID
	err = s.loanService.Update(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to update loan",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s Loanontroller) Delete(c *gin.Context) {

	loanIDParam := c.Param("id")
	loanID, err := strconv.Atoi(loanIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert loan id to int", err.Error())
		return
	}
	ctx := c.Request.Context()

	err = s.loanService.Delete(ctx, loanID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to delete loan",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s Loanontroller) CreateLoanType(c *gin.Context) {
	var request dto.CreateLoanTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	err := s.loanService.CreateLoanType(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to create loan type",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s Loanontroller) UpdateLoanType(c *gin.Context) {

	loanTypeIDParam := c.Param("id")
	loanTypeID, err := strconv.Atoi(loanTypeIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert loan id to int", err.Error())
		return
	}

	var request dto.UpdateLoanTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	fmt.Println("request")
	fmt.Println(request)

	ctx := c.Request.Context()

	request.ID = loanTypeID
	err = s.loanService.UpdateLoanType(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to update loan type",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s Loanontroller) ListLoanType(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := s.loanService.ListLoanType(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to update loan type",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, res, nil)
}

func (s Loanontroller) CreateLoanInstallment(c *gin.Context) {
	var request dto.CreateLoanInstallmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	err := s.loanService.CreateLoanInstallment(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to create loan",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s Loanontroller) ListLoanInstallment(c *gin.Context) {
	var request dto.GetLoanInstallmentListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	page, pagesize := paginator.GetPaginationQuery(c)
	request.Page = page
	request.PageSize = pagesize

	resp, err := s.loanService.ListLoanInstallment(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to get loan list",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, resp.LoanInstallments, &resp.Pagination)
}

func (s Loanontroller) GetLoanInstallment(c *gin.Context) {
	loanInstallmentIDParam := c.Param("id")
	loanInstallmentID, err := strconv.Atoi(loanInstallmentIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert loan id to int", err.Error())
		return
	}

	ctx := c.Request.Context()

	response, err := s.loanService.GetLoanInstallment(ctx, loanInstallmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to get loan by id",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s Loanontroller) UpdateLoanInstallment(c *gin.Context) {
	loanInstallmentIDParam := c.Param("id")
	loanInstallmentID, err := strconv.Atoi(loanInstallmentIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert loan id to int", err.Error())
		return
	}

	var request dto.UpdateLoanInstallmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	request.ID = loanInstallmentID
	err = s.loanService.UpdateLoanInstallment(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to update loan",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}

func (s Loanontroller) DeleteLoanInstallment(c *gin.Context) {

	loanInstallmentIDParam := c.Param("id")
	loanInstallmentID, err := strconv.Atoi(loanInstallmentIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert loan id to int", err.Error())
		return
	}
	ctx := c.Request.Context()

	err = s.loanService.DeleteInstallment(ctx, loanInstallmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to delete loan",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}
