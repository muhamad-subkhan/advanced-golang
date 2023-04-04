package main

import (
	"net/http"

	"simple-api/auth"
	"simple-api/middleware"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/login", auth.LoginHandler)

	r.GET("/test", middleware.AuthValid, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  "hello world",
		})
	})

	return r

}

func main() {
	r := setupRouter()

	r.Run()

}