package main

import (
	"examples.com/controller"
	"examples.com/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoServices      = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":3000")
}
