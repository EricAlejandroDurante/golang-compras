package main

import (
	"servidor_3/models"
	"time"
)

func main() {
	models.ConnectDatabase()
	for {
		models.DB.Exec("UPDATE despacho SET estado ='ENTREGADO' WHERE estado = 'EN_TRANSITO';")
		models.DB.Exec("UPDATE despacho SET estado ='EN_TRANSITO' WHERE estado = 'RECIBIDO';")
		time.Sleep(60 * time.Second)
	}
}
