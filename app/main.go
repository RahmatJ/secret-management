package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.New()

	api := r.Group("/api")
	internal := api.Group("/private")

	//	handler
	internal.GET("/ping", ping)

	err := r.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "pong",
	})
}
