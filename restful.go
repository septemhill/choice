package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getCurrentCoordinate(c *gin.Context) {
	fmt.Println("DDD")
}

func moveTo(c *gin.Context) {
	fmt.Println("GGG")
}

func RestfulService(router *gin.Engine) {
	v1 := router.Group("/api/v1/map")

	v1.GET("/currCord", getCurrentCoordinate)
	v1.POST("/moveTo", moveTo)

	router.Run(":9999")
}
