package rest

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/server/controller"
	"gorm.io/gorm"
)

func RegisterLoanRoute(router *gin.RouterGroup, dbConn *gorm.DB) {
	if router == nil {
		panic(errors.New("router is nil"))
	}

	ctrl := controller.InitLoanController(dbConn)

	routerAdmin := router.Group("/admin")
	{
		// TODO: add middleware
		routerAdmin.Use(middleware.AuthAdminAcess())
		routerAdmin.POST("/loan-type", ctrl.CreateLoanType)
		routerAdmin.GET("/loan-type", ctrl.ListLoanType)
		routerAdmin.PUT("/loan-type/:id", ctrl.UpdateLoanType)
		routerAdmin.POST("/loan", ctrl.Create)
		routerAdmin.GET("/loan", ctrl.List)
		routerAdmin.GET("/loan/:id", ctrl.Get)
		routerAdmin.PUT("/loan/:id", ctrl.Update)
		routerAdmin.DELETE("/loan/:id", ctrl.Delete)
		routerAdmin.POST("/loan-installment", ctrl.CreateLoanInstallment)
		routerAdmin.GET("/loan-installment", ctrl.ListLoanInstallment)
		routerAdmin.GET("/loan-installment/:id", ctrl.GetLoanInstallment)
		routerAdmin.PUT("/loan-installment/:id", ctrl.UpdateLoanInstallment)
		routerAdmin.DELETE("/loan-installment/:id", ctrl.DeleteLoanInstallment)
	}
}
