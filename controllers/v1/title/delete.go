package controllers

import (
	"chrisevett/pointy/v1/core"
	"chrisevett/pointy/v1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TitleDelete(c *gin.Context) {
	var title []core.Title
	if err := services.DB.Delete(&title).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to delete record"})
	}

	c.JSON(http.StatusOK, gin.H{"data": title})
}
