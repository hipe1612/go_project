package router

import (
	"crud-person/internal/controller"
	"github.com/gin-gonic/gin"
)

func initPersonRouter(router *gin.Engine) {
	personController := controller.NewPersonController()
	groupRouter := router.Group("/person/names")
	{
		groupRouter.POST("", personController.CreatePerson())
		groupRouter.GET("", personController.ListPerson())
		groupRouter.GET("/:id", personController.GetPerson())
		groupRouter.PATCH("/:id", personController.UpdatePerson())
		groupRouter.DELETE("/:id", personController.DeletePerson())
	}
}

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/check"))
	router.GET("/check", controller.HealthCheck())
	initPersonRouter(router)
	return router
}
