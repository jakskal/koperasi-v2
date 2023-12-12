package rest

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/server/controller"
	"gorm.io/gorm"
)

func RegisterUserRoute(router *gin.RouterGroup, dbConn *gorm.DB) {
	if router == nil {
		panic(errors.New("router is nil"))
	}

	ctrl := controller.InitUserController()
	router.POST("/register", ctrl.Register)
	router.POST("/login", ctrl.Login)

	routerUnauth := router.Group("")
	{
		// TODO: add middleware
		// routerUnauth.Use(middleware.CheckXClientId())
		routerUnauth.GET("/list", ctrl.List)
		routerUnauth.GET("/detail/:user_id", ctrl.GetUser)
	}
}
