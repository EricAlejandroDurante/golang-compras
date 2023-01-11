package controllers

import (
	"net/http"
	"servidor_3/models"

	"github.com/gin-gonic/gin"
)

func FindDespacho(c *gin.Context) {
	var despacho models.Despacho
	id := c.Param("id")
	if err := models.DB.Where("id_despacho = ?", id).First(&despacho).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, despacho)
}
