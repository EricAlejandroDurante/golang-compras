package models

type Despacho struct {
	Id_despacho int    `json:"id_producto" gorm:"primaryKey;auto_increment;not_null"`
	Estado      string `json:"nombre"`
	Id_compra   int    `json:"cantidad_disponible"`
	table       string `gorm:"-"`
}

func (p Despacho) TableName() string {
	return "despacho"
}
