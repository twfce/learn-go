package main

import (
	"simple-api/endpoints"
    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    router.GET("/albums", endpoints.GetAlbums)
	router.POST("/albums", endpoints.PostAlbums)
	router.GET("/albums/:id", endpoints.GetAlbumByID)
    router.Run("localhost:8080")
}