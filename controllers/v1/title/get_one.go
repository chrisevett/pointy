package controllers

import (
	"chrisevett/pointy/v1/core"
	"chrisevett/pointy/v1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TitleGetOne(c *gin.Context) {
	var title []core.Title

	//if err := services.DB.Find(&titles)
	if err := services.DB.Where("id = ?", c.Param("id")).First(&title).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": title})
}
