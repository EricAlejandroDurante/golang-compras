package models

type Detalle struct {
    Id_compra int `json:"id_compra" gorm:"primaryKey;auto_increment;not_null"`
    Id_producto int `json:"id_producto"`
    Cantidad int `json:"cantidad"`
    Fecha string `json:"fecha"`
    table string `gorm:"-"`
}

func (p Detalle) TableName() string {
    return "detalle"
}
