package controllers

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	pb "servidor_1/gen"
	"servidor_1/models"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type CompraInput struct {
	Id_cliente int            `json:"id_cliente" binding:"required"`
	Productos  []DetalleInput `json:"productos" binding:"required"`
}

type DetalleInput struct {
	Id_producto int `json:"id_producto" binding:"required"`
	Cantidad    int `json:"cantidad" binding:"required"`
}
type DespachoInput struct {
	Id_compra   int `json:"id_compra" binding:"required"`
	Id_despacho int `json:"id_despacho" binding:"required"`
}

func CreateCompra(c *gin.Context) {
	var input CompraInput
	var producto models.Producto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	compra := models.Compra{Id_cliente: input.Id_cliente}
	models.DB.Create(&compra)
	var requerido int
	var cantidad_actualizada int
	var detail models.Detalle
	for _, product := range input.Productos {
		models.DB.Where("id_producto = ?", product.Id_producto).First(&producto)
		requerido = producto.Cantidad_disponible - product.Cantidad
		cantidad_actualizada = producto.Cantidad_disponible - product.Cantidad
		if requerido < 0 {
			pedirProveedor(requerido, producto)
			cantidad_actualizada = 0
		}
		detail = models.Detalle{Id_compra: compra.Id_compra, Id_producto: product.Id_producto, Cantidad: product.Cantidad, Fecha: time.Now()}
		models.DB.Create(&detail)
		models.DB.Model(&producto).Where("id_producto = ?", product.Id_producto).Update("cantidad_disponible", cantidad_actualizada)
	}
	id_despacho := DespacharCompra(compra.Id_compra)
	c.JSON(http.StatusOK, gin.H{"id_compra": compra.Id_compra, "id_despacho": id_despacho})
}
func pedirProveedor(requerido int, product models.Producto) {
	conn, err := grpc.Dial("10.10.10.181:9000", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic("No se pudo conectar con el servidor")
	}
	serviceClient := pb.NewProveedorClient(conn)
	serviceClient.SuministrarProductos(context.Background(), &pb.Producto{Cantidad: int32((-1 * requerido)), Nombre: product.Nombre, Id: int32(product.Id_producto)})
}
func DespacharCompra(id_compra int) int {
	var despacho DespachoInput
	id_despacho := rand.Intn(8000000) + 1000000
	despacho.Id_compra = id_compra
	despacho.Id_despacho = id_despacho

	conn, err := amqp.Dial("amqp://usuario1:contrasena@10.10.10.230:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"despacho", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")
	body, _ := json.Marshal(despacho)

	err = ch.PublishWithContext(context.Background(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	return id_despacho
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
