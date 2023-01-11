package main

import (
	"encoding/json"
	"fmt"
	"servidor_3/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type DespachoInput struct {
	Id_compra   int `json:"id_compra" binding:"required"`
	Id_despacho int `json:"id_despacho" binding:"required"`
}

func main() {
	models.ConnectDatabase()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	var despacho DespachoInput
	var dsp models.Despacho

	for d := range msgs {
		json.Unmarshal(d.Body, &despacho)
		dsp.Id_despacho = despacho.Id_despacho
		dsp.Id_compra = despacho.Id_compra
		dsp.Estado = "RECIBIDO"
		models.DB.Create(&dsp)
	}
}
func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("%s: %s", msg, err)
	}
}
