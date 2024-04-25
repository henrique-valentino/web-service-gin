package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henrique-valentino/web-service-gin/models"
)

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Album invalid"})
		return
	}
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
