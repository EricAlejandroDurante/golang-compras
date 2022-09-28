package models

type Cliente struct {
    Id_cliente int `json:"id_cliente"`
    Nombre string `json:"nombre"`
    Contrasena string `json:"contrasena"`
    table string `gorm:"-"`
}

func (p Cliente) TableName() string {
    return "cliente"
}
