package routes

import (
	"github.com/chris097/gin-gorm-test/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/blogs", controller.GetUsers)
	router.POST("/blog", controller.CreateUsers)
	router.GET("blog/:id", controller.GetUserByID)
	router.DELETE("blog/:id", controller.DeleteUser)
	router.PUT("blog/:id", controller.UpdateUser)
}
