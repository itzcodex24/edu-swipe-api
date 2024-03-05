package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itzcodex24/edu-swipe-api/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.Get)
	r.Run(":8081")
}
