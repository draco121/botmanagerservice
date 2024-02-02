package routes

import (
	"github.com/draco121/botmanagerservice/controllers"
	"github.com/draco121/common/constants"
	"github.com/draco121/common/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(controllers controllers.Controllers, router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.POST("/bot", middlewares.AuthMiddleware(constants.Write), controllers.CreateBot)
	v1.GET("/bot", middlewares.AuthMiddleware(constants.Read), controllers.GetBot)
	v1.PATCH("/bot", middlewares.AuthMiddleware(constants.Write), controllers.UpdateBot)
	v1.DELETE("/bot", middlewares.AuthMiddleware(constants.Write), controllers.DeleteBot)
}
