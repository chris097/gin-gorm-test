package controller

import (
	"github.com/chris097/gin-gorm-test/config"
	"github.com/chris097/gin-gorm-test/models"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}
