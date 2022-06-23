package main

import (
	"io"
	"os"

	gindump "github.com/tpkeeper/gin-dump"

	"github.com/gin-gonic/gin"
	"github.com/leksyking/goGin/controller"
	"github.com/leksyking/goGin/middlewares"
	"github.com/leksyking/goGin/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setUpLog() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setUpLog()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":3000")
}
