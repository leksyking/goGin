package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leksyking/goGin/entity"
	"github.com/leksyking/goGin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(*gin.Context) entity.Video
	ShowAll(*gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
