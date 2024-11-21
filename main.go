package main

import (
	"fmt"
	"uservideoservice/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userHandler := handler.NewUserHandler()
	router.POST("/user", userHandler.AddUser)
	router.GET("/user", userHandler.GetUser)

	secureRouter := router.Group("/secure")
	secureRouter.Use(userHandler.VerifyUser())
	videoHandler := handler.NewVideoHandler()
	secureRouter.POST("/video", videoHandler.AddVideo)
	secureRouter.GET("/video", videoHandler.GetVideo)

	fmt.Println("running server")
	router.Run() // Listen and serve on 0.0.0.0:8080
}
