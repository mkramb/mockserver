package main

import (
	"mkramb/mockserver/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", pkg.GetAlbums)
	router.GET("/albums/:id", pkg.GetAlbumByID)
	router.POST("/albums", pkg.PostAlbums)

	router.Run("localhost:8080")
}
