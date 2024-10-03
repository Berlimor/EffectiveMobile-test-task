package routers

import (
	"EffectiveMobile/m/docs"
	"EffectiveMobile/m/middleware"
	orm "EffectiveMobile/m/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title EffectiveMobile Test Task API
// @version 1.0
// @description REST API for EffectiveMobile test task.

// InitRouter initializes base Gin router with middlewares
func InitRouter() *gin.Engine {
	r := gin.Default()
	db := orm.InitDB()
	r.Use(middleware.SetDBConnection(db))
	
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "EffectiveMobile Test Task API"
	docs.SwaggerInfo.Version = "1.0"

	// Swagger docs at /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return r
}