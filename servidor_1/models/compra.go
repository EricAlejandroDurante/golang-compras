package models

type Compra struct {
    Id_compra int `json:"id_compra" gorm:"primaryKey;auto_increment;not_null"`
    Id_cliente int `json:"id_cliente"`
    table string `gorm:"-"`
}

func (p Compra) TableName() string {
    return "compra"
}
