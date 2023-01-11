package main

import (
	"servidor_3/controllers"
	"servidor_3/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/api/clientes/estado_despacho/:id", controllers.FindDespacho)
	r.Run(":5000")
}
