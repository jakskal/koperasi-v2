package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/internal/service/user"
	"github.com/jakskal/koperasi-v2/pkg/dto"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"gorm.io/gorm"
)

type UserController struct {
	userService user.UserService
}

func InitUserController(db *gorm.DB) *UserController {
	return &UserController{
		userService: user.NewUserService(db),
	}
}

func (u UserController) Register(c *gin.Context) {
	print("register called")
}

func (u UserController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		log.Fatal("failed bind sruct", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	response, err := u.userService.Login(ctx, loginRequest)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (u UserController) List(c *gin.Context) {
	ctx := c.Request.Context()

	request := &dto.GetUsersRequest{
		PageSize: 10,
		Page:     1,
	}

	if err := c.BindQuery(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind query", err.Error())
		return
	}

	response, err := u.userService.GetUsers(ctx, *request)
	if err != nil {
		dto.ErrorResponse(c, http.StatusInternalServerError, "failed get users", err.Error())
		return
	}

	dto.SuccessResponse(c, http.StatusOK, response.Users, &response.Pagination)
}

func (u UserController) GetUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert user id to int", err.Error())
		return
	}

	ctx := c.Request.Context()

	user, err := u.userService.GetUser(ctx, userID)
	if err != nil {
		dto.ErrorResponse(c, http.StatusInternalServerError, "failed to get user with specified id", err.Error())
		return
	}

	dto.SuccessResponse(c, http.StatusOK, user, nil)
}

func (u UserController) GetProfile(c *gin.Context) {
	userIDParam, _ := c.Get(middleware.UserIDClaim)
	userID := userIDParam.(int)

	ctx := c.Request.Context()

	user, err := u.userService.GetUser(ctx, userID)
	if err != nil {
		dto.ErrorResponse(c, http.StatusInternalServerError, "failed to get user with specified id", err.Error())
		return
	}

	dto.SuccessResponse(c, http.StatusOK, user, nil)
}

func (u UserController) Create(c *gin.Context) {
	var request dto.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	ctx := c.Request.Context()

	response, err := u.userService.CreateUser(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, response, nil)
}

func (u UserController) Update(c *gin.Context) {
	var request dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded bind payload", err.Error())
		return
	}

	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert user id to int", err.Error())
		return
	}
	request.ID = userID

	ctx := c.Request.Context()

	response, err := u.userService.UpdateUser(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, response, nil)
}

func (u UserController) Delete(c *gin.Context) {

	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		dto.ErrorResponse(c, http.StatusUnprocessableEntity, "failded convert user id to int", err.Error())
		return
	}

	ctx := c.Request.Context()

	err = u.userService.DeleteUser(ctx, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	dto.SuccessResponse(c, http.StatusOK, nil, nil)
}
