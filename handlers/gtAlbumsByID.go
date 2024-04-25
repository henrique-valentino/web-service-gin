package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henrique-valentino/web-service-gin/models"
)

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range models.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
