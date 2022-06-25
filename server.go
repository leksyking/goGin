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

func setUpLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setUpLogOutput()

	server := gin.New()

	server.Static("/css", "./template/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.Save(ctx))
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll(ctx))
	}

	server.Run(":3000")
}
