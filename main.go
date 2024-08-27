package main

import (
	"github.com/chris097/gin-gorm-test/config"
	"github.com/chris097/gin-gorm-test/middleware"
	"github.com/chris097/gin-gorm-test/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	config.Coonect()
	routes.UserRoute(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
