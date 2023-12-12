package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/jakskal/koperasi-v2/config"
	"github.com/jakskal/koperasi-v2/pkg/middleware"
	"github.com/jakskal/koperasi-v2/server/rest"
	"github.com/pkg/errors"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/gorm"
)

func StartServer(cfg *config.Config, DB *gorm.DB) error {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.New()
	r.Use(gin.CustomRecovery(middleware.ErrorHandler))

	if err := registerRoutes(r, cfg, DB); err != nil {
		return errors.Wrap(err, "failed to register routes")
	}

	addr := cfg.ServerAddr
	if addr == "" {
		addr = ":8081"
	}
	return r.Run(addr)
}

func registerRoutes(r *gin.Engine, cfg *config.Config, DB *gorm.DB) error {
	// Ping test
	r.GET("/nnd-aruok", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	apiRouter := r.Group("/api")
	{
		v1 := apiRouter.Group("/v1")
		{
			rest.RegisterUserRoute(v1, DB)
			// rest.RegisterHelperRoute(v1, DB)

			// adminRoute := v1.Group("/admin")
			// {
			// 	rest.RegisterAdminRuleRoute(adminRoute, DB)
			// 	rest.RegisterAdminUsageRoute(adminRoute, DB)
			// 	rest.RegisterAdminVoucherRoute(adminRoute, DB)
			// }
		}
	}

	// Register Swagger
	if cfg.Stage != config.Prod {
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return nil
}
