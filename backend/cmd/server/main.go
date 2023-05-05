package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
)

var server *gin.Engine

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config, %v", err)
	}
	fmt.Println(config)

	server = gin.Default()
}

func main() {
	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is running!"})
	})

	log.Fatal(server.Run(":8080"))
}
