package controllers

import (
	"net/http"
	"servidor_1/models"

	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	var best_selling_product models.Detalle
	var least_selling_product models.Detalle
	var highest_profit_product models.Detalle
	var lower_profit_product models.Detalle
	models.DB.Raw("select id_producto from detalle group by id_producto order by sum(cantidad) desc limit 1;").Scan(&best_selling_product)
	models.DB.Raw("select id_producto from detalle group by id_producto order by sum(cantidad) asc limit 1;").Scan(&least_selling_product)
	models.DB.Raw("select producto.id_producto from detalle inner join producto on detalle.id_producto = producto.id_producto group by detalle.id_producto order by sum(detalle.cantidad * producto.precio_unitario) desc;").Scan(&highest_profit_product)
	models.DB.Raw("select producto.id_producto from detalle inner join producto on detalle.id_producto = producto.id_producto group by detalle.id_producto order by sum(detalle.cantidad * producto.precio_unitario) asc;").Scan(&lower_profit_product)
	c.JSON(http.StatusOK, gin.H{"producto_mas_vendido": best_selling_product.Id_producto, "producto_menos_vendido": least_selling_product.Id_producto, "producto_mayor_ganancia": highest_profit_product.Id_producto, "producto_menor_ganancia": lower_profit_product.Id_producto})
}
