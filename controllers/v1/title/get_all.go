package controllers

import (
	"chrisevett/pointy/v1/core"
	"chrisevett/pointy/v1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TitleGetAll(c *gin.Context) {
	var titles []core.Title

	// todo add error handling
	// todo add error handling common helper
	services.DB.Find(&titles)
	c.JSON(http.StatusOK, gin.H{"data": titles})
}
