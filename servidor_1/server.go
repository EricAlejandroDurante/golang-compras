package main

import (
	"servidor_1/controllers"
	"servidor_1/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/api/clientes/iniciar_sesion", controllers.IniciarSesion)
	r.POST("/api/compras", controllers.CreateCompra)
	r.POST("/api/productos", controllers.CreateProduct)
	r.GET("/api/productos", controllers.FindProducts)
	r.PUT("/api/productos/:id", controllers.UpdateProduct)
	r.DELETE("/api/productos/:id", controllers.DeleteProduct)
	r.GET("/api/estadisticas", controllers.GetStats)
	r.Run(":5000")
}
