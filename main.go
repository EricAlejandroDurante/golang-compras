package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	//"tarea_1_sds/models"
)

type ProductoInput struct {
	Nombre   string `json:"nombre"`
	Cantidad int    `json:"cantidad_disponible"`
	Precio   int    `json:"precio_unitario"`
}

type ClienteInput struct {
	Id_Cliente int    `json:"id_cliente"`
	Contrasena string `json:"contrasena"`
}
type AccesoResponse struct {
	Acceso_valido bool `json:"acceso_valido"`
}

type CrearProductResponse struct {
	response int `json:"response"`
}

type Producto []struct {
	Id_producto         int    `json:"id_producto" gorm:"primaryKey;auto_increment;not_null"`
	Nombre              string `json:"nombre"`
	Cantidad_disponible int    `json:"cantidad_disponible"`
	Precio_unitario     int    `json:"precio_unitario"`
}

type EstadisticasResponse struct {
	Producto_mas_vendido    int `json:"producto_mas_vendido"`
	Producto_menos_vendido  int `json:"producto_menos_vendido"`
	Producto_mayor_ganancia int `json:"producto_mayor_ganancia"`
	Producto_menor_ganancia int `json:"producto_menor_ganancia"`
}

type CompraInput struct {
	Id_cliente int            `json:"id_cliente" binding:"required"`
	Productos  []DetalleInput `json:"productos" binding:"required"`
}

type DetalleInput struct {
	Id_producto int `json:"id_producto" binding:"required"`
	Cantidad    int `json:"cantidad" binding:"required"`
}

type DeleteResponse struct {
	Id_producto int `json:"id_producto" binding:"required"`
}

func main() {
	fmt.Println("Bienvenido")
loopMain:
	for {
		fmt.Println("Opciones:")
		fmt.Println("1. Iniciar sesión como cliente")
		fmt.Println("2. Iniciar sesión como administrador")
		fmt.Println("3. Salir")
		fmt.Printf("Ingrese una opción: ")
		var opciones string
		fmt.Scan(&opciones)
		opciones = strings.TrimRight(opciones, "\n")
		switch opciones {
		case "1":
			UserOption()
		case "2":
			AdminLog("12345678")
		case "3":
			break loopMain
		}
	}
}
func UserOption() {
	var id int
	var password string

	fmt.Printf("Ingrese su id: ")
	fmt.Scan(&id)

	fmt.Printf("Ingrese su contraseña: ")
	fmt.Scan(&password)
	password = strings.TrimRight(password, "\n")

	input := ClienteInput{Id_Cliente: id, Contrasena: password}
	b, _ := json.Marshal(input)

	resp, _ := http.Post("http://127.0.0.1:5000/api/clientes/iniciar_sesion", "application/json", bytes.NewBuffer(b))
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	bytes := []byte(body)
	var accesoResponse AccesoResponse
	json.Unmarshal(bytes, &accesoResponse)
	if accesoResponse.Acceso_valido == true {
		fmt.Println("Inicio de sesion exitoso")
		OpcionesClientes(id)
	} else {
		fmt.Println("Error, no hay ninguna coincidencia con los datos ingresados.")
	}
}

func OpcionesClientes(id int) {
testLoop:
	for {
		fmt.Printf("\n")
		fmt.Println("Opciones:\n1. Ver lista de productos\n2. Hacer compra\n3 Salir")
		fmt.Printf("Ingrese una opcion: ")
		var opcion string
		fmt.Scan(&opcion)
		switch opcion {
		case "1":
			ListarProductos()
		case "2":
			hacerCompra(id)
		case "3":
			fmt.Println("")
			break testLoop
		}
	}
}

func ListarProductos() {
	resp, _ := http.Get("http://127.0.0.1:5000/api/productos")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	bytes := []byte(body)
	var producto Producto
	json.Unmarshal([]byte(bytes), &producto)
	for _, objeto := range producto {
		fmt.Printf("%d;%s;%d por unidad;%d disponibles\n", objeto.Id_producto, objeto.Nombre, objeto.Precio_unitario, objeto.Cantidad_disponible)
	}
}

