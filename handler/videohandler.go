package handler

import (
	"net/http"
	"uservideoservice/model"

	"github.com/gin-gonic/gin"
)

type IVideo interface {
	AddVideo(*gin.Context)
	GetVideo(*gin.Context)
}

func NewVideoHandler() IVideo {
	return &Video{
		videoData: make(map[string]model.Video),
	}
}

type Video struct {
	videoData map[string]model.Video
}

func (v *Video) AddVideo(ctx *gin.Context) {
	var video model.Video

	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v.videoData[video.Name] = video
	ctx.JSON(http.StatusCreated, gin.H{"message": "Video added"})
}

func (v *Video) GetVideo(ctx *gin.Context) {
	var videos = make([]model.Video, 0)
	for _, value := range v.videoData {
		videos = append(videos, value)
	}

	ctx.JSON(http.StatusCreated, videos)
}
