package main

import (
	"gin/controller"
	"gin/service"

	"github.com/gin-gonic/gin"
)

var (
	videoservice service.VideoService = service.New()
	videocontroller controller.VideoController = controller.New(&videoservice)
)

func main() {
	server := gin.Default()

	server.GET("/list-videos", func(ctx *gin.Context) {
		ctx.JSON(200, videocontroller.FindAll())
	})
	server.POST("/store-video", func(ctx *gin.Context) {
		ctx.JSON(201, videocontroller.Save(ctx))
	})
	server.Run(":8080")
}