package routes

import (
	"github.com/chris097/gin-gorm-test/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
