package main

import (
	"github.com/chris097/gin-gorm-test/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoute(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
