package main

import (
	"goorder/config"
	"goorder/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.GetOrderById)
	config.ConnectDatabase()

	r.Run()
}
