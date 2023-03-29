package routers

import (
	"web3-with-go/docs"

	"web3-with-go/pkg/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes := r.Group("/api/v1")
	{
		routes.GET("/current-block", controllers.BlockController.GetCurrentBlock)
	}

	return r
}
