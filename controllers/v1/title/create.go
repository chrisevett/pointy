package controllers

import (
	"chrisevett/pointy/v1/core"
	"chrisevett/pointy/v1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// for validation
// todo derive this from model
// todo retrieve from bgg

func TitleCreate(c *gin.Context) {
	var input CreateTitleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	title := core.Title{Name: input.Name, Designer: input.Designer}
	if err := services.DB.Create(&title).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add record to database"})
	}

	c.JSON(http.StatusOK, gin.H{"data": title})
}
