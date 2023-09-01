package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (srv *server) routes() http.Handler {
	v1 := srv.router.Group("/v1")
	{
		system := v1.Group("/system")
		{
			system.GET("/version", srv.Version)
		}

		swagger := v1.Group("/swagger")
		{
			swagger.GET("", func(c *gin.Context) {
				c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			})
		}
	}
	srv.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Declare web routing table at here.
	return srv.router
}
