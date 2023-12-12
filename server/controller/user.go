package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func InitUserController() *UserController {
	return &UserController{}
}

func (u UserController) Register(c *gin.Context) {
	print("register called")
}

func (u UserController) Login(c *gin.Context) {
	print("Login called")
}

func (u UserController) List(c *gin.Context) {
	print("List called")
}

func (u UserController) GetUser(c *gin.Context) {
	print("GetUser called")
}
