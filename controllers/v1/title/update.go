package controllers

import (
	"chrisevett/pointy/v1/core"
	"chrisevett/pointy/v1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TitleUpdate(c *gin.Context) {
	var input CreateTitleInput
	var title core.Title

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DB.First(&title).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	title = core.Title{Name: input.Name, Designer: input.Designer}
	if err := services.DB.Save(&title).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": title})
}
