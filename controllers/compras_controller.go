package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"tarea_1_sds/models"
)

type CompraInput struct{
	Id_cliente int `json:"id_cliente" binding:"required"`
	Productos []DetalleInput `json:"productos" binding:"required"`
}

type DetalleInput struct{
	Id_producto int `json:"id_producto" binding:"required"`
	Cantidad int `json:"cantidad" binding:"required"`
}

func CreateCompra(c *gin.Context){
	var input CompraInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	compra := models.Compra{Id_cliente: input.Id_cliente}
	models.DB.Create(&compra)
	for _, product := range input.Productos{
		detail := models.Detalle{Id_compra: compra.Id_compra, Id_producto: product.Id_producto, Cantidad: product.Cantidad, Fecha: time.Now().Format(time.RFC3339)}
		models.DB.Create(&detail)
	}
	c.JSON(http.StatusOK, gin.H{"id_compra": compra.Id_compra})
}
