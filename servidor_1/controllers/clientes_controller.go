package controllers

import (
	"net/http"
	"servidor_1/models"

	"github.com/gin-gonic/gin"
)

type ClienteInput struct {
	Id_Cliente int    `json:"id_cliente" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}

func IniciarSesion(c *gin.Context) {
	var cliente models.Cliente
	var input ClienteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("id_cliente = ?", input.Id_Cliente).First(&cliente).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if cliente.Contrasena == input.Contrasena {
		c.JSON(200, gin.H{"acceso_valido": true})
	} else {
		c.JSON(200, gin.H{"acceso_valido": false})
	}
}
