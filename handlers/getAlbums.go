package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henrique-valentino/web-service-gin/models"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}
