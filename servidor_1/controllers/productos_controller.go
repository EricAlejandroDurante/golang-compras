package controllers

import (
	"net/http"
	"servidor_1/models"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	Nombre              string `json:"nombre" binding:"required"`
	Cantidad_disponible int    `json: "cantidad_disponible" binding:"required"`
	Precio_unitario     int    `json:"precio_unitario" binding:"required"`
}

func CreateProduct(c *gin.Context) {
	// Validate input
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	product := models.Producto{Nombre: input.Nombre, Cantidad_disponible: input.Cantidad_disponible, Precio_unitario: input.Precio_unitario}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"id_producto": product.Id_producto})
}

func FindProducts(c *gin.Context) {
	var products []models.Producto
	models.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Producto
	id := c.Param("id")
	if err := models.DB.Where("id_producto = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&product).Where("id_producto = ?", id).Update(input)
	c.JSON(http.StatusOK, gin.H{"id_producto": product.Id_producto})
}

func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var product models.Producto
	id := c.Param("id")
	if err := models.DB.Where("id_producto = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Where("id_producto = ?", id).Delete(&product)
	c.JSON(http.StatusOK, gin.H{"id_producto": product.Id_producto})
}
