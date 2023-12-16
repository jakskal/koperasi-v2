package rest

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/server/controller"
	"gorm.io/gorm"
)

func RegisterUserRoute(router *gin.RouterGroup, dbConn *gorm.DB) {
	if router == nil {
		panic(errors.New("router is nil"))
	}

	ctrl := controller.InitUserController(dbConn)
	router.POST("/register", ctrl.Register)
	router.POST("/login", ctrl.Login)

	routerAdmin := router.Group("/admin")
	{
		// TODO: add middleware
		routerAdmin.Use(middleware.AuthAdminAcess())
		routerAdmin.POST("/user", ctrl.Create)
		routerAdmin.GET("/users", ctrl.List)
		routerAdmin.GET("/user/:user_id", ctrl.GetUser)
		routerAdmin.PUT("/user/:user_id", ctrl.Update)
		routerAdmin.DELETE("/user/:user_id", ctrl.Delete)
	}
	routerUser := router.Group("/user")
	routerUser.Use(middleware.AuthAcess())
	routerUser.GET("/profile", ctrl.GetProfile)

}
