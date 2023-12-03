package models

import "time"

type Servisio struct {
	IDProducto  int64  `json:"Id_Producto"`
	Nombre      string `json:"Nombre"`     
	Descripsion string `json:"Descripsion"`
	Unidad      int64  `json:"Unidad"`     
	Creado      time.Time `json:"Creado"`     
	Actualizado time.Time `json:"Actualizado"`
}
