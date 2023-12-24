package rest

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/server/controller"
	"gorm.io/gorm"
)

func RegisterSavingRoute(router *gin.RouterGroup, dbConn *gorm.DB) {
	if router == nil {
		panic(errors.New("router is nil"))
	}

	ctrl := controller.InitSavingController(dbConn)

	routerAdmin := router.Group("/admin")
	{
		// TODO: add middleware
		routerAdmin.Use(middleware.AuthAdminAcess())
		routerAdmin.POST("/saving", ctrl.Create)
		routerAdmin.GET("/saving", ctrl.List)
		routerAdmin.GET("/saving/:id", ctrl.Get)
		routerAdmin.PUT("/saving/:id", ctrl.Update)
		routerAdmin.DELETE("/saving/:id", ctrl.Delete)
		routerAdmin.POST("/saving-type", ctrl.CreateSavingType)
		routerAdmin.GET("/saving-type", ctrl.ListSavingType)
		routerAdmin.PUT("/saving-type/:id", ctrl.UpdateSavingType)
	}
}