func hacerCompra(id int) {
	var cant_productos int
	var opcion string
	var suma int
	var id_producto int
	var cantidad int
	var compra CompraInput
	var detalle DetalleInput
	///////////////////////////////////////////////////////////////
	resp2, _ := http.Get("http://127.0.0.1:5000/api/productos")
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	bytes1 := []byte(body2)
	var producto Producto
	json.Unmarshal([]byte(bytes1), &producto)
	///////////////////////////////////////////////////////////////
	var montototal int
	compra.Id_cliente = id
	fmt.Printf("Ingrese cantidad de productos a comprar: ")
	fmt.Scan(&cant_productos)
	for j := 0; j < cant_productos; j++ {
		var cantidad_nueva int
		var montototalAux int
		var prod ProductoInput
		fmt.Printf("Ingrese producto %d par id-cantidad: ", j+1)
		fmt.Scan(&opcion)
		comando := strings.Split(opcion, "-")
		_, _ = fmt.Sscan(comando[0], &id_producto)
		_, _ = fmt.Sscan(comando[1], &cantidad)

		for _, objeto := range producto {
			if objeto.Id_producto == id_producto {
				montototalAux += objeto.Precio_unitario * cantidad
				cantidad_nueva = objeto.Cantidad_disponible - cantidad
				if cantidad_nueva >= 0 {
					prod.Precio = objeto.Precio_unitario
					prod.Nombre = objeto.Nombre
					prod.Cantidad = cantidad_nueva
				}
			}
		}
		if cantidad_nueva < 0 {
			fmt.Println("La cantidad seleccionada esta excedida, selecciona nuevamente ")
			j--
			continue
		}

		montototal += montototalAux
		detalle.Id_producto = id_producto
		detalle.Cantidad = cantidad
		compra.Productos = append(compra.Productos, detalle)
		str := strconv.Itoa(id_producto)

		url := "http://127.0.0.1:5000/api/productos/" + str

		jsonProd, _ := json.Marshal(prod)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonProd))
		req.Header.Set("Accept", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		suma += cantidad
	}

	b, _ := json.Marshal(compra)
	_, _ = http.Post("http://127.0.0.1:5000/api/compras", "application/json", bytes.NewBuffer(b))
	fmt.Println("Gracias por su compra!")
	fmt.Printf("Cantidad de productos comprados: %d\n", suma)
	fmt.Printf("Monto total de la compra: %d\n", montototal)
}

func CrearProducto() {
	var nombre string
	var disponibilidad int
	var precio int
	fmt.Printf("Ingrese el nombre: ")
	fmt.Scan(&nombre)
	fmt.Printf("Ingrese la disponibilidad: ")
	fmt.Scan(&disponibilidad)
	fmt.Printf("Ingrese el precio unitario: ")
	fmt.Scan(&precio)

	input := ProductoInput{Nombre: nombre, Cantidad: disponibilidad, Precio: precio}
	b, _ := json.Marshal(input)
	resp, _ := http.Post("http://127.0.0.1:5000/api/productos", "application/json", bytes.NewBuffer(b))
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	bytes := []byte(body)

	var accesoResponse CrearProductResponse
	json.Unmarshal(bytes, &accesoResponse)
	if reflect.TypeOf(accesoResponse.response).Kind() == reflect.Int {
		fmt.Println("Producto Creado Exitosamente")
	} else {
		fmt.Println("No se ha podido crear producto")
	}
}

func EliminarProducto() {
	var id string
	//var accesoResponse CrearProductResponse
	fmt.Printf("Ingrese id producto: ")
	fmt.Scan(&id)
	url := "http://127.0.0.1:5000/api/productos/" + id
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var response DeleteResponse
	json.Unmarshal([]byte(body), &response)
	if response.Id_producto == 0 {
		fmt.Println("Producto no existe o no se ha podido eliminar")
	} else {
		fmt.Printf("Producto eliminado %d!\n", response.Id_producto)
	}
}

func Estadisticas() {
	var respuesta EstadisticasResponse
	resp, _ := http.Get("http://127.0.0.1:5000/api/estadisticas")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	bytes := []byte(body)
	json.Unmarshal(bytes, &respuesta)
	fmt.Printf("Producto mas vendido: %d\n", respuesta.Producto_mas_vendido)
	fmt.Printf("Producto menos vendido: %d\n", respuesta.Producto_menos_vendido)
	fmt.Printf("Producto mayor ganancia: %d\n", respuesta.Producto_mayor_ganancia)
	fmt.Printf("Producto menor ganancia: %d\n", respuesta.Producto_menor_ganancia)
}

func AdminLog(adminPass string) {
	var password string
	fmt.Printf("Ingrese contraseña de administrador: ")
	fmt.Scan(&password)
	password = strings.TrimRight(password, "\n")
	if password == adminPass {
		fmt.Println("Inicio de sesión exitoso")
		OpcionesAdmin()
	}
}

func OpcionesAdmin() {
opcionesLoop:
	for {
		fmt.Printf("\n")
		fmt.Println("Opciones:\n1. Ver lista de productos\n2. Crear producto\n3. Eliminar producto\n4. Ver estadísticas\n5. Salir")
		fmt.Printf("Ingrese una opcion: ")
		var opcion string
		fmt.Scan(&opcion)
		switch opcion {
		case "1":
			ListarProductos()
		case "2":
			CrearProducto()
		case "3":
			EliminarProducto()
		case "4":
			Estadisticas()
		case "5":
			fmt.Println("")
			break opcionesLoop
		}
	}
}
